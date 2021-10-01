package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"github.com/responserms/response/ent/schema/mixins"
)

// DeathManner holds the schema definition for the DeathManner entity.
type DeathManner struct {
	ent.Schema
}

// Fields of the DeathManner.
func (DeathManner) Fields() []ent.Field {
	return []ent.Field{}
}

func (DeathManner) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.Base{},
		mixins.Audit{},
		mixins.CustomizableField{},
		MetadataMixin{},
	}
}

// Edges of the DeathManner.
func (DeathManner) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("death_certificates", DeathCertificate.Type).
			Annotations(entgql.Bind()),
	}
}
