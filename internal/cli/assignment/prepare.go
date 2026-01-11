package assignment

import (
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/domain/assignment"
	"github.com/spf13/cobra"
)

var shouldForciblyPrepare bool

func prepareCmd(appCtx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "prepare <assignment-name>",
		Short: "Prepare a user assignment directory.",
		Long:  strings.TrimSpace(``),
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return assignment.PrepareAssignment(
				appCtx,
				args[0],
				shouldForciblyPrepare,
			)
		},
	}

	cmd.Flags().BoolVarP(
		&shouldForciblyPrepare,
		"force", "f", false,
		"re-prepare even if the user assignment directory "+
			"already exists",
	)

	return cmd
}
