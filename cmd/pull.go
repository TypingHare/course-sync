package cmd

import (
	"github.com/TypingHare/course-sync/internal/feature"
	"github.com/spf13/cobra"
)

var pullCommand = &cobra.Command{
	Use:   "pull",
	Short: "Pull changes from remote repository.",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		return feature.Pull()
	},
}

func init() {
	rootCommand.AddCommand(pullCommand)
}
