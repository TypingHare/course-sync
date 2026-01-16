package doc

import (
	"fmt"
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/support/exec"
	"github.com/spf13/cobra"
)

// openCmd builds the doc open subcommand.
func openCmd(ctx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "open [doc-name]",
		Short: "Open a documentation file.",
		Long: strings.TrimSpace(`
Open a documentation file.

This command opens a documentation file using the system’s default application
for the file type. You may specify the name of the document to open as an
argument. If no name is provided, the command attempts to open the default
documentation file, if one is configured.

Documentation files are stored in the documentation directory. To view the path
to this directory, run:

    csync path --docs

If the specified documentation file does not exist or cannot be opened, the
command returns an error.
        `),
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			docService := app.GetDocService(
				app.GetDocDataFile(app.GetDataDir(ctx.ProjectDir)),
			)

			var docName string
			if len(args) == 0 {
				defaultDoc, err := docService.GetDefaultDoc()
				if err != nil {
					return fmt.Errorf("failed to get default doc: %w", err)
				}
				if defaultDoc == nil {
					return fmt.Errorf("no default documentation found; " +
						"please specify a documentation name")
				}

				docName = defaultDoc.Name
			} else {
				docName = args[0]
			}

			// Get the documentation path.
			docToOpen, err := docService.GetDocByName(docName)
			if err != nil {
				return fmt.Errorf("failed to get doc by name: %w", err)
			}
			if docToOpen == nil {
				return fmt.Errorf("documentation '%s' not found", docName)
			}

			return exec.OpenFile(
				ctx.OutputMode,
				ctx.ProjectDir,
				app.GetDocFilePath(ctx.ProjectDir, docToOpen.Path),
			)
		},
	}

	return cmd
}
