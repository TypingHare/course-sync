package student

import "github.com/spf13/cobra"

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "student",
		Short: "Manage student records",
		Long: `A command-line tool to manage student records including adding, removing, and
listing students.`,
	}

	cmd.AddCommand(listCmd, registerCmd)

	return cmd
}
