package mixins

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/mixin"
)

type Base struct {
	mixin.Schema
	Entity string
}

func (b Base) Fields() []ent.Field {
	return []ent.Field{}
}
