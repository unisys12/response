package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/responserms/response/internal/ent"
	"github.com/responserms/response/internal/graphql/server"
)

func (r *ethnicityResolver) CreatedBy(ctx context.Context, obj *ent.Ethnicity) (*ent.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *ethnicityResolver) UpdatedBy(ctx context.Context, obj *ent.Ethnicity) (*ent.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateEthnicity(ctx context.Context, input ent.CreateEthnicity) (*ent.Ethnicity, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateEthnicity(ctx context.Context, id int, input ent.UpdateEthnicity) (*ent.Ethnicity, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteEthnicity(ctx context.Context, id int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Ethnicities(ctx context.Context) ([]*ent.Ethnicity, error) {
	panic(fmt.Errorf("not implemented"))
}

// Ethnicity returns server.EthnicityResolver implementation.
func (r *Resolver) Ethnicity() server.EthnicityResolver { return &ethnicityResolver{r} }

// Mutation returns server.MutationResolver implementation.
func (r *Resolver) Mutation() server.MutationResolver { return &mutationResolver{r} }

type ethnicityResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
