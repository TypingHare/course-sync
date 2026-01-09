package student

import (
	"os"
	"strconv"

	"github.com/TypingHare/course-sync/internal/config"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all students in the roster",

	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.Get()
		if err != nil {
			cmd.PrintErrf("Error loading config: %v\n", err)
		}
		if config == nil {
			cmd.PrintErrln("Config is nil.")
			return
		}

		studentInfoArray := config.Master.Roster

		table := tablewriter.NewWriter(os.Stdout)
		table.Header([]string{"ID", "Name", "Email", "Repository URL"})
		for _, studentInfo := range studentInfoArray {
			table.Append([]string{
				strconv.Itoa(studentInfo.Id),
				studentInfo.Name,
				studentInfo.Email,
				studentInfo.RepositoryUrl,
			})
		}

		table.Render()
	},
}
