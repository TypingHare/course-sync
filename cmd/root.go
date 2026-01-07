package cmd

import (
	"os"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/spf13/cobra"
)

var (
	verbose bool
	quiet   bool
)

var rootCmd = &cobra.Command{
	Use:   app.EXECUTABLE_NAME,
	Short: "Course Sync helps students and teachers synchronize course materials.",
	Long: `Course Sync is a Git-based CLI tool that helps students and teachers synchronize course
materials between local machines and remote repositories.`,
	Version: app.VERSION,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetVersionTemplate("{{.Version}}\n")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output.")
	rootCmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "Suppress non-error output.")
}
