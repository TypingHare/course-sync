package doc

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/domain/doc"
	"github.com/spf13/cobra"
)

func openCmd(appCtx *app.Context) *cobra.Command {
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
			docs, err := doc.GetDocs(appCtx)
			if err != nil {
				return err
			}

			// Determine the documentation name to open.
			var docName string
			if len(args) > 0 {
				docName = args[0]
			} else {
				defaultDoc := doc.GetDefaultDoc(docs)
				if defaultDoc == nil {
					return fmt.Errorf("no default documentation found; " +
						"please specify a documentation name")
				}

				docName = defaultDoc.Name
			}

			// Get the documentation path.
			docToOpen := doc.GetDocByName(docs, docName)
			if docToOpen == nil {
				return fmt.Errorf("documentation '%s' not found", docName)
			}

			return doc.OpenDoc(
				appCtx,
				filepath.Join(appCtx.DocsDir, docToOpen.Path),
			)
		},
	}

	return cmd
}
