package config

import (
	"encoding/base64"
	"fmt"

	"github.com/rs/zerolog"

	"github.com/hashicorp/hcl/v2"
)

type DatabaseType = string

const (
	DatabaseTypePostgres DatabaseType = "postgres"
	DatabaseTypeMySQL    DatabaseType = "mysql"
)

// root is the root of the Response declarative configuration files. This is converted into a
// Config to simplify usage in the Response application. This struct maps to the same declarative
// structure in the Response config files.
type root struct {
	EncryptionKey string `hcl:"encryption_key"`
	HideBanner    bool   `hcl:"hide_banner,optional"`
	HideServing   bool   `hcl:"hide_serving,optional"`
	//LogLevel      log.Level
	UIDirs       *uiDirsBlock    `hcl:"ui_dirs,block"`
	Database     *databaseBlock  `hcl:"database,block"`
	HTTPListener *httpBlock      `hcl:"http_listener,block"`
	GameListener *gameBlock      `hcl:"game_listener,block"`
	Events       *eventsBlock    `hcl:"events,block"`
	Developer    *developerBlock `hcl:"developer,block"`
	Auth         *authBlock      `hcl:"auth,block"`
	Remaining    hcl.Body        `hcl:",remain"`
}

func newRoot() *root {
	return &root{
		UIDirs:       &uiDirsBlock{},
		Database:     &databaseBlock{},
		HTTPListener: &httpBlock{},
		GameListener: &gameBlock{},
		Events:       &eventsBlock{},
		Developer:    &developerBlock{},
		Auth:         &authBlock{},
	}
}

func (r *root) ToConfig() (*Config, error) {
	encryptionKey, err := base64.StdEncoding.DecodeString(r.EncryptionKey)
	if err != nil {
		return nil, fmt.Errorf("base64...DecodeString(encryption_key): %w", err)
	}

	config := &Config{
		valid:         true,
		EncryptionKey: encryptionKey,
		HideBanner:    r.HideBanner,
		HideServing:   r.HideServing,
		LogLevel:      zerolog.InfoLevel,
	}

	if err := r.UIDirs.Apply(config); err != nil {
		return nil, err
	}

	if err := r.Database.Apply(config); err != nil {
		return nil, err
	}

	if err := r.HTTPListener.Apply(config); err != nil {
		return nil, err
	}

	if err := r.GameListener.Apply(config); err != nil {
		return nil, err
	}

	if err := r.Events.Apply(config); err != nil {
		return nil, err
	}

	if err := r.Developer.Apply(config); err != nil {
		return nil, err
	}

	if err := r.Auth.Apply(config); err != nil {
		return nil, err
	}

	return config, nil
}
