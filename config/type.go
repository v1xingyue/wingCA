package config

type configT struct {
	Wing wingConfigT `yaml:"Wing"`
}

type wingConfigT struct {
	RootCAPath string        `yaml:"RootCAPath"`
	Client     clientConfigT `yaml:"Client"`
	Site       siteConfigT   `yaml:"Site"`
}

type clientConfigT struct {
	DefaultLifeTimeSeconds int64 `yaml:"DefaultLifeTimeSeconds"`
}

type siteConfigT struct {
	DefaultLifeTimeSeconds int64 `yaml:"DefaultLifeTimeSeconds"`
}
