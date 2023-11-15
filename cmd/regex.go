/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	keygen "github.com/onee-only/keygen/pkg/gen"
	"github.com/spf13/cobra"
)

var (
	Regex     string
	MaxRepeat uint16

	regexCmd = &cobra.Command{
		Use:   "regex",
		Short: "Generates random string matching given regex.",
		Long: `Regex (keygen regex) generates new random string 
with given regular expression.`,
		Run: func(cmd *cobra.Command, args []string) {
			gen, err := keygen.NewRegexGenerator(&keygen.RegexConfig{
				BaseConfig: keygen.DefaultBaseConfig(),
				Regex:      Regex,
				MaxRepeat:  MaxRepeat,
			})
			cobra.CheckErr(err)

			stream, cancel := gen.GenerateStream()
			defer cancel()

			for i := uint8(0); i < Count; i++ {
				log.Println(string(<-stream))
			}
		},
	}
)

func init() {
	regexCmd.Flags().StringVarP(&Regex, "regex", "r", "", "regular expression to use")
	regexCmd.MarkFlagRequired("regex")

	regexCmd.Flags().Uint16VarP(&MaxRepeat, "max", "", 10, "max repeat count")
	regexCmd.MarkFlagRequired("max")

	rootCmd.AddCommand(regexCmd)
}
