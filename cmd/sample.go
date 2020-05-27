package cmd

import (
	"log"
	"wingCA/rootCA"

	"github.com/spf13/cobra"
)

var (
	sampleCmd = &cobra.Command{
		Use:   "sample",
		Short: "Start Sample Site",
		Run:   startSampleSite,
	}
)

func init() {
	sampleCmd.Flags().BoolVarP(&startDouble, "double", "", false, " Start site  double validate ")
	sampleCmd.Flags().StringVarP(&commonName, "common", "c", "", "Site Crt Common Name")
}

func startSampleSite(cmd *cobra.Command, args []string) {
	log.Println("start sample site -> ", commonName, " startDouble -> ", startDouble)
	serverCrtPath := rootCA.SiteCertPath(commonName)
	serverKeyPath := rootCA.PrivateKeyPath(commonName)
	caRootPath := rootCA.RootCACertPath
	if !startDouble {
		rootCA.SampleWeb(serverCrtPath, serverKeyPath)
	} else {
		log.Println(" validate command : ")
		log.Println("curl -v --cert ssl/client/xingyue.crt --key ssl/private/xingyue.key https://127.0.0.1")
		rootCA.SampleDoubleWeb(serverCrtPath, serverKeyPath, caRootPath)
	}
}
