package cmd

import (
	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/feature"
	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push changes to remote repository.",
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {
		app.RunAll(
			func() error { return feature.MakeSyncCommit() },
			func() error { return feature.GitPush() },
		)
	},
}
