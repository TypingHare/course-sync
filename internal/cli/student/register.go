package student

import (
	"fmt"
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/domain/student"
	"github.com/TypingHare/course-sync/internal/ui"
	"github.com/spf13/cobra"
)

func registerCmd(appCtx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register",
		Short: "Register a new student",
		Long: strings.TrimSpace(`
Add a new student to the course roster.

You will be prompted for the student's name, email, and repository URL. The
command assigns the next available numeric ID and saves the updated roster to
the configuration file.
        `),
		RunE: func(cmd *cobra.Command, args []string) error {
			config, err := appCtx.GetConfig()
			if err != nil {
				return fmt.Errorf("failed to get config: %w", err)
			}

			studentName, err := ui.ReadLine(cmd.OutOrStdout(), "Student Name: ")
			if err != nil {
				return fmt.Errorf("failed to read student name: %w", err)
			}

			studentEmail, err := ui.ReadLine(
				cmd.OutOrStdout(),
				"Student Email: ",
			)
			if err != nil {
				return fmt.Errorf("failed to read student email: %w", err)
			}

			repositoryURL, err := ui.ReadLine(
				cmd.OutOrStdout(),
				"Repository URL: ",
			)
			if err != nil {
				return fmt.Errorf("failed to read repository URL: %w", err)
			}

			// Determine the next student ID.
			largestID := 0
			for _, student := range config.Roster {
				if student.ID > largestID {
					largestID = student.ID
				}
			}

			newStudent := student.NewStudent(
				largestID+1,
				studentName,
				studentEmail,
				repositoryURL,
			)
			config.Roster = append(config.Roster, *newStudent)

			return appCtx.SaveConfig()
		},
	}

	return cmd
}
