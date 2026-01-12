package student

import (
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/spf13/cobra"
)

func Cmd(appCtx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "student",
		Short: "Manage students",
		Long:  strings.TrimSpace(``),
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	cmd.AddCommand(listCmd(appCtx), registerCmd(appCtx))

	return cmd
}
