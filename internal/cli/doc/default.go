package doc

import (
	"fmt"
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/spf13/cobra"
)

func defaultCmd(ctx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "default",
		Short: "Display the default documentation name.",
		Long: strings.TrimSpace(`
Display the default documentation name.

This command retrieves and displays the name of the default documentation file
configured in the application. If no default documentation is set, an error
message is displayed.
        `),
		RunE: func(cmd *cobra.Command, args []string) error {
			docService := app.GetDocService(
				app.GetDocDataFile(app.GetDataDir(ctx.ProjectDir)),
			)
			defaultDoc, err := docService.GetDefaultDoc()
			if err != nil {
				return fmt.Errorf("failed to retrieve default doc: %w", err)
			}
			if defaultDoc == nil {
				cmd.PrintErrf("no default documentation set\n")
				return nil
			}

			cmd.Println(defaultDoc.Name)

			return nil
		},
	}

	return cmd
}
