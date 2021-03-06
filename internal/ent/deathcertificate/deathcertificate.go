// Code generated by entc, DO NOT EDIT.

package deathcertificate

import (
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the deathcertificate type in the database.
	Label = "death_certificate"
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
	// FieldCause holds the string denoting the cause field in the database.
	FieldCause = "cause"
	// FieldComments holds the string denoting the comments field in the database.
	FieldComments = "comments"
	// FieldCertifiedAt holds the string denoting the certified_at field in the database.
	FieldCertifiedAt = "certified_at"
	// FieldCertifierComments holds the string denoting the certifier_comments field in the database.
	FieldCertifierComments = "certifier_comments"
	// FieldRequiresCertification holds the string denoting the requires_certification field in the database.
	FieldRequiresCertification = "requires_certification"
	// EdgeMetadata holds the string denoting the metadata edge name in mutations.
	EdgeMetadata = "metadata"
	// EdgeDecedent holds the string denoting the decedent edge name in mutations.
	EdgeDecedent = "decedent"
	// EdgeManner holds the string denoting the manner edge name in mutations.
	EdgeManner = "manner"
	// EdgePlace holds the string denoting the place edge name in mutations.
	EdgePlace = "place"
	// EdgeCertifier holds the string denoting the certifier edge name in mutations.
	EdgeCertifier = "certifier"
	// EdgeCertifiedBy holds the string denoting the certified_by edge name in mutations.
	EdgeCertifiedBy = "certified_by"
	// Table holds the table name of the deathcertificate in the database.
	Table = "death_certificates"
	// MetadataTable is the table that holds the metadata relation/edge.
	MetadataTable = "death_certificates"
	// MetadataInverseTable is the table name for the Metadata entity.
	// It exists in this package in order to avoid circular dependency with the "metadata" package.
	MetadataInverseTable = "metadata"
	// MetadataColumn is the table column denoting the metadata relation/edge.
	MetadataColumn = "death_certificate_metadata"
	// DecedentTable is the table that holds the decedent relation/edge.
	DecedentTable = "death_certificates"
	// DecedentInverseTable is the table name for the Person entity.
	// It exists in this package in order to avoid circular dependency with the "person" package.
	DecedentInverseTable = "persons"
	// DecedentColumn is the table column denoting the decedent relation/edge.
	DecedentColumn = "person_death_certificate"
	// MannerTable is the table that holds the manner relation/edge.
	MannerTable = "death_certificates"
	// MannerInverseTable is the table name for the DeathManner entity.
	// It exists in this package in order to avoid circular dependency with the "deathmanner" package.
	MannerInverseTable = "death_manners"
	// MannerColumn is the table column denoting the manner relation/edge.
	MannerColumn = "death_manner_death_certificates"
	// PlaceTable is the table that holds the place relation/edge.
	PlaceTable = "death_certificates"
	// PlaceInverseTable is the table name for the DeathPlace entity.
	// It exists in this package in order to avoid circular dependency with the "deathplace" package.
	PlaceInverseTable = "death_places"
	// PlaceColumn is the table column denoting the place relation/edge.
	PlaceColumn = "death_place_death_certificates"
	// CertifierTable is the table that holds the certifier relation/edge.
	CertifierTable = "death_certificates"
	// CertifierInverseTable is the table name for the DeathCertifier entity.
	// It exists in this package in order to avoid circular dependency with the "deathcertifier" package.
	CertifierInverseTable = "death_certifiers"
	// CertifierColumn is the table column denoting the certifier relation/edge.
	CertifierColumn = "death_certifier_death_certificates"
	// CertifiedByTable is the table that holds the certified_by relation/edge.
	CertifiedByTable = "death_certificates"
	// CertifiedByInverseTable is the table name for the Person entity.
	// It exists in this package in order to avoid circular dependency with the "person" package.
	CertifiedByInverseTable = "persons"
	// CertifiedByColumn is the table column denoting the certified_by relation/edge.
	CertifiedByColumn = "person_certified_deaths"
)

// Columns holds all SQL columns for deathcertificate fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldCreatedBy,
	FieldCreatedWith,
	FieldUpdatedAt,
	FieldUpdatedBy,
	FieldUpdatedWith,
	FieldCause,
	FieldComments,
	FieldCertifiedAt,
	FieldCertifierComments,
	FieldRequiresCertification,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "death_certificates"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"death_certificate_metadata",
	"death_certifier_death_certificates",
	"death_manner_death_certificates",
	"death_place_death_certificates",
	"person_death_certificate",
	"person_certified_deaths",
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
	// DefaultRequiresCertification holds the default value on creation for the "requires_certification" field.
	DefaultRequiresCertification bool
)
