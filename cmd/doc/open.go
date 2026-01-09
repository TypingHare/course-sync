package doc

import (
	"path/filepath"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/feature"
	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open [name]",
	Short: "Open a documentation file.",
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {
		docs, err := feature.GetDocs()
		if err != nil {
			cmd.PrintErrf("Error getting docs: %v\n", err)
		}

		// Determine which doc to open.
		var docName string
		if len(args) >= 1 {
			docName = args[0]
		} else {
			defaultDoc := feature.GetDefaultDoc(docs)
			if defaultDoc == nil {
				cmd.PrintErrln(
					"No default documentation found. Please specify a documentation name.",
				)
				return
			}
			docName = defaultDoc.Name
		}

		doc := feature.GetDocByName(docs, docName)
		if doc == nil {
			cmd.PrintErrf("Documentation with name '%s' not found.\n", docName)
		}

		docsDirPath, err := app.GetDocsDirPath()
		if err != nil {
			cmd.PrintErrf("Error getting docs directory path: %v\n", err)
		}

		err = feature.OpenDoc(filepath.Join(docsDirPath, doc.Path))
		if err != nil {
			cmd.PrintErrf("Error opening doc: %v\n", err)
		}
	},
}
