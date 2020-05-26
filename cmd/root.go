package cmd

import (
	"wingCA/rootCA"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use: "wingCA",
	}
	verbose bool
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	rootCmd.AddCommand(versionCmd, sampleCmd, dummpyCmd)
}

// Main 命令主入口
func Main() {
	rootCA.InitDir()
	rootCmd.Execute()
}
