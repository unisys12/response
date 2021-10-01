package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/responserms/response/ent/schema/mixins"
)

// PlayerIdentifier holds the schema definition for the PlayerIdentifier entity.
type PlayerIdentifier struct {
	ent.Schema
}

// Fields of the PlayerIdentifier.
func (PlayerIdentifier) Fields() []ent.Field {
	return []ent.Field{
		field.String("value").
			NotEmpty().
			Unique().
			Immutable(),
	}
}

func (PlayerIdentifier) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.Base{},
		mixins.Audit{},
	}
}

// Edges of the PlayerIdentifier.
func (PlayerIdentifier) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("player", Player.Type).
			Ref("identifiers").
			Unique(),
	}
}
