package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"github.com/responserms/response/ent/schema/mixins"
)

// State holds the schema definition for the State entity.
type State struct {
	ent.Schema
}

// Fields of the State.
func (State) Fields() []ent.Field {
	return nil
}

// Mixin of the State.
func (State) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.Base{},
		mixins.Audit{},
		mixins.CustomizableField{},
	}
}

// Edges of the State.
func (State) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("vehicle_registrations", VehicleRegistration.Type),
	}
}
