// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/responserms/response/internal/ent/metadata"
	"github.com/responserms/response/internal/ent/user"
)

// User is the model entity for the User schema.
type User struct {
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
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Password holds the value of the "password" field.
	Password *string `json:"password,omitempty"`
	// AvatarURL holds the value of the "avatar_url" field.
	AvatarURL *string `json:"avatar_url,omitempty"`
	// Permissions holds the value of the "permissions" field.
	Permissions []int `json:"permissions,omitempty"`
	// FirstSetupAt holds the value of the "first_setup_at" field.
	FirstSetupAt *time.Time `json:"first_setup_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges         UserEdges `json:"edges"`
	user_metadata *int
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Metadata holds the value of the metadata edge.
	Metadata *Metadata `json:"metadata,omitempty"`
	// OauthConnections holds the value of the oauth_connections edge.
	OauthConnections []*OAuthConnection `json:"oauth_connections,omitempty"`
	// People holds the value of the people edge.
	People []*Person `json:"people,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// MetadataOrErr returns the Metadata value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) MetadataOrErr() (*Metadata, error) {
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

// OauthConnectionsOrErr returns the OauthConnections value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) OauthConnectionsOrErr() ([]*OAuthConnection, error) {
	if e.loadedTypes[1] {
		return e.OauthConnections, nil
	}
	return nil, &NotLoadedError{edge: "oauth_connections"}
}

// PeopleOrErr returns the People value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) PeopleOrErr() ([]*Person, error) {
	if e.loadedTypes[2] {
		return e.People, nil
	}
	return nil, &NotLoadedError{edge: "people"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldPermissions:
			values[i] = new([]byte)
		case user.FieldID, user.FieldCreatedBy, user.FieldUpdatedBy:
			values[i] = new(sql.NullInt64)
		case user.FieldCreatedWith, user.FieldUpdatedWith, user.FieldName, user.FieldEmail, user.FieldPassword, user.FieldAvatarURL:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt, user.FieldFirstSetupAt:
			values[i] = new(sql.NullTime)
		case user.ForeignKeys[0]: // user_metadata
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				u.CreatedBy = int(value.Int64)
			}
		case user.FieldCreatedWith:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_with", values[i])
			} else if value.Valid {
				u.CreatedWith = value.String
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case user.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				u.UpdatedBy = int(value.Int64)
			}
		case user.FieldUpdatedWith:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_with", values[i])
			} else if value.Valid {
				u.UpdatedWith = value.String
			}
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = new(string)
				*u.Password = value.String
			}
		case user.FieldAvatarURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar_url", values[i])
			} else if value.Valid {
				u.AvatarURL = new(string)
				*u.AvatarURL = value.String
			}
		case user.FieldPermissions:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field permissions", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &u.Permissions); err != nil {
					return fmt.Errorf("unmarshal field permissions: %w", err)
				}
			}
		case user.FieldFirstSetupAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field first_setup_at", values[i])
			} else if value.Valid {
				u.FirstSetupAt = new(time.Time)
				*u.FirstSetupAt = value.Time
			}
		case user.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_metadata", value)
			} else if value.Valid {
				u.user_metadata = new(int)
				*u.user_metadata = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryMetadata queries the "metadata" edge of the User entity.
func (u *User) QueryMetadata() *MetadataQuery {
	return (&UserClient{config: u.config}).QueryMetadata(u)
}

// QueryOauthConnections queries the "oauth_connections" edge of the User entity.
func (u *User) QueryOauthConnections() *OAuthConnectionQuery {
	return (&UserClient{config: u.config}).QueryOauthConnections(u)
}

// QueryPeople queries the "people" edge of the User entity.
func (u *User) QueryPeople() *PersonQuery {
	return (&UserClient{config: u.config}).QueryPeople(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", created_by=")
	builder.WriteString(fmt.Sprintf("%v", u.CreatedBy))
	builder.WriteString(", created_with=")
	builder.WriteString(u.CreatedWith)
	builder.WriteString(", updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_by=")
	builder.WriteString(fmt.Sprintf("%v", u.UpdatedBy))
	builder.WriteString(", updated_with=")
	builder.WriteString(u.UpdatedWith)
	builder.WriteString(", name=")
	builder.WriteString(u.Name)
	builder.WriteString(", email=")
	builder.WriteString(u.Email)
	if v := u.Password; v != nil {
		builder.WriteString(", password=")
		builder.WriteString(*v)
	}
	if v := u.AvatarURL; v != nil {
		builder.WriteString(", avatar_url=")
		builder.WriteString(*v)
	}
	builder.WriteString(", permissions=")
	builder.WriteString(fmt.Sprintf("%v", u.Permissions))
	if v := u.FirstSetupAt; v != nil {
		builder.WriteString(", first_setup_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
