package cmd

import (
	"os"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   app.EXECUTABLE_NAME,
	Short: "Course Sync helps students and teachers synchronize course materials.",
	Long: `Course Sync is a Git-based CLI tool that helps students and teachers synchronize
course materials between local machines and remote repositories.`,
	Version: app.VERSION,
}

func Execute() {
	err := rootCommand.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCommand.SetVersionTemplate("{{.Version}}\n")
}
