package cmd

import (
	"fmt"
	"wingCA/config"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version  of wingCA",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("    wingCA v%s , compile time : %s ^_^ \n\n", config.Version, config.BuildTime)
		if verbose {
			fmt.Println(config.Item)
		}
	},
}
