package enthistory

import (
	"embed"
	_ "embed"
	"html/template"
	"os"
	"strings"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/entc/load"
)

var (
	//go:embed templates/*
	_templates embed.FS
)

// HistoryExtension implements entc.Extension.
type HistoryExtension struct {
	entc.DefaultExtension
	userIdKey UserIdKey
}

func NewHistoryExtension(userIdKey UserIdKey) *HistoryExtension {
	return &HistoryExtension{
		userIdKey: userIdKey,
	}
}

type templateInfo struct {
	Schema            *load.Schema
	TableName         string
	OriginalTableName string
}

func (*HistoryExtension) Templates() []*gen.Template {
	return []*gen.Template{
		parseTemplate("historyFromMutation", "templates/historyFromMutation.tmpl"),
		parseTemplate("historyQuery", "templates/historyQuery.tmpl"),
		parseTemplate("client", "templates/client.tmpl"),
	}
}

// Hooks of the HistoryExtension.
func (*HistoryExtension) Hooks() []gen.Hook {
	return []gen.Hook{
		generateHistorySchemas,
	}
}

func (s *HistoryExtension) Annotations() []entc.Annotation {
	return []entc.Annotation{
		s.userIdKey,
	}
}

var (
	schemaTemplate = template.Must(template.ParseFS(_templates, "templates/schema.tmpl"))
)

func generateHistorySchemas(next gen.Generator) gen.Generator {
	return gen.GenerateFunc(func(g *gen.Graph) error {
		err := removeOldGenerated(g.Schemas)
		if err != nil {
			return err
		}

		var schemas []*load.Schema
		for _, schema := range g.Schemas {
			// Old history schemas should be skipped
			if strings.HasSuffix(schema.Name, "History") {
				continue
			}

			// Load new base history schema
			historySchema, err := loadHistorySchema()
			if err != nil {
				return err
			}

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
				Schema:            historySchema,
				TableName:         tableName,
				OriginalTableName: schema.Name,
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
