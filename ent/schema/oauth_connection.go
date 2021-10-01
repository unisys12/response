package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/responserms/response/ent/schema/mixins"
)

// OAuthConnection holds the schema definition for the OAuthConnection entity.
type OAuthConnection struct {
	ent.Schema
}

// Fields of the OAuthConnection.
func (OAuthConnection) Fields() []ent.Field {
	return []ent.Field{
		field.String("provider"),
		field.String("provider_user_id"),
		field.String("name"),
	}
}

func (OAuthConnection) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.Base{},
		mixins.Audit{},
	}
}

// Edges of the OAuthConnection.
func (OAuthConnection) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("oauth_connections").
			Unique().
			Annotations(entgql.Bind()),
	}
}

func (OAuthConnection) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "oauth_connections",
		},
	}
}
