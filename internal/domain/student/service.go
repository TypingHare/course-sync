package student

import (
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"
)

// StudentRepoBaseDir is the base directory for local student repositories.
const StudentRepoBaseDir = "student-repos/"

// NewStudent creates a new Student instance.
func NewStudent(
	ID int,
	name string,
	email string,
	repositoryURL string,
) *Student {
	return &Student{
		ID:            ID,
		Name:          name,
		Email:         email,
		RepositoryURL: repositoryURL,
	}
}

// GetStudentDirName returns the directory name for a student by converting the
// student's name to kebab-case.
func GetStudentDirName(studentName string) string {
	return strings.ReplaceAll(strcase.ToSnake(studentName), "_", "-")
}

// GetStudentRepoDir constructs the local path for a student's repository
// based on the base directory and the student's directory name.
func GetStudentRepoDir(studentName string) string {
	return filepath.Join(StudentRepoBaseDir, GetStudentDirName(studentName))
}
