package doc

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func listCmd(ctx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all available documentation.",
		Long: strings.TrimSpace(`
List assignments available in the course.

This command reads assignment data from the assignments.json file in the
application data directory and displays the results in a table, including each
assignment’s name, title, release date, and due date.

By default, only assignments that have not been submitted are shown. Use the
--all flag to include assignments that have already been submitted.
        `),
		RunE: func(cmd *cobra.Command, args []string) error {
			docService := app.GetDocService(
				app.GetDocDataFile(app.GetDataDir(ctx.ProjectDir)),
			)
			docs, err := docService.GetAllDocs()
			if err != nil {
				return fmt.Errorf("failed to retrieve docs: %w", err)
			}

			table := tablewriter.NewWriter(os.Stdout)
			table.Header(
				[]string{
					"Name",
					"Title",
					"Version",
					"Released At",
					"Path",
					"Is Default",
				},
			)
			for _, doc := range docs {
				table.Append([]string{
					doc.Name,
					doc.Title,
					doc.Version,
					app.GetDateTimeString(doc.ReleasedAt),
					doc.Path,
					strconv.FormatBool(doc.IsDefault),
				})
			}

			return table.Render()
		},
	}

	return cmd
}
