package cli

import (
	"strconv"
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/spf13/cobra"
)

func contextCmd(ctx *app.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "context",
		Short: "Display CLI contexts",
		Long: strings.TrimSpace(`
Display Course Sync CLI contexts.

Course Sync maintains a set of contexts that influence command behavior and
output. This command displays the current runtime context, including global CLI
options and environment-derived settings.

Configuration is considered part of the application context but is loaded lazily
and therefore not shown here. To view or manage configuration values, use the
'config' command.
        `),
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Printf("verbose = %s\n", strconv.FormatBool(ctx.Verbose))
			cmd.Printf("quiet = %s\n", strconv.FormatBool(ctx.Quiet))
			cmd.Printf("plain = %s\n", strconv.FormatBool(ctx.Plain))
			cmd.Printf("working_directory = %s\n", ctx.WorkingDir)
			cmd.Printf("project_directory = %s\n", ctx.ProjectDir)
			cmd.Printf("user_role = %s\n", string(ctx.Role))
		},
	}
}
