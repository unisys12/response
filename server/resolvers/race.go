package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/responserms/response/internal/ent"
	"github.com/responserms/response/internal/graphql/server"
)

func (r *mutationResolver) CreateRace(ctx context.Context, input ent.CreateRace) (*ent.Race, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateRace(ctx context.Context, id int, input ent.UpdateRace) (*ent.Race, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteRace(ctx context.Context, id int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Races(ctx context.Context) ([]*ent.Race, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *raceResolver) CreatedBy(ctx context.Context, obj *ent.Race) (*ent.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *raceResolver) UpdatedBy(ctx context.Context, obj *ent.Race) (*ent.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Race returns server.RaceResolver implementation.
func (r *Resolver) Race() server.RaceResolver { return &raceResolver{r} }

type raceResolver struct{ *Resolver }
