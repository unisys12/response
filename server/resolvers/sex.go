package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/responserms/response/internal/ent"
	"github.com/responserms/response/internal/graphql/server"
)

func (r *mutationResolver) CreateSex(ctx context.Context, input ent.CreateSex) (*ent.Sex, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateSex(ctx context.Context, id int, input ent.UpdateSex) (*ent.Sex, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteSex(ctx context.Context, id int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Sexes(ctx context.Context) ([]*ent.Sex, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *sexResolver) CreatedBy(ctx context.Context, obj *ent.Sex) (*ent.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *sexResolver) UpdatedBy(ctx context.Context, obj *ent.Sex) (*ent.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Sex returns server.SexResolver implementation.
func (r *Resolver) Sex() server.SexResolver { return &sexResolver{r} }

type sexResolver struct{ *Resolver }
