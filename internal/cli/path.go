package cli

import (
	"fmt"
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/domain/workspace"
	"github.com/spf13/cobra"
)

var (
	shouldDisplayProjectRootDir        bool
	shouldDisplayAppDataDir            bool
	shouldDisplaySourceDir             bool
	shouldDisplayDocsDir               bool
	shouldDisplayPrototypeWorkspaceDir bool
	shouldDisplayUserWorkspaceDir      bool
)

func pathCmd(appCtx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "path",
		Short: "Display application related paths",
		Long: strings.TrimSpace(`
Display Course Sync application related paths.

Course Sync uses a small set of well-defined files and directories to store
application data. This command displays the resolved paths to those locations
within the current project.

Key paths include:

  - Project root directory  
    The top-level directory of the project, identified by the presence of a
    Git repository (.git).

  - Application data directory  
    A hidden directory located at the project root where Course Sync stores
    internal application data.

  - Source directory  
    The directory containing all course source files.

  - Documentation (docs) directory
    The directory containing all documentation files.

  - Prototype workspace directory
    A subdirectory of the source directory that holds prototype workspace
    files.

  - User workspace directory
    A subdirectory of the source directory that holds user-specific workspace
    files.

Use the flags below to display the paths to individual files or directories.
        `),
		RunE: func(cmd *cobra.Command, args []string) error {
			if shouldDisplayProjectRootDir {
				cmd.Println(appCtx.ProjectDir)
			} else if shouldDisplayAppDataDir {
				cmd.Println(appCtx.AppDataDir)
			} else if shouldDisplaySourceDir {
				cmd.Println(appCtx.SrcDir)
			} else if shouldDisplayDocsDir {
				cmd.Println(appCtx.DocsDir)
			} else if shouldDisplayPrototypeWorkspaceDir {
				cmd.Println(workspace.GetPrototypeWorkspaceDir(appCtx))
			} else if shouldDisplayUserWorkspaceDir {
				userWorkspaceDir, err := workspace.GetUserWorkspaceDir(appCtx)
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
		&shouldDisplayProjectRootDir,
		"project", "p", false,
		"display project root directory path",
	)

	cmd.Flags().BoolVarP(
		&shouldDisplayAppDataDir,
		"app-data", "a", false,
		"display application data directory path",
	)

	cmd.Flags().BoolVarP(
		&shouldDisplaySourceDir,
		"source", "s", false,
		"display source directory path",
	)

	cmd.Flags().BoolVarP(
		&shouldDisplayDocsDir,
		"docs", "d", false,
		"display documentation directory path",
	)

	cmd.Flags().BoolVarP(
		&shouldDisplayPrototypeWorkspaceDir,
		"prototype-workspace", "o", false,
		"display prototype workspace directory path",
	)

	cmd.Flags().BoolVarP(
		&shouldDisplayUserWorkspaceDir,
		"user-workspace", "u", false,
		"display user workspace directory path",
	)

	return cmd
}
