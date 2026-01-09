package grade

import (
	"github.com/TypingHare/course-sync/internal/feature"
	"github.com/spf13/cobra"
)

var submissionHash string

var feedbackCmd = &cobra.Command{
	Use:   "show",
	Short: "Display a grade entry in detail.",
	Long:  `Display a grade entry in detail.`,

	Run: func(cmd *cobra.Command, args []string) {
		gradeHistory, err := feature.GetGradeHistory()
		if err != nil {
			cmd.PrintErrf("Error retrieving grade history: %v\n", err)
			return
		}

		var grade *feature.Grade
		if submissionHash == "" {
			if len(args) == 0 {
				cmd.PrintErrln(
					"Please provide either a submission hash or an assignment name as an argument.",
				)
				return
			}

			assignmentName := args[0]
			grade = feature.FindLastGradeByAssignmentName(gradeHistory, assignmentName)
		} else {
			grade = feature.FindGradeBySubmissionHash(gradeHistory, submissionHash)
		}

		if grade == nil {
			cmd.PrintErrln("No grade found for the provided criteria.")
			return
		}

		cmd.PrintErrf("Feedback for assignment '%s':\n%s\n", grade.AssignmentName, grade.Feedback)
	},
}

func init() {
	feedbackCmd.Flags().StringVarP(
		&submissionHash,
		"submission-hash",
		"s",
		"",
		"Submission hash to get feedback for.",
	)
}
