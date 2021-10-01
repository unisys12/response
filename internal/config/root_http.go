package config

type httpBlock struct {
	BindAddress   string        `hcl:"bind_address,optional"`
	Port          int           `hcl:"port,optional"`
	MaxUploadSize int           `hcl:"max_upload_size,optional"`
	TLS           *httpTLSBlock `hcl:"tls,block"`
}

func (b *httpBlock) Apply(config *Config) error {
	if b.BindAddress == "" {
		b.BindAddress = "127.0.0.1"
	}

	if b.Port == 0 {
		b.Port = 9080
	}

	if b.MaxUploadSize == 0 {
		b.MaxUploadSize = 10000
	}

	config.HTTPBindAddress = b.BindAddress
	config.HTTPPort = b.Port
	config.HTTPMaxUploadSize = b.MaxUploadSize

	if b.TLS != nil {
		if err := b.TLS.Apply(config); err != nil {
			return err
		}
	}

	return nil
}

type httpTLSBlock struct {
	Port     int               `hcl:"port,optional"`
	CertPath string            `hcl:"cert_path,optional"`
	KeyPath  string            `hcl:"key_path,optional"`
	Auto     *httpTLSAutoBlock `hcl:"auto,block"`
}

func (b *httpTLSBlock) Apply(config *Config) error {
	if b.Port == 0 && ((b.CertPath != "" && b.KeyPath != "") || (b.Auto != nil)) {
		b.Port = 9443
	}

	if b.Auto != nil {
		if err := b.Auto.Apply(config); err != nil {
			return err
		}
	}

	config.HTTPTLSPort = b.Port
	config.HTTPTLSCertPath = b.CertPath
	config.HTTPTLSCertKey = b.KeyPath

	return nil
}

type httpTLSAutoBlock struct {
	Production bool                 `hcl:"production,optional"`
	Email      string               `hcl:"email,attr"`
	Domains    []string             `hcl:"domains,attr"`
	DNS        *httpTLSAutoDNSBlock `hcl:"dns,block"`
}

func (b *httpTLSAutoBlock) Apply(config *Config) error {
	config.HTTPTLSAutoProduction = b.Production
	config.HTTPTLSAutoEmail = b.Email
	config.HTTPTLSAutoDomains = b.Domains

	if b.DNS != nil {
		if err := b.DNS.Apply(config); err != nil {
			return err
		}
	}

	return nil
}

type httpTLSAutoDNSBlock struct {
	Provider string `hcl:"solver,label"`
	APIToken string `hcl:"token,attr"`
}

func (b *httpTLSAutoDNSBlock) Apply(config *Config) error {
	config.HTTPTLSAutoDNSEnabled = true
	config.HTTPTLSAutoDNSProvider = b.Provider
	config.HTTPTLSAutoDNSAPIToken = b.APIToken

	return nil
}
