package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var wingVersion = "0.0.2"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version  of wingCA",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("    wingCA v%s ^_^ \n", wingVersion)
	},
}
