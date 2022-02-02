package cmd

import (
	"github.com/spf13/cobra"

	"github.com/force4760/quotedopai/src/printer"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of generated code example",
	Long:  `All software has versions. This is generated code example`,
	Run: func(cmd *cobra.Command, args []string) {
		printer.ColorVersion()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
