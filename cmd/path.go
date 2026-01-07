package cmd

import (
	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/feature"
	"github.com/spf13/cobra"
)

var (
	shouldDisplayProjectPath   bool
	shouldDisplayDocsPath      bool
	shouldDisplaySrcPath       bool
	shouldDisplayPrototypePath bool
	shouldDisplayUserPath      bool
	shouldDisplayAppPath       bool
)

var pathCmd = &cobra.Command{
	Use:   "path",
	Short: "Display paths used by Course Sync.",
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {
		if shouldDisplayProjectPath {
			projectDirPath, err := app.GetProjectDirPath()
			if err != nil {
				cmd.Println("Error retrieving project path:", err)
				return
			}

			cmd.Println(projectDirPath)
		} else if shouldDisplayDocsPath {
			docsDirPath, err := app.GetDocsDirPath()
			if err != nil {
				cmd.Println("Error retrieving docs path:", err)
				return
			}
			cmd.Println(docsDirPath)
		} else if shouldDisplaySrcPath {
			srcDirPath, err := app.GetSrcDirPath()
			if err != nil {
				cmd.Println("Error retrieving src path:", err)
				return
			}
			cmd.Println(srcDirPath)
		} else if shouldDisplayPrototypePath {
			prototypeDirPath, err := app.GetPrototypeDirPath()
			if err != nil {
				cmd.Println("Error retrieving prototype path:", err)
				return
			}
			cmd.Println(prototypeDirPath)
		} else if shouldDisplayUserPath {
			userDirPath, err := feature.GetUserDirPath()
			if err != nil {
				cmd.Println("Error retrieving user path:", err)
				return
			}
			cmd.Println(userDirPath)
		} else if shouldDisplayAppPath {
			appDirPath, err := app.GetAppDirPath()
			if err != nil {
				cmd.Println("Error retrieving app path:", err)
				return
			}
			cmd.Println(appDirPath)
		} else {
			cmd.Println("Please specify a flag to display a specific path.")
		}
	},
}

func init() {
	pathCmd.PersistentFlags().BoolVarP(
		&shouldDisplayProjectPath,
		"project-root",
		"r",
		false,
		"Display the Course Sync project root directory path.",
	)

	pathCmd.PersistentFlags().BoolVarP(
		&shouldDisplayDocsPath,
		"docs",
		"d",
		false,
		"Display the Course Sync documentation directory path.",
	)

	pathCmd.PersistentFlags().BoolVarP(
		&shouldDisplaySrcPath,
		"src",
		"s",
		false,
		"Display the Course Sync source code directory path.",
	)

	pathCmd.PersistentFlags().BoolVarP(
		&shouldDisplayPrototypePath,
		"prototype",
		"p",
		false,
		"Display the Course Sync prototype directory path.",
	)

	pathCmd.PersistentFlags().BoolVarP(
		&shouldDisplayUserPath,
		"user",
		"u",
		false,
		"Display the Course Sync user directory path.",
	)

	pathCmd.PersistentFlags().BoolVarP(
		&shouldDisplayAppPath,
		"app",
		"a",
		false,
		"Display the Course Sync application directory path.",
	)
}
