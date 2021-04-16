package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

var (
	//Item 配置实例
	Item = &configT{}

	// Version 软件版本
	Version = "0.0.5"

	BuildTime = ""
)

// InitConfigFile 使用配置文件初始化配置
func InitConfigFile(configFilePath string) {
	if f, err := os.Open(configFilePath); err == nil {
		yaml.NewDecoder(f).Decode(Item)
	} else {
		Item.Wing.Client.DefaultLifeTimeSeconds = 86400 * 30
		Item.Wing.Site.DefaultLifeTimeSeconds = 86400 * 90
		Item.Wing.RootCAPath = "./ssl"
	}

	itemToDefault(Item)
}

// 根据item 的状态更新 Default 的内容
func itemToDefault(item *configT) {

}
