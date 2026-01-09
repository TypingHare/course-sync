package doc

import "github.com/spf13/cobra"

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "doc",
		Short: "Manage documentation.",
		Long:  ``,
	}

	cmd.AddCommand(listCmd, openCmd)

	return cmd
}
