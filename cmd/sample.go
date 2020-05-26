package cmd

import (
	"io/ioutil"
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
	startDouble bool
)

func init() {
	sampleCmd.Flags().BoolVarP(&startDouble, "double", "", false, " Start site  double validate ")
}

func startSampleSite(cmd *cobra.Command, args []string) {
	log.Println("start sample wite : ", startDouble)
	if !startDouble {
		rootCA.SampleWeb("ssl/site/server.crt", "ssl/private/server.key")
	} else {
		cert, key, err := rootCA.IssueClient("xingyue")
		if err == nil {
			ioutil.WriteFile("ssl/client/client.crt", cert, 0700)
			ioutil.WriteFile("ssl/private/client.key", key, 0700)
			pkBytes, err := rootCA.MakePKCS12("ssl/client/client.crt", "ssl/private/client.key", "super")
			if err != nil {
				log.Println(err)
			} else {
				ioutil.WriteFile("ssl/p12/client.p12", pkBytes, 0600)
			}
			rootCA.SampleDoubleWeb("ssl/site/server.crt", "ssl/private/server.key", "ssl/root/rootCA.crt")
		} else {
			log.Println(err)
		}
	}
}
