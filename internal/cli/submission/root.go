package submission

import (
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/spf13/cobra"
)

func Cmd(appCtx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submission",
		Short: "Manage submissions",
		Long: strings.TrimSpace(`
Manage submissions for assignments.

This command groups all submission-related actions such as listing submissions
and creating new submissions.
        `),
	}

	cmd.AddCommand(listCmd(appCtx), createCmd(appCtx))

	return cmd
}
