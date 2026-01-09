package feature

import (
	"errors"
	"fmt"
	"runtime"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/execx"
)

// Doc represents a documentation item with a title and path.
type Doc struct {
	Name      string `json:"name"`
	Title     string `json:"title"`
	Path      string `json:"path"`
	IsDefault bool   `json:"isDefault"`
}

// GetDocs reads and returns the documentation items from the docs.json file.
func GetDocs() ([]Doc, error) {
	docs, err := ReadJSONSlice[[]Doc](app.DOCS_FILE_NAME)
	if err != nil {
		return nil, err
	}

	return docs, nil
}

// OpenDoc opens the documentation file located at the given path (absolute path).
func OpenDoc(docPath string) error {
	if docPath == "" {
		return errors.New("doc path is empty")
	}

	var command string
	var args []string
	switch runtime.GOOS {
	case "darwin":
		// macOS
		command = "open"
		args = []string{docPath}

	case "windows":
		// Native Windows
		command = "cmd"
		args = []string{"/c", "start", "", docPath}

	case "linux":
		// Linux & WSL
		// xdg-open works for most desktop environments and WSLg
		command = "xdg-open"
		args = []string{docPath}

	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	relativeDocPath, err := app.GetRelativePath(docPath)
	if err != nil {
		return err
	}

	commandTask := execx.CommandTask{
		Command:        command,
		Args:           args,
		OngoingMessage: fmt.Sprintf("Opening documentation at <%s>...", relativeDocPath),
		DoneMessage:    fmt.Sprintf("Documentation opened at <%s>.", relativeDocPath),
		ErrorMessage:   fmt.Sprintf("Failed to open documentation at <%s>.", relativeDocPath),
		Quiet:          app.Quiet,
		PrintCommand:   app.Verbose,
		PrintStdout:    app.Verbose,
		PrintStderr:    app.Verbose,
	}

	return commandTask.Start()
}

// GetDefaultDoc returns the default documentation item from the provided slice.
func GetDefaultDoc(docs []Doc) *Doc {
	for _, doc := range docs {
		if doc.IsDefault {
			return &doc
		}
	}

	return nil
}

// GetDocByName returns the documentation item with the specified name from the provided slice.
func GetDocByName(docs []Doc, name string) *Doc {
	for _, doc := range docs {
		if doc.Name == name {
			return &doc
		}
	}

	return nil
}
