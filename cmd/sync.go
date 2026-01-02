/* Copyright © 2026 James Chen <jameschan312.cn@gmail.com> */

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var syncCommand = &cobra.Command{
	Use:   "sync",
	Short: "Synchronize local repository with the remote repository.",
	Long: `This command sequentially performs a pull followed by a push to ensure
    both local and remote repositories are up to date.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sync called")
	},
}

func init() {
	rootCommand.AddCommand(syncCommand)
}
