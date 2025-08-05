package enthistory

import (
	"go/ast"
	"go/token"
	"strconv"

	"entgo.io/ent"

	"entgo.io/ent/schema"
	"github.com/go-viper/mapstructure/v2"
)

type Annotations struct {
	IsHistory bool `json:"isHistory,omitempty"`
	// If you would like to add custom annotations to the history table,
	// otherwise it will default to the same annotations as the original table
	Annotations []schema.Annotation
	// if you would like to add custom mixins to the history table,
	// otherwise it will default to the same mixins as the original table
	Mixins []ent.Mixin
	// Global triggers for tracking history, can be any combination of OpTypeInsert, OpTypeUpdate, OpTypeDelete,
	// nil value will default to all triggers, to exclude all triggers set to an empty slice,
	// schema specific triggers will override these global triggers
	Triggers []OpType `json:"triggers,omitempty"`

	// Deprecated: Has no effect anymore, models must be tracked manually in the entc config
	Exclude bool `json:"exclude,omitempty"`
}

func (Annotations) Name() string {
	return "History"
}

func (m Annotations) Merge(other schema.Annotation) schema.Annotation {
	var ant Annotations
	switch o := other.(type) {
	case Annotations:
		ant = o
	case *Annotations:
		if o != nil {
			ant = *o
		}
	default:
		return m
	}
	if !m.IsHistory {
		m.IsHistory = ant.IsHistory
	}
	if len(ant.Annotations) > 0 {
		m.Annotations = append(m.Annotations, ant.Annotations...)
		cleaned, err := cleanAnnotations(m.Annotations)
		if err == nil {
			m.Annotations = cleaned
		}
	}
	if len(ant.Mixins) > 0 {
		m.Mixins = append(m.Mixins, ant.Mixins...)
	}
	if m.Triggers == nil {
		m.Triggers = ant.Triggers
	}
	return m
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
	if m.Triggers != nil {
		c.Elts = append(c.Elts, structAttr("Triggers", optypeSliceLit(m.Triggers)))
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

func optypeSliceLit(opTypes []OpType) ast.Expr {
	c := &ast.CompositeLit{
		Type: &ast.ArrayType{
			Elt: ast.NewIdent("enthistory.OpType"),
		},
	}
	for _, v := range opTypes {
		name, err := v.Name()
		if err != nil {
			continue
		}
		c.Elts = append(c.Elts, selectorLit("enthistory", name))
	}
	return c
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
