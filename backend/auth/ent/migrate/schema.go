// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// SessionsColumns holds the columns for the "sessions" table.
	SessionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "session_id", Type: field.TypeUUID},
		{Name: "access_token", Type: field.TypeString},
		{Name: "access_token_expiration_at", Type: field.TypeTime},
		{Name: "refresh_token", Type: field.TypeString},
		{Name: "refresh_token_expiration_at", Type: field.TypeTime},
		{Name: "ip", Type: field.TypeString, SchemaType: map[string]string{"postgres": "inet"}},
		{Name: "agent", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "last_activity_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt, Nullable: true},
	}
	// SessionsTable holds the schema information for the "sessions" table.
	SessionsTable = &schema.Table{
		Name:       "sessions",
		Columns:    SessionsColumns,
		PrimaryKey: []*schema.Column{SessionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sessions_users_sessions",
				Columns:    []*schema.Column{SessionsColumns[10]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "active", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		SessionsTable,
		UsersTable,
	}
)

func init() {
	SessionsTable.ForeignKeys[0].RefTable = UsersTable
}
