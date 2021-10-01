package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/mixin"
)

type MetadataMixin struct {
	mixin.Schema
}

// Fields registers one or more fields on the including schema.
func (MetadataMixin) Fields() []ent.Field {
	return nil
}

// Edges registers the metadata edges on the including schema.
func (MetadataMixin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("metadata", Metadata.Type).
			Unique().
			Annotations(entgql.Bind()),
	}
}
