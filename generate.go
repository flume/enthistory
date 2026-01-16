package enthistory

import (
	"context"
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

	"golang.org/x/sync/errgroup"

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
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type HistoryOptions struct {
	InheritIdType    bool
	UpdatedBy        *UpdatedBy
	FieldProperties  *FieldProperties
	HistoryTimeIndex bool
	Triggers         []OpType
	ReverseEdge      bool
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

// WithTriggers allows you to set the triggers for tracking history, can be any combination of OpTypeInsert, OpTypeUpdate, OpTypeDelete,
// nil value will default to all triggers, to exclude all triggers set to an empty slice
func WithTriggers(triggers ...OpType) Option {
	return func(config *HistoryOptions) {
		config.Triggers = triggers
	}
}

// WithReverseEdge enables generation of reverse edges from history schemas back to their original entities.
// When enabled, history schemas will include an edge like:
//
//	edge.From("character", Character.Type).Ref("history").Field("ref").Unique().Immutable()
//
// This allows traversal from history records back to the original entity.
// Note: The original schema must define a matching edge.To("history", CharacterHistory.Type) for this to work.
func WithReverseEdge() Option {
	return func(config *HistoryOptions) {
		config.ReverseEdge = true
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

	filenames, ferr := getFileNames(schemaPath)
	if ferr != nil {
		return ferr
	}

	g, _ := errgroup.WithContext(context.Background())
	mutations := make([]schemast.Mutator, len(schemas))
	for i, s := range schemas {
		g.Go(func() error {
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

			if opts.ReverseEdge {
				upsert.Edges = append(upsert.Edges, reverseEdge(schemaName))
			}

			if len(s.Mixin()) > 0 {
				upsert.Mixins = s.Mixin()
			}

			annotations, gerr := handleAnnotation(schemaName, s.Annotations(), opts.Triggers)
			if gerr != nil {
				return gerr
			}
			historyAnt := reduce(annotations, func(agg Annotations, item schema.Annotation) Annotations {
				if item.Name() == "History" {
					ant := agg.Merge(item)
					agg = ant.(Annotations)
				}
				return agg
			}, Annotations{})
			if historyAnt.Mixins != nil {
				upsert.Mixins = historyAnt.Mixins
			}
			upsert.Annotations = annotations
			if historyAnt.Annotations != nil {
				triggers := []OpType{OpTypeInsert, OpTypeUpdate, OpTypeDelete}
				if historyAnt.Triggers != nil {
					triggers = historyAnt.Triggers
				}
				withIsHist := historyAnt.Merge(Annotations{Annotations: []schema.Annotation{Annotations{IsHistory: true, Triggers: triggers}}})
				upsert.Annotations = withIsHist.(Annotations).Annotations
			}
			mutations[i] = &upsert
			return nil
		})

		err = g.Wait()
		if err != nil {
			return err
		}
	}

	if err = schemast.Mutate(ctx, mutations...); err != nil {
		return err
	}
	return ctx.Print(schemaPath, schemast.Header("Code generated by enthistory, DO NOT EDIT."))
}

func getFileNames(schemaPath string) (map[string]string, error) {
	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedName | packages.NeedTypes | packages.NeedFiles | packages.NeedCompiledGoFiles | packages.NeedSyntax,
	}, schemaPath)
	if err != nil {
		return nil, err
	}
	if len(pkgs) < 1 {
		return nil, fmt.Errorf("missing package information for: %s", schemaPath)
	}
	filenames := make(map[string]string)
	for _, f := range pkgs[0].GoFiles {
		readFile, rerr := os.ReadFile(f)
		if rerr != nil {
			return nil, rerr
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
	return filenames, nil
}

func handleAnnotation(schemaName string, ants []schema.Annotation, triggers []OpType) ([]schema.Annotation, error) {
	if triggers == nil {
		triggers = []OpType{OpTypeInsert, OpTypeUpdate, OpTypeDelete}
	}
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
			ant.Table = fmt.Sprintf("%s_history", strings.ToLower(ant.Table))
			annotations[idx] = ant
		}
	}

	idx = slices.IndexFunc(annotations, func(sc schema.Annotation) bool {
		_, ok := sc.(Annotations)
		return ok
	})
	if idx == -1 {
		annotations = append(annotations, Annotations{IsHistory: true, Triggers: triggers})
	} else {
		ant, ok := annotations[idx].(Annotations)
		if ok {
			ant.IsHistory = true
			if ant.Triggers == nil {
				ant.Triggers = triggers
			}
			annotations[idx] = ant
		}
	}

	return cleanAnnotations(annotations)
}

func cleanAnnotations(annotations []schema.Annotation) ([]schema.Annotation, error) {
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
		case "History":
			typed, err := typedSliceToType[schema.Annotation, Annotations](a)
			if err != nil {
				return nil, err
			}
			mergedAnnotations = append(mergedAnnotations, mergeAnnotations[Annotations](typed...))
		case "Fields":
			merged := reduce(a, func(agg field.Annotation, item schema.Annotation) field.Annotation {
				merged := agg.Merge(item)
				return merged.(field.Annotation)
			}, field.Annotation{})
			if len(merged.ID) > 0 {
				merged.ID = nil
			}
			if len(merged.StructTag) > 0 {
				mergedAnnotations = append(mergedAnnotations, merged)
			}
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
	var t = *new(T)
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
	entgqlEnabled := false

	// Collect fields and parse out id type
	for _, f := range schema.Fields() {
		descriptor := f.Descriptor()
		if descriptor.Name == "id" {
			managedId = true
			ref = f
			if opts.InheritIdType {
				idField = f
			}
		}
		for _, ant := range descriptor.Annotations {
			if _, ok := ant.(entgql.Annotation); ok {
				entgqlEnabled = true
				break
			}
		}
	}

	// Look for id field in mixins
	for _, m := range schema.Mixin() {
		for _, f := range m.Fields() {
			descriptor := f.Descriptor()
			if descriptor.Name == "id" {
				managedId = true
				ref = f
				if opts.InheritIdType {
					idField = f
				}
			}
		}
	}

	if managedId {
		fields = append(fields, idField)
	}
	fields = append(fields, field.Time("history_time").Default(time.Now).Immutable())
	fields = append(fields, field.Enum("operation").GoType(OpType("")).Immutable())
	ref, err := refField(ref.Descriptor().Info, entgqlEnabled)
	if err != nil {
		return nil, err
	}
	fields = append(fields, ref)

	if opts.UpdatedBy != nil {
		ubf, uerr := updatedByField(opts.UpdatedBy.valueType, entgqlEnabled)
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
	descriptor.Unique = false

	if descriptor.Info.Type != field.TypeJSON {
		if opts.FieldProperties.Nillable {
			descriptor.Nillable = true
			descriptor.Optional = true
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

// historyEdge is a custom ent.Edge implementation for generating reverse edges.
type historyEdge struct {
	desc *edge.Descriptor
}

func (h *historyEdge) Descriptor() *edge.Descriptor {
	return h.desc
}

// reverseEdge creates an edge from the history schema back to the original entity.
// This enables traversal like historyRecord.QueryCharacter().
// Note: We intentionally do NOT use Field("ref") because that would create a FK constraint,
// which would prevent creating history records for DELETE operations (the original entity
// no longer exists when the delete hook runs).
func reverseEdge(schemaName string) ent.Edge {
	edgeName := strings.ToLower(schemaName)
	return &historyEdge{
		desc: &edge.Descriptor{
			Name:   edgeName,
			Type:   schemaName,
			Unique: true,
		},
	}
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
