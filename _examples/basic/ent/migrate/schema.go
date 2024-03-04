// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CharacterColumns holds the columns for the "character" table.
	CharacterColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "age", Type: field.TypeInt},
		{Name: "name", Type: field.TypeString},
		{Name: "nicknames", Type: field.TypeJSON, Nullable: true},
		{Name: "info", Type: field.TypeJSON, Nullable: true},
		{Name: "residence_occupants", Type: field.TypeUUID, Nullable: true},
	}
	// CharacterTable holds the schema information for the "character" table.
	CharacterTable = &schema.Table{
		Name:       "character",
		Columns:    CharacterColumns,
		PrimaryKey: []*schema.Column{CharacterColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "character_residence_occupants",
				Columns:    []*schema.Column{CharacterColumns[7]},
				RefColumns: []*schema.Column{ResidenceColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CharacterHistoryColumns holds the columns for the "character_history" table.
	CharacterHistoryColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "history_time", Type: field.TypeTime},
		{Name: "operation", Type: field.TypeEnum, Enums: []string{"INSERT", "UPDATE", "DELETE"}},
		{Name: "ref", Type: field.TypeInt, Nullable: true},
		{Name: "updated_by", Type: field.TypeInt, Nullable: true},
		{Name: "age", Type: field.TypeInt},
		{Name: "name", Type: field.TypeString},
		{Name: "nicknames", Type: field.TypeJSON, Nullable: true},
		{Name: "info", Type: field.TypeJSON, Nullable: true},
	}
	// CharacterHistoryTable holds the schema information for the "character_history" table.
	CharacterHistoryTable = &schema.Table{
		Name:       "character_history",
		Columns:    CharacterHistoryColumns,
		PrimaryKey: []*schema.Column{CharacterHistoryColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "characterhistory_history_time",
				Unique:  false,
				Columns: []*schema.Column{CharacterHistoryColumns[3]},
			},
		},
	}
	// FriendshipColumns holds the columns for the "friendship" table.
	FriendshipColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "character_id", Type: field.TypeInt},
		{Name: "friend_id", Type: field.TypeInt},
	}
	// FriendshipTable holds the schema information for the "friendship" table.
	FriendshipTable = &schema.Table{
		Name:       "friendship",
		Columns:    FriendshipColumns,
		PrimaryKey: []*schema.Column{FriendshipColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "friendship_character_character",
				Columns:    []*schema.Column{FriendshipColumns[3]},
				RefColumns: []*schema.Column{CharacterColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "friendship_character_friend",
				Columns:    []*schema.Column{FriendshipColumns[4]},
				RefColumns: []*schema.Column{CharacterColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "friendship_character_id_friend_id",
				Unique:  true,
				Columns: []*schema.Column{FriendshipColumns[3], FriendshipColumns[4]},
			},
		},
	}
	// FriendshipHistoryColumns holds the columns for the "friendship_history" table.
	FriendshipHistoryColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "history_time", Type: field.TypeTime},
		{Name: "operation", Type: field.TypeEnum, Enums: []string{"INSERT", "UPDATE", "DELETE"}},
		{Name: "ref", Type: field.TypeString, Nullable: true},
		{Name: "updated_by", Type: field.TypeInt, Nullable: true},
		{Name: "character_id", Type: field.TypeInt},
		{Name: "friend_id", Type: field.TypeInt},
	}
	// FriendshipHistoryTable holds the schema information for the "friendship_history" table.
	FriendshipHistoryTable = &schema.Table{
		Name:       "friendship_history",
		Columns:    FriendshipHistoryColumns,
		PrimaryKey: []*schema.Column{FriendshipHistoryColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "friendshiphistory_history_time",
				Unique:  false,
				Columns: []*schema.Column{FriendshipHistoryColumns[3]},
			},
		},
	}
	// ResidenceColumns holds the columns for the "residence" table.
	ResidenceColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
	}
	// ResidenceTable holds the schema information for the "residence" table.
	ResidenceTable = &schema.Table{
		Name:       "residence",
		Columns:    ResidenceColumns,
		PrimaryKey: []*schema.Column{ResidenceColumns[0]},
	}
	// ResidenceHistoryColumns holds the columns for the "residence_history" table.
	ResidenceHistoryColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "history_time", Type: field.TypeTime},
		{Name: "operation", Type: field.TypeEnum, Enums: []string{"INSERT", "UPDATE", "DELETE"}},
		{Name: "ref", Type: field.TypeUUID, Nullable: true},
		{Name: "updated_by", Type: field.TypeInt, Nullable: true},
		{Name: "name", Type: field.TypeString},
	}
	// ResidenceHistoryTable holds the schema information for the "residence_history" table.
	ResidenceHistoryTable = &schema.Table{
		Name:       "residence_history",
		Columns:    ResidenceHistoryColumns,
		PrimaryKey: []*schema.Column{ResidenceHistoryColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "residencehistory_history_time",
				Unique:  false,
				Columns: []*schema.Column{ResidenceHistoryColumns[3]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CharacterTable,
		CharacterHistoryTable,
		FriendshipTable,
		FriendshipHistoryTable,
		ResidenceTable,
		ResidenceHistoryTable,
	}
)

func init() {
	CharacterTable.ForeignKeys[0].RefTable = ResidenceTable
	CharacterTable.Annotation = &entsql.Annotation{
		Table: "character",
	}
	CharacterHistoryTable.Annotation = &entsql.Annotation{
		Table: "character_history",
	}
	FriendshipTable.ForeignKeys[0].RefTable = CharacterTable
	FriendshipTable.ForeignKeys[1].RefTable = CharacterTable
	FriendshipTable.Annotation = &entsql.Annotation{
		Table: "friendship",
	}
	FriendshipHistoryTable.Annotation = &entsql.Annotation{
		Table: "friendship_history",
	}
	ResidenceTable.Annotation = &entsql.Annotation{
		Table: "residence",
	}
	ResidenceHistoryTable.Annotation = &entsql.Annotation{
		Table: "residence_history",
	}
}
