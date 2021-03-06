// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// FileMiddlewaresColumns holds the columns for the "file_middlewares" table.
	FileMiddlewaresColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "url_id", Type: field.TypeString},
		{Name: "file_path", Type: field.TypeString},
		{Name: "accessed", Type: field.TypeBool, Default: false},
		{Name: "file_middleware_file_middleware_to_user", Type: field.TypeUUID, Nullable: true},
	}
	// FileMiddlewaresTable holds the schema information for the "file_middlewares" table.
	FileMiddlewaresTable = &schema.Table{
		Name:       "file_middlewares",
		Columns:    FileMiddlewaresColumns,
		PrimaryKey: []*schema.Column{FileMiddlewaresColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "file_middlewares_users_FileMiddlewareToUser",
				Columns:    []*schema.Column{FileMiddlewaresColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"ADMIN", "USER"}},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		FileMiddlewaresTable,
		UsersTable,
	}
)

func init() {
	FileMiddlewaresTable.ForeignKeys[0].RefTable = UsersTable
}
