package cli

import (
	"fmt"
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/domain/model"
	"github.com/spf13/cobra"
)

var (
	shouldDisplayStudentFiles    bool
	shouldDisplayInstructorFiles bool
)

// filesCmd builds the files subcommand.
func filesCmd(ctx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "files",
		Short: "Show files managed by each role",
		Long: strings.TrimSpace(`
List the paths that Course Sync treats as role-managed files.

Student-managed files are the ones students are expected to modify and sync.
Instructor-managed files are owned by the instructor and are restored for
students during sync operations.

Each line is a path relative to the project root. If a path is a directory,
all files under that directory are considered managed.

Use --student or --instructor to show the list for a specific role.
        `),
		RunE: func(cmd *cobra.Command, args []string) error {
			role := ctx.Role
			if shouldDisplayStudentFiles {
				role = model.RoleStudent
			} else if shouldDisplayInstructorFiles {
				role = model.RoleInstructor
			}

			var files []string
			switch role {
			case model.RoleStudent:
				files = app.GetStudentFiles(ctx.ProjectDir)
				cmd.Printf("Student files:\n")
			case model.RoleInstructor:
				files = app.GetInstructorFiles(ctx.ProjectDir)
				cmd.Printf("Instructor files:\n")
			case model.RoleUnknown:
				return fmt.Errorf("Unknown role")
			}

			for _, file := range files {
				cmd.Println(file)
			}

			return nil
		},
	}

	cmd.Flags().BoolVarP(
		&shouldDisplayStudentFiles,
		"student", "s", false,
		"display student files",
	)

	cmd.Flags().BoolVarP(
		&shouldDisplayInstructorFiles,
		"instructor", "i", false,
		"display instructor files",
	)

	return cmd
}
