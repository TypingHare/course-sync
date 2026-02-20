package exec

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/TypingHare/course-sync/internal/support/io"
)

// OpenFile opens the given file path using the OS default handler.
func OpenFile(
	outputMode *io.OutputMode,
	projectDir string,
	absPath string,
) error {
	if absPath == "" {
		return fmt.Errorf("file path is empty")
	}

	var args []string
	switch runtime.GOOS {
	case "darwin":
		// macOS
		args = []string{"open", absPath}

	case "windows":
		// Native Windows
		args = []string{"cmd", "/c", "start", "", absPath}

	case "linux":
		// Linux & WSL
		// xdg-open works for most desktop environments and WSLg
		args = []string{"xdg-open", absPath}

	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	relPath, err := filepath.Rel(projectDir, absPath)
	if err != nil {
		relPath = absPath
	}

	return NewCommandRunner(
		outputMode,
		args,
		fmt.Sprintf("Opening documentation at %q...", relPath),
		fmt.Sprintf("Opened documentation at %q.", relPath),
		fmt.Sprintf("Failed to open documentation at %q.", relPath),
	).StartE()
}
