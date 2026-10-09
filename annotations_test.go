package enthistory

import (
	"testing"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/stretchr/testify/require"
)

type TestFieldOverrideSchema struct {
	ent.Schema
}

func (TestFieldOverrideSchema) Fields() []ent.Field {
	return []ent.Field{
		field.String("plain"),
		field.String("nillable_override").
			Annotations(FieldAnnotations{Nillable: Bool(true)}),
		field.String("immutable_off").
			Annotations(FieldAnnotations{Immutable: Bool(false)}),
		field.String("not_nillable").
			Annotations(FieldAnnotations{Nillable: Bool(false)}),
		field.String("orig_nillable").Optional().Nillable().
			Annotations(FieldAnnotations{Nillable: Bool(false)}),
		field.JSON("json_field", []string{}).
			Annotations(FieldAnnotations{Nillable: Bool(true)}),
		field.String("with_default").Default("x").
			Annotations(FieldAnnotations{Nillable: Bool(true)}),
	}
}

func historyDescriptors(t *testing.T, props FieldProperties) map[string]*field.Descriptor {
	t.Helper()
	fields, err := historyFields(TestFieldOverrideSchema{}, HistoryOptions{FieldProperties: &props})
	require.NoError(t, err)
	out := map[string]*field.Descriptor{}
	for _, f := range fields {
		out[f.Descriptor().Name] = f.Descriptor()
	}
	return out
}

func TestFieldAnnotationsNillableOverride(t *testing.T) {
	d := historyDescriptors(t, FieldProperties{})
	require.False(t, d["plain"].Nillable)
	require.False(t, d["plain"].Optional)
	require.True(t, d["nillable_override"].Nillable)
	require.True(t, d["nillable_override"].Optional)
	require.True(t, d["with_default"].Nillable)
	require.True(t, d["with_default"].Optional)
	require.NotNil(t, d["with_default"].Default)
	require.False(t, d["json_field"].Nillable, "JSON fields are never made nillable")
	require.False(t, d["orig_nillable"].Nillable, "explicit false overrides an original nillable field")
	for _, name := range []string{"plain", "nillable_override"} {
		require.Empty(t, d[name].Annotations, "override annotation must not leak into the history schema")
	}
}

func TestFieldAnnotationsOverrideGlobalOptions(t *testing.T) {
	d := historyDescriptors(t, FieldProperties{Nillable: true, Immutable: true})
	require.True(t, d["plain"].Nillable)
	require.True(t, d["plain"].Immutable)
	require.False(t, d["not_nillable"].Nillable)
	require.False(t, d["not_nillable"].Optional)
	require.True(t, d["not_nillable"].Immutable, "unset override keeps the global value")
	require.False(t, d["immutable_off"].Immutable)
	require.True(t, d["immutable_off"].Nillable, "unset override keeps the global value")
}
