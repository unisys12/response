package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/responserms/response/ent/schema/mixins"
)

// Person holds the schema definition for the Person entity.
type Person struct {
	ent.Schema
}

// Fields of the Person.
func (Person) Fields() []ent.Field {
	return []ent.Field{
		field.String("first_name").
			Optional().
			Annotations(entgql.OrderField("FIRST_NAME")),
		field.String("middle_name").
			Optional(),
		field.String("last_name").
			Optional().
			Annotations(entgql.OrderField("LAST_NAME")),
		field.String("suffix").
			Optional(),
		field.Time("date_of_birth").
			Optional().
			Annotations(entgql.OrderField("DATE_OF_BIRTH")),
		field.String("birthplace").
			Optional(),
		field.Time("deceased_at").
			Optional().
			Nillable(),
		field.Time("archived_at").
			Optional().
			Nillable(),
	}
}

func (Person) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.Base{},
		mixins.Audit{},
		MetadataMixin{},
	}
}

// Edges of the Person.
func (Person) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("vehicle_registrations", VehicleRegistration.Type),
		edge.From("race", Race.Type).
			Ref("people").
			Unique(),
		edge.From("ethnicity", Ethnicity.Type).
			Ref("people").
			Unique(),
		edge.From("sex", Sex.Type).
			Ref("people").
			Unique(),
		edge.From("user", User.Type).
			Ref("people").
			Unique(),
		edge.To("death_certificate", DeathCertificate.Type).
			Unique(),
		edge.To("certified_deaths", DeathCertificate.Type),
	}
}
