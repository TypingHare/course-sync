package submission

import (
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submission",
		Short: "Manage submissions.",
		Long:  `Commands to manage submissions in the system.`,
	}

	cmd.AddCommand(listCmd, createCmd)

	return cmd
}
