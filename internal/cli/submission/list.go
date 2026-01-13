package submission

import (
	"fmt"
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func listCmd(ctx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all submissions",
		Long: strings.TrimSpace(`
List all submissions.

This command retrieves and displays all submissions in a tabular format,
including their hash, assignment name, and submission date.
        `),
		RunE: func(cmd *cobra.Command, args []string) error {
			submissionService := app.GetSubmissionService(
				app.GetSubmissionDataFile(app.GetDataDir(ctx.ProjectDir)),
			)
			submissions, err := submissionService.GetAllSubmissions()
			if err != nil {
				return fmt.Errorf("failed to get submissions: %w", err)
			}

			table := tablewriter.NewWriter(cmd.OutOrStdout())
			table.Header([]string{"Hash", "Assignment Name", "Submitted At"})
			for _, submission := range submissions {
				table.Append([]string{
					submission.Hash,
					submission.AssignmentName,
					app.GetDateTimeString(submission.SubmittedAt),
				})
			}

			return table.Render()
		},
	}

	return cmd
}
