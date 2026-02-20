package doc

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// listCmd builds the doc list subcommand.
func listCmd(ctx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all available documentation.",
		Long: strings.TrimSpace(`
List documentation available in the course.

This command reads documentation metadata from docs.json in the application data
directory and displays the results in a table, including each document's name,
title, version, release date, path, and whether it is the default document.
        `),
		RunE: func(cmd *cobra.Command, args []string) error {
			docService := app.GetDocService(
				app.GetDocDataFile(app.GetDataDir(ctx.ProjectDir)),
			)
			docs, err := docService.GetAllDocs()
			if err != nil {
				return fmt.Errorf("failed to retrieve docs: %w", err)
			}

			table := tablewriter.NewWriter(cmd.OutOrStdout())
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
