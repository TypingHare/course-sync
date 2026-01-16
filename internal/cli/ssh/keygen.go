package ssh

import (
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/spf13/cobra"
)

var shouldForceKeygen bool = false

// keygenCmd builds the ssh keygen subcommand.
func keygenCmd(ctx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "keygen",
		Short: "Generate a new instructor SSH key pair",
		Long: strings.TrimSpace(`
Generate a new instructor SSH key pair and save it to the application data
directory.

By default, this command will not overwrite an existing instructor SSH key pair.
Use the --force flag to regenerate the key pair even if one already exists.
        `),
		RunE: func(cmd *cobra.Command, args []string) error {
			return app.GenerateInstructorKeyPair(
				ctx.OutputMode,
				ctx.ProjectDir,
				app.GetDataDir(ctx.ProjectDir),
				shouldForceKeygen,
			)
		},
	}

	cmd.Flags().BoolVarP(
		&shouldForceKeygen,
		"force", "f", false,
		"regenerate and overwrite if it already exists",
	)

	return cmd
}
