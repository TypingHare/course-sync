package cli

import (
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/infra/fs"
	"github.com/spf13/cobra"
)

func initCmd(appCtx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize Course Sync folders and configuration",
		Long: strings.TrimSpace(`
Initialize the Course Sync application.

This command creates the required directory structure and default configuration
files needed for Course Sync to operate. It ensures that the application data,
documentation, and source directories exist, and creates a default configuration
file if one is not already present.
        `),
		RunE: func(cmd *cobra.Command, args []string) error {
			// Ensures the application data directory exists.
			if err := fs.EnsureDirExists(appCtx.AppDataDir); err != nil {
				return err
			}

			// Create a new configuration file.
			appCtx.SaveConfig()

			// Ensures the documentation directory exists.
			if err := fs.EnsureDirExists(appCtx.DocsDir); err != nil {
				return err
			}

			// Ensures the source directory exists.
			if err := fs.EnsureDirExists(appCtx.SrcDir); err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}
