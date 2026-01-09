package student

import "github.com/spf13/cobra"

var registerCmd = &cobra.Command{
	Use:   "register <name> <email>",
	Short: "Register a new student",
	Long:  `Register a new student by providing necessary details.`,

	Run: func(cmd *cobra.Command, args []string) {
	},
}
