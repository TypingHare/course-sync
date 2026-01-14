package student

import (
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/spf13/cobra"
)

func Cmd(ctx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "student",
		Short: "Manage students",
		Long: strings.TrimSpace(`
Manage the course roster stored in the Course Sync configuration.

Use this command to list registered students or to add new students to the
roster. Changes are saved to the project configuration file and used by other
commands that need student metadata.
        `),
	}

	cmd.AddCommand(listCmd(ctx))
	cmd.AddCommand(registerCmd(ctx))

	return cmd
}
