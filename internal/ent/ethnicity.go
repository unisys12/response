// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/responserms/response/internal/ent/ethnicity"
	"github.com/responserms/response/internal/ent/metadata"
)

// Ethnicity is the model entity for the Ethnicity schema.
type Ethnicity struct {
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
	// Short holds the value of the "short" field.
	Short string `json:"short,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EthnicityQuery when eager-loading is set.
	Edges              EthnicityEdges `json:"edges"`
	ethnicity_metadata *int
}

// EthnicityEdges holds the relations/edges for other nodes in the graph.
type EthnicityEdges struct {
	// Metadata holds the value of the metadata edge.
	Metadata *Metadata `json:"metadata,omitempty"`
	// People holds the value of the people edge.
	People []*Person `json:"people,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// MetadataOrErr returns the Metadata value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EthnicityEdges) MetadataOrErr() (*Metadata, error) {
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

// PeopleOrErr returns the People value or an error if the edge
// was not loaded in eager-loading.
func (e EthnicityEdges) PeopleOrErr() ([]*Person, error) {
	if e.loadedTypes[1] {
		return e.People, nil
	}
	return nil, &NotLoadedError{edge: "people"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Ethnicity) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case ethnicity.FieldID, ethnicity.FieldCreatedBy, ethnicity.FieldUpdatedBy:
			values[i] = new(sql.NullInt64)
		case ethnicity.FieldCreatedWith, ethnicity.FieldUpdatedWith, ethnicity.FieldShort, ethnicity.FieldTitle, ethnicity.FieldDescription:
			values[i] = new(sql.NullString)
		case ethnicity.FieldCreatedAt, ethnicity.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case ethnicity.ForeignKeys[0]: // ethnicity_metadata
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Ethnicity", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Ethnicity fields.
func (e *Ethnicity) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case ethnicity.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			e.ID = int(value.Int64)
		case ethnicity.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				e.CreatedAt = value.Time
			}
		case ethnicity.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				e.CreatedBy = int(value.Int64)
			}
		case ethnicity.FieldCreatedWith:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_with", values[i])
			} else if value.Valid {
				e.CreatedWith = value.String
			}
		case ethnicity.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				e.UpdatedAt = value.Time
			}
		case ethnicity.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				e.UpdatedBy = int(value.Int64)
			}
		case ethnicity.FieldUpdatedWith:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_with", values[i])
			} else if value.Valid {
				e.UpdatedWith = value.String
			}
		case ethnicity.FieldShort:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field short", values[i])
			} else if value.Valid {
				e.Short = value.String
			}
		case ethnicity.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				e.Title = value.String
			}
		case ethnicity.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				e.Description = value.String
			}
		case ethnicity.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field ethnicity_metadata", value)
			} else if value.Valid {
				e.ethnicity_metadata = new(int)
				*e.ethnicity_metadata = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryMetadata queries the "metadata" edge of the Ethnicity entity.
func (e *Ethnicity) QueryMetadata() *MetadataQuery {
	return (&EthnicityClient{config: e.config}).QueryMetadata(e)
}

// QueryPeople queries the "people" edge of the Ethnicity entity.
func (e *Ethnicity) QueryPeople() *PersonQuery {
	return (&EthnicityClient{config: e.config}).QueryPeople(e)
}

// Update returns a builder for updating this Ethnicity.
// Note that you need to call Ethnicity.Unwrap() before calling this method if this Ethnicity
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Ethnicity) Update() *EthnicityUpdateOne {
	return (&EthnicityClient{config: e.config}).UpdateOne(e)
}

// Unwrap unwraps the Ethnicity entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Ethnicity) Unwrap() *Ethnicity {
	tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Ethnicity is not a transactional entity")
	}
	e.config.driver = tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Ethnicity) String() string {
	var builder strings.Builder
	builder.WriteString("Ethnicity(")
	builder.WriteString(fmt.Sprintf("id=%v", e.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(e.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", created_by=")
	builder.WriteString(fmt.Sprintf("%v", e.CreatedBy))
	builder.WriteString(", created_with=")
	builder.WriteString(e.CreatedWith)
	builder.WriteString(", updated_at=")
	builder.WriteString(e.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_by=")
	builder.WriteString(fmt.Sprintf("%v", e.UpdatedBy))
	builder.WriteString(", updated_with=")
	builder.WriteString(e.UpdatedWith)
	builder.WriteString(", short=")
	builder.WriteString(e.Short)
	builder.WriteString(", title=")
	builder.WriteString(e.Title)
	builder.WriteString(", description=")
	builder.WriteString(e.Description)
	builder.WriteByte(')')
	return builder.String()
}

// Ethnicities is a parsable slice of Ethnicity.
type Ethnicities []*Ethnicity

func (e Ethnicities) config(cfg config) {
	for _i := range e {
		e[_i].config = cfg
	}
}
