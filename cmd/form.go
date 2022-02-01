/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/force4760/quotedopai/src/suggest"
	"github.com/spf13/cobra"
)

// formCmd represents the form command
var formCmd = &cobra.Command{
	Use:   "form",
	Short: "Suggest a new Quote",
	Long:  `Open, in your default browser, a google form for suggesting a new Quote`,
	Run: func(cmd *cobra.Command, args []string) {
		suggest.OpenURL()
		fmt.Println("Thank you for contributing")
	},
}

func init() {
	rootCmd.AddCommand(formCmd)
}
