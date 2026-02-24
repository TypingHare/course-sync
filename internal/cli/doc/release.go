package doc

import (
	"fmt"
	"strings"
	"time"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/domain/model"
	"github.com/spf13/cobra"
)

// releaseCmd builds the doc release subcommand.
func releaseCmd(ctx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "release <doc-name> [path] [title] [version]",
		Short: "Release a documentation file.",
		Long: strings.TrimSpace(`
Release a documentation file.

This command creates or updates a document record and distributes the document
file to students' repositories. Instructors should use "sync" command to
commit all the changes.
        `),
		Args: cobra.RangeArgs(1, 4),
		RunE: func(cmd *cobra.Command, args []string) error {
			docName := args[0]

			docService := app.GetDocService(
				app.GetDocDataFile(app.GetDataDir(ctx.ProjectDir)),
			)
			existingDoc, err := docService.GetDocByName(docName)
			if err != nil {
				return fmt.Errorf("failed to get doc by name: %w", err)
			}

			now := time.Now().UTC()

			path := ""
			title := ""
			version := ""
			releasedAt := now
			isDefault := false

			if existingDoc != nil {
				path = existingDoc.Path
				title = existingDoc.Title
				version = existingDoc.Version
				releasedAt = existingDoc.ReleasedAt
				isDefault = existingDoc.IsDefault
			}

			if len(args) > 1 {
				path = args[1]
			}
			if len(args) > 2 {
				title = args[2]
			}
			if len(args) > 3 {
				version = args[3]
			}

			if path == "" {
				return fmt.Errorf(
					"path is required when releasing new documentation %q",
					docName,
				)
			}

			newDoc := &model.Doc{
				Name:       docName,
				Path:       path,
				Title:      title,
				Version:    version,
				ReleasedAt: releasedAt,
				UpdatedAt:  now,
				IsDefault:  isDefault,
			}
			err = app.ReleaseDoc(ctx.OutputMode, ctx.ProjectDir, newDoc)
			if err != nil {
				return fmt.Errorf("failed to release documentation: %w", err)
			}

			return nil
		},
	}

	return cmd
}
