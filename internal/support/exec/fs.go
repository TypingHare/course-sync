package exec

import (
	"fmt"
	"path/filepath"

	"github.com/TypingHare/course-sync/internal/support/io"
)

// ShellDeleteDir deletes the specified directory using a shell command.
func ShellDeleteDir(
	outputMode *io.OutputMode,
	projectDir string,
	absDir string,
) error {
	relDir, err := filepath.Rel(projectDir, absDir)
	if err != nil {
		return fmt.Errorf("get relative path: %w", err)
	}

	return NewCommandRunner(
		outputMode,
		[]string{"rm", "-rf", absDir},
		fmt.Sprintf("Deleting directory %q...", relDir),
		fmt.Sprintf("Deleted directory %q.", relDir),
		fmt.Sprintf("Failed to delete directory %q.", relDir),
	).StartE()
}

// ShellEnsureDir creates the specified directory using a shell command.
func ShellEnsureDir(
	outputMode *io.OutputMode,
	projectDir string,
	absDir string,
) error {
	relDir, err := filepath.Rel(projectDir, absDir)
	if err != nil {
		return fmt.Errorf("get relative path: %w", err)
	}

	return NewCommandRunner(
		outputMode,
		[]string{"mkdir", "-p", absDir},
		fmt.Sprintf("Creating directory %q...", relDir),
		fmt.Sprintf("Created directory %q.", relDir),
		fmt.Sprintf("Failed to create directory %q.", relDir),
	).StartE()
}

// ShellCopyDir copies a directory from source to destination using a shell
// command.
func ShellCopyDir(
	outputMode *io.OutputMode,
	projectDir string,
	absSrcDir string,
	absDestDir string,
) error {
	relSrcDir, err := filepath.Rel(projectDir, absSrcDir)
	if err != nil {
		return fmt.Errorf("get relative source path: %w", err)
	}

	relDestDir, err := filepath.Rel(projectDir, absDestDir)
	if err != nil {
		return fmt.Errorf("get relative destination path: %w", err)
	}

	return NewCommandRunner(
		outputMode,
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
	).StartE()
}
