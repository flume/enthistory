package _examples

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
	"strings"
	"testing"

	"entgo.io/ent/entc/gen"
	"github.com/stretchr/testify/assert"
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

	g := gen.Snapshot{}

	snapshot, err := trim([]byte(schemaValue))
	assert.NoError(t, err)
	err = json.Unmarshal(snapshot, &g)
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

func trim(line []byte) ([]byte, error) {
	start := bytes.IndexByte(line, '"')
	end := bytes.LastIndexByte(line, '"')
	if start == -1 || start >= end {
		return nil, fmt.Errorf("unexpected snapshot line %s", line)
	}
	l, err := strconv.Unquote(string(line[start : end+1]))
	if err != nil {
		return nil, err
	}
	return []byte(l), nil
}
