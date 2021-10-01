package server

import (
	"fmt"

	"github.com/go-chi/chi/v5"
)

func (s *server) initAuthRoutes() error {
	s.routes.Route("/auth", func(r chi.Router) {
		fmt.Println("/auth")

		s.auth.Register(r)
	})

	//group := s.http.Group("/auth", []echo.MiddlewareFunc{
	//	// Enable CSRF protection for all /auth routes to prevent request
	//	// forgery. This config requires the CSRF token in all forms as
	//	// _csrf.
	//	middleware.CSRFWithConfig(middleware.CSRFConfig{
	//		ContextKey:     "auth_csrf",
	//		TokenLookup:    "form:_csrf",
	//		CookieName:     "response_auth_csrf",
	//		CookiePath:     "/auth",
	//		CookieHTTPOnly: true,
	//	}),
	//}...)

	return nil
}
