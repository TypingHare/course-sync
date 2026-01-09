package submission

import (
	"os"

	"github.com/TypingHare/course-sync/internal/feature"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all submissions",
	Long:  `List all submissions in the system with their details.`,

	Run: func(cmd *cobra.Command, args []string) {
		submissions, err := feature.GetSubmissionHistory()
		if err != nil {
			cmd.PrintErrln("Error fetching submission history:", err)
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.Header([]string{"Hash", "Assignment Name", "Submitted At"})
		for _, submission := range submissions {
			table.Append([]string{
				submission.Hash,
				submission.AssignmentName,
				submission.SubmittedAt.Format("2006-01-02 15:04"),
			})
		}

		table.Render()
	},
}
