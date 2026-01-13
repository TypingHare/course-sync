package cli

import (
	"fmt"
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/spf13/cobra"
)

var (
	shouldCommitOnly bool
	shouldPullOnly   bool
	shouldPushOnly   bool
)

func syncCmd(ctx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sync",
		Short: "Synchronize local and remote repositories",
		Long:  strings.TrimSpace(``),
		RunE: func(cmd *cobra.Command, args []string) error {
			dataDir := app.GetDataDir(ctx.ProjectDir)

			if shouldCommitOnly {
				return app.Commit(ctx.Role, ctx.OutputMode, dataDir)
			}

			if shouldPullOnly {
				return app.Pull(ctx.Role, ctx.OutputMode, dataDir)
			}

			if shouldPushOnly {
				return app.Push(ctx.Role, ctx.OutputMode, dataDir)
			}

			err := app.Commit(ctx.Role, ctx.OutputMode, dataDir)
			if err != nil {
				return fmt.Errorf("failed to commit local changes: %w", err)
			}

			err = app.Pull(ctx.Role, ctx.OutputMode, dataDir)
			if err != nil {
				return fmt.Errorf("failed to pull remote changes: %w", err)
			}

			err = app.Push(ctx.Role, ctx.OutputMode, dataDir)
			if err != nil {
				return fmt.Errorf("failed to push local commits: %w", err)
			}

			return nil
		},
	}

	cmd.Flags().BoolVarP(
		&shouldCommitOnly,
		"commit", "c", false,
		"Only commit local changes",
	)

	cmd.Flags().BoolVarP(
		&shouldPullOnly,
		"pull", "p", false,
		"Only pull remote changes",
	)

	cmd.Flags().BoolVarP(
		&shouldPushOnly,
		"push", "u", false,
		"Only push local commits",
	)

	return cmd
}
