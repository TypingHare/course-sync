package assignment

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/domain/assignment"
	"github.com/TypingHare/course-sync/internal/ui"
	"github.com/spf13/cobra"
)

func releaseCmd(appCtx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "release <name> <title> <due-at>",
		Short: "Release an assignment to students",
		Long: strings.TrimSpace(`
Release a new assignment with the specified name, title, and due date/time.
        `),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 3 {
				return fmt.Errorf("expected 3 arguments, got %d", len(args))
			}

			assignmentName := args[0]
			title := args[1]
			dueAtStr := args[2]

			dueAt, err := ui.ParseUTCTime(dueAtStr)
			if err != nil {
				return err
			}

			newAssignment := assignment.CreateAssignment(
				assignmentName,
				title,
				dueAt.UTC(),
			)

			err = assignment.SaveAssignmentToFile(appCtx, newAssignment)
			if err != nil {
				return fmt.Errorf("save assignment: %w", err)
			}

			err = app.MoveFileToStudentRepos(
				appCtx,
				filepath.Join(
					appCtx.AppDataDir,
					assignment.AssignmentsFileName,
				),
			)
			if err != nil {
				return fmt.Errorf(
					"distribute assignment to student repos: %w",
					err,
				)
			}

			return nil
		},
	}

	return cmd
}
