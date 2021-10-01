package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/responserms/response/internal/graphql/server"
	"github.com/responserms/response/internal/graphql/types"
	"github.com/responserms/response/response"
)

func (r *mutationResolver) RegisterWithPassword(ctx context.Context, input types.RegistrationData) (*types.RegistrationResult, error) {
	result := &types.RegistrationResult{}

	user, err := r.core.RegisterWithPassword(ctx, input.Name, input.Email, input.Password)
	if err != nil {
		switch {
		case errors.Is(err, response.ErrPasswordAuthDisabled):
			result.Errors = append(result.Errors, &types.PasswordAuthDisabled{
				Path:    []string{graphql.GetPath(ctx).String()},
				Message: "Password authentication is disabled in Response.",
			})
		case errors.Is(err, response.ErrEmailExists):
			result.Errors = append(result.Errors, &types.EmailAlreadyExists{
				Path:    []string{graphql.GetPath(ctx).String()},
				Message: "Your email is already registered with Response.",
			})
		default:
			return result, nil
		}
	}

	result.User = user
	return result, nil
}

// Mutation returns server.MutationResolver implementation.
func (r *Resolver) Mutation() server.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
