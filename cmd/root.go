package cmd

import (
	"fmt"
	"os"

	"github.com/TypingHare/course-sync/cmd/assignment"
	"github.com/TypingHare/course-sync/cmd/config"
	"github.com/TypingHare/course-sync/cmd/doc"
	"github.com/TypingHare/course-sync/cmd/grade"
	"github.com/TypingHare/course-sync/cmd/ssh"
	"github.com/TypingHare/course-sync/cmd/student"
	"github.com/TypingHare/course-sync/cmd/submission"
	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/feature"
	"github.com/spf13/cobra"
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
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetVersionTemplate("{{.Version}}\n")
	rootCmd.PersistentFlags().
		BoolVarP(&app.Verbose, "verbose", "v", false, "Enable verbose output.")
	rootCmd.PersistentFlags().
		BoolVarP(&app.Quiet, "quiet", "q", false, "Suppress non-error output.")

	rootCmd.AddCommand(
		roleCmd,
		pathCmd,
		pullCmd,
		pushCmd,
		syncCmd,
		config.Command(),
		ssh.Command(),
		assignment.Command(),
		submission.Command(),
		grade.Command(),
		doc.Command(),
	)

	role, err := feature.GetRole()
	if err == nil && role == feature.RoleMaster {
		rootCmd.AddCommand(
			student.Command(),
		)
	}
}
