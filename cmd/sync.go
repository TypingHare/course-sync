package cmd

import (
	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/feature"
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Synchronize local repository with the remote repository.",
	Long: `This command sequentially performs a pull followed by a push to ensure both local and
remote repositories are up to date.`,

	Run: func(cmd *cobra.Command, args []string) {
		app.RunAll(
			func() error { return feature.GitPull() },
			func() error { return feature.MakeSyncCommit() },
			func() error { return feature.GitPush() },
		)
	},
}
