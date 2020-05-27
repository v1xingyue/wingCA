package cmd

import "net"

var (
	confirmInitCA bool
	// 证书CA 需要的参数
	name, org, province, locality, street, postcode string

	// 签发证书需要的参数
	siteNames  []string
	siteIPStr  []string
	email      string
	commonName string
	issueType  string
	siteIPs    []net.IP
	withP12    bool
	password   string

	// 启动示例站点需要的参数
	startDouble bool
)
