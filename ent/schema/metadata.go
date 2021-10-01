package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/responserms/response/ent/schema/mixins"
)

// Metadata holds the schema definition for the Metadata entity.
type Metadata struct {
	ent.Schema
}

// Fields of the Metadata.
func (Metadata) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("data", map[string]interface{}{}),
	}
}

func (Metadata) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.Base{},
		mixins.Audit{},
	}
}

// Edges of the Metadata.
func (Metadata) Edges() []ent.Edge {
	return []ent.Edge{

		// Back-references to various metadata-supporting entities
		//edge.From("user", User.Type).
		//	Ref("metadata").
		//	Unique().
		//	Annotations(entgql.Bind()),
	}
}
