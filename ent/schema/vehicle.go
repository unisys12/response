package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/responserms/response/ent/schema/mixins"
)

// Vehicle holds the schema definition for the Vehicle entity.
type Vehicle struct {
	ent.Schema
}

// Fields of the Vehicle.
func (Vehicle) Fields() []ent.Field {
	return []ent.Field{
		field.String("vin").
			Optional(),
		field.String("style").
			Optional(),
		field.Time("extra_features").
			Optional(),
		field.Text("private_notes").
			Optional(),
	}
}

func (Vehicle) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.Base{},
		mixins.Audit{},
		MetadataMixin{},
	}
}

// Edges of the Vehicle.
func (Vehicle) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("registrations", VehicleRegistration.Type),
		edge.From("make", VehicleMake.Type).
			Ref("vehicles").
			Unique(),
		edge.From("model", VehicleModel.Type).
			Ref("vehicles").
			Unique(),
		edge.From("major_color", VehicleColor.Type).
			Ref("vehicles_major").
			Unique(),
		edge.From("minor_color", VehicleColor.Type).
			Ref("vehicles_minor").
			Unique(),
		edge.From("class", VehicleClass.Type).
			Ref("vehicles").
			Unique(),
	}
}
