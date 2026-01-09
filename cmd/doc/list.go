package doc

import (
	"os"
	"strconv"

	"github.com/TypingHare/course-sync/internal/feature"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available documentation.",
	Long:  `List all available documentation in the project.`,

	Run: func(cmd *cobra.Command, args []string) {
		docs, err := feature.GetDocs()
		if err != nil {
			cmd.PrintErrf("Error retrieving documentation: %v\n", err)
			return
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
	},
}
