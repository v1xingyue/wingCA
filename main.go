package main

import (
	"crypto/x509/pkix"
	"io/ioutil"
	"log"
	"net"
	"wingCA/rootCA"
)

func main() {
	name := pkix.Name{
		Organization:  []string{"SomeBody CA"},
		Country:       []string{"CN"},
		Province:      []string{"Beijing"},
		Locality:      []string{"Haidian"},
		StreetAddress: []string{"NoWhere Road"},
		PostalCode:    []string{"100093"},
		CommonName:    "SomeBodySuperCA",
	}
	rootCA.InitRootCA(name)

	ca, _ := rootCA.ParseCertificate("ssl/rootCA.pem")
	log.Println(ca.IsCA)

	cert, key, err := rootCA.IssueOneCert("debug.ssl.com.cn", []net.IP{}, []string{"localhost"})
	if err == nil {
		ioutil.WriteFile("ssl/server.cert", cert, 0700)
		ioutil.WriteFile("ssl/server.key", key, 0700)
	}

	rootCA.SampleWeb("ssl/server.cert", "ssl/server.key")

}
