package cmd

import (
	"crypto/x509/pkix"
	"log"
	"wingCA/config"
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
	initCmd.Flags().StringVarP(&name, "name", "n", config.Default.Name, "CA Name You Will Created.")
	initCmd.Flags().StringVarP(&org, "org", "o", config.Default.Org, "Organization Name In Your CA.")
	initCmd.Flags().StringVarP(&province, "province", "p", config.Default.Province, "Province In Your CA.")
	initCmd.Flags().StringVarP(&locality, "locality", "l", config.Default.Locality, "Locality In Your CA.")
	initCmd.Flags().StringVarP(&street, "street", "s", config.Default.Street, "Locality In Your CA.")
	initCmd.Flags().StringVarP(&postcode, "postcode", "", config.Default.Postcode, "Postcode In Your CA.")
	initCmd.Flags().BoolVarP(&withMiddle, "withMiddle", "", false, "Switch to init with middle certificate .")
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

		if withMiddle {
			name.CommonName += "_Middle"
			log.Println("Make middle certificate : ", name)
			if rootCA.InitMiddle(name) == nil {
				log.Printf("Your CA Middle Root Have Been Put  %s/middle/rootCAMiddle.crt \n ", rootCA.RootCAPath)
			}
		}
	}
}
