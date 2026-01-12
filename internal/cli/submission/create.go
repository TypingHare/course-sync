package submission

import (
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/domain/submission"
	"github.com/spf13/cobra"
)

func createCmd(appCtx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new submission",
		Long: strings.TrimSpace(`
Create a new submission for an assignment.

This command creates a new submission entry for the specified assignment. A
unique hash is generated from the files included in the submission, and the
submission is recorded in the submissions file in the application data
directory.

On success, the command prints the hash of the newly created submission.
        `),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			assignmentName := args[0]

			s, err := submission.CreateSubmission(appCtx, assignmentName)
			if err != nil {
				return err
			}

			cmd.Printf("Created submission: %s\n", s.Hash)
			return nil
		},
	}

	return cmd
}
