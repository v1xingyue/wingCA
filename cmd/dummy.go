package cmd

import (
	"crypto/x509/pkix"
	"io/ioutil"
	"log"
	"net"
	"wingCA/rootCA"

	"github.com/spf13/cobra"
)

var (
	dummpyCmd = &cobra.Command{
		Use:   "dummy",
		Short: "Start dummy Command",
		Run: func(cmd *cobra.Command, args []string) {
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

			cert, key, err := rootCA.IssueSite("debug.ssl.com.cn", []net.IP{net.IPv4(127, 0, 0, 1), net.IPv4(192, 168, 100, 87)}, []string{"localhost"})
			if err == nil {
				ioutil.WriteFile("ssl/site/server.crt", cert, 0700)
				ioutil.WriteFile("ssl/private/server.key", key, 0700)
			}

			crlBytes, err := rootCA.CrlBytes()
			if err == nil {
				ioutil.WriteFile("crl.list", crlBytes, 0600)
			}
		},
	}
)
