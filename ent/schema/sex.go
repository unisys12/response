package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"github.com/responserms/response/ent/schema/mixins"
)

// Sex holds the schema definition for the Sex entity.
type Sex struct {
	ent.Schema
}

// Fields of the Sex.
func (Sex) Fields() []ent.Field {
	return []ent.Field{}
}

func (Sex) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.Base{},
		mixins.Audit{},
		mixins.CustomizableField{},
		MetadataMixin{},
	}
}

// Edges of the Sex.
func (Sex) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("people", Person.Type),
	}
}
