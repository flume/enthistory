package enthistory

import (
	"go/ast"
	"go/token"
	"strconv"

	"entgo.io/ent/schema"
	"github.com/mitchellh/mapstructure"
)

type Annotations struct {
	Exclude   bool `json:"exclude,omitempty"`   // Will exclude history tracking for this schema
	IsHistory bool `json:"isHistory,omitempty"` // DO NOT APPLY TO ANYTHING EXCEPT HISTORY SCHEMAS
}

func (Annotations) Name() string {
	return "History"
}

func entHistory(annot schema.Annotation) (ast.Expr, bool, error) {
	m := &Annotations{}
	if err := mapstructure.Decode(annot, m); err != nil {
		return nil, false, err
	}
	c := &ast.CompositeLit{
		Type: selectorLit("enthistory", "Annotations"),
	}
	if m.IsHistory {
		c.Elts = append(c.Elts, structAttr("IsHistory", boolLit(m.IsHistory)))
	}
	if m.Exclude {
		c.Elts = append(c.Elts, structAttr("Exclude", boolLit(m.Exclude)))
	}
	return c, true, nil
}

func selectorLit(x, sel string) *ast.SelectorExpr {
	return &ast.SelectorExpr{
		X:   ast.NewIdent(x),
		Sel: ast.NewIdent(sel),
	}
}

func boolLit(lit bool) ast.Expr {
	return &ast.Ident{
		Name: strconv.FormatBool(lit),
	}
}

func structAttr(name string, val ast.Expr) ast.Expr {
	return &ast.KeyValueExpr{
		Key: &ast.BasicLit{
			Kind:  token.STRING,
			Value: name,
		},
		Value: val,
	}
}
