package config

type configT struct {
	Wing        wingConfigT   `yaml:"Wing"`
	BasicConfig basicCconfigT `yaml:"BasicConfig"`
	SMTPConfig  smtpConfigT   `yaml:"SMTPConfig"`
}

type basicCconfigT struct {
	BindHost string `yaml:"BindHost"`
	BindPort int    `yaml:"BindPort"`
}

type wingConfigT struct {
	RootCAPath string        `yaml:"RootCAPath"`
	Domain     string        `yaml:"Domain"`
	Client     clientConfigT `yaml:"Client"`
	Site       siteConfigT   `yaml:"Site"`
}

type clientConfigT struct {
	DefaultLifeTimeSeconds int64 `yaml:"DefaultLifeTimeSeconds"`
}

type siteConfigT struct {
	DefaultLifeTimeSeconds int64 `yaml:"DefaultLifeTimeSeconds"`
}

type smtpConfigT struct {
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
	Email    string `yaml:"Email"`
	Password string `yaml:"Password"`
	Nick     string `yaml:"Nick"`
}
