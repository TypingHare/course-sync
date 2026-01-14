package student

import (
	"fmt"
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/domain/model"
	"github.com/spf13/cobra"
)

func registerCmd(ctx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register <name> <email> <repository-url>",
		Short: "Register a new student",
		Long: strings.TrimSpace(`
Add a new student to the course roster.

You will be prompted for the student's name, email, and repository URL. The
command assigns the next available numeric ID and saves the updated roster to
the configuration file.
        `),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			email := args[1]
			repositoryURL := args[2]

			studentService := app.GetStudentService(
				app.GetStudentDataFile(app.GetDataDir(ctx.ProjectDir)),
			)
			newStudentID, err := studentService.GetNextStudentID()
			if err != nil {
				return fmt.Errorf("failed to get new student ID: %w", err)
			}

			newStudent := &model.Student{
				ID:            newStudentID,
				Name:          name,
				Email:         email,
				RepositoryURL: repositoryURL,
			}

			err = studentService.AddStudent(newStudent)
			if err != nil {
				return fmt.Errorf("failed to add new student: %w", err)
			}

			return nil
		},
	}

	return cmd
}
