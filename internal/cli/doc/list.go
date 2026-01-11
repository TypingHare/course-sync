package doc

import (
	"os"
	"strconv"
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/domain/doc"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func listCmd(appCtx *app.Context) *cobra.Command {
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
			docs, err := doc.GetDocs(appCtx)
			if err != nil {
				return err
			}

			table := tablewriter.NewWriter(os.Stdout)
			table.Header([]string{"Name", "Title", "Path", "Is Default"})
			for _, doc := range docs {
				table.Append([]string{
					doc.Name,
					doc.Title,
					doc.Path,
					strconv.FormatBool(doc.IsDefault),
				})
			}

			table.Render()

			return nil
		},
	}

	return cmd
}
