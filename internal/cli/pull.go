package cli

import (
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/infra/git"
	"github.com/spf13/cobra"
)

func pullCmd(appCtx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pull",
		Short: "Pull the latest changes from the remote repository",
		Long: strings.TrimSpace(`
Pull the latest changes from the remote repository.

This command fetches and integrates updates from the remote Git repository into
the local repository. Before pulling, it restores instructor-managed files to
prevent local changes from being overwritten.

Instructor-managed files include:

    - Documentation files
    - Prototype workspace files
    - Public key files
    - Documentation records
    - Assignment records
    - Grade records

Restoring these files ensures that any local modifications are preserved during
the pull operation.
        `),
		RunE: func(cmd *cobra.Command, args []string) error {
			return pull(appCtx)
		},
	}

	return cmd
}

func pull(appCtx *app.Context) error {
	// Restore all instructor files before pulling to avoid overwriting them.
	for _, file := range getInstructorFiles(appCtx) {
		err := git.Restore(appCtx, file)
		if err != nil {
			return err
		}
	}

	return git.Pull(appCtx, true)
}
