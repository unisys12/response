package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"github.com/responserms/response/ent/schema/mixins"
)

// VehicleClass holds the schema definition for the VehicleClass entity.
type VehicleClass struct {
	ent.Schema
}

// Fields of the VehicleClass.
func (VehicleClass) Fields() []ent.Field {
	return nil
}

func (VehicleClass) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.Base{},
		mixins.Audit{},
		mixins.CustomizableField{},
		MetadataMixin{},
	}
}

// Edges of the VehicleClass.
func (VehicleClass) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("vehicles", Vehicle.Type),
	}
}
