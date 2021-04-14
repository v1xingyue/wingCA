package rootCA

import (
	"os"
	"time"
	"wingCA/config"
)

var (
	// RootCAPath 根证书文件夹路径
	RootCAPath     = "./ssl"
	rootCAKeyPath  = RootCAPath + "/private/rootCA.key"
	rootCACertPath = RootCAPath + "/root/rootCA.crt"

	serialFile                = RootCAPath + "/serial"
	revokeListPath            = RootCAPath + "/revokelist"
	crlLifetime               = 3600 * time.Second
	defaultCertLifetime       = time.Second * 86400 * 90
	defaultClientCertLifetime = time.Second * 86400 * 30

	// RootCACertPath 根证书路径
	RootCACertPath = rootCACertPath

	newFileMode = os.FileMode(0755)
)

/**
证书类型的枚举类型
*/
const (
	CertTypeSite = iota
	CertTypeClient
)

// InitConfigParamas 根据配置文件初始化参数
func InitConfigParamas() {
	RootCAPath = config.Item.Wing.RootCAPath
	defaultCertLifetime = time.Second * time.Duration(config.Item.Wing.Site.DefaultLifeTimeSeconds)
	defaultClientCertLifetime = time.Second * time.Duration(config.Item.Wing.Client.DefaultLifeTimeSeconds)
}
