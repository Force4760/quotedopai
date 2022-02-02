package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/force4760/quotedopai/src/imageMaker"
	"github.com/force4760/quotedopai/src/quotes"
)

// Flags
var isXorImage bool = false

// withCmd represents the with command
var xorCmd = &cobra.Command{
	Use:   "xor",
	Short: "A quote containing at one of the provided words",
	Long: `Returns a quote from Nunes da Silva containing at one of the provided words
	word1 XOR word2
	`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		word1 := args[0]
		word2 := args[1]

		quote := quotes.QuotesDoPai.WordXorIn(word1, word2)

		if isXorImage {
			imageMaker.MakeImage(string(quote))
		} else {
			fmt.Println(string(quote))
		}

	},
}

func init() {
	rootCmd.AddCommand(xorCmd)

	xorCmd.Flags().BoolVarP(&isXorImage, "image", "i", false, "If the quote should be exported as an image")
}
