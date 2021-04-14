package cmd

import (
	"log"
	"wingCA/config"

	"github.com/spf13/cobra"
)

var (
	adminCommand = &cobra.Command{
		Use:   "admin",
		Short: "Start admin Application",
		Run:   startApp,
	}
)

func startApp(cmd *cobra.Command, args []string) {
	log.Printf(
		"start admin application at %s:%d \n", config.Item.APIConfig.BindHost, config.Item.APIConfig.BindPort)
}
