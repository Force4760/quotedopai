package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/force4760/quotedopai/src/imageMaker"
	"github.com/force4760/quotedopai/src/quotes"
)

// Flags
var isAndImage bool = false

// withCmd represents the with command
var andCmd = &cobra.Command{
	Use:   "and",
	Short: "A quote containing both of the provided words",
	Long: `Returns a quote from Nunes da Silva containing both of the provided words
	word1 && word2
	`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		word1 := args[0]
		word2 := args[1]

		quote := quotes.QuotesDoPai.WordAndIn(word1, word2)

		if isAndImage {
			imageMaker.MakeImage(string(quote))
		} else {
			fmt.Println(string(quote))
		}

	},
}

func init() {
	rootCmd.AddCommand(andCmd)

	andCmd.Flags().BoolVarP(&isAndImage, "image", "i", false, "If the quote should be exported as an image")
}
