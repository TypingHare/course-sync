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

func filesCmd(ctx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "files",
		Short: "Display files users manage",
		Long:  strings.TrimSpace(``),
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
