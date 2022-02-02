package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/force4760/quotedopai/src/quotes"
)

// quoteCmd represents the quote command
var quoteCmd = &cobra.Command{
	Use:   "quote",
	Short: "Default command",
	Long:  `Returns a random quote from Nunes da Silva, without any filter`,
	Run: func(cmd *cobra.Command, args []string) {
		quote := quotes.QuotesDoPai.Random()

		fmt.Println(string(quote))
	},
}

func init() {
	rootCmd.AddCommand(quoteCmd)
}
