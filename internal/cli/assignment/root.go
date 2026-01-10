package assignment

import (
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/spf13/cobra"
)

func AssignmentCmd(appCtx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "assignment",
		Short: "Manage assignments",
		Long: strings.TrimSpace(`
Manage course assignments.

This command groups all assignment-related actions such as listing available
assignments and preparing assignments for submission.
        `),
	}

	cmd.AddCommand(ListCmd(appCtx))

	return cmd
}
