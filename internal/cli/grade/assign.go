package grade

import (
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/spf13/cobra"
)

func assignCmd(ctx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "assign",
		Short: "Assign a grade to a student",
		Long: strings.TrimSpace(`
        `),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			println(ctx)
			return nil
		},
	}

	return cmd
}
