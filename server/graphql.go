package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/apollotracing"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gorilla/websocket"
	"github.com/responserms/response/server/resolvers"
)

func (s *server) initAPIRoutes() error {
	graphql := handler.New(resolvers.NewSchema(s.core))

	// Extensions
	graphql.Use(extension.Introspection{})
	graphql.Use(apollotracing.Tracer{})

	// Add Transports
	graphql.AddTransport(transport.GET{})
	graphql.AddTransport(transport.POST{})
	graphql.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})

	// Register the GraphQL API
	s.routes.Route("/api", func(r chi.Router) {
		r.Get("/", s.handleGraphQLPlayground())
		r.Mount("/graphql", graphql)
	})

	return nil
}
