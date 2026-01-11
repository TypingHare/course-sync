package ssh

import (
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/spf13/cobra"
)

func Cmd(appCtx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ssh",
		Short: "Manage SSH",
		Long: strings.TrimSpace(`
Manage SSH keys used by Course Sync.

This command provides access to ssh-related actions, such as generating new
master SSH keys.
        `),
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	cmd.AddCommand(keygenCmd(appCtx))

	return cmd
}
