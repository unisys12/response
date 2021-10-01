package auth

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/labstack/echo/v4"
	"github.com/responserms/response/internal/config"
	"github.com/responserms/response/response"
	"github.com/rs/zerolog"
	"github.com/volatiletech/authboss/v3"

	// authboss modules
	_ "github.com/volatiletech/authboss/v3/auth"
	_ "github.com/volatiletech/authboss/v3/oauth2"
	_ "github.com/volatiletech/authboss/v3/register"
)

type provider struct {
	Type string
	Name string
	Icon string
	URL  string
}

type Service struct {
	core      *response.Core
	config    *config.Config
	logger    zerolog.Logger
	boss      *authboss.Authboss
	providers []*provider
}

func New(core *response.Core) (*Service, error) {
	svc := &Service{
		core:      core,
		config:    core.RunningConfig(),
		logger:    core.Logger("server.auth"),
		boss:      authboss.New(),
		providers: []*provider{},
	}

	// Initialize Paths config in Authboss
	svc.initPathsConfig()

	// Initialize Modules config in Authboss
	svc.initModulesConfig()

	// Initialize Storage config in Authboss
	svc.initStorageConfig()

	// Initialize Core config in Authboss
	if err := svc.initCoreConfig(); err != nil {
		return nil, fmt.Errorf("svc.initCoreConfig: %w", err)
	}

	if err := svc.boss.Init(); err != nil {
		return nil, err
	}

	return svc, nil
}

func (s *Service) Register(g chi.Router) {
	g.Use(authboss.ModuleListMiddleware(s.boss))
	g.Mount("/", http.StripPrefix("/auth", s.boss.Config.Core.Router))
}

func (s *Service) RequireFullAuth() []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{
		echo.WrapMiddleware(authboss.Middleware(s.boss, true, true, true)),
	}
}

func (s *Service) RequireAuth() []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{
		echo.WrapMiddleware(authboss.Middleware(s.boss, true, false, true)),
	}
}

// Global provides middleware that should be added to all routes in Response.
func (s *Service) Global() chi.Middlewares {
	return chi.Middlewares{
		s.boss.LoadClientStateMiddleware,
	}
}
