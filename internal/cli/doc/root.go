package doc

import (
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/spf13/cobra"
)

// Cmd builds the doc command group.
func Cmd(ctx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "doc",
		Short: "Manage documentation",
		Long: strings.TrimSpace(`
Manage course documentation.

This command provides access to documentation-related actions, such as listing
available documents and opening specific documents.
        `),
	}

	cmd.AddCommand(listCmd(ctx))
	cmd.AddCommand(defaultCmd(ctx))
	cmd.AddCommand(openCmd(ctx))

	return cmd
}
