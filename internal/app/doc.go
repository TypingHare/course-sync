package app

import (
	"path/filepath"

	"github.com/TypingHare/course-sync/internal/adapter/repo/jsonstore"
	"github.com/TypingHare/course-sync/internal/domain/service"
)

const DocDataFileName = "docs.json"

// GetDocDataFile returns the docs data file path under the given data
// directory.
func GetDocDataFile(dataDir string) string {
	return filepath.Join(dataDir, DocDataFileName)
}

// GetDocService constructs a DocService backed by the given docs data file.
func GetDocService(docDataFile string) *service.DocService {
	return service.NewDocService(jsonstore.NewDocRepo(docDataFile))
}

// GetDocFilePath returns the full path to a doc file under the project docs
// directory.
func GetDocFilePath(projectDir string, relPath string) string {
	return filepath.Join(GetDocsDir(projectDir), relPath)
}
