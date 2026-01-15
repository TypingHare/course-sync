package exec

import (
	"fmt"
	"path/filepath"

	"github.com/TypingHare/course-sync/internal/support/filesystem"
	"github.com/TypingHare/course-sync/internal/support/io"
)

func ShellDeleteFile(
	outputMode *io.OutputMode,
	projectDir string,
	absFile string,
) error {
	relFile, err := filepath.Rel(projectDir, absFile)
	if err != nil {
		return fmt.Errorf("get relative path: %w", err)
	}

	return NewCommandRunner(
		outputMode,
		[]string{"rm", "-f", absFile},
		fmt.Sprintf("Deleting file %q...", relFile),
		fmt.Sprintf("Deleted file %q.", relFile),
		fmt.Sprintf("Failed to delete file %q.", relFile),
	).StartE()
}

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
		fmt.Sprintf("Ensuring directory  %q...", relDir),
		fmt.Sprintf("Ensured directory %q.", relDir),
		fmt.Sprintf("Failed to ensure directory %q.", relDir),
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
	relSrcDir := filesystem.RelOrOriginal(projectDir, absSrcDir)
	relDestDir := filesystem.RelOrOriginal(projectDir, absDestDir)

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

func ShellCopyFile(
	outputMode *io.OutputMode,
	projectDir string,
	absSrcFile string,
	absDestFile string,
) error {
	relSrcFile := filesystem.RelOrOriginal(projectDir, absSrcFile)
	relDestFile := filesystem.RelOrOriginal(projectDir, absDestFile)

	return NewCommandRunner(
		outputMode,
		[]string{"cp", absSrcFile, absDestFile},
		fmt.Sprintf(
			"Copying file from %q to %q...",
			relSrcFile,
			relDestFile,
		),
		fmt.Sprintf("Copied file from %q to %q.", relSrcFile, relDestFile),
		fmt.Sprintf(
			"Failed to copy file from %q to %q.",
			relSrcFile,
			relDestFile,
		),
	).StartE()
}
