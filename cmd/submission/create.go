package submission

import (
	"github.com/TypingHare/course-sync/internal/feature"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create <assignment-name>",
	Short: "Create a submission for an assignment.",
	Long:  `Create a submission for an assignment.`,
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.PrintErrln("Error: Assignment name is required.")
		}

		assignmentName := args[0]

		submission, err := feature.CreateSubmission(assignmentName)
		if err != nil {
			cmd.PrintErrln("Error creating submission:", err)
		}

		feature.AppendSubmissionToHistory(*submission)
	},
}
