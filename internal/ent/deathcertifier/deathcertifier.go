// Code generated by entc, DO NOT EDIT.

package deathcertifier

import (
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the deathcertifier type in the database.
	Label = "death_certifier"
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
	// FieldShort holds the string denoting the short field in the database.
	FieldShort = "short"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeMetadata holds the string denoting the metadata edge name in mutations.
	EdgeMetadata = "metadata"
	// EdgeDeathCertificates holds the string denoting the death_certificates edge name in mutations.
	EdgeDeathCertificates = "death_certificates"
	// Table holds the table name of the deathcertifier in the database.
	Table = "death_certifiers"
	// MetadataTable is the table that holds the metadata relation/edge.
	MetadataTable = "death_certifiers"
	// MetadataInverseTable is the table name for the Metadata entity.
	// It exists in this package in order to avoid circular dependency with the "metadata" package.
	MetadataInverseTable = "metadata"
	// MetadataColumn is the table column denoting the metadata relation/edge.
	MetadataColumn = "death_certifier_metadata"
	// DeathCertificatesTable is the table that holds the death_certificates relation/edge.
	DeathCertificatesTable = "death_certificates"
	// DeathCertificatesInverseTable is the table name for the DeathCertificate entity.
	// It exists in this package in order to avoid circular dependency with the "deathcertificate" package.
	DeathCertificatesInverseTable = "death_certificates"
	// DeathCertificatesColumn is the table column denoting the death_certificates relation/edge.
	DeathCertificatesColumn = "death_certifier_death_certificates"
)

// Columns holds all SQL columns for deathcertifier fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldCreatedBy,
	FieldCreatedWith,
	FieldUpdatedAt,
	FieldUpdatedBy,
	FieldUpdatedWith,
	FieldShort,
	FieldTitle,
	FieldDescription,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "death_certifiers"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"death_certifier_metadata",
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
	// ShortValidator is a validator for the "short" field. It is called by the builders before save.
	ShortValidator func(string) error
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
)
