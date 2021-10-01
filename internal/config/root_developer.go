package config

type developerBlock struct {
	Profiling               bool `hcl:"profiling,optional"`
	UIProxyPort             int  `hcl:"ui_proxy_port,optional"`
	DisableCachingAuthPages bool `hcl:"disable_caching_auth_pages,optional"`
}

func (b *developerBlock) Apply(config *Config) error {
	config.DeveloperProfiling = b.Profiling
	config.DeveloperUIProxyPort = b.UIProxyPort
	config.DeveloperDisableCachingAuthPages = b.DisableCachingAuthPages

	return nil
}
