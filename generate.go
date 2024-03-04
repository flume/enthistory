package enthistory

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime/debug"
	"slices"
	"strings"
	"time"

	"golang.org/x/tools/go/packages"

	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"

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

var updatedBy *UpdatedBy

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
	updatedBy = &UpdatedBy{
		key:       key,
		valueType: valueType,
	}
	return func(config *HistoryOptions) {
		config.UpdatedBy = updatedBy
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
	for _, opt := range options {
		opt(opts)
	}

	if abs, aerr := filepath.Abs(schemaPath); aerr == nil {
		schemaPath = abs
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

	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedName | packages.NeedTypes | packages.NeedFiles | packages.NeedCompiledGoFiles | packages.NeedSyntax,
	}, schemaPath)
	if err != nil {
		return err
	}
	if len(pkgs) < 1 {
		return fmt.Errorf("missing package information for: %s", schemaPath)
	}

	filenames := make(map[string]string)
	for _, f := range pkgs[0].GoFiles {
		readFile, rerr := os.ReadFile(f)
		if rerr != nil {
			return rerr
		}
		filecontent := string(readFile)
		if strings.Contains(filecontent, "ent.Schema") {
			submatch := regexp.MustCompile(`type ([a-zA-Z0-9_]+) struct`).FindAllStringSubmatch(filecontent, -1)
			if len(submatch) > 0 {
				if len(submatch[0]) > 1 {
					split := strings.Split(f, "/")
					if len(split) > 0 {
						filename := strings.Split(split[len(split)-1], ".")[0]
						filenames[submatch[0][len(submatch[0])-1]] = fmt.Sprintf("%s_history.go", filename)
					}
				}
			}
		}
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

		if filename, ok := filenames[schemaName]; ok {
			upsert.FileName = filename
		}

		if opts.HistoryTimeIndex {
			upsert.Indexes = append(upsert.Indexes, index.Fields("history_time"))
		}

		if len(s.Mixin()) > 0 {
			upsert.Mixins = s.Mixin()
		}

		annotations, gerr := handleAnnotation(schemaName, s.Annotations())
		if gerr != nil {
			return gerr
		}
		upsert.Annotations = annotations
		mutations = append(mutations, &upsert)
	}

	if err = schemast.Mutate(ctx, mutations...); err != nil {
		return err
	}
	return ctx.Print(schemaPath, schemast.Header("Code generated by enthistory, DO NOT EDIT."))
}

func handleAnnotation(schemaName string, ants []schema.Annotation) ([]schema.Annotation, error) {
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

	var mergedAnnotations []schema.Annotation
	antMap := make(map[string][]schema.Annotation)
	for _, a := range annotations {
		name := a.Name()
		if named, ok := antMap[name]; !ok {
			antMap[name] = []schema.Annotation{a}
		} else {
			antMap[name] = append(named, a)
		}
	}

	for name, a := range antMap {
		if len(a) == 0 {
			continue
		}

		switch name {
		case "EntSQL":
			typed, err := typedSliceToType[schema.Annotation, entsql.Annotation](a)
			if err != nil {
				return nil, err
			}
			mergedAnnotations = append(mergedAnnotations, mergeAnnotations[entsql.Annotation](typed...))
		case "EntGQL":
			merged := reduce(a, func(agg entgql.Annotation, item schema.Annotation) entgql.Annotation {
				merged := agg.Merge(item)
				return merged.(entgql.Annotation)
			}, entgql.Annotation{})
			mergedAnnotations = append(mergedAnnotations, merged)
		default:
			mergedAnnotations = append(mergedAnnotations, a...)
		}
	}

	return mergedAnnotations, nil
}

type mergeable interface {
	schema.Annotation
	schema.Merger
}

func mergeAnnotations[T mergeable](from ...T) schema.Annotation {
	var t T
	if _, ok := any(t).(schema.Merger); !ok {
		return t
	}
	if len(from) == 0 {
		return t
	}
	for _, f := range from {
		var e = t.Merge(f)
		t = e.(T)
	}
	return t
}

func historyFields(schema ent.Interface, opts HistoryOptions) ([]ent.Field, error) {
	var fields []ent.Field
	var idField ent.Field = field.Int("id").Immutable()
	var ref = idField

	var managedId bool
	for _, f := range schema.Fields() {
		if f.Descriptor().Name == "id" {
			managedId = true
			ref = f
			if opts.InheritIdType {
				idField = f
			}
		}
	}
	if managedId {
		fields = append(fields, idField)
	}
	fields = append(fields, field.Time("history_time").Default(time.Now).Immutable())
	fields = append(fields, field.Enum("operation").GoType(OpType("")).Immutable())
	ref, err := refField(ref.Descriptor().Info, true)
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

	if descriptor.Info.Type != field.TypeJSON {
		if opts.FieldProperties.Nillable {
			descriptor.Nillable = true
		}
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
