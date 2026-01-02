/* Copyright © 2026 James Chen <jameschan312.cn@gmail.com> */

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var pushCommand = &cobra.Command{
	Use:   "push",
	Short: "Push local changes to the remote repository",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sync called")
	},
}

func init() {
	rootCommand.AddCommand(pushCommand)
}
