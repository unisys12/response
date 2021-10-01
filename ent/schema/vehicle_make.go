package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"github.com/responserms/response/ent/schema/mixins"
)

// VehicleMake holds the schema definition for the VehicleMake entity.
type VehicleMake struct {
	ent.Schema
}

// Fields of the VehicleMake.
func (VehicleMake) Fields() []ent.Field {
	return nil
}

func (VehicleMake) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.Base{},
		mixins.Audit{},
		mixins.CustomizableField{},
	}
}

// Edges of the VehicleMake.
func (VehicleMake) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("vehicles", Vehicle.Type),
	}
}
