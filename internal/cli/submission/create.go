package submission

import (
	"fmt"
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/spf13/cobra"
)

// createCmd builds the submission create subcommand.
func createCmd(ctx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create <assignment-name>",
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

			submission, err := app.CreateSubmission(
				ctx.OutputMode,
				app.GetDataDir(ctx.ProjectDir),
				app.GetSrcDir(ctx.ProjectDir),
				assignmentName,
			)
			if err != nil {
				return fmt.Errorf("failed to create submission: %w", err)
			}

			submissionService := app.GetSubmissionService(
				app.GetSubmissionDataFile(app.GetDataDir(ctx.ProjectDir)),
			)
			err = submissionService.AddSubmission(submission)
			if err != nil {
				return fmt.Errorf("failed to add submission: %w", err)
			}

			cmd.Printf("Created submission: %s\n", submission.Hash)
			return nil
		},
	}

	return cmd
}
