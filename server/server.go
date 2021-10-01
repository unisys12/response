package server

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/responserms/response/internal/auth"

	"github.com/labstack/echo/v4"
	"github.com/responserms/response/internal/config"
	"github.com/responserms/response/response"
	"github.com/rs/zerolog"
)

type Server interface {
	MonitoredService(ctx context.Context) (func() error, func(error))
}

type server struct {
	core   *response.Core
	config *config.Config
	logger zerolog.Logger
	routes *chi.Mux
	http2  *echo.Echo
	http   *http.Server
	https  *http.Server
	auth   *auth.Service
}

// New creates a new instance of the Response Server.
func New(core *response.Core) (Server, error) {
	if core == nil {
		return nil, errors.New("provided Core is nil")
	}

	authService, err := auth.New(core)
	if err != nil {
		return nil, err
	}

	server := &server{
		core:   core,
		config: core.RunningConfig(),
		logger: core.Logger("server"),
		routes: chi.NewRouter(),
		auth:   authService,
	}

	//
	//server.http2.HideBanner = true
	//server.http2.HidePort = true
	server.routes.Use(server.auth.Global()...)

	// Global auth middleware.
	//server.http2.Use(server.auth.Global()...)

	//server.http.Use(middleware.Recover())
	//server.http.Use(middleware.RequestID())

	// Apply our internal context.Context values to the http.Request Context.
	//server.http.Use(func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
	//	return func(c echo.Context) error {
	//		req := c.Request()
	//		ctx := server.core.ApplyContext(req.Context())
	//		ctx = contx.SetEchoContext(ctx, c)
	//		c.SetRequest(req.WithContext(ctx))
	//
	//		return handlerFunc(c)
	//	}
	//})

	// Register our Go Template renderer on the echo.Echo instance.
	server.registerRenderer()

	//api := server.http.Group("/api")
	//api.GET("", server.handleGraphQLPlayground)
	//api.GET("/graphql", server.graphql.Handler)
	//api.POST("/graphql", server.graphql.Handler)

	// initialize auth services
	if err := server.initAuthRoutes(); err != nil {
		return nil, err
	}

	if err := server.initAPIRoutes(); err != nil {
		return nil, err
	}

	server.registerDeveloperRoutes()

	// Response Console (root /) is served using one of the options below
	if server.isDevProxyEnabled() {
		server.registerDevProxyRoute()

		return server, nil
	}

	server.routes.Get("/*", server.handleConsoleRoutes())

	return server, nil
}

func (s *server) MonitoredService(ctx context.Context) (func() error, func(error)) {
	ctx, cancel := context.WithCancel(ctx)

	runFunc := func() error {
		return s.Serve(ctx)
	}

	intFunc := func(err error) {
		defer cancel()

		if err := s.Shutdown(ctx); err != nil {
			s.logger.Error().
				Err(err).
				Msg("failed to Shutdown server")
		}
	}

	return runFunc, intFunc
}
