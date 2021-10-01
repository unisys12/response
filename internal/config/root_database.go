package config

import (
	"fmt"
	"net"
	"net/url"
	"strconv"
)

type databaseBlock struct {
	Type     DatabaseType      `hcl:"type,optional"`
	Path     string            `hcl:"path,optional"`
	URL      string            `hcl:"url,optional"`
	Name     string            `hcl:"name,optional"`
	Host     string            `hcl:"host,optional"`
	Port     int               `hcl:"port,optional"`
	Username string            `hcl:"username,optional"`
	Password string            `hcl:"password,optional"`
	Options  map[string]string `hcl:"options,optional"`
}

func (b *databaseBlock) Apply(config *Config) error {
	str, err := url.Parse(b.URL)
	if err != nil {
		return fmt.Errorf("url.Parse: %w", err)
	}

	if b.Host == "" && str.Hostname() == "" {
		b.Host = "localhost"
	}

	if b.Name == "" && (str.Path == "" || str.Path == "/") {
		b.Name = "response"
	}

	// apply defaults based on the type
	switch b.Type {
	case DatabaseTypePostgres:
		if b.Port == 0 && str.Port() == "" {
			b.Port = 5432
		}

		if b.Username == "" && str.User == nil {
			b.Username = "postgres"
		}
	case DatabaseTypeMySQL:
		if b.Port == 0 && str.Port() == "" {
			b.Port = 3306
		}

		if b.Username == "" && str.User == nil {
			b.Username = "root"
		}
	}

	connUrl, err := b.MergeToConnString()
	if err != nil {
		return err
	}

	config.DatabaseConnURL = connUrl

	return nil
}

func (b *databaseBlock) MergeToConnString() (*url.URL, error) {
	str, err := url.Parse(b.URL)
	if err != nil {
		return nil, fmt.Errorf("unable to parse database url: %w", err)
	}

	switch b.Type {
	case DatabaseTypePostgres:
		str.Scheme = "postgresql"
	case DatabaseTypeMySQL:
		str.Scheme = "mysql"
	}

	if str.User == nil {
		str.User = url.UserPassword(b.Username, b.Password)
	}

	if b.Host != "" || b.Port != 0 {
		if b.Host != "" {
			str.Host = net.JoinHostPort(b.Host, str.Port())
		}

		if b.Port != 0 {
			str.Host = net.JoinHostPort(str.Hostname(), strconv.Itoa(b.Port))
		}
	}

	if b.Name != "" {
		str.Path = fmt.Sprintf("/%s", b.Name)
	}

	if len(b.Options) > 0 {
		values := url.Values{}
		for k, v := range b.Options {
			values.Set(k, v)
		}

		str.RawQuery = values.Encode()
	}

	if str.User == nil {
		str.User = url.UserPassword(b.Username, b.Password)

		return str, nil
	}

	if b.Username != "" {
		if password, set := str.User.Password(); set {
			str.User = url.UserPassword(b.Username, password)

			return str, nil
		}

		str.User = url.User(b.Username)
	}

	if b.Password != "" {
		str.User = url.UserPassword(str.User.Username(), b.Password)
	}

	return str, nil
}
