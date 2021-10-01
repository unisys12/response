package mixins

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type CustomizableField struct {
	mixin.Schema
}

func (CustomizableField) Fields() []ent.Field {
	return []ent.Field{
		field.String("short").
			MaxLen(5),
		field.String("title").
			MaxLen(64),
		field.String("description").
			Optional().
			MaxLen(160),
	}
}
