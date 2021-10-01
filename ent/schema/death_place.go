package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"github.com/responserms/response/ent/schema/mixins"
)

// DeathPlace holds the schema definition for the DeathPlace entity.
type DeathPlace struct {
	ent.Schema
}

// Fields of the DeathPlace.
func (DeathPlace) Fields() []ent.Field {
	return nil
}

func (DeathPlace) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.Base{},
		mixins.Audit{},
		mixins.CustomizableField{},
		MetadataMixin{},
	}
}

// Edges of the DeathPlace.
func (DeathPlace) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("death_certificates", DeathCertificate.Type).
			Annotations(entgql.Bind()),
	}
}
