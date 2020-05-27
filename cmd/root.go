package cmd

import (
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
	rootCmd.AddCommand(versionCmd, initCmd, issueCmd, sampleCmd, dummpyCmd)
}

// Main 命令主入口
func Main() {
	rootCmd.Execute()
}
