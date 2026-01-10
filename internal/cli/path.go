package cli

import (
	"fmt"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/domain/workspace"
	"github.com/spf13/cobra"
)

var shouldDisplayUserWorkspacePath bool

func PathCmd(appCtx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "path",
		Short: "Display application related paths",
		Long:  `A command to manage and display system paths.`,

		RunE: func(cmd *cobra.Command, args []string) error {
			if shouldDisplayUserWorkspacePath {
				userWorkspaceDir, err := workspace.GetUserWorkspaceDir(*appCtx)
				if err != nil {
					return err
				}

				cmd.Println(userWorkspaceDir)
			} else {
				return fmt.Errorf("no flag specified")
			}

			return nil
		},
	}

	cmd.Flags().BoolVarP(
		&shouldDisplayUserWorkspacePath,
		"user-workspace",
		"u",
		false,
		"Display user workspace directory path.",
	)

	return cmd
}
