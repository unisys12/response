package resolvers

import (
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/responserms/response/internal/graphql/server"
	"github.com/responserms/response/response"
	"github.com/rs/zerolog"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	core   *response.Core
	logger zerolog.Logger
}

func NewSchema(c *response.Core) graphql.ExecutableSchema {
	return server.NewExecutableSchema(server.Config{
		Resolvers: &Resolver{
			core:   c,
			logger: c.Logger("graphql"),
		},
	})
}

func newTestSchema(t *testing.T, c *response.Core) graphql.ExecutableSchema {
	return server.NewExecutableSchema(server.Config{
		Resolvers: &Resolver{
			core:   c,
			logger: c.Logger("graphql"),
		},
	})
}
