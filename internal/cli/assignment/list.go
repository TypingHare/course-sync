package assignment

import (
	"os"
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/domain/assignment"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var shouldDisplayAll bool

func ListCmd(appCtx *app.Context) *cobra.Command {
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
			// TODO: Implement filtering based on shouldDisplayAll flag
			assignments, err := assignment.GetAssignments(appCtx)
			if err != nil {
				return err
			}

			table := tablewriter.NewWriter(os.Stdout)
			table.Header([]string{"Name", "Title", "Released At", "Due At"})
			for _, assignment := range assignments {
				table.Append([]string{
					assignment.Name,
					assignment.Title,
					assignment.ReleasedAt.Format("2006-01-02 15:04"),
					assignment.DueAt.Format("2006-01-02 15:04"),
				})
			}

			table.Render()

			return nil
		},
	}

	cmd.Flags().BoolVarP(
		&shouldDisplayAll,
		"all", "a", false,
		"Display all assignments, including submitted ones",
	)

	return cmd
}
