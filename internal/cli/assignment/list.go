package assignment

import (
	"fmt"
	"os"
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func listCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all assignments",
		Long: strings.TrimSpace(`
List the assignments available in the course.

This command reads assignment information from the assignments.json file in the
application data directory and displays the results in a table, showing each
assignment’s name, title, release date, and due date.

By default, only assignments that have not yet been submitted are listed. Use
the --all flag to include submitted assignments as well.
        `),
		RunE: func(cmd *cobra.Command, args []string) error {
			assignments, err := app.GetAllAssignment()
			if err != nil {
				return fmt.Errorf("failed to get assignments: %w", err)
			}

			table := tablewriter.NewWriter(os.Stdout)
			table.Header([]string{"Name", "Title", "Released At", "Due At"})
			for _, assignment := range assignments {
				table.Append([]string{
					assignment.Name,
					assignment.Title,
					assignment.ReleasedAt.Local().
						Format("2006-01-02 15:04 MST"),
					assignment.DueAt.Local().Local().
						Format("2006-01-02 15:04 MST"),
				})
			}

			return table.Render()
		},
	}

	return cmd
}
