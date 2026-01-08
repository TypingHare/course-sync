/* Copyright © 2026 James Chen <jameschan312.cn@gmail.com> */

package cmd

import (
	"os"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/config"
	"github.com/spf13/cobra"
)

var (
	appConfigStore *config.Store
	appConfig      config.Config
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

func initStore() error {
	store, err := config.NewStore()
	if err != nil {
		return err
	}

	appConfigStore = store

	_config, err := appConfigStore.Load()
	if err != nil {
		return err
	}

	appConfig = _config

	return nil
}
