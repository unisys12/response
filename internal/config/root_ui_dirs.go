package config

type uiDirsBlock struct {
	ConsoleDist string `hcl:"console_dist,optional"`
	AuthPages   string `hcl:"auth_pages,optional"`
}

func (b *uiDirsBlock) Apply(config *Config) error {
	config.UIDirConsoleDist = b.ConsoleDist
	config.UIDirAuthPages = b.AuthPages

	return nil
}
