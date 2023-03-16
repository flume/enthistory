package enthistory

import (
	"embed"
	"html/template"
	"os"

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
	UpdatedBy UpdatedBy
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
		ex.config.UpdatedBy = UpdatedBy{
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

			updatedByValueType := "String"
			if h.config != nil {
				if h.config.UpdatedBy.valueType == ValueTypeInt {
					updatedByValueType = "Int"
				}
			}

			// Load new base history schema
			historySchema, err := loadHistorySchema()
			if err != nil {
				return err
			}

			updatedByField, err := getUpdatedByField(updatedByValueType)
			if err != nil {
				return err
			}

			historySchema.Fields = append(historySchema.Fields, updatedByField)

			// merge the original schema onto the history schema
			tableName := mergeSchemaAndHistorySchema(historySchema, schema)

			// Get path to write new history schema file
			path, err := getHistorySchemaPath(schema)
			if err != nil {
				return err
			}
			// Create history schema file
			create, err := os.Create(path)
			if err != nil {
				return err
			}
			defer create.Close()

			templateInfo := templateInfo{
				Schema:             historySchema,
				TableName:          tableName,
				OriginalTableName:  schema.Name,
				UpdatedByValueType: updatedByValueType,
			}
			// execute schemaTemplate at the history schema path
			if err = schemaTemplate.Execute(create, templateInfo); err != nil {
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
