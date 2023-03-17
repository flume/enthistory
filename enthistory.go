package enthistory

import (
	"embed"
	"fmt"
	"os"
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

type Config struct {
	UpdatedBy *UpdatedBy
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

func NewHistoryExtension(opts ...ExtensionOption) *HistoryExtension {
	extension := &HistoryExtension{
		config: &Config{},
	}
	for _, opt := range opts {
		opt(extension)
	}
	return extension
}

type templateInfo struct {
	Schema             *load.Schema
	TableName          string
	OriginalTableName  string
	WithUpdatedBy      bool
	UpdatedByValueType string
}

func (*HistoryExtension) Templates() []*gen.Template {
	return []*gen.Template{
		parseTemplate("historyFromMutation", "templates/historyFromMutation.tmpl"),
		parseTemplate("historyQuery", "templates/historyQuery.tmpl"),
		parseTemplate("client", "templates/client.tmpl"),
	}
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
	templateInfo := templateInfo{
		TableName:         fmt.Sprintf("%v_history", getSchemaTableName(schema)),
		OriginalTableName: schema.Name,
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

	// merge the original schema onto the history schema
	historySchema.Name = fmt.Sprintf("%vHistory", schema.Name)
	historySchema.Fields = append(historySchema.Fields, createHistoryFields(schema.Fields)...)
	historySchema.Annotations = map[string]any{
		"EntSQL": map[string]any{
			"table": templateInfo.TableName,
		},
	}

	templateInfo.Schema = historySchema
	// Get path to write new history schema file
	path, err := getHistorySchemaPath(schema)
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
		err := removeOldGenerated(g.Schemas)
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
