package assignment

import (
	"github.com/TypingHare/course-sync/internal/feature"
	"github.com/spf13/cobra"
)

var submitCmd = &cobra.Command{
	Use:   "submit",
	Short: "Submit your assignment.",
	Long:  `Submit your assignment.`,

	Run: func(cmd *cobra.Command, args []string) {
		assignmentName := args[0]
		feature.SubmitAssignment(assignmentName)
	},
}
