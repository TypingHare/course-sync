package grade

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/domain/model"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var submissionHash string

// showCmd builds the grade show subcommand.
func showCmd(ctx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show [assignment-name]",
		Short: "Show grade details",
		Long: strings.TrimSpace(`
Show detailed grade information for a submission.

This command displays the grade and feedback for a specific assignment
submission. If an assignment name is provided, the grade for the most recent
submission of that assignment is shown.

If no assignment name is provided, use the --submission-hash flag to display
the grade for a specific submission.
        `),
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			gradeService := app.GetGradeService(
				app.GetGradeDataFile(app.GetDataDir(ctx.ProjectDir)),
			)

			var err error
			var grade *model.Grade
			if submissionHash == "" {
				if len(args) == 0 {
					return fmt.Errorf(
						"please provide either a submission hash or an " +
							"assignment name as an argument",
					)
				}

				assignmentName := args[0]
				grade, err = gradeService.GetLastGradeByAssignmentName(
					assignmentName,
				)
				if err != nil {
					return fmt.Errorf("failed to retrieve grade: %w", err)
				}
			} else {
				grade, err = gradeService.GetGradeBySubmissionHash(
					submissionHash,
				)
				if err != nil {
					return fmt.Errorf("failed to retrieve grade: %w", err)
				}
			}

			if grade == nil {
				return fmt.Errorf("no grade found for the provided criteria")
			}

			colorAssignmentNameFunc := color.New(color.FgHiMagenta).SprintFunc()
			colorGradeFunc := color.New(color.FgHiGreen).SprintFunc()
			cmd.Printf(
				"Your grade for assignment %s is %s\n",
				colorAssignmentNameFunc(grade.AssignmentName),
				colorGradeFunc(
					strconv.FormatFloat(grade.Score, 'f', 2, 64),
				),
			)

			feedbackColorFunc := color.New(color.FgHiYellow).SprintFunc()
			cmd.Printf("\n%s\n", feedbackColorFunc(grade.Feedback))

			return nil
		},
	}

	return cmd
}
