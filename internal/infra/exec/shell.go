package exec

import (
	"fmt"

	"github.com/TypingHare/course-sync/internal/app"
)

// ShellDeleteDir deletes the specified directory using a shell command.
func ShellDeleteDir(appCtx *app.Context, absDir string) error {
	relDir, err := appCtx.GetRelPath(absDir)
	if err != nil {
		return fmt.Errorf("get relative path: %w", err)
	}

	commandTask := NewCommandTask(
		appCtx,
		[]string{"rm", "-rf", absDir},
		fmt.Sprintf("Deleting directory %q...", relDir),
		fmt.Sprintf("Deleted directory %q.", relDir),
		fmt.Sprintf("Failed to delete directory %q.", relDir),
	)

	_, err = commandTask.Start()

	return err
}

// ShellEnsureDir creates the specified directory using a shell command.
func ShellEnsureDir(appCtx *app.Context, absDir string) error {
	relDir, err := appCtx.GetRelPath(absDir)
	if err != nil {
		return fmt.Errorf("get relative path: %w", err)
	}

	commandTask := NewCommandTask(
		appCtx,
		[]string{"mkdir", "-p", absDir},
		fmt.Sprintf("Creating directory %q...", relDir),
		fmt.Sprintf("Created directory %q.", relDir),
		fmt.Sprintf("Failed to create directory %q.", relDir),
	)

	_, err = commandTask.Start()

	return err
}

func ShellCopyDir(
	appCtx *app.Context,
	absSrcDir string,
	absDestDir string,
) error {
	relSrcDir, err := appCtx.GetRelPath(absSrcDir)
	if err != nil {
		return fmt.Errorf("get relative source path: %w", err)
	}

	relDestDir, err := appCtx.GetRelPath(absDestDir)
	if err != nil {
		return fmt.Errorf("get relative destination path: %w", err)
	}

	commandTask := NewCommandTask(
		appCtx,
		[]string{"cp", "-r", absSrcDir, absDestDir},
		fmt.Sprintf(
			"Copying directory from %q to %q...",
			relSrcDir,
			relDestDir,
		),
		fmt.Sprintf("Copied directory from %q to %q.", relSrcDir, relDestDir),
		fmt.Sprintf(
			"Failed to copy directory from %q to %q.",
			relSrcDir,
			relDestDir,
		),
	)

	_, err = commandTask.Start()

	return err
}
