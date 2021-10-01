// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/responserms/response/internal/ent/metadata"
	"github.com/responserms/response/internal/ent/player"
)

// Player is the model entity for the Player schema.
type Player struct {
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
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PlayerQuery when eager-loading is set.
	Edges           PlayerEdges `json:"edges"`
	player_metadata *int
}

// PlayerEdges holds the relations/edges for other nodes in the graph.
type PlayerEdges struct {
	// Metadata holds the value of the metadata edge.
	Metadata *Metadata `json:"metadata,omitempty"`
	// Servers holds the value of the servers edge.
	Servers []*GameServer `json:"servers,omitempty"`
	// Identifiers holds the value of the identifiers edge.
	Identifiers []*PlayerIdentifier `json:"identifiers,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// MetadataOrErr returns the Metadata value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PlayerEdges) MetadataOrErr() (*Metadata, error) {
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

// ServersOrErr returns the Servers value or an error if the edge
// was not loaded in eager-loading.
func (e PlayerEdges) ServersOrErr() ([]*GameServer, error) {
	if e.loadedTypes[1] {
		return e.Servers, nil
	}
	return nil, &NotLoadedError{edge: "servers"}
}

// IdentifiersOrErr returns the Identifiers value or an error if the edge
// was not loaded in eager-loading.
func (e PlayerEdges) IdentifiersOrErr() ([]*PlayerIdentifier, error) {
	if e.loadedTypes[2] {
		return e.Identifiers, nil
	}
	return nil, &NotLoadedError{edge: "identifiers"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Player) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case player.FieldID, player.FieldCreatedBy, player.FieldUpdatedBy:
			values[i] = new(sql.NullInt64)
		case player.FieldCreatedWith, player.FieldUpdatedWith, player.FieldName:
			values[i] = new(sql.NullString)
		case player.FieldCreatedAt, player.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case player.ForeignKeys[0]: // player_metadata
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Player", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Player fields.
func (pl *Player) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case player.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pl.ID = int(value.Int64)
		case player.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pl.CreatedAt = value.Time
			}
		case player.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				pl.CreatedBy = int(value.Int64)
			}
		case player.FieldCreatedWith:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_with", values[i])
			} else if value.Valid {
				pl.CreatedWith = value.String
			}
		case player.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pl.UpdatedAt = value.Time
			}
		case player.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				pl.UpdatedBy = int(value.Int64)
			}
		case player.FieldUpdatedWith:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_with", values[i])
			} else if value.Valid {
				pl.UpdatedWith = value.String
			}
		case player.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pl.Name = value.String
			}
		case player.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field player_metadata", value)
			} else if value.Valid {
				pl.player_metadata = new(int)
				*pl.player_metadata = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryMetadata queries the "metadata" edge of the Player entity.
func (pl *Player) QueryMetadata() *MetadataQuery {
	return (&PlayerClient{config: pl.config}).QueryMetadata(pl)
}

// QueryServers queries the "servers" edge of the Player entity.
func (pl *Player) QueryServers() *GameServerQuery {
	return (&PlayerClient{config: pl.config}).QueryServers(pl)
}

// QueryIdentifiers queries the "identifiers" edge of the Player entity.
func (pl *Player) QueryIdentifiers() *PlayerIdentifierQuery {
	return (&PlayerClient{config: pl.config}).QueryIdentifiers(pl)
}

// Update returns a builder for updating this Player.
// Note that you need to call Player.Unwrap() before calling this method if this Player
// was returned from a transaction, and the transaction was committed or rolled back.
func (pl *Player) Update() *PlayerUpdateOne {
	return (&PlayerClient{config: pl.config}).UpdateOne(pl)
}

// Unwrap unwraps the Player entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pl *Player) Unwrap() *Player {
	tx, ok := pl.config.driver.(*txDriver)
	if !ok {
		panic("ent: Player is not a transactional entity")
	}
	pl.config.driver = tx.drv
	return pl
}

// String implements the fmt.Stringer.
func (pl *Player) String() string {
	var builder strings.Builder
	builder.WriteString("Player(")
	builder.WriteString(fmt.Sprintf("id=%v", pl.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(pl.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", created_by=")
	builder.WriteString(fmt.Sprintf("%v", pl.CreatedBy))
	builder.WriteString(", created_with=")
	builder.WriteString(pl.CreatedWith)
	builder.WriteString(", updated_at=")
	builder.WriteString(pl.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_by=")
	builder.WriteString(fmt.Sprintf("%v", pl.UpdatedBy))
	builder.WriteString(", updated_with=")
	builder.WriteString(pl.UpdatedWith)
	builder.WriteString(", name=")
	builder.WriteString(pl.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Players is a parsable slice of Player.
type Players []*Player

func (pl Players) config(cfg config) {
	for _i := range pl {
		pl[_i].config = cfg
	}
}