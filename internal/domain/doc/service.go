package doc

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/infra/exec"
	"github.com/TypingHare/course-sync/internal/infra/jsonstore"
)

const DOCS_FILE_NAME = "docs.json"

// GetDocs reads and returns the documentation items from the docs JSON file
// file in the application data directory.
func GetDocs(appCtx *app.Context) ([]Doc, error) {
	docs, err := jsonstore.ReadJSONFile[[]Doc](
		filepath.Join(appCtx.AppDataDir, DOCS_FILE_NAME),
	)
	if err != nil {
		return nil, fmt.Errorf("retrieve docs: %w", err)
	}

	return docs, nil
}

// GetDefaultDoc returns the default documentation item from the provided slice.
// It returns nil if no default documentation item is found.
func GetDefaultDoc(docs []Doc) *Doc {
	for i := range docs {
		if docs[i].IsDefault {
			return &docs[i]
		}
	}

	return nil
}

// GetDocByName returns the documentation item with the specified name from the
// provided slice.
func GetDocByName(docs []Doc, name string) *Doc {
	for i := range docs {
		if docs[i].Name == name {
			return &docs[i]
		}
	}

	return nil
}

// OpenDoc opens the documentation file located at the given path. This function
// uses the default application associated with the file type on the user's
// operating system.
func OpenDoc(appCtx *app.Context, docAbsFile string) error {
	if docAbsFile == "" {
		return fmt.Errorf("documentation file path is empty")
	}

	var args []string
	switch runtime.GOOS {
	case "darwin":
		// macOS
		args = []string{"open", docAbsFile}

	case "windows":
		// Native Windows
		args = []string{"cmd", "/c", "start", "", docAbsFile}

	case "linux":
		// Linux & WSL
		// xdg-open works for most desktop environments and WSLg
		args = []string{"xdg-open", docAbsFile}

	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	docRelPath, err := appCtx.GetRelPath(docAbsFile)
	if err != nil {
		return err
	}

	commandTask := exec.NewCommandTask(
		appCtx,
		args,
		fmt.Sprintf("Opening documentation at %q...", docRelPath),
		fmt.Sprintf("Opened documentation at %q.", docRelPath),
		fmt.Sprintf("Failed to opened documentation at %q.", docRelPath),
	)

	_, err = commandTask.Start()

	return err
}
