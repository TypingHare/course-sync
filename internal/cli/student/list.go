package student

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func listCmd(ctx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all students",
		Long: strings.TrimSpace(`
Display the current student roster in a table.

This reads the roster from the configuration file and prints each student's
ID, name, email, and repository URL. Use this to verify the roster before
syncing or grading operations.
        `),
		RunE: func(cmd *cobra.Command, args []string) error {
			studentService := app.GetStudentService(
				app.GetStudentDataFile(app.GetDataDir(ctx.ProjectDir)),
			)

			students, err := studentService.GetAllStudents()
			if err != nil {
				return fmt.Errorf("failed to get students: %w", err)
			}

			table := tablewriter.NewWriter(cmd.OutOrStdout())
			table.Header([]string{"ID", "Name", "Email", "Repository URL"})
			for _, student := range students {
				table.Append([]string{
					strconv.Itoa(student.ID),
					student.Name,
					student.Email,
					student.RepositoryURL,
				})
			}

			return table.Render()
		},
	}

	return cmd
}
