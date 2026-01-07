package cmd

import (
	"github.com/TypingHare/course-sync/internal/feature"
	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull changes from remote repository.",
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {
		feature.Pull()
	},
}
