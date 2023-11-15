/*
Copyright © 2023 김원욱 <kimww0306@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:     "keygen",
		Version: "0.0.1",
		Long:    `Keygen is a CLI tool to generate keys/passwords/etc.`,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().Uint8VarP(&Count, "count", "", 1, "determine output count")
	rootCmd.MarkFlagRequired("count")
}
