package app

import (
	"path/filepath"
	"strings"

	"github.com/TypingHare/course-sync/internal/adapter/repo/jsonstore"
	"github.com/TypingHare/course-sync/internal/domain/service"
	"github.com/iancoleman/strcase"
)

const StudentDataFileName = "students.json"

func GetStudentDataFile(dataDir string) string {
	return filepath.Join(dataDir, StudentDataFileName)
}

func GetStudentService(
	studentDataFile string,
) *service.StudentService {
	return service.NewStudentService(
		jsonstore.NewStudentRepo(studentDataFile),
	)
}

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
