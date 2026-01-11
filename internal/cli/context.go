package cli

import (
	"strconv"
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/spf13/cobra"
)

func ContextCmd(appCtx *app.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "context",
		Short: "Display CLI contexts",
		Long: strings.TrimSpace(`
Course Sync maintains a set of contexts that influence command behavior and
output. This command displays the current runtime context, including global CLI
options and environment-derived settings.

Configuration is considered part of the application context but is loaded lazily
and therefore not shown here. To view or manage configuration values, use the
'config' command.
        `),
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Printf("verbose = %s\n", strconv.FormatBool(appCtx.Verbose))
			cmd.Printf("quiet = %s\n", strconv.FormatBool(appCtx.Quiet))
			cmd.Printf("working directory = %s\n", appCtx.WorkingDir)
			cmd.Printf("project directory = %s\n", appCtx.ProjectDir)
			cmd.Printf("application data directory = %s\n", appCtx.AppDataDir)
			cmd.Printf("source directory = %s\n", appCtx.SrcDir)
			cmd.Printf("user role = %s\n", string(appCtx.Role))
		},
	}
}
