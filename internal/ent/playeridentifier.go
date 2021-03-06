// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/responserms/response/internal/ent/player"
	"github.com/responserms/response/internal/ent/playeridentifier"
)

// PlayerIdentifier is the model entity for the PlayerIdentifier schema.
type PlayerIdentifier struct {
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
	// Value holds the value of the "value" field.
	Value string `json:"value,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PlayerIdentifierQuery when eager-loading is set.
	Edges              PlayerIdentifierEdges `json:"edges"`
	player_identifiers *int
}

// PlayerIdentifierEdges holds the relations/edges for other nodes in the graph.
type PlayerIdentifierEdges struct {
	// Player holds the value of the player edge.
	Player *Player `json:"player,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PlayerOrErr returns the Player value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PlayerIdentifierEdges) PlayerOrErr() (*Player, error) {
	if e.loadedTypes[0] {
		if e.Player == nil {
			// The edge player was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: player.Label}
		}
		return e.Player, nil
	}
	return nil, &NotLoadedError{edge: "player"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PlayerIdentifier) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case playeridentifier.FieldID, playeridentifier.FieldCreatedBy, playeridentifier.FieldUpdatedBy:
			values[i] = new(sql.NullInt64)
		case playeridentifier.FieldCreatedWith, playeridentifier.FieldUpdatedWith, playeridentifier.FieldValue:
			values[i] = new(sql.NullString)
		case playeridentifier.FieldCreatedAt, playeridentifier.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case playeridentifier.ForeignKeys[0]: // player_identifiers
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type PlayerIdentifier", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PlayerIdentifier fields.
func (pi *PlayerIdentifier) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case playeridentifier.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pi.ID = int(value.Int64)
		case playeridentifier.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pi.CreatedAt = value.Time
			}
		case playeridentifier.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				pi.CreatedBy = int(value.Int64)
			}
		case playeridentifier.FieldCreatedWith:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_with", values[i])
			} else if value.Valid {
				pi.CreatedWith = value.String
			}
		case playeridentifier.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pi.UpdatedAt = value.Time
			}
		case playeridentifier.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				pi.UpdatedBy = int(value.Int64)
			}
		case playeridentifier.FieldUpdatedWith:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_with", values[i])
			} else if value.Valid {
				pi.UpdatedWith = value.String
			}
		case playeridentifier.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				pi.Value = value.String
			}
		case playeridentifier.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field player_identifiers", value)
			} else if value.Valid {
				pi.player_identifiers = new(int)
				*pi.player_identifiers = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryPlayer queries the "player" edge of the PlayerIdentifier entity.
func (pi *PlayerIdentifier) QueryPlayer() *PlayerQuery {
	return (&PlayerIdentifierClient{config: pi.config}).QueryPlayer(pi)
}

// Update returns a builder for updating this PlayerIdentifier.
// Note that you need to call PlayerIdentifier.Unwrap() before calling this method if this PlayerIdentifier
// was returned from a transaction, and the transaction was committed or rolled back.
func (pi *PlayerIdentifier) Update() *PlayerIdentifierUpdateOne {
	return (&PlayerIdentifierClient{config: pi.config}).UpdateOne(pi)
}

// Unwrap unwraps the PlayerIdentifier entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pi *PlayerIdentifier) Unwrap() *PlayerIdentifier {
	tx, ok := pi.config.driver.(*txDriver)
	if !ok {
		panic("ent: PlayerIdentifier is not a transactional entity")
	}
	pi.config.driver = tx.drv
	return pi
}

// String implements the fmt.Stringer.
func (pi *PlayerIdentifier) String() string {
	var builder strings.Builder
	builder.WriteString("PlayerIdentifier(")
	builder.WriteString(fmt.Sprintf("id=%v", pi.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(pi.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", created_by=")
	builder.WriteString(fmt.Sprintf("%v", pi.CreatedBy))
	builder.WriteString(", created_with=")
	builder.WriteString(pi.CreatedWith)
	builder.WriteString(", updated_at=")
	builder.WriteString(pi.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_by=")
	builder.WriteString(fmt.Sprintf("%v", pi.UpdatedBy))
	builder.WriteString(", updated_with=")
	builder.WriteString(pi.UpdatedWith)
	builder.WriteString(", value=")
	builder.WriteString(pi.Value)
	builder.WriteByte(')')
	return builder.String()
}

// PlayerIdentifiers is a parsable slice of PlayerIdentifier.
type PlayerIdentifiers []*PlayerIdentifier

func (pi PlayerIdentifiers) config(cfg config) {
	for _i := range pi {
		pi[_i].config = cfg
	}
}
