package cli

import (
	"fmt"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/support/exec"
	"github.com/spf13/cobra"
)

var shouldDisplayUserDirname bool

func userCmd(ctx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			gitUsername, err := exec.GitGetUsername(&ctx.OutputMode)
			if err != nil {
				return fmt.Errorf("failed to get git username: %w", err)
			}

			if shouldDisplayUserDirname {
				cmd.Println(app.GetStudentDirName(gitUsername))
			} else {
				cmd.Println(gitUsername)
			}

			return nil
		},
	}

	cmd.Flags().BoolVarP(
		&shouldDisplayUserDirname,
		"dirname", "d", false,
		"Display the user's directory name",
	)

	return cmd
}
