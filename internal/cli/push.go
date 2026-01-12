package cli

import (
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/infra/git"
	"github.com/spf13/cobra"
)

func pushCmd(appCtx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "push",
		Short: "Push the local changes to the remote repository",
		Long: strings.TrimSpace(`
Push the committed changes from the local Git repository to the remote
repository.

This command simply executes 'git push' in the context of the course
repository.
        `),
		RunE: func(cmd *cobra.Command, args []string) error {
			return push(appCtx)
		},
	}

	return cmd
}

// push executes the git push command in the context of the course repository.
func push(appCtx *app.Context) error {
	return git.Push(appCtx)
}
