package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"github.com/responserms/response/ent/schema/mixins"
)

// Ethnicity holds the schema definition for the Ethnicity entity.
type Ethnicity struct {
	ent.Schema
}

// Fields of the Ethnicity.
func (Ethnicity) Fields() []ent.Field {
	return nil
}

func (Ethnicity) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.Base{},
		mixins.Audit{},
		mixins.CustomizableField{},
		MetadataMixin{},
	}
}

// Edges of the Ethnicity.
func (Ethnicity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("people", Person.Type),
	}
}
