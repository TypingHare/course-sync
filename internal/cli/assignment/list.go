package assignment

import (
	"os"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/domain/assignment"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func ListCmd(appCtx *app.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all assignments",
		RunE: func(cmd *cobra.Command, args []string) error {
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
}
