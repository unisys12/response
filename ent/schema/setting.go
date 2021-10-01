package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/responserms/response/ent/schema/mixins"
	"github.com/responserms/response/internal/structs"
)

// Setting holds the schema definition for the Setting entity.
type Setting struct {
	ent.Schema
}

// Fields of the Setting.
func (Setting) Fields() []ent.Field {
	return []ent.Field{
		field.String("key").
			Unique(),
		field.JSON("data", structs.SettingData{}).
			Optional(),
	}
}

func (Setting) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.Base{},
		mixins.Audit{},
	}
}

// Edges of the Setting.
func (Setting) Edges() []ent.Edge {
	return nil
}
