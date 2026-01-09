package grade

import (
	"os"
	"strconv"

	"github.com/TypingHare/course-sync/internal/feature"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var shouldDisplayAll bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display a list of grades.",
	Long:  `Display a list of grades.`,

	Run: func(cmd *cobra.Command, args []string) {
		grades, err := feature.GetGradeHistory()
		if err != nil {
			cmd.PrintErrf("Error retrieving grade history: %v\n", err)
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.Header([]string{"Assignment Name", "Submission Hash", "Score", "Graded At"})
		for _, grade := range grades {
			table.Append([]string{
				grade.AssignmentName,
				grade.SubmissionHash,
				strconv.FormatFloat(grade.Score, 'f', 2, 64),
				grade.GradedAt.Format("2006-01-02 15:04"),
			})
		}

		table.Render()
	},
}

func init() {
	listCmd.Flags().BoolVarP(
		&shouldDisplayAll,
		"all",
		"a",
		false,
		"Display the entire grade history.",
	)
}
