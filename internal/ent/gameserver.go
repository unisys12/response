// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/responserms/response/internal/ent/gameserver"
	"github.com/responserms/response/internal/ent/metadata"
)

// GameServer is the model entity for the GameServer schema.
type GameServer struct {
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
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Secret holds the value of the "secret" field.
	Secret *string `json:"secret,omitempty"`
	// IPAddress holds the value of the "ip_address" field.
	IPAddress string `json:"ip_address,omitempty"`
	// Monitoring holds the value of the "monitoring" field.
	Monitoring bool `json:"monitoring,omitempty"`
	// DisabledAt holds the value of the "disabled_at" field.
	DisabledAt *time.Time `json:"disabled_at,omitempty"`
	// LastContactAt holds the value of the "last_contact_at" field.
	LastContactAt *time.Time `json:"last_contact_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GameServerQuery when eager-loading is set.
	Edges                GameServerEdges `json:"edges"`
	game_server_metadata *int
}

// GameServerEdges holds the relations/edges for other nodes in the graph.
type GameServerEdges struct {
	// Metadata holds the value of the metadata edge.
	Metadata *Metadata `json:"metadata,omitempty"`
	// Players holds the value of the players edge.
	Players []*Player `json:"players,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// MetadataOrErr returns the Metadata value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GameServerEdges) MetadataOrErr() (*Metadata, error) {
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

// PlayersOrErr returns the Players value or an error if the edge
// was not loaded in eager-loading.
func (e GameServerEdges) PlayersOrErr() ([]*Player, error) {
	if e.loadedTypes[1] {
		return e.Players, nil
	}
	return nil, &NotLoadedError{edge: "players"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GameServer) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case gameserver.FieldMonitoring:
			values[i] = new(sql.NullBool)
		case gameserver.FieldID, gameserver.FieldCreatedBy, gameserver.FieldUpdatedBy:
			values[i] = new(sql.NullInt64)
		case gameserver.FieldCreatedWith, gameserver.FieldUpdatedWith, gameserver.FieldName, gameserver.FieldDescription, gameserver.FieldSecret, gameserver.FieldIPAddress:
			values[i] = new(sql.NullString)
		case gameserver.FieldCreatedAt, gameserver.FieldUpdatedAt, gameserver.FieldDisabledAt, gameserver.FieldLastContactAt:
			values[i] = new(sql.NullTime)
		case gameserver.ForeignKeys[0]: // game_server_metadata
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type GameServer", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GameServer fields.
func (gs *GameServer) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case gameserver.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			gs.ID = int(value.Int64)
		case gameserver.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				gs.CreatedAt = value.Time
			}
		case gameserver.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				gs.CreatedBy = int(value.Int64)
			}
		case gameserver.FieldCreatedWith:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_with", values[i])
			} else if value.Valid {
				gs.CreatedWith = value.String
			}
		case gameserver.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				gs.UpdatedAt = value.Time
			}
		case gameserver.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				gs.UpdatedBy = int(value.Int64)
			}
		case gameserver.FieldUpdatedWith:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_with", values[i])
			} else if value.Valid {
				gs.UpdatedWith = value.String
			}
		case gameserver.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				gs.Name = value.String
			}
		case gameserver.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				gs.Description = value.String
			}
		case gameserver.FieldSecret:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field secret", values[i])
			} else if value.Valid {
				gs.Secret = new(string)
				*gs.Secret = value.String
			}
		case gameserver.FieldIPAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ip_address", values[i])
			} else if value.Valid {
				gs.IPAddress = value.String
			}
		case gameserver.FieldMonitoring:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field monitoring", values[i])
			} else if value.Valid {
				gs.Monitoring = value.Bool
			}
		case gameserver.FieldDisabledAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field disabled_at", values[i])
			} else if value.Valid {
				gs.DisabledAt = new(time.Time)
				*gs.DisabledAt = value.Time
			}
		case gameserver.FieldLastContactAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_contact_at", values[i])
			} else if value.Valid {
				gs.LastContactAt = new(time.Time)
				*gs.LastContactAt = value.Time
			}
		case gameserver.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field game_server_metadata", value)
			} else if value.Valid {
				gs.game_server_metadata = new(int)
				*gs.game_server_metadata = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryMetadata queries the "metadata" edge of the GameServer entity.
func (gs *GameServer) QueryMetadata() *MetadataQuery {
	return (&GameServerClient{config: gs.config}).QueryMetadata(gs)
}

// QueryPlayers queries the "players" edge of the GameServer entity.
func (gs *GameServer) QueryPlayers() *PlayerQuery {
	return (&GameServerClient{config: gs.config}).QueryPlayers(gs)
}

// Update returns a builder for updating this GameServer.
// Note that you need to call GameServer.Unwrap() before calling this method if this GameServer
// was returned from a transaction, and the transaction was committed or rolled back.
func (gs *GameServer) Update() *GameServerUpdateOne {
	return (&GameServerClient{config: gs.config}).UpdateOne(gs)
}

// Unwrap unwraps the GameServer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gs *GameServer) Unwrap() *GameServer {
	tx, ok := gs.config.driver.(*txDriver)
	if !ok {
		panic("ent: GameServer is not a transactional entity")
	}
	gs.config.driver = tx.drv
	return gs
}

// String implements the fmt.Stringer.
func (gs *GameServer) String() string {
	var builder strings.Builder
	builder.WriteString("GameServer(")
	builder.WriteString(fmt.Sprintf("id=%v", gs.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(gs.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", created_by=")
	builder.WriteString(fmt.Sprintf("%v", gs.CreatedBy))
	builder.WriteString(", created_with=")
	builder.WriteString(gs.CreatedWith)
	builder.WriteString(", updated_at=")
	builder.WriteString(gs.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_by=")
	builder.WriteString(fmt.Sprintf("%v", gs.UpdatedBy))
	builder.WriteString(", updated_with=")
	builder.WriteString(gs.UpdatedWith)
	builder.WriteString(", name=")
	builder.WriteString(gs.Name)
	builder.WriteString(", description=")
	builder.WriteString(gs.Description)
	if v := gs.Secret; v != nil {
		builder.WriteString(", secret=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ip_address=")
	builder.WriteString(gs.IPAddress)
	builder.WriteString(", monitoring=")
	builder.WriteString(fmt.Sprintf("%v", gs.Monitoring))
	if v := gs.DisabledAt; v != nil {
		builder.WriteString(", disabled_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := gs.LastContactAt; v != nil {
		builder.WriteString(", last_contact_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// GameServers is a parsable slice of GameServer.
type GameServers []*GameServer

func (gs GameServers) config(cfg config) {
	for _i := range gs {
		gs[_i].config = cfg
	}
}
