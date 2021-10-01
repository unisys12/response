package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/responserms/response/ent/schema/mixins"
)

// DeathCertificate holds the schema definition for the DeathCertificate entity.
type DeathCertificate struct {
	ent.Schema
}

func (DeathCertificate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.Base{},
		mixins.Audit{},
		MetadataMixin{},
	}
}

// Fields of the DeathCertificate.
func (DeathCertificate) Fields() []ent.Field {
	return []ent.Field{
		field.String("cause"),
		field.Text("comments").
			Optional(),
		field.Time("certified_at").
			Optional().
			Nillable(),
		field.Text("certifier_comments").
			Optional(),
		field.Bool("requires_certification").
			Default(false),
	}
}

// Edges of the DeathCertificate.
func (DeathCertificate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("decedent", Person.Type).
			Ref("death_certificate").
			Unique().
			Annotations(entgql.Bind()).
			Required(),
		edge.From("manner", DeathManner.Type).
			Ref("death_certificates").
			Unique().
			Annotations(entgql.Bind()),
		edge.From("place", DeathPlace.Type).
			Ref("death_certificates").
			Unique().
			Annotations(entgql.Bind()),
		edge.From("certifier", DeathCertifier.Type).
			Ref("death_certificates").
			Unique().
			Annotations(entgql.Bind()),
		edge.From("certified_by", Person.Type).
			Ref("certified_deaths").
			Unique().
			Annotations(entgql.Bind()),
	}
}
