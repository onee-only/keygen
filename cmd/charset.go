/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/onee-only/keygen/internal/util"
	keygen "github.com/onee-only/keygen/pkg/gen"
	"github.com/spf13/cobra"
)

var (
	Charset string

	charsetCmd = &cobra.Command{
		Use:   "charset",
		Short: "Generates new random key using given charset",
		Long:  `Charset (keygen new charset) generates new random string using given charset.`,
		Run: func(cmd *cobra.Command, args []string) {
			gen := keygen.NewCustomGenerator(&keygen.CustomConfig{
				BaseConfig: keygen.DefaultBaseConfig(),
				Len:        Len,
				Chartset:   util.StrToBytes(Charset),
			})

			stream, cancel := gen.GenerateStream()
			defer cancel()

			for i := uint8(0); i < Count; i++ {
				log.Println(string(<-stream))
			}
		},
	}
)

func init() {
	charsetCmd.Flags().StringVarP(&Charset, "charset", "c", "", "charset to use for generation")
	charsetCmd.MarkFlagRequired("charset")

	newCmd.AddCommand(charsetCmd)
}
