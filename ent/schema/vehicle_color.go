package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"github.com/responserms/response/ent/schema/mixins"
)

// VehicleColor holds the schema definition for the VehicleColor entity.
type VehicleColor struct {
	ent.Schema
}

// Fields of the VehicleColor.
func (VehicleColor) Fields() []ent.Field {
	return nil
}

func (VehicleColor) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.Base{},
		mixins.Audit{},
		mixins.CustomizableField{},
		MetadataMixin{},
	}
}

// Edges of the VehicleColor.
func (VehicleColor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("vehicles_major", Vehicle.Type),
		edge.To("vehicles_minor", Vehicle.Type),
	}
}
