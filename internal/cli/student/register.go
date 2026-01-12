package student

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/domain/student"
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

			scanner := bufio.NewScanner(os.Stdin)
			promptLine := func(label string) (string, error) {
				fmt.Print(label)
				if !scanner.Scan() {
					if err := scanner.Err(); err != nil {
						return "", err
					}
					return "", io.EOF
				}
				return strings.TrimSpace(scanner.Text()), nil
			}

			studentName, err := promptLine("Student Name: ")
			if err != nil {
				return fmt.Errorf("failed to read student name: %w", err)
			}

			studentEmail, err := promptLine("Student Email: ")
			if err != nil {
				return fmt.Errorf("failed to read student email: %w", err)
			}

			repositoryURL, err := promptLine("Repository URL: ")
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
