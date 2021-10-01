package authold

import (
	"fmt"
	"net/http"

	"github.com/go-pkgz/auth"
	"github.com/go-pkgz/auth/avatar"
	"github.com/go-pkgz/auth/provider"
	"github.com/go-pkgz/auth/token"
	"github.com/labstack/echo/v4"
	"github.com/responserms/response/internal/config"
	"github.com/responserms/response/response"
)

const (
	tokenCookieName = "response_sess"
	tokenHeaderName = "X-Token"
	csrfCookieName  = "response_csrf"
	csrfHeaderName  = "X-CSRF-Token"
)

type Handler struct {
	core   *response.Core
	config *config.Config
	auth   *auth.Service
	oauth  *oauthManager
}

type oauthMapperData struct {
	User provider.UserData
}

func New(core *response.Core) (*Handler, error) {
	h := &Handler{
		core:   core,
		config: core.RunningConfig(),
	}

	if err := h.initAuthService(); err != nil {
		return nil, err
	}

	return h, nil
}

func (h *Handler) initAuthService() error {
	opts := auth.Opts{
		Logger:         &authLogger{logger: h.core.Logger("server.auth")},
		URL:            h.config.AuthSessionRootURL,
		JWTCookieName:  tokenCookieName,
		JWTHeaderKey:   tokenHeaderName,
		XSRFCookieName: csrfCookieName,
		XSRFHeaderKey:  csrfHeaderName,
		SecureCookies:  h.config.HTTPTLSPort != 0,
		AvatarStore:    avatar.NewLocalFS("./.store"),
		UseGravatar:    true,
		SecretReader: token.SecretFunc(func(aud string) (string, error) {
			return h.config.AuthSessionSecret, nil
		}),
	}

	h.auth = auth.NewService(opts)
	h.oauth = &oauthManager{
		core:           h.core,
		auth:           h.auth,
		config:         h.config,
		oauthProviders: map[string]*config.OAuth2Provider{},
	}

	if err := h.oauth.Configure(); err != nil {
		return fmt.Errorf("auth.Configure: %w", err)
	}

	if err := h.initPasswordAuthProvider(); err != nil {
		return fmt.Errorf("auth.initPasswordAuthProvider: %w", err)
	}

	return nil
}

func (h *Handler) WithRouteGroups(auth *echo.Group, avatars *echo.Group) {
	authHandler, avatarHandler := h.auth.Handlers()

	auth.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"password": !h.config.AuthPasswordDisabled,
			"oauth2":   h.oauth.AuthableOAuthProviders(c.Request().Context()),
		})
	})

	auth.Any("*", echo.WrapHandler(authHandler))
	avatars.Any("*", echo.WrapHandler(avatarHandler))
}

func (h *Handler) Middleware() echo.MiddlewareFunc {
	middleware := h.auth.Middleware()
	return echo.WrapMiddleware(middleware.Trace)
}
