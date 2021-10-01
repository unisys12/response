package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/responserms/response/ent/schema/annotations"
	"github.com/responserms/response/ent/schema/mixins"
)

// GameServer holds the schema definition for the GameServer entity.
type GameServer struct {
	ent.Schema
}

func (GameServer) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.Base{},
		mixins.Audit{},
		MetadataMixin{},
	}
}

// Fields of the GameServer.
func (GameServer) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Annotations(annotations.Validator{
				Create: "required",
			}),
		field.String("description").
			Annotations(annotations.Validator{
				Both: "lte:200",
			}),
		field.String("secret").
			Optional().
			Nillable(),
		field.String("ip_address").
			Annotations(annotations.Validator{
				Both: "ip",
			}),
		field.Bool("monitoring").
			Default(false).
			Optional(),
		field.Time("disabled_at").
			Optional().
			Nillable(),
		field.Time("last_contact_at").
			Optional().
			Nillable(),
	}
}

// Edges of the GameServer.
func (GameServer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("players", Player.Type).
			Ref("servers"),
	}
}
