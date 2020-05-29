package cmd

import (
	"crypto/x509/pkix"
	"log"
	"wingCA/rootCA"

	"github.com/spf13/cobra"
)

var (
	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Init CA Root",
		Run:   initRootCA,
	}
)

func init() {
	initCmd.Flags().BoolVarP(&confirmInitCA, "confirm", "", false, "confirm init CA Structure.")
	initCmd.Flags().StringVarP(&name, "name", "n", "ArkRootCA", "CA Name You Will Created.")
	initCmd.Flags().StringVarP(&org, "org", "o", "Ryan Ark Center", "Organization Name In Your CA.")
	initCmd.Flags().StringVarP(&province, "province", "p", "Beijing", "Province In Your CA.")
	initCmd.Flags().StringVarP(&locality, "locality", "l", "Beijing", "Locality In Your CA.")
	initCmd.Flags().StringVarP(&street, "street", "s", "NoWhere Road 9+3/4 Site Corner", "Locality In Your CA.")
	initCmd.Flags().StringVarP(&postcode, "postcode", "", "061219", "Postcode In Your CA.")
}

func initRootCA(cmd *cobra.Command, args []string) {
	if !confirmInitCA {
		log.Println("Must Add --confirm flag confirm init CA Structure.")
	} else {
		log.Println("Begin Init Root CA")
		rootCA.InitDir()
		name := pkix.Name{
			Country:       []string{"CN"},
			Organization:  []string{org},
			Province:      []string{province},
			Locality:      []string{locality},
			StreetAddress: []string{street},
			PostalCode:    []string{postcode},
			CommonName:    name,
		}
		log.Println("Your CA Name Info : ", name)
		if rootCA.InitRootCA(name) == nil {
			log.Printf("Your CA Root Have Been Put  %s/root/rootCA.crt \n ", rootCA.RootCAPath)
		}
	}
}
