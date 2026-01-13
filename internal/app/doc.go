package app

import (
	"path/filepath"

	"github.com/TypingHare/course-sync/internal/adapter/repo/jsonstore"
	"github.com/TypingHare/course-sync/internal/domain/service"
)

const DocDataFileName = "docs.json"

func GetDocDataFile(dataDir string) string {
	return filepath.Join(dataDir, DocDataFileName)
}

func GetDocService(docDataFile string) *service.DocService {
	return service.NewDocService(jsonstore.NewDocRepo(docDataFile))
}

func GetDocFilePath(projectDir string, relPath string) string {
	return filepath.Join(GetDocsDir(projectDir), relPath)
}
