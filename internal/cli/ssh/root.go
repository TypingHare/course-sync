package ssh

import (
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/spf13/cobra"
)

func Cmd(ctx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ssh",
		Short: "Manage SSH",
		Long: strings.TrimSpace(`
Manage SSH keys used by Course Sync.

This command provides access to ssh-related actions, such as generating new
instructor SSH keys.
        `),
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	cmd.AddCommand(keygenCmd(ctx))

	return cmd
}
