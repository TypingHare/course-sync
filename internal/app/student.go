package app

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/TypingHare/course-sync/internal/adapter/repo/jsonstore"
	"github.com/TypingHare/course-sync/internal/domain/model"
	"github.com/TypingHare/course-sync/internal/domain/service"
	"github.com/TypingHare/course-sync/internal/support/exec"
	"github.com/TypingHare/course-sync/internal/support/io"
	"github.com/iancoleman/strcase"
)

const StudentDataFileName = "students.json"

// StudentRepoBaseDir is the base directory for local student repositories in
// the master repository
const StudentRepoBaseDir = "student-repos/"

// GetStudentDataFile returns the students data file path.
func GetStudentDataFile(dataDir string) string {
	return filepath.Join(dataDir, StudentDataFileName)
}

// GetStudentRepoDir constructs the local path for a student's repository
// based on the base directory and the student's directory name.
func GetStudentRepoDir(projectDir string, studentName string) string {
	return filepath.Join(
		projectDir,
		StudentRepoBaseDir,
		GetStudentDirName(studentName),
	)
}

// GetStudentService constructs a StudentService backed by the data file.
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
func GetStudentFiles(projectDir string) []string {
	dataDir := GetDataDir(projectDir)

	return []string{
		filepath.Join(dataDir, SubmissionDataFileName),
	}
}

// DistributeFileToStudentRepos copies a file into each student repository.
func DistributeFileToStudentRepos(
	outputMode *io.OutputMode,
	students []model.Student,
	projectDir string,
	relPath string,
) error {
	var studentRepoDirs []string
	for i := range students {
		studentRepoDirs = append(
			studentRepoDirs,
			GetStudentRepoDir(projectDir, students[i].Name),
		)
	}

	srcAbsPath := filepath.Join(projectDir, relPath)

	for i := range studentRepoDirs {
		destAbsPath := filepath.Join(studentRepoDirs[i], relPath)
		err := exec.ShellEnsureDir(
			outputMode,
			projectDir,
			filepath.Dir(destAbsPath),
		)
		if err != nil {
			return fmt.Errorf("ensure dir for %s: %w", destAbsPath, err)
		}

		err = exec.ShellCopyFile(
			outputMode,
			projectDir,
			srcAbsPath,
			destAbsPath,
		)
		if err != nil {
			return fmt.Errorf("copy file to %s: %w", destAbsPath, err)
		}
	}

	return nil
}

// DistributeDirToStudentRepos replaces a directory in each student repository.
func DistributeDirToStudentRepos(
	outputMode *io.OutputMode,
	students []model.Student,
	projectDir string,
	relPath string,
) error {
	var studentRepoDirs []string
	for i := range students {
		studentRepoDirs = append(
			studentRepoDirs,
			GetStudentRepoDir(projectDir, students[i].Name),
		)
	}

	srcAbsPath := filepath.Join(projectDir, relPath)

	for i := range studentRepoDirs {
		destAbsPath := filepath.Join(studentRepoDirs[i], relPath)
		err := exec.ShellDeleteDir(
			outputMode,
			projectDir,
			destAbsPath,
		)
		if err != nil {
			return fmt.Errorf("delete dir for %s: %w", destAbsPath, err)
		}

		err = exec.ShellEnsureDir(
			outputMode,
			projectDir,
			filepath.Dir(destAbsPath),
		)
		if err != nil {
			return fmt.Errorf("ensure dir for %s: %w", destAbsPath, err)
		}

		err = exec.ShellCopyDir(
			outputMode,
			projectDir,
			srcAbsPath,
			destAbsPath,
		)
		if err != nil {
			return fmt.Errorf("copy dir to %s: %w", destAbsPath, err)
		}
	}

	return nil
}
