package assignment

import (
	"os"

	"github.com/TypingHare/course-sync/internal/feature"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var shouldDisplayAllAssignments bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display a list of assignments.",
	Long:  `Display a list of assignments.`,

	Run: func(cmd *cobra.Command, args []string) {
		assignments, err := feature.GetAssignments()
		if err != nil {
			cmd.PrintErrf("Error retrieving assignments: %v\n", err)
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.Header([]string{"Name", "Title", "Release", "Due"})
		for _, assignment := range assignments {
			table.Append([]string{
				assignment.Name,
				assignment.Title,
				assignment.ReleaseTime.Format("2006-01-02 15:04"),
				assignment.DueTime.Format("2006-01-02 15:04"),
			})
		}

		table.Render()
	},
}

func init() {
	listCmd.Flags().BoolVarP(
		&shouldDisplayAllAssignments,
		"all",
		"a",
		false,
		"Display all assignments, including completed ones.",
	)
}
