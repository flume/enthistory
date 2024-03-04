// Code generated by ent, DO NOT EDIT.

package organization

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the organization type in the database.
	Label = "organization"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldInfo holds the string denoting the info field in the database.
	FieldInfo = "info"
	// EdgeOrganizationStores holds the string denoting the organization_stores edge name in mutations.
	EdgeOrganizationStores = "organization_stores"
	// Table holds the table name of the organization in the database.
	Table = "organization"
	// OrganizationStoresTable is the table that holds the organization_stores relation/edge.
	OrganizationStoresTable = "store"
	// OrganizationStoresInverseTable is the table name for the Store entity.
	// It exists in this package in order to avoid circular dependency with the "store" package.
	OrganizationStoresInverseTable = "store"
	// OrganizationStoresColumn is the table column denoting the organization_stores relation/edge.
	OrganizationStoresColumn = "organization_id"
)

// Columns holds all SQL columns for organization fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldInfo,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Organization queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByOrganizationStoresCount orders the results by organization_stores count.
func ByOrganizationStoresCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOrganizationStoresStep(), opts...)
	}
}

// ByOrganizationStores orders the results by organization_stores terms.
func ByOrganizationStores(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrganizationStoresStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newOrganizationStoresStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrganizationStoresInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, OrganizationStoresTable, OrganizationStoresColumn),
	)
}
