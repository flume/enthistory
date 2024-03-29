// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// OrganizationColumns holds the columns for the "organization" table.
	OrganizationColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "info", Type: field.TypeJSON, Nullable: true},
	}
	// OrganizationTable holds the schema information for the "organization" table.
	OrganizationTable = &schema.Table{
		Name:       "organization",
		Columns:    OrganizationColumns,
		PrimaryKey: []*schema.Column{OrganizationColumns[0]},
	}
	// OrganizationHistoryColumns holds the columns for the "organization_history" table.
	OrganizationHistoryColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "history_time", Type: field.TypeTime},
		{Name: "operation", Type: field.TypeEnum, Enums: []string{"INSERT", "UPDATE", "DELETE"}},
		{Name: "ref", Type: field.TypeUUID, Nullable: true},
		{Name: "updated_by", Type: field.TypeUUID, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "info", Type: field.TypeJSON, Nullable: true},
	}
	// OrganizationHistoryTable holds the schema information for the "organization_history" table.
	OrganizationHistoryTable = &schema.Table{
		Name:       "organization_history",
		Columns:    OrganizationHistoryColumns,
		PrimaryKey: []*schema.Column{OrganizationHistoryColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "organizationhistory_history_time",
				Unique:  false,
				Columns: []*schema.Column{OrganizationHistoryColumns[3]},
			},
		},
	}
	// StoreColumns holds the columns for the "store" table.
	StoreColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "region", Type: field.TypeString},
		{Name: "organization_id", Type: field.TypeUUID},
	}
	// StoreTable holds the schema information for the "store" table.
	StoreTable = &schema.Table{
		Name:       "store",
		Columns:    StoreColumns,
		PrimaryKey: []*schema.Column{StoreColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "store_organization_organization_stores",
				Columns:    []*schema.Column{StoreColumns[5]},
				RefColumns: []*schema.Column{OrganizationColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// StoreHistoryColumns holds the columns for the "store_history" table.
	StoreHistoryColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "history_time", Type: field.TypeTime},
		{Name: "operation", Type: field.TypeEnum, Enums: []string{"INSERT", "UPDATE", "DELETE"}},
		{Name: "ref", Type: field.TypeUUID, Nullable: true},
		{Name: "updated_by", Type: field.TypeUUID, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "region", Type: field.TypeString},
		{Name: "organization_id", Type: field.TypeUUID},
	}
	// StoreHistoryTable holds the schema information for the "store_history" table.
	StoreHistoryTable = &schema.Table{
		Name:       "store_history",
		Columns:    StoreHistoryColumns,
		PrimaryKey: []*schema.Column{StoreHistoryColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "storehistory_history_time",
				Unique:  false,
				Columns: []*schema.Column{StoreHistoryColumns[3]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		OrganizationTable,
		OrganizationHistoryTable,
		StoreTable,
		StoreHistoryTable,
	}
)

func init() {
	OrganizationTable.Annotation = &entsql.Annotation{
		Table: "organization",
	}
	OrganizationHistoryTable.Annotation = &entsql.Annotation{
		Table: "organization_history",
	}
	StoreTable.ForeignKeys[0].RefTable = OrganizationTable
	StoreTable.Annotation = &entsql.Annotation{
		Table: "store",
	}
	StoreHistoryTable.Annotation = &entsql.Annotation{
		Table: "store_history",
	}
}
