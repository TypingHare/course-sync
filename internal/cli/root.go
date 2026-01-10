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
	cmd.AddCommand(ContextCmd(*appCtx))

	return cmd
}
