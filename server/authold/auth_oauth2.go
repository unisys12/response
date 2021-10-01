package authold

import (
	"bytes"
	"context"
	"fmt"
	"strconv"
	"sync"

	"github.com/go-pkgz/auth"
	"github.com/go-pkgz/auth/provider"
	"github.com/go-pkgz/auth/token"
	"github.com/responserms/response/internal/config"
	"github.com/responserms/response/internal/ptr"
	"github.com/responserms/response/response"
	"github.com/rs/zerolog"
)

type Provider struct {
	Key  string `json:"key"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type oauthManager struct {
	logger             zerolog.Logger
	core               *response.Core
	auth               *auth.Service
	config             *config.Config
	oauthProviders     map[string]*config.OAuth2Provider
	oauthProvidersLock sync.RWMutex
}

func (m *oauthManager) Configure() error {
	m.oauthProvidersLock.Lock()
	defer m.oauthProvidersLock.Unlock()

	m.logger = m.core.Logger("server.auth.oauth2")

	for _, oauth := range m.config.AuthOAuth2Providers {
		m.oauthProviders[oauth.Provider] = oauth

		if oauth.Client == nil {
			return fmt.Errorf("oauth.Client is nil for %q", oauth.Provider)
		}

		if oauth.Endpoint == nil {
			return fmt.Errorf("oauth.Endpoint is nil for %q", oauth.Provider)
		}

		m.auth.AddCustomProvider(oauth.Provider, *oauth.Client, provider.CustomHandlerOpt{
			Endpoint:  *oauth.Endpoint,
			Scopes:    oauth.Scopes,
			InfoURL:   oauth.UserinfoURL,
			MapUserFn: m.userDataMapper(oauth.Provider),
		})
	}

	return nil
}

func (m *oauthManager) userDataMapper(name string) func(data provider.UserData, bytes []byte) token.User {
	return func(data provider.UserData, bytes []byte) token.User {
		var ctx = m.core.ApplyContext(context.Background())

		entry, err := m.makeOAuthConnectionEntry(name, data)
		if err != nil {
			m.logger.Error().
				Err(err).
				Str("provider", name).
				Msg("Unable to make an OAuthConnectionEntry using the provided UserData.")

			return token.User{}
		}

		if entry == nil {
			m.logger.Error().
				Str("provider", name).
				Msg("OAuthConnectionEntry is nil.")

			return token.User{}
		}

		conn, err := m.core.EnsureOAuthConnection(ctx, *entry)
		if err != nil {
			m.logger.Error().
				Err(err).
				Str("provider", name).
				Msg("Unable to ensure OAuthConnection.")

			return token.User{}
		}

		user, err := conn.User(ctx)
		if err != nil {
			m.logger.Error().
				Err(err).
				Str("provider", name).
				Msg("Unable to load User from OAuthConnection.")

			return token.User{}
		}

		return token.User{
			ID:      strconv.Itoa(user.ID),
			Name:    user.Name,
			Email:   user.Email,
			Picture: ptr.String(user.AvatarURL),
		}
	}
}

func (m *oauthManager) AuthableOAuthProviders(ctx context.Context) []*Provider {
	m.oauthProvidersLock.RLock()
	defer m.oauthProvidersLock.RUnlock()

	var authable []*Provider
	for _, provider := range m.oauthProviders {
		authable = append(authable, &Provider{
			Key:  provider.Provider,
			Name: provider.Provider,
			Icon: "fab-discord",
		})
	}

	return authable
}

func (m *oauthManager) makeOAuthConnectionEntry(provider string, data provider.UserData) (*response.OAuthConnectionEntry, error) {
	m.oauthProvidersLock.RLock()
	defer m.oauthProvidersLock.RUnlock()

	oauthProvider, ok := m.oauthProviders[provider]
	if !ok {
		return nil, fmt.Errorf("%q provider has no registered OAuth2Provider", provider)
	}

	if oauthProvider == nil {
		return nil, fmt.Errorf("%q provider has a nil OAuth2Provider", provider)
	}

	mapTemplates := oauthProvider.UserinfoMapper
	if mapTemplates == nil {
		return nil, fmt.Errorf("%q provider has no registered UserinfoMapper", provider)
	}

	vars := &oauthMapperData{
		User: data,
	}

	var providerUserID bytes.Buffer
	if err := mapTemplates.ID.ExecuteTemplate(&providerUserID, "id", vars); err != nil {
		return nil, fmt.Errorf("ID.ExecuteTemplate: %w", err)
	}

	var name bytes.Buffer
	if err := mapTemplates.Name.ExecuteTemplate(&name, "name", vars); err != nil {
		return nil, fmt.Errorf("Name.ExecuteTemplate: %w", err)
	}

	var email bytes.Buffer
	if err := mapTemplates.Email.ExecuteTemplate(&email, "email", vars); err != nil {
		return nil, fmt.Errorf("Email.ExecuteTemplate: %w", err)
	}

	var avatarURL bytes.Buffer
	if err := mapTemplates.AvatarURL.ExecuteTemplate(&avatarURL, "avatar_url", vars); err != nil {
		return nil, fmt.Errorf("AvatarURL.ExecuteTemplate: %w", err)
	}

	var metadata map[string]string
	for prop, tmpl := range mapTemplates.Metadata {
		var buf bytes.Buffer
		if err := tmpl.ExecuteTemplate(&buf, prop, vars); err != nil {
			return nil, fmt.Errorf("Metadata[%s].ExecuteTemplate: %w", prop, err)
		}

		if metadata == nil {
			metadata = map[string]string{}
		}

		metadata[prop] = buf.String()
	}

	return &response.OAuthConnectionEntry{
		ProviderKey:    provider,
		ProviderUserID: providerUserID.String(),
		UpdateProfile:  oauthProvider.UpdateProfileOnLogin,
		Name:           name.String(),
		Email:          email.String(),
		AvatarURL:      avatarURL.String(),
		Metadata:       metadata,
	}, nil
}
