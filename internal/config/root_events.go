package config

type eventsBlock struct {
	Type string `hcl:"type,optional"`
}

func (b *eventsBlock) Apply(config *Config) error {
	config.NATSConnURL = ""

	return nil
}
