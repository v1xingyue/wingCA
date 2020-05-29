package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

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

var (
	//Item 配置实例
	Item = configT{}
)

// InitConfigFile 使用配置文件初始化配置
func InitConfigFile(configFilePath string) {
	if f, err := os.Open(configFilePath); err == nil {
		yaml.NewDecoder(f).Decode(&Item)
	} else {
		Item.Wing.Client.DefaultLifeTimeSeconds = 86400 * 30
		Item.Wing.Site.DefaultLifeTimeSeconds = 86400 * 90
		Item.Wing.RootCAPath = "./ssl"
	}
}
