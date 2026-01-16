package submission

import (
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/spf13/cobra"
)

// Cmd builds the submission command group.
func Cmd(ctx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submission",
		Short: "Manage submissions",
		Long: strings.TrimSpace(`
Manage submissions for assignments.

This command groups all submission-related actions such as listing submissions
and creating new submissions.
        `),
	}

	cmd.AddCommand(listCmd(ctx))
	cmd.AddCommand(createCmd(ctx))

	return cmd
}
