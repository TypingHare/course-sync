package grade

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// listCmd builds the grade list subcommand.
func listCmd(ctx *app.Context) *cobra.Command {
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
			gradeService := app.GetGradeService(
				app.GetGradeDataFile(app.GetDataDir(ctx.ProjectDir)),
			)
			grades, err := gradeService.GetAllGrades()
			if err != nil {
				return fmt.Errorf("failed to get grades: %w", err)
			}

			table := tablewriter.NewWriter(cmd.OutOrStdout())
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
					app.GetDateTimeString(grade.GradedAt),
				})
			}

			return table.Render()
		},
	}

	return cmd
}
