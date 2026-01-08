/* Copyright © 2026 James Chen <jameschan312.cn@gmail.com> */

package cmd

import (
	"github.com/spf13/cobra"
)

var force bool

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "Initialize the application configuration.",
	Long: `This command initializes the application configuration file with default 
    configurations.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		initStore()
		return appConfigStore.InitFile(force)
	},
}

func init() {
	initCommand.Flags().
		BoolVar(&force, "force", false, "Overwrite the existing configuration file if it exists.")

	rootCommand.AddCommand(initCommand)
}
