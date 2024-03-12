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
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "other", Type: field.TypeString},
		{Name: "age", Type: field.TypeInt},
		{Name: "name", Type: field.TypeString},
		{Name: "nicknames", Type: field.TypeJSON, Nullable: true},
		{Name: "info", Type: field.TypeJSON, Nullable: true},
	}
	// CharacterTable holds the schema information for the "character" table.
	CharacterTable = &schema.Table{
		Name:       "character",
		Columns:    CharacterColumns,
		PrimaryKey: []*schema.Column{CharacterColumns[0]},
	}
	// CharacterHistoryColumns holds the columns for the "character_history" table.
	CharacterHistoryColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "history_time", Type: field.TypeTime},
		{Name: "operation", Type: field.TypeEnum, Enums: []string{"INSERT", "UPDATE", "DELETE"}},
		{Name: "ref", Type: field.TypeUUID, Nullable: true},
		{Name: "updated_by", Type: field.TypeUUID, Nullable: true},
		{Name: "age", Type: field.TypeInt, Nullable: true},
		{Name: "name", Type: field.TypeString, Nullable: true},
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
		{Name: "id", Type: field.TypeUUID},
		{Name: "character_id", Type: field.TypeUUID},
		{Name: "friend_id", Type: field.TypeUUID},
	}
	// FriendshipTable holds the schema information for the "friendship" table.
	FriendshipTable = &schema.Table{
		Name:       "friendship",
		Columns:    FriendshipColumns,
		PrimaryKey: []*schema.Column{FriendshipColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "friendship_character_character",
				Columns:    []*schema.Column{FriendshipColumns[1]},
				RefColumns: []*schema.Column{CharacterColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "friendship_character_friend",
				Columns:    []*schema.Column{FriendshipColumns[2]},
				RefColumns: []*schema.Column{CharacterColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "friendship_character_id_friend_id",
				Unique:  true,
				Columns: []*schema.Column{FriendshipColumns[1], FriendshipColumns[2]},
			},
		},
	}
	// FriendshipHistoryColumns holds the columns for the "friendship_history" table.
	FriendshipHistoryColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "history_time", Type: field.TypeTime},
		{Name: "operation", Type: field.TypeEnum, Enums: []string{"INSERT", "UPDATE", "DELETE"}},
		{Name: "ref", Type: field.TypeUUID, Nullable: true},
		{Name: "updated_by", Type: field.TypeUUID, Nullable: true},
		{Name: "character_id", Type: field.TypeUUID, Nullable: true},
		{Name: "friend_id", Type: field.TypeUUID, Nullable: true},
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
				Columns: []*schema.Column{FriendshipHistoryColumns[1]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CharacterTable,
		CharacterHistoryTable,
		FriendshipTable,
		FriendshipHistoryTable,
	}
)

func init() {
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
}
