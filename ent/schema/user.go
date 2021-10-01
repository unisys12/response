package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/responserms/response/ent/schema/mixins"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("email").
			Unique(),
		field.String("password").
			NotEmpty().
			Optional().
			Nillable(),
		field.String("avatar_url").
			Optional().
			Nillable(),
		field.Ints("permissions").
			Optional(),
		field.Time("first_setup_at").
			Optional().
			Nillable(),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.Base{},
		mixins.Audit{},
		MetadataMixin{},
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("oauth_connections", OAuthConnection.Type).
			Annotations(entgql.Bind()),
		edge.To("people", Person.Type).
			Annotations(entgql.Bind()),
	}
}
