// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/responserms/response/internal/ent/state"
)

// State is the model entity for the State schema.
type State struct {
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
	// The values are being populated by the StateQuery when eager-loading is set.
	Edges StateEdges `json:"edges"`
}

// StateEdges holds the relations/edges for other nodes in the graph.
type StateEdges struct {
	// VehicleRegistrations holds the value of the vehicle_registrations edge.
	VehicleRegistrations []*VehicleRegistration `json:"vehicle_registrations,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// VehicleRegistrationsOrErr returns the VehicleRegistrations value or an error if the edge
// was not loaded in eager-loading.
func (e StateEdges) VehicleRegistrationsOrErr() ([]*VehicleRegistration, error) {
	if e.loadedTypes[0] {
		return e.VehicleRegistrations, nil
	}
	return nil, &NotLoadedError{edge: "vehicle_registrations"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*State) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case state.FieldID, state.FieldCreatedBy, state.FieldUpdatedBy:
			values[i] = new(sql.NullInt64)
		case state.FieldCreatedWith, state.FieldUpdatedWith, state.FieldShort, state.FieldTitle, state.FieldDescription:
			values[i] = new(sql.NullString)
		case state.FieldCreatedAt, state.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type State", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the State fields.
func (s *State) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case state.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case state.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case state.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				s.CreatedBy = int(value.Int64)
			}
		case state.FieldCreatedWith:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_with", values[i])
			} else if value.Valid {
				s.CreatedWith = value.String
			}
		case state.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		case state.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				s.UpdatedBy = int(value.Int64)
			}
		case state.FieldUpdatedWith:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_with", values[i])
			} else if value.Valid {
				s.UpdatedWith = value.String
			}
		case state.FieldShort:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field short", values[i])
			} else if value.Valid {
				s.Short = value.String
			}
		case state.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				s.Title = value.String
			}
		case state.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				s.Description = value.String
			}
		}
	}
	return nil
}

// QueryVehicleRegistrations queries the "vehicle_registrations" edge of the State entity.
func (s *State) QueryVehicleRegistrations() *VehicleRegistrationQuery {
	return (&StateClient{config: s.config}).QueryVehicleRegistrations(s)
}

// Update returns a builder for updating this State.
// Note that you need to call State.Unwrap() before calling this method if this State
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *State) Update() *StateUpdateOne {
	return (&StateClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the State entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *State) Unwrap() *State {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: State is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *State) String() string {
	var builder strings.Builder
	builder.WriteString("State(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", created_by=")
	builder.WriteString(fmt.Sprintf("%v", s.CreatedBy))
	builder.WriteString(", created_with=")
	builder.WriteString(s.CreatedWith)
	builder.WriteString(", updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_by=")
	builder.WriteString(fmt.Sprintf("%v", s.UpdatedBy))
	builder.WriteString(", updated_with=")
	builder.WriteString(s.UpdatedWith)
	builder.WriteString(", short=")
	builder.WriteString(s.Short)
	builder.WriteString(", title=")
	builder.WriteString(s.Title)
	builder.WriteString(", description=")
	builder.WriteString(s.Description)
	builder.WriteByte(')')
	return builder.String()
}

// States is a parsable slice of State.
type States []*State

func (s States) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
