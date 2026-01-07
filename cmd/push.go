package cmd

import (
	"github.com/TypingHare/course-sync/internal/feature"
	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push changes to remote repository.",
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {
		feature.Push(quiet, verbose)
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
}
