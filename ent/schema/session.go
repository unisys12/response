package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/responserms/response/ent/schema/mixins"
)

// Session holds the schema definition for the Session entity.
type Session struct {
	ent.Schema
}

// Fields of the Session.
func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.Time("started_at").
			Annotations(entgql.OrderField("STARTED_AT")),
		field.String("started_from").
			Optional().
			Nillable(),
		field.String("ip_address").
			Optional().
			Nillable(),
		field.Time("ended_at").
			Annotations(entgql.OrderField("ENDED_AT")),
	}
}

func (Session) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.Base{},
		mixins.Audit{},
	}
}

// Edges of the Session.
func (Session) Edges() []ent.Edge {
	return []ent.Edge{}
}
