package grade

import (
	"os"
	"strconv"
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/domain/grade"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func listCmd(appCtx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all grades",
		Long: strings.TrimSpace(`
List grade entries.

This command retrieves all recorded grades and displays them in a table,
including the assignment name, submission hash, score, and the time each
submission was graded.
        `),
		RunE: func(cmd *cobra.Command, args []string) error {
			grades, err := grade.GetGrades(appCtx)
			if err != nil {
				return err
			}

			table := tablewriter.NewWriter(os.Stdout)
			table.Header(
				[]string{
					"Assignment Name",
					"Submission Hash",
					"Score",
					"Graded At",
				},
			)
			for _, grade := range grades {
				table.Append([]string{
					grade.AssignmentName,
					grade.SubmissionHash,
					strconv.FormatFloat(grade.Score, 'f', 2, 64),
					grade.GradedAt.Format("2006-01-02 15:04"),
				})
			}

			return table.Render()
		},
	}

	return cmd
}
