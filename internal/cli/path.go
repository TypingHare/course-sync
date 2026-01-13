package cli

import (
	"fmt"
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/spf13/cobra"
)

var (
	shouldDisplayProjectDir            bool
	shouldDisplayDataDir               bool
	shouldDisplaySrcDir                bool
	shouldDisplayDocsDir               bool
	shouldDisplayPrototypeWorkspaceDir bool
	shouldDisplayStudentWorkspaceDir   bool
)

func pathCmd(ctx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "path",
		Short: "Display paths",
		Long: strings.TrimSpace(`
Display Course Sync application related paths.

Course Sync uses a small set of well-defined files and directories to store
application data. This command displays the resolved paths to those locations
within the current project.

Key paths include:

    - Project (root) directory  
      The top-level directory of the project, identified by the presence of a
      Git repository (i.e., the hidden folder ".git").

    - Data directory  
      A hidden directory located at the project root directory where Course Sync
      stores internal application data.

    - Documentation (docs) directory
      The directory containing all documentation files.

    - Source directory  
      The directory containing all source files.

    - Prototype workspace directory
      A subdirectory of the source directory that holds prototype workspace
      files. This directory is managed by instructors and should not be modified
      by students.

    - Student workspace directory
      A subdirectory of the source directory that holds student-specific
      workspace files. The name of the directory can be checked by the
      "user --dirname" command. Students should avoid modifying files in other
      workspace directories.

Use the flags below to display the paths to individual files or directories.
        `),
		RunE: func(cmd *cobra.Command, args []string) error {
			switch {
			case shouldDisplayProjectDir:
				cmd.Println(ctx.ProjectDir)
			case shouldDisplayDataDir:
				cmd.Println(app.GetDataDir(ctx.ProjectDir))
			case shouldDisplayDocsDir:
				cmd.Println(app.GetDocsDir(ctx.ProjectDir))
			case shouldDisplaySrcDir:
				cmd.Println(app.GetSrcDir(ctx.ProjectDir))
			case shouldDisplayPrototypeWorkspaceDir:
				cmd.Println(app.GetPrototypeWorkspaceDir(ctx.ProjectDir))
			case shouldDisplayStudentWorkspaceDir:
				studentWorkspaceDir, err := app.GetStudentWorkspaceDir(
					&ctx.OutputMode,
					ctx.ProjectDir,
				)
				if err != nil {
					return fmt.Errorf(
						"failed to get student workspace dir: %w",
						err,
					)
				}

				cmd.Println(studentWorkspaceDir)
			}

			return nil
		},
	}

	cmd.Flags().BoolVar(
		&shouldDisplayProjectDir,
		"project", false,
		"display project root directory path",
	)

	cmd.Flags().BoolVar(
		&shouldDisplayDataDir,
		"data", false,
		"display application data directory path",
	)

	cmd.Flags().BoolVar(
		&shouldDisplaySrcDir,
		"src", false,
		"display source directory path",
	)

	cmd.Flags().BoolVar(
		&shouldDisplayDocsDir,
		"docs", false,
		"display documentation directory path",
	)

	cmd.Flags().BoolVar(
		&shouldDisplayPrototypeWorkspaceDir,
		"prototype-workspace", false,
		"display prototype workspace directory path",
	)

	cmd.Flags().BoolVar(
		&shouldDisplayStudentWorkspaceDir,
		"student-workspace", false,
		"display student workspace directory path",
	)

	return cmd
}
