package mixins

import (
	"context"
	"fmt"
	"time"

	"entgo.io/contrib/entgql"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/responserms/response/internal/contx"
)

type Audit struct {
	mixin.Schema
}

func (Audit) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Annotations(entgql.OrderField("CREATED_AT")).
			Immutable().
			Default(time.Now),
		field.Int("created_by").
			Optional(),
		field.String("created_with").
			Optional(),
		field.Time("updated_at").
			Annotations(entgql.OrderField("UPDATED_AT")).
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Int("updated_by").
			Optional(),
		field.String("updated_with").
			Optional(),
	}
}

func (Audit) Hooks() []ent.Hook {
	return []ent.Hook{
		AuditHook,
	}
}

func AuditHook(next ent.Mutator) ent.Mutator {
	// AuditLogger wraps the methods that are shared between all mutations of
	// schemas that embed the AuditLog mixin. The variable "exists" is true, if
	// the field already exists in the mutation (e.g. was set by a different hook).
	type AuditLogger interface {
		SetCreatedAt(time.Time)
		CreatedAt() (value time.Time, exists bool)
		SetCreatedBy(int)
		CreatedBy() (id int, exists bool)
		SetUpdatedAt(time.Time)
		UpdatedAt() (value time.Time, exists bool)
		SetUpdatedBy(int)
		UpdatedBy() (id int, exists bool)
	}
	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
		ml, ok := m.(AuditLogger)
		if !ok {
			return nil, fmt.Errorf("unexpected audit-log call from mutation type %T", m)
		}

		userID := contx.GetUserID(ctx)

		switch op := m.Op(); {
		case op.Is(ent.OpCreate):
			ml.SetCreatedAt(time.Now())

			if _, exists := ml.CreatedBy(); !exists && userID != 0 {
				ml.SetCreatedBy(userID)
			}
		case op.Is(ent.OpUpdateOne | ent.OpUpdate):
			ml.SetUpdatedAt(time.Now())

			if _, exists := ml.UpdatedBy(); !exists && userID != 0 {
				ml.SetUpdatedBy(userID)
			}
		}
		return next.Mutate(ctx, m)
	})
}
