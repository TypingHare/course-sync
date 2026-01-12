package cli

import (
	"path/filepath"
	"strings"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/domain/assignment"
	"github.com/TypingHare/course-sync/internal/domain/doc"
	"github.com/TypingHare/course-sync/internal/domain/grade"
	"github.com/TypingHare/course-sync/internal/domain/submission"
	"github.com/TypingHare/course-sync/internal/domain/workspace"
	"github.com/spf13/cobra"
)

func syncCommand(appCtx *app.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sync",
		Short: "Synchronize local and remote repositories",
		Long: strings.TrimSpace(`
Synchronize the local repository with the remote repository.

This command pulls the latest changes from the remote repository, commits local
changes, and then pushes the resulting commit back to the remote. It is
equivalent to running the pull, commit, and push commands in sequence.
        `),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := pull(appCtx)
			if err != nil {
				return err
			}

			err = commit(appCtx)
			if err != nil {
				return nil
			}

			return push(appCtx)
		},
	}

	return cmd
}

// STUDENT_FILES is a list of files that should always be committed by
// students.
func getStudentFiles(appCtx *app.Context) []string {
	return []string{
		filepath.Join(appCtx.AppDataDir, app.CONFIG_FILE_NAME),
		filepath.Join(appCtx.AppDataDir, submission.SUBMISSIONS_FILE_NAME),
	}
}

// INSTRUCTOR_FILES is a list of files that should always be committed by
// instructors.
func getInstructorFiles(appCtx *app.Context) []string {
	return []string{
		appCtx.DocsDir,
		filepath.Join(appCtx.SrcDir, workspace.PROTOTYPE_WORKSPACE),
		filepath.Join(appCtx.AppDataDir, app.INSTRUCTOR_PUBLIC_KEY_FILE_NAME),
		filepath.Join(appCtx.AppDataDir, doc.DOCS_FILE_NAME),
		filepath.Join(appCtx.AppDataDir, assignment.ASSIGNMENTS_FILE_NAME),
		filepath.Join(appCtx.AppDataDir, grade.GRADES_FILE_NAME),
	}
}
