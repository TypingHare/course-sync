/* Copyright © 2026 James Chen <jameschan312.cn@gmail.com> */

package cmd

import (
	"github.com/spf13/cobra"
)

var (
	initConfig  bool
	forceConfig bool
)

var configCommand = &cobra.Command{
	Use:   "config",
	Short: "Display application configuration.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCommand.AddCommand(configCommand)
}
