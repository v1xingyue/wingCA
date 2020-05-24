package main

import (
	"crypto/x509/pkix"
	"fmt"
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

	if ca, err := rootCA.ParseCertificate("ssl/rootCA.crt"); err != nil {
		rootCA.InitRootCA(name)
	} else {
		log.Println(ca.IsCA)
	}

	cert, key, err := rootCA.IssueOneCert("debug.ssl.com.cn", []net.IP{net.IPv4(127, 0, 0, 1), net.IPv4(192, 168, 100, 87)}, []string{"localhost"})
	if err == nil {
		ioutil.WriteFile("ssl/server.crt", cert, 0700)
		ioutil.WriteFile("ssl/server.key", key, 0700)
	}

	startDouble := true
	if !startDouble {
		rootCA.SampleWeb("ssl/server.crt", "ssl/server.key")
	} else {
		fmt.Println(" start double")
		cert, key, err := rootCA.IssueClient("xingyue")
		if err == nil {
			ioutil.WriteFile("ssl/client.crt", cert, 0700)
			ioutil.WriteFile("ssl/client.key", key, 0700)
			pkBytes, err := rootCA.MakePKCS12("ssl/client.crt", "ssl/client.key", "super")
			if err != nil {
				log.Println(err)
			} else {
				ioutil.WriteFile("ssl/client.p12", pkBytes, 0600)
			}
			rootCA.SampleDoubleWeb("ssl/server.crt", "ssl/server.key", "ssl/rootCA.crt")
		} else {
			log.Println(err)
		}
	}

}
