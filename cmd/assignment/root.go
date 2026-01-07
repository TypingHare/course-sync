package assignment

import (
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "assignment",
		Short: "",
		Long:  ``,
	}

	cmd.AddCommand(listCmd, submitCmd)

	return cmd
}
