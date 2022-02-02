package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/force4760/quotedopai/src/quotes"
)

// withCmd represents the with command
var orCmd = &cobra.Command{
	Use:   "with",
	Short: "A quote containing at least one of the provided words",
	Long: `Returns a quote from Nunes da Silva containing at least one of the provided words
	word1 || word2
	`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		word1 := args[0]
		word2 := args[1]

		quote := quotes.QuotesDoPai.WordOrIn(word1, word2)

		fmt.Println(string(quote))

	},
}

func init() {
	rootCmd.AddCommand(orCmd)
}
