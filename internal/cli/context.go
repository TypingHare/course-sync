package cli

import (
	"strconv"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/spf13/cobra"
)

func ContextCmd(appCtx app.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "context",
		Short: "Display CLI contexts.",

		Run: func(cmd *cobra.Command, args []string) {
			cmd.Printf("verbose = %s\n", strconv.FormatBool(appCtx.Verbose))
			cmd.Printf("quiet = %s\n", strconv.FormatBool(appCtx.Quiet))
			cmd.Printf("working directory = %s\n", appCtx.WorkingDir)
			cmd.Printf("project directory = %s\n", appCtx.ProjectDir)
			cmd.Printf("user role = %s\n", string(appCtx.Role))
		},
	}
}
