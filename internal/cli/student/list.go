package student

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func listCmd(appCtx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all students",
		Long:  strings.TrimSpace(``),
		RunE: func(cmd *cobra.Command, args []string) error {
			config, err := appCtx.GetConfig()
			if err != nil {
				return fmt.Errorf("failed to get config: %w", err)
			}
			if config == nil {
				return fmt.Errorf("config is nil")
			}

			roster := config.Roster

			table := tablewriter.NewWriter(os.Stdout)
			table.Header([]string{"ID", "Name", "Email", "Repository URL"})
			for _, student := range roster {
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
