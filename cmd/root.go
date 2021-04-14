package cmd

import (
	"fmt"
	"wingCA/config"
	"wingCA/rootCA"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:     "wingCA",
		Version: "0.0.0",
	}
	verbose    bool
	configFile string
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "wing.yaml", "config file path")
	rootCmd.AddCommand(versionCmd, initCmd, issueCmd, sampleCmd, dummpyCmd, adminCommand)

}

// Main 命令主入口
func Main() {
	fmt.Println("")
	config.InitConfigFile(configFile)
	rootCA.InitConfigParamas()
	rootCmd.Execute()
}
