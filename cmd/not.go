package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/force4760/quotedopai/src/imageMaker"
	"github.com/force4760/quotedopai/src/quotes"
)

// Flags
var isNotImage bool = false

// withCmd represents the with command
var notCmd = &cobra.Command{
	Use:   "not",
	Short: "A quote not containing the provided word",
	Long:  `Returns a quote from Nunes da Silva not containing the provided word`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		word := args[0]

		quote := quotes.QuotesDoPai.WordNotIn(word)

		if isNotImage {
			imageMaker.MakeImage(string(quote))
		} else {
			fmt.Println(string(quote))
		}

	},
}

func init() {
	rootCmd.AddCommand(notCmd)

	notCmd.Flags().BoolVarP(&isNotImage, "image", "i", false, "If the quote should be exported as an image")
}
