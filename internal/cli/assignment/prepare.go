package assignment

import (
	"fmt"
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/spf13/cobra"
)

var shouldForcePrepare bool

func prepareCmd(ctx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "prepare <assignment-name>",
		Short: "Prepare a user assignment directory",
		Long: strings.TrimSpace(`
This command creates a user assignment directory by copying the corresponding
prototype assignment directory into the user workspace.

The prototype assignment directory is located under the prototype workspace.
You can view its path with:

    csync path --prototype-workspace

The user assignment directory is created under the user workspace, which can be
viewed with:
    
    csync path --user-workspace -q

The assignment directory name matches the assignment name, which is provided as
the first argument to this command.

In effect, this command copies:

    <prototype-workspace>/<assignment-name>
    
to:

    <user-workspace>/<assignment-name>

If the destination directory already exists, it will not be overwritten unless
the --force flag is specified.
        `),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			assignmentName := args[0]
			err := app.PrepareAssignment(
				ctx.OutputMode,
				ctx.ProjectDir,
				assignmentName,
				shouldForcePrepare,
			)
			if err != nil {
				return fmt.Errorf("failed to prepare assignment: %w", err)
			}

			return nil
		},
	}

	cmd.Flags().BoolVarP(
		&shouldForcePrepare,
		"force", "f", false,
		"re-prepare even if the user assignment directory "+
			"already exists",
	)

	return cmd
}
