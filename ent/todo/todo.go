// Code generated by entc, DO NOT EDIT.

package todo

import (
	"time"
)

const (
	// Label holds the string label denoting the todo type in the database.
	Label = "todo"
	// FieldID holds the string denoting the id field in the database.
	FieldID        = "id"         // FieldCreatedAt holds the string denoting the created_at vertex property in the database.
	FieldCreatedAt = "created_at" // FieldUpdatedAt holds the string denoting the updated_at vertex property in the database.
	FieldUpdatedAt = "updated_at" // FieldText holds the string denoting the text vertex property in the database.
	FieldText      = "text"       // FieldDone holds the string denoting the done vertex property in the database.
	FieldDone      = "done"

	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"

	// Table holds the table name of the todo in the database.
	Table = "todos"
	// UserTable is the table the holds the user relation/edge.
	UserTable = "todos"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_todos"
)

// Columns holds all SQL columns for todo fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldText,
	FieldDone,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Todo type.
var ForeignKeys = []string{
	"user_todos",
}

var (
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the updated_at field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	UpdateDefaultUpdatedAt func() time.Time
	// TextValidator is a validator for the "text" field. It is called by the builders before save.
	TextValidator func(string) error
	// DefaultDone holds the default value on creation for the done field.
	DefaultDone bool
)
