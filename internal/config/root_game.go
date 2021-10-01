package config

type gameBlock struct {
	BindAddress string        `hcl:"bind_address,optional"`
	Port        int           `hcl:"port,optional"`
	TLS         *gameTLSBlock `hcl:"tls,block"`
}

func (b *gameBlock) Apply(config *Config) error {
	if b.BindAddress == "" {
		b.BindAddress = "127.0.0.1"
	}

	if b.Port == 0 {
		b.Port = 9000
	}

	config.RESPBindAddress = b.BindAddress
	config.RESPPort = b.Port

	if b.TLS != nil {
		if err := b.TLS.Apply(config); err != nil {
			return err
		}
	}

	return nil
}

type gameTLSBlock struct {
	Port     int               `hcl:"port,optional"`
	CertPath string            `hcl:"cert_path,optional"`
	KeyPath  string            `hcl:"key_path,optional"`
	Auto     *gameTLSAutoBlock `hcl:"auto,block"`
}

func (b *gameTLSBlock) Apply(config *Config) error {
	if b.Port == 0 && ((b.CertPath != "" && b.KeyPath != "") || (b.Auto != nil)) {
		b.Port = 9443
	}

	if b.Auto != nil {
		if err := b.Auto.Apply(config); err != nil {
			return err
		}
	}

	config.RESPTLSPort = b.Port
	config.RESPTLSCertPath = b.CertPath
	config.RESPTLSCertKey = b.KeyPath

	return nil
}

type gameTLSAutoBlock struct {
	Production bool                 `hcl:"production,optional"`
	Email      string               `hcl:"email,attr"`
	Domains    []string             `hcl:"domains,attr"`
	DNS        *gameTLSAutoDNSBlock `hcl:"dns,block"`
}

func (b *gameTLSAutoBlock) Apply(config *Config) error {
	config.RESPTLSAutoProduction = b.Production
	config.RESPTLSAutoEmail = b.Email
	config.RESPTLSAutoDomains = b.Domains

	if b.DNS != nil {
		if err := b.DNS.Apply(config); err != nil {
			return err
		}
	}

	return nil
}

type gameTLSAutoDNSBlock struct {
	Provider string `hcl:"solver,label"`
	APIToken string `hcl:"token,attr"`
}

func (b *gameTLSAutoDNSBlock) Apply(config *Config) error {
	config.RESPTLSAutoDNSEnabled = true
	config.RESPTLSAutoDNSProvider = b.Provider
	config.RESPTLSAutoDNSAPIToken = b.APIToken

	return nil
}
