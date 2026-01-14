package assignment

import (
	"fmt"
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/domain/model"
	"github.com/spf13/cobra"
)

func assignCmd(ctx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "assign <name> <title> <due-at>",
		Short: "Assign an assignment",
		Long: strings.TrimSpace(`
Assign a new assignment.
        `),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO: Options for updating other fields like MaxScore,
			// PassingScore, Description, etc.
			if len(args) != 3 {
				return fmt.Errorf("expected 3 arguments, got %d", len(args))
			}

			name := args[0]
			title := args[1]
			dueAtStr := args[2]

			dueAt, err := app.ParseDateTimeString(dueAtStr)
			if err != nil {
				return err
			}

			newAssignment := &model.Assignment{
				Name:         name,
				Title:        title,
				Description:  "",
				DueAt:        dueAt.UTC(),
				MaxScore:     100.0,
				PassingScore: 60.0,
			}

			err = app.Assign(ctx.OutputMode, ctx.ProjectDir, newAssignment)
			if err != nil {
				return fmt.Errorf("failed to assign assignment: %w", err)
			}

			// Note: app.Assign doesn't commit all the changes, only create the
			// assignment and distribute files and directories. Instructors
			// should use "sync" command to commit all the changes.

			return nil
		},
	}

	return cmd
}
