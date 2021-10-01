// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/responserms/response/internal/ent/race"
)

// Race is the model entity for the Race schema.
type Race struct {
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
	// The values are being populated by the RaceQuery when eager-loading is set.
	Edges RaceEdges `json:"edges"`
}

// RaceEdges holds the relations/edges for other nodes in the graph.
type RaceEdges struct {
	// People holds the value of the people edge.
	People []*Person `json:"people,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PeopleOrErr returns the People value or an error if the edge
// was not loaded in eager-loading.
func (e RaceEdges) PeopleOrErr() ([]*Person, error) {
	if e.loadedTypes[0] {
		return e.People, nil
	}
	return nil, &NotLoadedError{edge: "people"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Race) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case race.FieldID, race.FieldCreatedBy, race.FieldUpdatedBy:
			values[i] = new(sql.NullInt64)
		case race.FieldCreatedWith, race.FieldUpdatedWith, race.FieldShort, race.FieldTitle, race.FieldDescription:
			values[i] = new(sql.NullString)
		case race.FieldCreatedAt, race.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Race", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Race fields.
func (r *Race) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case race.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case race.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = value.Time
			}
		case race.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				r.CreatedBy = int(value.Int64)
			}
		case race.FieldCreatedWith:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_with", values[i])
			} else if value.Valid {
				r.CreatedWith = value.String
			}
		case race.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				r.UpdatedAt = value.Time
			}
		case race.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				r.UpdatedBy = int(value.Int64)
			}
		case race.FieldUpdatedWith:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_with", values[i])
			} else if value.Valid {
				r.UpdatedWith = value.String
			}
		case race.FieldShort:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field short", values[i])
			} else if value.Valid {
				r.Short = value.String
			}
		case race.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				r.Title = value.String
			}
		case race.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				r.Description = value.String
			}
		}
	}
	return nil
}

// QueryPeople queries the "people" edge of the Race entity.
func (r *Race) QueryPeople() *PersonQuery {
	return (&RaceClient{config: r.config}).QueryPeople(r)
}

// Update returns a builder for updating this Race.
// Note that you need to call Race.Unwrap() before calling this method if this Race
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Race) Update() *RaceUpdateOne {
	return (&RaceClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Race entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Race) Unwrap() *Race {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Race is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Race) String() string {
	var builder strings.Builder
	builder.WriteString("Race(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(r.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", created_by=")
	builder.WriteString(fmt.Sprintf("%v", r.CreatedBy))
	builder.WriteString(", created_with=")
	builder.WriteString(r.CreatedWith)
	builder.WriteString(", updated_at=")
	builder.WriteString(r.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_by=")
	builder.WriteString(fmt.Sprintf("%v", r.UpdatedBy))
	builder.WriteString(", updated_with=")
	builder.WriteString(r.UpdatedWith)
	builder.WriteString(", short=")
	builder.WriteString(r.Short)
	builder.WriteString(", title=")
	builder.WriteString(r.Title)
	builder.WriteString(", description=")
	builder.WriteString(r.Description)
	builder.WriteByte(')')
	return builder.String()
}

// Races is a parsable slice of Race.
type Races []*Race

func (r Races) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}