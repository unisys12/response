package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"github.com/responserms/response/ent/schema/mixins"
)

// VehicleModel holds the schema definition for the VehicleModel entity.
type VehicleModel struct {
	ent.Schema
}

// Fields of the VehicleModel.
func (VehicleModel) Fields() []ent.Field {
	return nil
}

func (VehicleModel) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.Base{},
		mixins.Audit{},
		mixins.CustomizableField{},
		MetadataMixin{},
	}
}

// Edges of the VehicleModel.
func (VehicleModel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("vehicles", Vehicle.Type),
	}
}
