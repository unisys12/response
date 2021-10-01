package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/responserms/response/ent/schema/mixins"
)

// VehicleRegistration holds the schema definition for the VehicleRegistration entity.
type VehicleRegistration struct {
	ent.Schema
}

// Fields of the VehicleRegistration.
func (VehicleRegistration) Fields() []ent.Field {
	return []ent.Field{
		field.String("plate"),
		field.Time("expired_at"),
	}
}

func (VehicleRegistration) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.Base{},
		mixins.Audit{},
		MetadataMixin{},
	}
}

// Edges of the VehicleRegistration.
func (VehicleRegistration) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("state", State.Type).
			Required().
			Ref("vehicle_registrations").
			Unique(),
		edge.From("person", Person.Type).
			Required().
			Ref("vehicle_registrations").
			Unique(),
		edge.From("vehicle", Vehicle.Type).
			Ref("registrations").
			Required().
			Unique(),
	}
}
