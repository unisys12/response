package config

import (
	"net/url"
	"text/template"
	"time"

	"github.com/rs/zerolog"

	"golang.org/x/oauth2"
)

// Config represents all the Config needed for Response to function. This is the struct used by
// the application to configure internal services.
type Config struct {
	valid bool

	// General
	EncryptionKey    []byte
	UIDirConsoleDist string
	UIDirAuthPages   string
	LogLevel         zerolog.Level

	// Terminal Output
	HideBanner  bool
	HideServing bool

	// Database
	DatabaseConnURL *url.URL

	// HTTP
	HTTPBindAddress   string
	HTTPPort          int
	HTTPMaxUploadSize int

	// HTTP TLS
	HTTPTLSPort     int
	HTTPTLSCertPath string
	HTTPTLSCertKey  string

	// HTTP Automatic TLS (LetsEncrypt)
	HTTPTLSAutoProduction  bool
	HTTPTLSAutoEmail       string
	HTTPTLSAutoDomains     []string
	HTTPTLSAutoDNSEnabled  bool
	HTTPTLSAutoDNSProvider string
	HTTPTLSAutoDNSAPIToken string

	// RESP
	RESPBindAddress string
	RESPPort        int

	// RESP TLS
	RESPTLSPort     int
	RESPTLSCertPath string
	RESPTLSCertKey  string

	// RESP Automatic TLS (LetsEncrypt)
	RESPTLSAutoProduction  bool
	RESPTLSAutoEmail       string
	RESPTLSAutoDomains     []string
	RESPTLSAutoDNSEnabled  bool
	RESPTLSAutoDNSProvider string
	RESPTLSAutoDNSAPIToken string

	// High Availability
	NATSConnURL string

	// Developer
	DeveloperProfiling               bool
	DeveloperUIProxyPort             int
	DeveloperDisableCachingAuthPages bool

	// Auth
	AuthSessionRootURL  string
	AuthSessionDomain   string
	AuthSessionDuration time.Duration
	AuthSessionSecret   []byte

	// Password Auth
	AuthPasswordDisabled   bool
	AuthPasswordBcryptCost int

	// OAuth 2.0 Auth
	AuthOAuth2Providers []*OAuth2Provider
}

func (c *Config) Valid() bool {
	return c.valid
}

type OAuth2Provider struct {
	Name                 string
	Icon                 string
	Provider             string
	UpdateProfileOnLogin bool
	Config               *oauth2.Config
	UserinfoURL          string
	UserinfoExtraHeaders map[string]*template.Template
	UserinfoMapper       *UserinfoMapper
}

type UserinfoExtraHeaders map[string]*template.Template

type UserinfoMapper struct {
	ID        *template.Template
	Name      *template.Template
	Email     *template.Template
	AvatarURL *template.Template
	Metadata  map[string]*template.Template
}
