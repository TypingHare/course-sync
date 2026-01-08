package grade

import (
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "grade",
		Short: "Manage grades.",
		Long:  `Manage grades.`,
	}

	cmd.AddCommand(listCmd)

	return cmd
}
