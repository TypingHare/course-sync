package app

import (
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"
)

// GetStudentDirName returns the directory name for a student by converting the
// student's name to kebab-case.
func GetStudentDirName(studentName string) string {
	return strings.ReplaceAll(strcase.ToSnake(studentName), "_", "-")
}

// GetStudentFiles returns a list of file paths related to student data.
func GetStudentFiles(dataDir string) []string {
	return []string{
		filepath.Join(dataDir, SubmissionDataFileName),
	}
}
