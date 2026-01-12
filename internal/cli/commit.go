package cli

import (
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/infra/git"
	"github.com/spf13/cobra"
)

func commitCmd(appCtx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "commit",
		Short: "Commit the application data to the local repository",
		Long: strings.TrimSpace(`
Commit staged student files to the local repository.

This command stages all student-managed files tracked by Course Sync and creates
a commit in the local Git repository using a standardized commit message.

Student-managed files include:

    - Configuration files
    - Student submissions

No other files are staged or committed by this command.
        `),
		RunE: func(cmd *cobra.Command, args []string) error {
			return commit(appCtx)
		},
	}

	return cmd
}

// commit stages the student files and creates a commit in the local repository.
func commit(appCtx *app.Context) error {
	studentFiles := getStudentFiles(appCtx)

	for _, file := range studentFiles {
		err := git.Add(appCtx, file)
		if err != nil {
			return err
		}
	}

	return git.Commit(appCtx, "csync: update student files")
}
