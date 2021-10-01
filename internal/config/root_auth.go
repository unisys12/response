package config

import (
	"encoding/base64"
	"fmt"
	"text/template"
	"time"

	"github.com/Masterminds/sprig"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
)

type authBlock struct {
	SessionRootURL  string        `hcl:"session_root_url,optional"`
	SessionDomain   string        `hcl:"session_domain,optional"`
	SessionDuration string        `hcl:"session_duration,optional"`
	SessionSecret   string        `hcl:"session_secret,optional"`
	Password        *authPassword `hcl:"password,block"`
	OAuth2          []*authOAuth2 `hcl:"oauth2,block"`
}

func (b *authBlock) Apply(config *Config) error {
	if b.SessionDuration == "" {
		b.SessionDuration = "8h"
	}

	dur, err := time.ParseDuration(b.SessionDuration)
	if err != nil {
		return err
	}

	config.AuthSessionDuration = dur
	config.AuthSessionDomain = b.SessionDomain
	config.AuthSessionRootURL = b.SessionRootURL

	if b.SessionSecret == "" {
		sessionSecret, err := GenerateSessionSecret()
		if err != nil {
			return fmt.Errorf("SessionSecret GenerateEncryptionKey: %w", err)
		}

		b.SessionSecret = sessionSecret
	}

	sessionSecret, err := base64.StdEncoding.DecodeString(b.SessionSecret)
	if err != nil {
		return fmt.Errorf("base64...DecodeString(session_secret): %w", err)
	}

	config.AuthSessionSecret = sessionSecret

	for _, oauth := range b.OAuth2 {
		if err := oauth.Apply(config); err != nil {
			return fmt.Errorf("oauth.Apply for %q: %w", oauth.Provider, err)
		}
	}

	if b.Password == nil {
		b.Password = &authPassword{}
	}

	if err := b.Password.Apply(config); err != nil {
		return err
	}

	return nil
}

type authPassword struct {
	Disabled   bool `hcl:"disabled,optional"`
	BcryptCost int  `hcl:"bcrypt_cost,optional"`
}

func (b *authPassword) Apply(config *Config) error {
	config.AuthPasswordDisabled = b.Disabled

	if b.BcryptCost == 0 {
		b.BcryptCost = bcrypt.DefaultCost
	}

	config.AuthPasswordBcryptCost = b.BcryptCost

	return nil
}

type authOAuth2 struct {
	Provider             string                  `hcl:"provider,label"`
	Name                 string                  `hcl:"name,optional"`
	Icon                 string                  `hcl:"icon,optional"`
	UpdateProfileOnLogin bool                    `hcl:"update_profile_on_login,optional"`
	ClientID             string                  `hcl:"client_id,attr"`
	ClientSecret         string                  `hcl:"client_secret,attr"`
	Scopes               []string                `hcl:"scopes,optional"`
	AuthorizationURL     string                  `hcl:"authorization_url,optional"`
	TokenURL             string                  `hcl:"token_url,optional"`
	UserinfoURL          string                  `hcl:"userinfo_url,optional"`
	UserinfoMap          *authOAuth2Map          `hcl:"userinfo_map,block"`
	UserinfoExtraHeaders *authOAuth2ExtraHeaders `hcl:"userinfo_extra_headers,attr"`
}

func (b *authOAuth2) Apply(config *Config) error {
	oauth, err := b.getKnownOAuthProvider(b.Provider)
	if err != nil {
		return fmt.Errorf("getKnownOAuthProvider: %w", err)
	}

	oauth.Config.ClientID = b.ClientID
	oauth.Config.ClientSecret = b.ClientSecret

	oauth.UpdateProfileOnLogin = b.UpdateProfileOnLogin

	if len(b.Scopes) > 0 {
		oauth.Config.Scopes = append(oauth.Config.Scopes, b.Scopes...)
	}

	if b.Name != "" {
		oauth.Name = b.Name
	}

	if b.Icon != "" {
		oauth.Icon = b.Icon
	}

	if b.AuthorizationURL != "" {
		oauth.Config.Endpoint.AuthStyle = oauth2.AuthStyleAutoDetect
		oauth.Config.Endpoint.AuthURL = b.AuthorizationURL
	}

	if b.TokenURL != "" {
		oauth.Config.Endpoint.TokenURL = b.TokenURL
	}

	if b.UserinfoURL != "" {
		oauth.UserinfoURL = b.UserinfoURL
	}

	if b.UserinfoMap != nil {
		mapper, err := b.UserinfoMap.MakeTemplates()
		if err != nil {
			return fmt.Errorf("MakeTemplates: %w", err)
		}

		oauth.UserinfoMapper = mapper
	}

	if b.UserinfoExtraHeaders != nil {
		headers, err := b.UserinfoExtraHeaders.MakeTemplates()
		if err != nil {
			return fmt.Errorf("MakeTemplates: %w", err)
		}

		oauth.UserinfoExtraHeaders = headers
	}

	config.AuthOAuth2Providers = append(config.AuthOAuth2Providers, oauth)

	return nil
}

func (b *authOAuth2) getKnownOAuthProvider(key string) (*OAuth2Provider, error) {
	var provider *OAuth2Provider

	switch key {
	case "discord":
		maps := &authOAuth2Map{
			ID:        "{{.User.id}}",
			Name:      "{{.User.username}}",
			Email:     "{{.User.email}}",
			AvatarURL: `{{if .User.avatar}}https://cdn.discordapp.com/avatars/{{.User.id}}/{{.User.avatar}}{{if hasPrefix "a_" .User.avatar}}.gif{{else}}.png{{end}}{{end}}`,
		}

		mapper, err := maps.MakeTemplates()
		if err != nil {
			return nil, fmt.Errorf("MakeTemplates: %w", err)
		}

		provider = &OAuth2Provider{
			Name:     "Discord",
			Icon:     "discord",
			Provider: key,
			Config: &oauth2.Config{
				Scopes: []string{"identify", "email"},
				Endpoint: oauth2.Endpoint{
					AuthStyle: oauth2.AuthStyleAutoDetect,
					AuthURL:   "https://discord.com/api/v9/oauth2/authorize",
					TokenURL:  "https://discord.com/api/v9/oauth2/token",
				},
			},
			UserinfoURL:    "https://discord.com/api/v9/users/@me",
			UserinfoMapper: mapper,
		}
	case "twitch":
		maps := &authOAuth2Map{
			ID:        `{{ get .User "data.0.id" }}`,
			Name:      `{{ get .User "data.0.display_name" }}`,
			Email:     `{{ get .User "data.0.email" }}`,
			AvatarURL: `{{ get .User "data.0.profile_image_url" }}`,
			Metadata: map[string]string{
				"twitch_login": `{{ get .User "data.0.login" }}`,
			},
		}

		mapper, err := maps.MakeTemplates()
		if err != nil {
			return nil, fmt.Errorf("MakeTemplates: %w", err)
		}

		extraHeaders := &authOAuth2ExtraHeaders{
			"Client-Id": `{{ .ClientID }}`,
		}

		extraHeaderTemplates, err := extraHeaders.MakeTemplates()
		if err != nil {
			return nil, fmt.Errorf("extraHeaders.MakeTemplates: %w", err)
		}

		provider = &OAuth2Provider{
			Name:     "Twitch",
			Icon:     "twitch",
			Provider: key,
			Config: &oauth2.Config{
				Scopes: []string{"user:read:email"},
				Endpoint: oauth2.Endpoint{
					AuthStyle: oauth2.AuthStyleAutoDetect,
					AuthURL:   "https://id.twitch.tv/oauth2/authorize",
					TokenURL:  "https://id.twitch.tv/oauth2/token",
				},
			},
			UserinfoURL:          "https://api.twitch.tv/helix/users",
			UserinfoExtraHeaders: extraHeaderTemplates,
			UserinfoMapper:       mapper,
		}
	}

	// Initialize a blank provider here so that we can add on to it
	// elsewhere.
	if provider == nil {
		provider = &OAuth2Provider{
			Provider: key,
			Config: &oauth2.Config{
				Endpoint: oauth2.Endpoint{},
			},
		}
	}

	return provider, nil
}

type authOAuth2ExtraHeaders map[string]string

func (b authOAuth2ExtraHeaders) MakeTemplates() (UserinfoExtraHeaders, error) {
	headers := UserinfoExtraHeaders{}

	for prop, rawTemplate := range b {
		prepared, err := makeTemplate(prop, rawTemplate)
		if err != nil {
			return nil, fmt.Errorf("makeTemplate(userinfo_extra_headers/%s): %w", prop, err)
		}

		headers[prop] = prepared
	}

	return headers, nil
}

type authOAuth2Map struct {
	ID        string            `hcl:"id,attr"`
	Name      string            `hcl:"name,attr"`
	Email     string            `hcl:"email,attr"`
	AvatarURL string            `hcl:"avatar_url,attr"`
	Metadata  map[string]string `hcl:"metadata,optional"`
}

func (b *authOAuth2Map) MakeTemplates() (*UserinfoMapper, error) {
	mapper := &UserinfoMapper{}

	id, err := b.makeTemplate("id", b.ID)
	if err != nil {
		return nil, fmt.Errorf("makeTemplate(id): %w", err)
	}

	mapper.ID = id

	name, err := b.makeTemplate("name", b.Name)
	if err != nil {
		return nil, fmt.Errorf("makeTemplate(name): %w", err)
	}

	mapper.Name = name

	email, err := b.makeTemplate("email", b.Email)
	if err != nil {
		return nil, fmt.Errorf("makeTemplate(email): %w", err)
	}

	mapper.Email = email

	avatarURL, err := b.makeTemplate("avatar_url", b.AvatarURL)
	if err != nil {
		return nil, fmt.Errorf("makeTemplate(avatar_url): %w", err)
	}

	mapper.AvatarURL = avatarURL

	mapper.Metadata = map[string]*template.Template{}
	for prop, rawTemplate := range b.Metadata {
		prepared, err := b.makeTemplate(prop, rawTemplate)
		if err != nil {
			return nil, fmt.Errorf("makeTemplate(metadata/%s): %w", prop, err)
		}

		mapper.Metadata[prop] = prepared
	}

	return mapper, nil
}

func (b *authOAuth2Map) makeTemplate(name string, raw string) (*template.Template, error) {
	return makeTemplate(name, raw)
}

func makeTemplate(name, raw string) (*template.Template, error) {
	funcs := template.FuncMap{
		"get": func(data map[string]string, key string) string {
			return data[key]
		},
	}

	return template.New(name).
		Funcs(funcs).
		Funcs(sprig.TxtFuncMap()).
		Parse(raw)
}
