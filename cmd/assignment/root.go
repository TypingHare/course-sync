package assignment

import (
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "assignment",
		Short: "Manage assignments.",
		Long:  `Manage assignments.`,
	}

	cmd.AddCommand(listCmd, prepareCmd)

	return cmd
}
