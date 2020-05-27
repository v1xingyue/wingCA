package rootCA

import (
	"os"
	"time"
)

var (
	// RootCAPath 根证书文件夹路径
	RootCAPath     = "./ssl"
	rootCAKeyPath  = RootCAPath + "/private/rootCA.key"
	rootCACertPath = RootCAPath + "/root/rootCA.crt"

	rootCAKeyPassword         = ""
	serialFile                = RootCAPath + "/serial"
	revokeListPath            = RootCAPath + "/revokelist"
	crlLifetime               = 3600 * time.Second
	defaultCertLifetime       = time.Hour * 24 * 90
	defaultClientCertLifetime = time.Hour * 24 * 14

	// RootCACertPath 根证书路径
	RootCACertPath = rootCACertPath

	newFileMode = os.FileMode(0755)
)
