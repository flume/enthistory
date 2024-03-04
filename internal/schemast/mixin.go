package schemast

import (
	"fmt"
	"go/ast"
	"reflect"
	"strings"

	"entgo.io/ent"
)

// AppendMixin adds a mixin to the returned values of the Mixins method of type typeName.
func (c *Context) AppendMixin(typeName string, mix ent.Mixin) error {
	mixType := reflect.TypeOf(mix)
	_ = mixType
	name := mixType.Name()
	path := mixType.PkgPath()

	split := strings.Split(path, "/")
	pkg := split[len(split)-1]

	var lit *ast.CompositeLit
	if pkg != c.SchemaPackage.Name {
		c.appendImport(typeName, fmt.Sprintf("\"%s\"", path))
		lit = structLit(&ast.SelectorExpr{
			X:   ast.NewIdent(pkg),
			Sel: ast.NewIdent(name),
		})
	} else {
		lit = structLit(ast.NewIdent(name))
	}

	return c.appendReturnItem(kindMixin, typeName, lit)
}
