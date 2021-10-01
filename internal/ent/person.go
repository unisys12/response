// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/responserms/response/internal/ent/deathcertificate"
	"github.com/responserms/response/internal/ent/ethnicity"
	"github.com/responserms/response/internal/ent/metadata"
	"github.com/responserms/response/internal/ent/person"
	"github.com/responserms/response/internal/ent/race"
	"github.com/responserms/response/internal/ent/sex"
	"github.com/responserms/response/internal/ent/user"
)

// Person is the model entity for the Person schema.
type Person struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy int `json:"created_by,omitempty"`
	// CreatedWith holds the value of the "created_with" field.
	CreatedWith string `json:"created_with,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy int `json:"updated_by,omitempty"`
	// UpdatedWith holds the value of the "updated_with" field.
	UpdatedWith string `json:"updated_with,omitempty"`
	// FirstName holds the value of the "first_name" field.
	FirstName string `json:"first_name,omitempty"`
	// MiddleName holds the value of the "middle_name" field.
	MiddleName string `json:"middle_name,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// Suffix holds the value of the "suffix" field.
	Suffix string `json:"suffix,omitempty"`
	// DateOfBirth holds the value of the "date_of_birth" field.
	DateOfBirth time.Time `json:"date_of_birth,omitempty"`
	// Birthplace holds the value of the "birthplace" field.
	Birthplace string `json:"birthplace,omitempty"`
	// DeceasedAt holds the value of the "deceased_at" field.
	DeceasedAt *time.Time `json:"deceased_at,omitempty"`
	// ArchivedAt holds the value of the "archived_at" field.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PersonQuery when eager-loading is set.
	Edges            PersonEdges `json:"edges"`
	ethnicity_people *int
	person_metadata  *int
	race_people      *int
	sex_people       *int
	user_people      *int
}

// PersonEdges holds the relations/edges for other nodes in the graph.
type PersonEdges struct {
	// Metadata holds the value of the metadata edge.
	Metadata *Metadata `json:"metadata,omitempty"`
	// VehicleRegistrations holds the value of the vehicle_registrations edge.
	VehicleRegistrations []*VehicleRegistration `json:"vehicle_registrations,omitempty"`
	// Race holds the value of the race edge.
	Race *Race `json:"race,omitempty"`
	// Ethnicity holds the value of the ethnicity edge.
	Ethnicity *Ethnicity `json:"ethnicity,omitempty"`
	// Sex holds the value of the sex edge.
	Sex *Sex `json:"sex,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// DeathCertificate holds the value of the death_certificate edge.
	DeathCertificate *DeathCertificate `json:"death_certificate,omitempty"`
	// CertifiedDeaths holds the value of the certified_deaths edge.
	CertifiedDeaths []*DeathCertificate `json:"certified_deaths,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [8]bool
}

// MetadataOrErr returns the Metadata value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PersonEdges) MetadataOrErr() (*Metadata, error) {
	if e.loadedTypes[0] {
		if e.Metadata == nil {
			// The edge metadata was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: metadata.Label}
		}
		return e.Metadata, nil
	}
	return nil, &NotLoadedError{edge: "metadata"}
}

// VehicleRegistrationsOrErr returns the VehicleRegistrations value or an error if the edge
// was not loaded in eager-loading.
func (e PersonEdges) VehicleRegistrationsOrErr() ([]*VehicleRegistration, error) {
	if e.loadedTypes[1] {
		return e.VehicleRegistrations, nil
	}
	return nil, &NotLoadedError{edge: "vehicle_registrations"}
}

// RaceOrErr returns the Race value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PersonEdges) RaceOrErr() (*Race, error) {
	if e.loadedTypes[2] {
		if e.Race == nil {
			// The edge race was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: race.Label}
		}
		return e.Race, nil
	}
	return nil, &NotLoadedError{edge: "race"}
}

// EthnicityOrErr returns the Ethnicity value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PersonEdges) EthnicityOrErr() (*Ethnicity, error) {
	if e.loadedTypes[3] {
		if e.Ethnicity == nil {
			// The edge ethnicity was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: ethnicity.Label}
		}
		return e.Ethnicity, nil
	}
	return nil, &NotLoadedError{edge: "ethnicity"}
}

// SexOrErr returns the Sex value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PersonEdges) SexOrErr() (*Sex, error) {
	if e.loadedTypes[4] {
		if e.Sex == nil {
			// The edge sex was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: sex.Label}
		}
		return e.Sex, nil
	}
	return nil, &NotLoadedError{edge: "sex"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PersonEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[5] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// DeathCertificateOrErr returns the DeathCertificate value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PersonEdges) DeathCertificateOrErr() (*DeathCertificate, error) {
	if e.loadedTypes[6] {
		if e.DeathCertificate == nil {
			// The edge death_certificate was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: deathcertificate.Label}
		}
		return e.DeathCertificate, nil
	}
	return nil, &NotLoadedError{edge: "death_certificate"}
}

// CertifiedDeathsOrErr returns the CertifiedDeaths value or an error if the edge
// was not loaded in eager-loading.
func (e PersonEdges) CertifiedDeathsOrErr() ([]*DeathCertificate, error) {
	if e.loadedTypes[7] {
		return e.CertifiedDeaths, nil
	}
	return nil, &NotLoadedError{edge: "certified_deaths"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Person) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case person.FieldID, person.FieldCreatedBy, person.FieldUpdatedBy:
			values[i] = new(sql.NullInt64)
		case person.FieldCreatedWith, person.FieldUpdatedWith, person.FieldFirstName, person.FieldMiddleName, person.FieldLastName, person.FieldSuffix, person.FieldBirthplace:
			values[i] = new(sql.NullString)
		case person.FieldCreatedAt, person.FieldUpdatedAt, person.FieldDateOfBirth, person.FieldDeceasedAt, person.FieldArchivedAt:
			values[i] = new(sql.NullTime)
		case person.ForeignKeys[0]: // ethnicity_people
			values[i] = new(sql.NullInt64)
		case person.ForeignKeys[1]: // person_metadata
			values[i] = new(sql.NullInt64)
		case person.ForeignKeys[2]: // race_people
			values[i] = new(sql.NullInt64)
		case person.ForeignKeys[3]: // sex_people
			values[i] = new(sql.NullInt64)
		case person.ForeignKeys[4]: // user_people
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Person", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Person fields.
func (pe *Person) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case person.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pe.ID = int(value.Int64)
		case person.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pe.CreatedAt = value.Time
			}
		case person.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				pe.CreatedBy = int(value.Int64)
			}
		case person.FieldCreatedWith:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_with", values[i])
			} else if value.Valid {
				pe.CreatedWith = value.String
			}
		case person.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pe.UpdatedAt = value.Time
			}
		case person.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				pe.UpdatedBy = int(value.Int64)
			}
		case person.FieldUpdatedWith:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_with", values[i])
			} else if value.Valid {
				pe.UpdatedWith = value.String
			}
		case person.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field first_name", values[i])
			} else if value.Valid {
				pe.FirstName = value.String
			}
		case person.FieldMiddleName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field middle_name", values[i])
			} else if value.Valid {
				pe.MiddleName = value.String
			}
		case person.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				pe.LastName = value.String
			}
		case person.FieldSuffix:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field suffix", values[i])
			} else if value.Valid {
				pe.Suffix = value.String
			}
		case person.FieldDateOfBirth:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field date_of_birth", values[i])
			} else if value.Valid {
				pe.DateOfBirth = value.Time
			}
		case person.FieldBirthplace:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field birthplace", values[i])
			} else if value.Valid {
				pe.Birthplace = value.String
			}
		case person.FieldDeceasedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deceased_at", values[i])
			} else if value.Valid {
				pe.DeceasedAt = new(time.Time)
				*pe.DeceasedAt = value.Time
			}
		case person.FieldArchivedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field archived_at", values[i])
			} else if value.Valid {
				pe.ArchivedAt = new(time.Time)
				*pe.ArchivedAt = value.Time
			}
		case person.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field ethnicity_people", value)
			} else if value.Valid {
				pe.ethnicity_people = new(int)
				*pe.ethnicity_people = int(value.Int64)
			}
		case person.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field person_metadata", value)
			} else if value.Valid {
				pe.person_metadata = new(int)
				*pe.person_metadata = int(value.Int64)
			}
		case person.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field race_people", value)
			} else if value.Valid {
				pe.race_people = new(int)
				*pe.race_people = int(value.Int64)
			}
		case person.ForeignKeys[3]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field sex_people", value)
			} else if value.Valid {
				pe.sex_people = new(int)
				*pe.sex_people = int(value.Int64)
			}
		case person.ForeignKeys[4]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_people", value)
			} else if value.Valid {
				pe.user_people = new(int)
				*pe.user_people = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryMetadata queries the "metadata" edge of the Person entity.
func (pe *Person) QueryMetadata() *MetadataQuery {
	return (&PersonClient{config: pe.config}).QueryMetadata(pe)
}

// QueryVehicleRegistrations queries the "vehicle_registrations" edge of the Person entity.
func (pe *Person) QueryVehicleRegistrations() *VehicleRegistrationQuery {
	return (&PersonClient{config: pe.config}).QueryVehicleRegistrations(pe)
}

// QueryRace queries the "race" edge of the Person entity.
func (pe *Person) QueryRace() *RaceQuery {
	return (&PersonClient{config: pe.config}).QueryRace(pe)
}

// QueryEthnicity queries the "ethnicity" edge of the Person entity.
func (pe *Person) QueryEthnicity() *EthnicityQuery {
	return (&PersonClient{config: pe.config}).QueryEthnicity(pe)
}

// QuerySex queries the "sex" edge of the Person entity.
func (pe *Person) QuerySex() *SexQuery {
	return (&PersonClient{config: pe.config}).QuerySex(pe)
}

// QueryUser queries the "user" edge of the Person entity.
func (pe *Person) QueryUser() *UserQuery {
	return (&PersonClient{config: pe.config}).QueryUser(pe)
}

// QueryDeathCertificate queries the "death_certificate" edge of the Person entity.
func (pe *Person) QueryDeathCertificate() *DeathCertificateQuery {
	return (&PersonClient{config: pe.config}).QueryDeathCertificate(pe)
}

// QueryCertifiedDeaths queries the "certified_deaths" edge of the Person entity.
func (pe *Person) QueryCertifiedDeaths() *DeathCertificateQuery {
	return (&PersonClient{config: pe.config}).QueryCertifiedDeaths(pe)
}

// Update returns a builder for updating this Person.
// Note that you need to call Person.Unwrap() before calling this method if this Person
// was returned from a transaction, and the transaction was committed or rolled back.
func (pe *Person) Update() *PersonUpdateOne {
	return (&PersonClient{config: pe.config}).UpdateOne(pe)
}

// Unwrap unwraps the Person entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pe *Person) Unwrap() *Person {
	tx, ok := pe.config.driver.(*txDriver)
	if !ok {
		panic("ent: Person is not a transactional entity")
	}
	pe.config.driver = tx.drv
	return pe
}

// String implements the fmt.Stringer.
func (pe *Person) String() string {
	var builder strings.Builder
	builder.WriteString("Person(")
	builder.WriteString(fmt.Sprintf("id=%v", pe.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(pe.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", created_by=")
	builder.WriteString(fmt.Sprintf("%v", pe.CreatedBy))
	builder.WriteString(", created_with=")
	builder.WriteString(pe.CreatedWith)
	builder.WriteString(", updated_at=")
	builder.WriteString(pe.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_by=")
	builder.WriteString(fmt.Sprintf("%v", pe.UpdatedBy))
	builder.WriteString(", updated_with=")
	builder.WriteString(pe.UpdatedWith)
	builder.WriteString(", first_name=")
	builder.WriteString(pe.FirstName)
	builder.WriteString(", middle_name=")
	builder.WriteString(pe.MiddleName)
	builder.WriteString(", last_name=")
	builder.WriteString(pe.LastName)
	builder.WriteString(", suffix=")
	builder.WriteString(pe.Suffix)
	builder.WriteString(", date_of_birth=")
	builder.WriteString(pe.DateOfBirth.Format(time.ANSIC))
	builder.WriteString(", birthplace=")
	builder.WriteString(pe.Birthplace)
	if v := pe.DeceasedAt; v != nil {
		builder.WriteString(", deceased_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := pe.ArchivedAt; v != nil {
		builder.WriteString(", archived_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Persons is a parsable slice of Person.
type Persons []*Person

func (pe Persons) config(cfg config) {
	for _i := range pe {
		pe[_i].config = cfg
	}
}
