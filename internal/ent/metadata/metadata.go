// Code generated by entc, DO NOT EDIT.

package metadata

import (
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the metadata type in the database.
	Label = "metadata"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldCreatedWith holds the string denoting the created_with field in the database.
	FieldCreatedWith = "created_with"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldUpdatedWith holds the string denoting the updated_with field in the database.
	FieldUpdatedWith = "updated_with"
	// FieldData holds the string denoting the data field in the database.
	FieldData = "data"
	// Table holds the table name of the metadata in the database.
	Table = "metadata"
)

// Columns holds all SQL columns for metadata fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldCreatedBy,
	FieldCreatedWith,
	FieldUpdatedAt,
	FieldUpdatedBy,
	FieldUpdatedWith,
	FieldData,
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

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/responserms/response/internal/ent/runtime"
//
var (
	Hooks [1]ent.Hook
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)
