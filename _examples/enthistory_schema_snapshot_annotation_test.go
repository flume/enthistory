package _examples

import (
	"encoding/json"
	"entgo.io/ent/entc/gen"
	"github.com/stretchr/testify/assert"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
	"testing"
)

func TestSchemaSnapshotAnnotations(t *testing.T) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "./basic/ent/internal/schema.go", nil, parser.ParseComments)
	assert.NoError(t, err)

	assert.Len(t, node.Decls, 1)
	schemaDecl := node.Decls[0].(*ast.GenDecl)
	assert.Len(t, schemaDecl.Specs, 1)
	schemaSpec := schemaDecl.Specs[0].(*ast.ValueSpec)

	assert.Len(t, schemaSpec.Values, 1)
	assert.Len(t, schemaSpec.Names, 1)

	assert.Equal(t, "Schema", schemaSpec.Names[0].Name)

	schemaValue := schemaSpec.Values[0].(*ast.BasicLit).Value
	assert.NotEmpty(t, schemaValue)

	var g = struct {
		gen.Graph
		Features []string
	}{}
	err = json.Unmarshal([]byte(strings.Trim(schemaValue, "`")), &g)
	assert.NoError(t, err)

	for _, schema := range g.Schemas {
		if strings.Contains(strings.ToLower(schema.Name), "history") {
			historyAnnotations := schema.Annotations["History"]
			assert.NotEmpty(t, historyAnnotations)
			historyAnnotationsMap, ok := historyAnnotations.(map[string]any)
			assert.True(t, ok)
			assert.True(t, historyAnnotationsMap["exclude"].(bool))
			assert.True(t, historyAnnotationsMap["isHistory"].(bool))
		} else {
			assert.Empty(t, schema.Annotations["History"])
		}
	}
}
