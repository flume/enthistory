package enthistory

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

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
	UpdatedBy           *UpdatedBy
	Auditing            bool
	OriginSchemaPath    string
	HisotrySchemaPath   string
	OriginSchemaFullPkg string
	FieldProperties     *FieldProperties
	HistoryTimeIndex    bool
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
		ex.config.OriginSchemaPath = schemaPath
	}
}

func WithOriginSchemaPath(schemaPath string) ExtensionOption {
	return func(ex *HistoryExtension) {
		ex.config.OriginSchemaPath = schemaPath
	}
}

// WithHistorySchemaPath allows you to set an alternative for history schemaPath
// Defaults to "./schema"
func WithHisotrySchemaPath(schemaPath string) ExtensionOption {
	return func(ex *HistoryExtension) {
		ex.config.HisotrySchemaPath = schemaPath
	}
}

func WithOriginSchemaFullPkg(schemaPkg string) ExtensionOption {
	return func(ex *HistoryExtension) {
		ex.config.OriginSchemaFullPkg = schemaPkg
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
			OriginSchemaPath:  "./schema",
			HisotrySchemaPath: "./schema",
			Auditing:          false,
			FieldProperties:   &FieldProperties{},
		},
	}
	for _, opt := range opts {
		opt(extension)
	}

	return extension
}

type templateInfo struct {
	Schema               *load.Schema
	OriginSchemaPath     string
	SchemaPkg            string
	OriginSchemaFullPkg  string
	TableName            string
	OriginalTableName    string
	WithUpdatedBy        bool
	UpdatedByValueType   string
	WithHistoryTimeIndex bool
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

func (h *HistoryExtension) generateHistorySchema(schema *load.Schema) (*load.Schema, error) {
	pkg, err := getPkgFromSchemaPath(h.config.HisotrySchemaPath)
	if err != nil {
		return nil, err
	}
	origin_pkg, err := getPkgFromSchemaPath(h.config.OriginSchemaPath)
	if err != nil {
		return nil, err
	}

	templateInfo := templateInfo{
		TableName:         fmt.Sprintf("%v_history", getSchemaTableName(schema)),
		OriginalTableName: schema.Name,
		SchemaPkg:         pkg,
	}

	if h.config.HisotrySchemaPath != h.config.OriginSchemaPath {
		otablename := fmt.Sprintf("%s.%s", origin_pkg, schema.Name)
		templateInfo.OriginalTableName = otablename
		templateInfo.OriginSchemaPath = h.config.OriginSchemaPath
		templateInfo.OriginSchemaFullPkg = h.config.OriginSchemaFullPkg
	}

	if h.config != nil {
		if h.config.UpdatedBy != nil {
			valueType := h.config.UpdatedBy.valueType
			if valueType == ValueTypeInt {
				templateInfo.UpdatedByValueType = "Int"
			} else if valueType == ValueTypeString {
				templateInfo.UpdatedByValueType = "String"
			}
			templateInfo.WithUpdatedBy = true
		}
		templateInfo.WithHistoryTimeIndex = h.config.HistoryTimeIndex

	}

	// Load new base history schema
	historySchema, err := loadHistorySchema()
	if err != nil {
		return nil, err
	}

	updatedByField, err := getUpdatedByField(templateInfo.UpdatedByValueType)
	if err != nil {
		return nil, err
	}

	if updatedByField != nil {
		historySchema.Fields = append(historySchema.Fields, updatedByField)
	}

	if templateInfo.WithHistoryTimeIndex {
		historySchema.Indexes = append(historySchema.Indexes, &load.Index{Fields: []string{"history_time"}})
	}

	// merge the original schema onto the history schema
	historySchema.Name = fmt.Sprintf("%vHistory", schema.Name)
	historySchema.Fields = append(historySchema.Fields, h.createHistoryFields(schema.Fields)...)
	historySchema.Annotations = map[string]any{
		"EntSQL": map[string]any{
			"table": templateInfo.TableName,
		},
	}

	templateInfo.Schema = historySchema
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
	if err = schemaTemplate.Execute(create, templateInfo); err != nil {
		return nil, err
	}
	return historySchema, nil
}

func (h *HistoryExtension) generateHistorySchemas(next gen.Generator) gen.Generator {
	return gen.GenerateFunc(func(g *gen.Graph) error {
		err := h.removeOldGenerated(g.Schemas)
		if err != nil {
			return err
		}

		var schemas []*load.Schema
		for _, schema := range g.Schemas {
			annotations := getHistoryAnnotations(schema)

			if annotations.Exclude {
				if !annotations.IsHistory {
					schemas = append(schemas, schema)
				}
				continue
			}

			historySchema, err := h.generateHistorySchema(schema)
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
		return next.Generate(graph)
	})
}

func (h *HistoryExtension) getHistorySchemaPath(schema *load.Schema) (string, error) {
	abs, err := filepath.Abs(h.config.HisotrySchemaPath)
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
	i := 4
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
