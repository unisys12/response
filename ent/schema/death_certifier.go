package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"github.com/responserms/response/ent/schema/mixins"
)

// DeathCertifier holds the schema definition for the DeathCertifier entity.
type DeathCertifier struct {
	ent.Schema
}

// Fields of the DeathCertifier.
func (DeathCertifier) Fields() []ent.Field {
	return nil
}

func (DeathCertifier) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.Base{},
		mixins.Audit{},
		mixins.CustomizableField{},
		MetadataMixin{},
	}
}

// Edges of the DeathCertifier.
func (DeathCertifier) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("death_certificates", DeathCertificate.Type).
			Annotations(entgql.Bind()),
	}
}
