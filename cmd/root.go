/* Copyright © 2026 James Chen <jameschan312.cn@gmail.com> */

package cmd

import (
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "course-sync",
	Short: "Course Sync helps students and teachers sync course materials.",
	Long: `Course Sync is a Git-based CLI tool for syncing course materials.`,
    Version: "2026.0",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
    rootCmd.SetVersionTemplate("{{.Version}}\n")
}
