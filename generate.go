package enthistory

import (
	"errors"
	"fmt"
	"go/ast"
	"go/token"
	"os"
	"path/filepath"
	"reflect"
	"runtime/debug"
	"slices"
	"strconv"
	"strings"
	"time"

	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/mitchellh/mapstructure"

	"github.com/flume/enthistory/internal/schemast"

	"entgo.io/contrib/entgql"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/entc/load"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type HistoryOptions struct {
	InheritIdType    bool
	UpdatedBy        *UpdatedBy
	FieldProperties  *FieldProperties
	HistoryTimeIndex bool
}

type UpdatedBy struct {
	key       string
	valueType ValueType
}

type FieldProperties struct {
	Nillable  bool
	Immutable bool
}

const (
	ValueTypeInt ValueType = iota
	ValueTypeString
	ValueTypeUUID
)

type ValueType uint

func (ValueType) ValueType() string {
	return "ValueType"
}

type Option = func(*HistoryOptions)

// WithNillableFields allows you to set all tracked fields in history to Nillable
// except enthistory managed fields (history_time, ref, operation, & updated_by)
func WithNillableFields() Option {
	return func(config *HistoryOptions) {
		config.FieldProperties.Nillable = true
	}
}

// WithImmutableFields allows you to set all tracked fields in history to Immutable
func WithImmutableFields() Option {
	return func(config *HistoryOptions) {
		config.FieldProperties.Immutable = true
	}
}

// WithUpdatedBy sets the key and type for pulling updated_by from the context,
// usually done via a middleware to track which users are making which changes
func WithUpdatedBy(key string, valueType ValueType) Option {
	return func(config *HistoryOptions) {
		config.UpdatedBy = &UpdatedBy{
			key:       key,
			valueType: valueType,
		}
	}
}

// WithHistoryTimeIndex allows you to add an index to the "history_time" fields
func WithHistoryTimeIndex() Option {
	return func(config *HistoryOptions) {
		config.HistoryTimeIndex = true
	}
}

// WithInheritIdType allows you to set the history schema id type to match the original schema id type,
// instead of defaulting to int. Otherwise, the history schema id type will default to int.
func WithInheritIdType() Option {
	return func(config *HistoryOptions) {
		config.InheritIdType = true
	}
}

func Generate(schemaPath string, schemas []ent.Interface, options ...Option) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered from panic: %v\n\nstack:%s", r, debug.Stack())
		}
	}()

	opts := new(HistoryOptions)
	opts.FieldProperties = new(FieldProperties)
	opts.UpdatedBy = new(UpdatedBy)
	for _, opt := range options {
		opt(opts)
	}

	graph, err := entc.LoadGraph(schemaPath, &gen.Config{})
	if err != nil {
		return fmt.Errorf("failed loading ent graph: %v", err)
	}

	err = removeOldGenerated(schemaPath, graph.Schemas)
	if err != nil {
		return err
	}

	ctx, err := schemast.Load(schemaPath, map[string]schemast.Annotator{
		"History": entHistory,
	})
	if err != nil {
		return fmt.Errorf("failed loading schema: %v", err)
	}
	if ctx.SchemaPackage == nil {
		return fmt.Errorf("failed loading schema package: %v", err)
	}

	var mutations []schemast.Mutator
	for _, s := range schemas {
		hfields, herr := historyFields(s, deref(opts))
		if herr != nil {
			return herr
		}

		typeof := reflect.TypeOf(s)
		var schemaName string
		if typeof.Kind() == reflect.Ptr {
			schemaName = typeof.Elem().Name()
		} else {
			schemaName = typeof.Name()
		}
		upsert := schemast.UpsertSchema{
			Name:   fmt.Sprintf("%sHistory", schemaName),
			Fields: hfields,
		}

		if opts.HistoryTimeIndex {
			upsert.Indexes = append(upsert.Indexes, index.Fields("history_time"))
		}

		upsert.Annotations = handleAnnotation(schemaName, s.Annotations())
		mutations = append(mutations, &upsert)
	}

	if err = schemast.Mutate(ctx, mutations...); err != nil {
		return err
	}
	return ctx.Print(schemaPath, schemast.Header("// Code generated by enthistory, DO NOT EDIT."))
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

func handleAnnotation(schemaName string, ants []schema.Annotation) []schema.Annotation {
	annotations := slices.Clone(ants)

	idx := slices.IndexFunc(annotations, func(sc schema.Annotation) bool {
		_, ok := sc.(entsql.Annotation)
		return ok
	})

	if idx == -1 {
		annotations = append(annotations, entsql.Annotation{Table: fmt.Sprintf("%s_history", strings.ToLower(schemaName))})
	} else {
		ant, ok := annotations[idx].(entsql.Annotation)
		if ok {
			ant.Table = fmt.Sprintf("%s_history", strings.ToLower(schemaName))
			annotations[idx] = ant
		}
	}

	idx = slices.IndexFunc(annotations, func(sc schema.Annotation) bool {
		_, ok := sc.(entsql.Annotation)
		return ok
	})
	if idx == -1 {
		annotations = append(annotations, Annotations{IsHistory: true})
	} else {
		ant, ok := annotations[idx].(Annotations)
		if ok {
			ant.IsHistory = true
			annotations[idx] = ant
		}
	}
	return annotations
}

func historyFields(schema ent.Interface, opts HistoryOptions) ([]ent.Field, error) {
	var fields []ent.Field
	var idField ent.Field = field.Int("ref").Immutable().Optional()
	for _, f := range schema.Fields() {
		if f.Descriptor().Name == "id" {
			if opts.InheritIdType {
				idField = f
			}
		}
	}
	fields = append(fields, idField)
	fields = append(fields, field.Time("history_time").Default(time.Now).Immutable())
	fields = append(fields, field.Enum("operation").GoType(OpType("")).Immutable())
	ref, err := refField(idField.Descriptor().Info, true)
	if err != nil {
		return nil, err
	}
	fields = append(fields, ref)

	if opts.UpdatedBy != nil {
		ubf, uerr := updatedByField(opts.UpdatedBy.valueType, true)
		if uerr != nil {
			return nil, uerr
		}
		fields = append(fields, ubf)
	}

	for _, f := range schema.Fields() {
		if f.Descriptor().Name == "id" {
			continue
		}

		fields = append(fields, prepareField(opts, f))
	}

	for _, m := range schema.Mixin() {
		fields = append(fields, m.Fields()...)
	}

	return fields, nil
}

type cleanedField struct {
	descriptor *field.Descriptor
}

func (c *cleanedField) Descriptor() *field.Descriptor {
	return c.descriptor
}

func prepareField(opts HistoryOptions, f ent.Field) ent.Field {
	descriptor := f.Descriptor()
	descriptor.Validators = nil

	if opts.FieldProperties.Nillable {
		descriptor.Nillable = true
	}
	if opts.FieldProperties.Immutable {
		descriptor.Immutable = true
	}

	return &cleanedField{descriptor: descriptor}
}

func refField(IdType *field.TypeInfo, entqlEnabled bool) (ent.Field, error) {
	switch IdType.String() {
	case "int":
		f := field.Int("ref").Immutable().Optional()
		if entqlEnabled {
			f = f.Annotations(entgql.Annotation{Type: "ID"})
		}
		return f, nil
	case "string":
		f := field.String("ref").Immutable().Optional()
		if entqlEnabled {
			f = f.Annotations(entgql.Annotation{Type: "ID"})
		}
		return f, nil
	case "uuid.UUID":
		equal := IdType.RType.TypeEqual(reflect.TypeOf(uuid.UUID{}))
		if !equal {
			return nil, errors.New("unsupported uuid type")
		}
		f := field.UUID("ref", uuid.UUID{}).Immutable().Optional()
		if entqlEnabled {
			f = f.Annotations(entgql.Annotation{Type: "ID"})
		}
		return f, nil
	}
	return nil, errors.New("unsupported ref type")
}

func updatedByField(valueType ValueType, entgqlEnabled bool) (ent.Field, error) {
	switch valueType {
	case ValueTypeInt:
		f := field.Int("updated_by").Optional().Nillable().Immutable()
		if entgqlEnabled {
			f = f.Annotations(entgql.Annotation{Type: "ID"})
		}
		return f, nil
	case ValueTypeString:
		f := field.String("updated_by").Optional().Nillable().Immutable()
		if entgqlEnabled {
			f = f.Annotations(entgql.Annotation{Type: "ID"})
		}
		return f, nil
	case ValueTypeUUID:
		f := field.UUID("updated_by", uuid.UUID{}).Optional().Nillable().Immutable()
		if entgqlEnabled {
			f = f.Annotations(entgql.Annotation{Type: "ID"})
		}
		return f, nil
	}
	return nil, errors.New("unsupported updated_by type")
}

func removeOldGenerated(schemaPath string, schemas []*load.Schema) error {
	for _, schema := range schemas {
		abs, err := filepath.Abs(schemaPath)
		if err != nil {
			return err
		}

		path := fmt.Sprintf("%v/%v.go", abs, fmt.Sprintf("%s_history", strings.ToLower(schema.Name)))

		err = os.RemoveAll(path)
		if err != nil {
			return err
		}
	}
	return nil
}
