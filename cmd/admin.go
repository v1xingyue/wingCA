package cmd

import (
	"wingCA/api"

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
	api.StartDownloadRootCA()
	api.StartSSL()
}
