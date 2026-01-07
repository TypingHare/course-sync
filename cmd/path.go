package cmd

import (
	"github.com/TypingHare/course-sync/internal/app"
	"github.com/spf13/cobra"
)

var shouldShowProjectPath bool

var pathCmd = &cobra.Command{
	Use:   "path",
	Short: "Display paths used by Course Sync.",
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {
		if shouldShowProjectPath {
			projectPath, err := app.FindProjectDir("")
			if err != nil {
				cmd.PrintErrln("Error retrieving project path:", err)
				return
			}
			cmd.Println(projectPath)
		} else {
			cmd.Println("Please specify a flag to display a specific path.")
		}
	},
}

func init() {
	rootCmd.AddCommand(pathCmd)
	pathCmd.PersistentFlags().BoolVarP(
		&shouldShowProjectPath,
		"project",
		"p",
		false,
		"Show the Course Sync project directory path.",
	)
}
