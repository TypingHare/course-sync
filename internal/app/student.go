package app

import (
	"strings"

	"github.com/iancoleman/strcase"
)

// GetStudentDirName returns the directory name for a student by converting the
// student's name to kebab-case.
func GetStudentDirName(studentName string) string {
	return strings.ReplaceAll(strcase.ToSnake(studentName), "_", "-")
}
