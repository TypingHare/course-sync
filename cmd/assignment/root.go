package assignment

import (
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "assignment",
		Short:   "Manage assignments.",
		Long:    `Manage assignments.`,
		Aliases: []string{"am"},
	}

	cmd.AddCommand(listCmd, prepareCmd)

	return cmd
}
