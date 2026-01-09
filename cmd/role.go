package cmd

import (
	"github.com/TypingHare/course-sync/internal/feature"
	"github.com/spf13/cobra"
)

var roleCmd = &cobra.Command{
	Use:   "role",
	Short: "Display the current user role.",
	Long:  `This command displays the current user role, either 'student' or 'master'.`,

	Run: func(cmd *cobra.Command, args []string) {
		role, err := feature.GetRole()
		if err != nil {
			cmd.PrintErrln("Error determining role:", err)
		}

		cmd.Println(role)
	},
}
