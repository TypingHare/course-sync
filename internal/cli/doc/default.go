package doc

import (
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/domain/doc"
	"github.com/spf13/cobra"
)

func defaultCmd(appCtx *app.Context) *cobra.Command {
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
			docs, err := doc.GetDocs(appCtx)
			if err != nil {
				return err
			}

			defaultDoc := doc.GetDefaultDoc(docs)
			if defaultDoc != nil {
				cmd.Println(defaultDoc.Name)
			} else {
				cmd.PrintErrf("No default documentation set.")
			}

			return nil
		},
	}

	return cmd
}
