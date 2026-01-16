package enthistory

import (
	"testing"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

// TestAnonFuncDefaultSchema is a test schema with an anonymous function Default.
// This reproduces the bug from the GitHub issue.
type TestAnonFuncDefaultSchema struct {
	ent.Schema
}

func (TestAnonFuncDefaultSchema) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(func() uuid.UUID {
				return uuid.Must(uuid.NewV7())
			}),
		field.String("name"),
	}
}

// TestNamedFuncDefaultSchema is a test schema with a named function Default.
// This should work both before and after the fix.
type TestNamedFuncDefaultSchema struct {
	ent.Schema
}

func (TestNamedFuncDefaultSchema) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("name"),
	}
}

// TestHistoryFieldsWithAnonFuncDefault tests that historyFields() handles
// ID fields with anonymous function Defaults when WithInheritIdType is enabled.
// This reproduces the bug: "schemast: only selector exprs are supported for default func"
func TestHistoryFieldsWithAnonFuncDefault(t *testing.T) {
	opts := HistoryOptions{
		InheritIdType:   true,
		FieldProperties: &FieldProperties{},
	}

	// This should succeed and the returned ID field should inherit the Default function
	fields, err := historyFields(TestAnonFuncDefaultSchema{}, opts)
	require.NoError(t, err, "historyFields should succeed with anonymous function Default")
	require.NotEmpty(t, fields)

	// The first field should be the inherited ID field
	idField := fields[0]
	require.Equal(t, "id", idField.Descriptor().Name)
	require.Equal(t, field.TypeUUID, idField.Descriptor().Info.Type)

	// The Default should be inherited from the original schema
	require.NotNil(t, idField.Descriptor().Default,
		"ID field Default should be inherited from the original schema")
}

// TestHistoryFieldsWithNamedFuncDefault tests that historyFields() handles
// ID fields with named function Defaults when WithInheritIdType is enabled.
// This is a regression test to ensure we don't break existing behavior.
func TestHistoryFieldsWithNamedFuncDefault(t *testing.T) {
	opts := HistoryOptions{
		InheritIdType:   true,
		FieldProperties: &FieldProperties{},
	}

	fields, err := historyFields(TestNamedFuncDefaultSchema{}, opts)
	require.NoError(t, err, "historyFields should succeed with named function Default")
	require.NotEmpty(t, fields)

	// The first field should be the inherited ID field
	idField := fields[0]
	require.Equal(t, "id", idField.Descriptor().Name)
	require.Equal(t, field.TypeUUID, idField.Descriptor().Info.Type)

	// The Default should be inherited from the original schema
	require.NotNil(t, idField.Descriptor().Default,
		"ID field Default should be inherited from the original schema")
}

// TestHistoryFieldsWithoutInheritIdType tests that historyFields() uses
// the default int ID type when WithInheritIdType is not enabled.
func TestHistoryFieldsWithoutInheritIdType(t *testing.T) {
	opts := HistoryOptions{
		InheritIdType:   false,
		FieldProperties: &FieldProperties{},
	}

	fields, err := historyFields(TestAnonFuncDefaultSchema{}, opts)
	require.NoError(t, err, "historyFields should succeed without InheritIdType")
	require.NotEmpty(t, fields)

	// The first field should be the default int ID field, not the UUID field
	idField := fields[0]
	require.Equal(t, "id", idField.Descriptor().Name)
	require.Equal(t, field.TypeInt, idField.Descriptor().Info.Type,
		"Without InheritIdType, ID field should be int type")
}

// TestAnonFuncDefaultMixin is a test mixin with an anonymous function Default on the ID field.
type TestAnonFuncDefaultMixin struct {
	ent.Schema
}

func (TestAnonFuncDefaultMixin) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(func() uuid.UUID {
				return uuid.Must(uuid.NewV7())
			}),
	}
}

// TestSchemaWithAnonFuncDefaultMixin is a test schema that uses a mixin with an anonymous function Default.
type TestSchemaWithAnonFuncDefaultMixin struct {
	ent.Schema
}

func (TestSchemaWithAnonFuncDefaultMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

func (TestSchemaWithAnonFuncDefaultMixin) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TestAnonFuncDefaultMixin{},
	}
}

// TestHistoryFieldsWithMixinAnonFuncDefault tests that historyFields() handles
// ID fields with anonymous function Defaults from mixins when WithInheritIdType is enabled.
func TestHistoryFieldsWithMixinAnonFuncDefault(t *testing.T) {
	opts := HistoryOptions{
		InheritIdType:   true,
		FieldProperties: &FieldProperties{},
	}

	fields, err := historyFields(TestSchemaWithAnonFuncDefaultMixin{}, opts)
	require.NoError(t, err, "historyFields should succeed with mixin anonymous function Default")
	require.NotEmpty(t, fields)

	// The first field should be the inherited ID field from the mixin
	idField := fields[0]
	require.Equal(t, "id", idField.Descriptor().Name)
	require.Equal(t, field.TypeUUID, idField.Descriptor().Info.Type)

	// The Default should be inherited from the original mixin
	require.NotNil(t, idField.Descriptor().Default,
		"Mixin ID field Default should be inherited from the original schema")
}
