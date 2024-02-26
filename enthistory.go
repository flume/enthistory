package enthistory

import (
	"embed"
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"entgo.io/ent/schema/field"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/entc/load"
)

var (
	//go:embed templates/*
	_templates embed.FS
)

type UpdatedBy struct {
	key       string
	valueType ValueType
}

type FieldProperties struct {
	Nillable  bool
	Immutable bool
}

type Config struct {
	InheritIdType    bool
	UpdatedBy        *UpdatedBy
	Auditing         bool
	SchemaPath       string
	FieldProperties  *FieldProperties
	HistoryTimeIndex bool
}

func (c Config) Name() string {
	return "HistoryConfig"
}

// HistoryExtension implements entc.Extension.
type HistoryExtension struct {
	entc.DefaultExtension
	config *Config
}

type ExtensionOption = func(*HistoryExtension)

// WithInheritIdType allows you to set the history schema id type to match the original schema id type,
// instead of defaulting to int. Otherwise, the history schema id type will default to int.
func WithInheritIdType() ExtensionOption {
	return func(ex *HistoryExtension) {
		ex.config.InheritIdType = true
	}
}

// WithUpdatedBy sets the key and type for pulling updated_by from the context,
// usually done via a middleware to track which users are making which changes
func WithUpdatedBy(key string, valueType ValueType) ExtensionOption {
	return func(ex *HistoryExtension) {
		ex.config.UpdatedBy = &UpdatedBy{
			key:       key,
			valueType: valueType,
		}
	}
}

// WithAuditing allows you to turn on the code generation for the `.Audit()` method
func WithAuditing() ExtensionOption {
	return func(ex *HistoryExtension) {
		ex.config.Auditing = true
	}
}

// WithSchemaPath allows you to set an alternative schemaPath
// Defaults to "./schema"
func WithSchemaPath(schemaPath string) ExtensionOption {
	return func(ex *HistoryExtension) {
		ex.config.SchemaPath = schemaPath
	}
}

// WithNillableFields allows you to set all tracked fields in history to Nillable
// except enthistory managed fields (history_time, ref, operation, & updated_by)
func WithNillableFields() ExtensionOption {
	return func(ex *HistoryExtension) {
		ex.config.FieldProperties.Nillable = true
	}
}

// WithImmutableFields allows you to set all tracked fields in history to Immutable
func WithImmutableFields() ExtensionOption {
	return func(ex *HistoryExtension) {
		ex.config.FieldProperties.Immutable = true
	}
}

// WithHistoryTimeIndex allows you to add an index to the "history_time" fields
func WithHistoryTimeIndex() ExtensionOption {
	return func(ex *HistoryExtension) {
		ex.config.HistoryTimeIndex = true
	}
}

func NewHistoryExtension(opts ...ExtensionOption) *HistoryExtension {
	extension := &HistoryExtension{
		// Set configuration defaults that can get overridden with ExtensionOption
		config: &Config{
			SchemaPath:      "./schema",
			Auditing:        false,
			FieldProperties: &FieldProperties{},
		},
	}
	for _, opt := range opts {
		opt(extension)
	}

	return extension
}

type templateInfo struct {
	Schema               *load.Schema
	EntqlEnabled         bool
	IdType               string
	SchemaPkg            string
	TableName            string
	OriginalTableName    string
	WithUpdatedBy        bool
	UpdatedByValueType   string
	WithHistoryTimeIndex bool
	InheritIdType        bool
}

func (h *HistoryExtension) Templates() []*gen.Template {
	templates := []*gen.Template{
		parseTemplate("historyFromMutation", "templates/historyFromMutation.tmpl"),
		parseTemplate("historyQuery", "templates/historyQuery.tmpl"),
		parseTemplate("client", "templates/client.tmpl"),
	}
	if h.config.Auditing {
		templates = append(templates, parseTemplate("auditing", "templates/auditing.tmpl"))
	}
	return templates
}

// Hooks of the HistoryExtension.
func (h *HistoryExtension) Hooks() []gen.Hook {
	return []gen.Hook{
		h.generateHistorySchemas,
	}
}

func (h *HistoryExtension) Annotations() []entc.Annotation {
	return []entc.Annotation{
		h.config,
	}
}

var (
	schemaTemplate = template.Must(template.ParseFS(_templates, "templates/schema.tmpl"))
)

func (h *HistoryExtension) generateHistorySchema(info templateInfo, schema *load.Schema, IdType *field.TypeInfo) (*load.Schema, error) {
	pkg, err := getPkgFromSchemaPath(h.config.SchemaPath)
	if err != nil {
		return nil, err
	}
	info.TableName = fmt.Sprintf("%v_history", getSchemaTableName(schema))
	info.SchemaPkg = pkg
	info.InheritIdType = h.config.InheritIdType

	if h.config != nil {
		if h.config.UpdatedBy != nil {
			valueType := h.config.UpdatedBy.valueType
			if valueType == ValueTypeInt {
				info.UpdatedByValueType = "Int"
			} else if valueType == ValueTypeString {
				info.UpdatedByValueType = "String"
			} else if valueType == ValueTypeUUID {
				info.UpdatedByValueType = "UUID"
			}
			info.WithUpdatedBy = true
		}
		info.WithHistoryTimeIndex = h.config.HistoryTimeIndex
	}

	switch IdType.String() {
	case "int":
		info.IdType = "Int"
	case "string":
		info.IdType = "String"
	case "uuid.UUID":
		info.IdType = "UUID"
	default:
		return nil, fmt.Errorf("unsupported id type: %s", IdType)
	}

	// Load new base history schema
	historySchema, err := loadHistorySchema(IdType, info.EntqlEnabled)
	if err != nil {
		return nil, err
	}

	updatedByField, err := getUpdatedByField(info.UpdatedByValueType, info.EntqlEnabled)
	if err != nil {
		return nil, err
	}

	if updatedByField != nil {
		historySchema.Fields = append(historySchema.Fields, updatedByField)
	}

	if info.WithHistoryTimeIndex {
		historySchema.Indexes = append(historySchema.Indexes, &load.Index{Fields: []string{"history_time"}})
	}

	var historyFields []*load.Field
	for _, f := range h.createHistoryFields(schema.Fields) {
		if f.Name == "id" && !info.InheritIdType {
			f.Default = false
			f.Info = &field.TypeInfo{Type: field.TypeInt}
		}
		historyFields = append(historyFields, f)
	}

	// merge the original schema onto the history schema
	historySchema.Name = fmt.Sprintf("%vHistory", schema.Name)
	historySchema.Fields = append(historySchema.Fields, historyFields...)
	historySchema.Annotations = map[string]any{
		"EntSQL": map[string]any{
			"table": info.TableName,
		},
		"History": map[string]any{
			"exclude":   true,
			"isHistory": true,
		},
	}
	historySchema.Annotations = mergeAnnotations(maps.Clone(schema.Annotations), maps.Clone(historySchema.Annotations))
	info.Schema = historySchema
	// Get path to write new history schema file
	path, err := h.getHistorySchemaPath(schema)
	if err != nil {
		return nil, err
	}
	// Create history schema file
	create, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	defer create.Close()
	// execute schemaTemplate at the history schema path
	if err = schemaTemplate.Execute(create, info); err != nil {
		return nil, err
	}
	return historySchema, nil
}

func mergeAnnotations(dest, src map[string]any) map[string]any {
	merged := maps.Clone(src)
	for k, v := range dest {
		if _, ok := merged[k]; !ok {
			merged[k] = v
		} else {
			destMap, destOk := v.(map[string]any)
			srcMap, srcOk := merged[k].(map[string]any)
			if destOk && srcOk {
				merged[k] = mergeAnnotations(destMap, srcMap)
			}
		}
	}
	return merged
}

func (h *HistoryExtension) generateHistorySchemas(next gen.Generator) gen.Generator {
	return gen.GenerateFunc(func(g *gen.Graph) error {
		err := h.removeOldGenerated(g.Schemas)
		if err != nil {
			return err
		}

		entqlEnabled, _ := g.FeatureEnabled("entql")
		var schemas []*load.Schema
		for _, schema := range g.Schemas {
			annotations := getHistoryAnnotations(schema)

			if annotations.Exclude {
				if !annotations.IsHistory {
					schemas = append(schemas, schema)
				}
				continue
			}

			var IdType *field.TypeInfo
			for _, node := range g.Nodes {
				if schema.Name == node.Name {
					IdType = node.ID.Type
				}
			}

			if IdType == nil {
				return fmt.Errorf("could not get id type for schema: %s", schema.Name)
			}

			info := templateInfo{
				OriginalTableName: schema.Name,
				EntqlEnabled:      entqlEnabled,
			}

			historySchema, err := h.generateHistorySchema(info, schema, IdType)
			if err != nil {
				return err
			}

			// add history schema to list of schemas in the graph
			schemas = append(schemas, schema, historySchema)
		}

		// Create a new graph
		graph, err := gen.NewGraph(g.Config, schemas...)
		if err != nil {
			return err
		}
		*g = *graph
		return next.Generate(g)
	})
}

func (h *HistoryExtension) getHistorySchemaPath(schema *load.Schema) (string, error) {
	abs, err := filepath.Abs(h.config.SchemaPath)
	if err != nil {
		return "", err
	}

	path := fmt.Sprintf("%v/%v.go", abs, fmt.Sprintf("%s_history", strings.ToLower(schema.Name)))
	return path, nil
}

func (h *HistoryExtension) removeOldGenerated(schemas []*load.Schema) error {
	for _, schema := range schemas {
		path, err := h.getHistorySchemaPath(schema)
		if err != nil {
			return err
		}

		err = os.RemoveAll(path)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *HistoryExtension) createHistoryFields(schemaFields []*load.Field) []*load.Field {
	historyFields := make([]*load.Field, len(schemaFields))
	fieldPropertiesSet := h.config.FieldProperties != nil
	i := len(history{}.Fields())
	if h.config.UpdatedBy != nil {
		i = i + 1
	}
	for j, field := range schemaFields {
		nillable := field.Nillable
		immutable := field.Immutable
		optional := field.Optional
		if fieldPropertiesSet && field.Name != "id" {
			nillable = h.config.FieldProperties.Nillable || nillable
			optional = h.config.FieldProperties.Nillable || optional
			immutable = h.config.FieldProperties.Immutable || immutable
		}

		newField := load.Field{
			Name:          field.Name,
			Info:          copyRef(field.Info),
			Tag:           field.Tag,
			Size:          copyRef(field.Size),
			Enums:         field.Enums,
			Unique:        false,
			Nillable:      nillable,
			Optional:      optional,
			Default:       field.Default,
			DefaultValue:  field.DefaultValue,
			DefaultKind:   field.DefaultKind,
			UpdateDefault: field.UpdateDefault,
			Immutable:     immutable,
			Validators:    field.Validators,
			StorageKey:    field.StorageKey,
			Position:      copyRef(field.Position),
			Sensitive:     field.Sensitive,
			SchemaType:    field.SchemaType,
			Annotations:   field.Annotations,
			Comment:       field.Comment,
		}
		if !field.Position.MixedIn {
			newField.Position = &load.Position{
				Index:      i,
				MixedIn:    false,
				MixinIndex: 0,
			}
			i += 1
		}
		historyFields[j] = &newField
	}
	return historyFields
}
