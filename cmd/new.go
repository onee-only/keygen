/*
Copyright © 2023 김원욱
*/
package cmd

import (
	"log"

	keygen "github.com/onee-only/keygen/pkg/gen"
	"github.com/spf13/cobra"
)

var (
	Len       uint64
	Count     uint8
	UseUpper  bool
	UseLower  bool
	UseNumber bool
	UseSymbol bool

	newCmd = &cobra.Command{
		Use:   "new",
		Short: "Generates new random key.",
		Long:  `New (keygen new) generates new random key with config.`,
		Run: func(cmd *cobra.Command, args []string) {
			gen := keygen.NewConvGenerator(&keygen.ConvConfig{
				BaseConfig: keygen.DefaultBaseConfig(),
				Len:        Len,
				UseUpper:   UseUpper,
				UseLower:   UseLower,
				UseNumber:  UseNumber,
				UseSymbol:  UseSymbol,
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
	newCmd.PersistentFlags().Uint64VarP(&Len, "len", "", 0, "determine length of the string")
	newCmd.MarkFlagRequired("len")

	newCmd.Flags().BoolVarP(&UseUpper, "upper", "", false, "wheter string will contain uppercase letters")
	newCmd.Flags().BoolVarP(&UseLower, "lower", "", false, "wheter string will contain lowercase letters")
	newCmd.Flags().BoolVarP(&UseNumber, "number", "n", false, "wheter string will contain numbers")
	newCmd.Flags().BoolVarP(&UseSymbol, "symbol", "", false, "wheter string will contain symbols")
	newCmd.MarkFlagsOneRequired("upper", "lower", "number", "symbol")

	rootCmd.AddCommand(newCmd)
}
