/* Copyright © 2026 James Chen <jameschan312.cn@gmail.com> */

package cmd

import (
	"os"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:     app.EXECUTABLE_NAME,
	Short:   "Course Sync helps students and teachers sync course materials.",
	Long:    `Course Sync is a Git-based CLI tool for syncing course materials.`,
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
