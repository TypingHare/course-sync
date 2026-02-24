package app

import (
	"fmt"
	"path/filepath"

	"github.com/TypingHare/course-sync/internal/adapter/repo/jsonstore"
	"github.com/TypingHare/course-sync/internal/domain/model"
	"github.com/TypingHare/course-sync/internal/domain/service"
	"github.com/TypingHare/course-sync/internal/support/io"
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

// ReleaseDoc creates a doc record and distribute files and directories to
// students' repositories.
func ReleaseDoc(
	outputMode *io.OutputMode,
	projectDir string,
	newDoc *model.Doc,
) error {
	docDataFile := GetDocDataFile(GetDataDir(projectDir))
	docService := GetDocService(docDataFile)

	// Add or update the doc in the doc data file.
	err := docService.UpsertDoc(newDoc)
	if err != nil {
		return fmt.Errorf("upsert doc: %w", err)
	}

	// Get the document path.
	docToRelease, err := docService.GetDocByName(newDoc.Name)
	if err != nil {
		return fmt.Errorf("failed to get doc by name: %w", err)
	}
	if docToRelease == nil {
		return fmt.Errorf("documentation '%s' not found", newDoc.Name)
	}

	// Get all students.
	studentService := GetStudentService(
		GetStudentDataFile(GetDataDir(projectDir)),
	)
	students, err := studentService.GetAllStudents()
	if err != nil {
		return fmt.Errorf("get students: %w", err)
	}

	// Distribute the document data file to students' repositories.
	docDataFileRel, err := filepath.Rel(projectDir, docDataFile)
	if err != nil {
		return fmt.Errorf(
			"get relative path of doc data file: %w",
			err,
		)
	}

	err = DistributeFileToStudentRepos(
		outputMode,
		students,
		projectDir,
		docDataFileRel,
	)
	if err != nil {
		return fmt.Errorf(
			"distribute doc data file to student repositories: %w",
			err,
		)
	}

	// Distribute the document files to students' repositories.
	docFilePath := docToRelease.Path
	docFileRelPath, err := filepath.Rel(
		projectDir,
		GetDocFilePath(projectDir, docFilePath),
	)
	if err != nil {
		return fmt.Errorf(
			"get relative path of doc file: %w",
			err,
		)
	}

	err = DistributeFileToStudentRepos(
		outputMode,
		students,
		projectDir,
		docFileRelPath,
	)
	if err != nil {
		return fmt.Errorf(
			"distribute doc file to student repositories: %w",
			err,
		)
	}

	return nil
}
