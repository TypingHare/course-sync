package submission

import (
	"os"
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/domain/submission"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func listCmd(appCtx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all submissions",
		Long: strings.TrimSpace(`
List all submissions.

This command retrieves and displays all submissions in a tabular format,
including their hash, assignment name, and submission date.
        `),
		RunE: func(cmd *cobra.Command, args []string) error {
			submissions, err := submission.GetSubmissions(appCtx)
			if err != nil {
				return err
			}

			table := tablewriter.NewWriter(os.Stdout)
			table.Header([]string{"Hash", "Assignment Name", "Submitted At"})
			for _, submission := range submissions {
				table.Append([]string{
					submission.Hash,
					submission.AssignmentName,
					submission.SubmittedAt.Format("2006-01-02 15:04"),
				})
			}

			return table.Render()
		},
	}

	return cmd
}
