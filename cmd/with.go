package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/force4760/quotedopai/src/quotes"
)

// withCmd represents the with command
var withCmd = &cobra.Command{
	Use:   "with",
	Short: "A quote containing the provided word",
	Long:  `Returns a quote from Nunes da Silva containing the provided word`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		word := args[0]

		quote := quotes.QuotesDoPai.WordIn(word)

		fmt.Println(string(quote))

	},
}

func init() {
	rootCmd.AddCommand(withCmd)
}
