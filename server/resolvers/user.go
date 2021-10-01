package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/responserms/response/internal/ent"
	"github.com/responserms/response/internal/graphql/server"
)

func (r *userResolver) FirstSetupCompleted(ctx context.Context, obj *ent.User) (bool, error) {
	if obj == nil {
		return false, nil
	}

	return obj.FirstSetupAt != nil, nil
}

// User returns server.UserResolver implementation.
func (r *Resolver) User() server.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
