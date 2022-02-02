package cmd

import (
	"github.com/spf13/cobra"

	"github.com/force4760/quotedopai/src/suggest"
)

// formCmd represents the form command
var formCmd = &cobra.Command{
	Use:   "form",
	Short: "Suggest a new Quote",
	Long:  `Open, in your default browser, a google form for suggesting a new Quote`,
	Run: func(cmd *cobra.Command, args []string) {
		suggest.OpenURL()
	},
}

func init() {
	rootCmd.AddCommand(formCmd)
}
