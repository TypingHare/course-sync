package cli

import (
	"github.com/TypingHare/course-sync/internal/app"
	"github.com/spf13/cobra"
)

func Cmd(appCtx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use: "csync",
		Short: "Course Sync (csync) helps students and teachers " +
			"synchronize course materials.",
		Long:    ``,
		Version: "2026.1",
	}

	cmd.SetVersionTemplate("{{.Version}}\n")
	cmd.PersistentFlags().BoolVarP(
		&appCtx.Verbose,
		"verbose",
		"v",
		false,
		"Enable verbose output.",
	)
	cmd.PersistentFlags().BoolVarP(
		&appCtx.Quiet,
		"quiet",
		"q",
		false,
		"Suppress non-error output.",
	)

	cmd.AddCommand(
		ContextCmd(*appCtx),
		PathCmd(appCtx),
	)

	return cmd
}
