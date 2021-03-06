// Code generated by entc, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
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
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldAvatarURL holds the string denoting the avatar_url field in the database.
	FieldAvatarURL = "avatar_url"
	// FieldPermissions holds the string denoting the permissions field in the database.
	FieldPermissions = "permissions"
	// FieldFirstSetupAt holds the string denoting the first_setup_at field in the database.
	FieldFirstSetupAt = "first_setup_at"
	// EdgeMetadata holds the string denoting the metadata edge name in mutations.
	EdgeMetadata = "metadata"
	// EdgeOauthConnections holds the string denoting the oauth_connections edge name in mutations.
	EdgeOauthConnections = "oauth_connections"
	// EdgePeople holds the string denoting the people edge name in mutations.
	EdgePeople = "people"
	// Table holds the table name of the user in the database.
	Table = "users"
	// MetadataTable is the table that holds the metadata relation/edge.
	MetadataTable = "users"
	// MetadataInverseTable is the table name for the Metadata entity.
	// It exists in this package in order to avoid circular dependency with the "metadata" package.
	MetadataInverseTable = "metadata"
	// MetadataColumn is the table column denoting the metadata relation/edge.
	MetadataColumn = "user_metadata"
	// OauthConnectionsTable is the table that holds the oauth_connections relation/edge.
	OauthConnectionsTable = "oauth_connections"
	// OauthConnectionsInverseTable is the table name for the OAuthConnection entity.
	// It exists in this package in order to avoid circular dependency with the "oauthconnection" package.
	OauthConnectionsInverseTable = "oauth_connections"
	// OauthConnectionsColumn is the table column denoting the oauth_connections relation/edge.
	OauthConnectionsColumn = "user_oauth_connections"
	// PeopleTable is the table that holds the people relation/edge.
	PeopleTable = "persons"
	// PeopleInverseTable is the table name for the Person entity.
	// It exists in this package in order to avoid circular dependency with the "person" package.
	PeopleInverseTable = "persons"
	// PeopleColumn is the table column denoting the people relation/edge.
	PeopleColumn = "user_people"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldCreatedBy,
	FieldCreatedWith,
	FieldUpdatedAt,
	FieldUpdatedBy,
	FieldUpdatedWith,
	FieldName,
	FieldEmail,
	FieldPassword,
	FieldAvatarURL,
	FieldPermissions,
	FieldFirstSetupAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "users"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_metadata",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
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
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
)
