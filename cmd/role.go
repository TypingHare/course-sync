package cmd

import "github.com/spf13/cobra"

var roleCmd = &cobra.Command{
	Use:   "role",
	Short: "Display the current user role.",
	Long:  `This command displays the current user role, either 'student' or 'master'.`,

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("student")
	},
}
