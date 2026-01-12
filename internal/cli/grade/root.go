package grade

import (
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/spf13/cobra"
)

func Cmd(appCtx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "grade",
		Short: "Manage grades",
		Long: strings.TrimSpace(`
Manage submission grades.

This command groups all grade-related actions, including listing grades and
viewing feedback for specific submissions.
        `),
	}

	cmd.AddCommand(listCmd(appCtx), showCmd(appCtx))

	return cmd
}
