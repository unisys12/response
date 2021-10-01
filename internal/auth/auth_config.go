package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/volatiletech/authboss/v3/defaults"

	abclientstate "github.com/volatiletech/authboss-clientstate"
	"github.com/volatiletech/authboss/v3"
)

func (s *Service) initPathsConfig() {
	s.boss.Config.Paths.Mount = "/auth"
	s.boss.Config.Paths.RootURL = s.config.AuthSessionRootURL
}

func (s *Service) initModulesConfig() {
	providers := map[string]authboss.OAuth2Provider{}
	for _, oAuth2Provider := range s.config.AuthOAuth2Providers {
		s.providers = append(s.providers, &provider{
			Type: "oauth2",
			Name: oAuth2Provider.Name,
			Icon: oAuth2Provider.Icon,
			URL:  fmt.Sprintf("/auth/oauth2/%s", oAuth2Provider.Provider),
		})

		providers[oAuth2Provider.Provider] = authboss.OAuth2Provider{
			OAuth2Config:    oAuth2Provider.Config,
			FindUserDetails: handleOAuth2UserDetails(s.logger, *oAuth2Provider),
		}
	}

	s.boss.Config.Modules.BCryptCost = s.config.AuthPasswordBcryptCost
	s.boss.Config.Modules.ExpireAfter = 1 * time.Hour
	s.boss.Config.Modules.RegisterPreserveFields = []string{"name", "email"}
	s.boss.Config.Modules.RecoverTokenDuration = 1 * time.Hour
	s.boss.Config.Modules.RecoverLoginAfterRecovery = true
	s.boss.Config.Modules.OAuth2Providers = providers
	s.boss.Config.Modules.TOTP2FAIssuer = "Response"
	s.boss.Config.Modules.ResponseOnUnauthed = authboss.RespondRedirect
}

func (s *Service) initStorageConfig() {
	s.boss.Config.Storage.Server = newServerStorer(s.core)
	s.boss.Config.Storage.CookieState = abclientstate.NewCookieStorer(s.config.AuthSessionSecret, s.config.EncryptionKey)
	s.boss.Config.Storage.SessionState = abclientstate.NewSessionStorer("response_sess", s.config.AuthSessionSecret, s.config.EncryptionKey)
	s.boss.Config.Storage.SessionStateWhitelistKeys = []string{}
}

func (s *Service) initCoreConfig() error {
	viewRenderer, err := s.newAuthPagesRenderer()
	if err != nil {
		return fmt.Errorf("newAuthPagesRenderer: %w", err)
	}

	s.boss.Config.Core.Router = defaults.NewRouter()
	s.boss.Config.Core.ErrorHandler = defaults.NewErrorHandler(defaults.NewLogger(os.Stderr))
	s.boss.Config.Core.Responder = defaults.NewResponder(viewRenderer)
	s.boss.Config.Core.Redirector = defaults.NewRedirector(viewRenderer, "redirect")
	s.boss.Config.Core.BodyReader = defaults.NewHTTPBodyReader(false, false)
	s.boss.Config.Core.ViewRenderer = viewRenderer
	s.boss.Config.Core.Logger = defaults.NewLogger(os.Stderr)

	return nil
}
