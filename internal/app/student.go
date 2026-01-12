package app

import (
	"fmt"
	"path/filepath"

	"github.com/TypingHare/course-sync/internal/domain/student"
	"github.com/TypingHare/course-sync/internal/infra/fs"
)

func GetStudentRepoDirs(appCtx *Context) ([]string, error) {
	config, err := appCtx.GetConfig()
	if err != nil {
		return nil, fmt.Errorf("get config: %w", err)
	}

	students := config.Roster
	var studentRepoDirs []string
	for i := range students {
		studentRepoDirs = append(
			studentRepoDirs,
			filepath.Join(
				appCtx.ProjectDir,
				student.GetStudentRepoDir(students[i].Name),
			),
		)
	}

	return studentRepoDirs, nil
}

func EnsureStudentRepos(appCtx *Context) error {
	studentRepoDirs, err := GetStudentRepoDirs(appCtx)
	if err != nil {
		return fmt.Errorf("get student repo dirs: %w", err)
	}

	for i := range studentRepoDirs {
		studentRepoExists, err := fs.DirExists(studentRepoDirs[i])
		if err != nil {
			return fmt.Errorf("check if student repo dir exists: %w", err)
		}

		if !studentRepoExists {
			// err = git
		}
	}

	return nil
}

// MoveFileToStudentRepos copies a file from the given absolute path to each
// student's repository, preserving the relative path within the project.
func MoveFileToStudentRepos(appCtx *Context, absPath string) error {
	config, err := appCtx.GetConfig()
	if err != nil {
		return fmt.Errorf("get config: %w", err)
	}

	students := config.Roster
	var studentRepoDirs []string
	for i := range students {
		studentRepoDirs = append(
			studentRepoDirs,
			filepath.Join(
				appCtx.ProjectDir,
				student.GetStudentRepoDir(students[i].Name),
			),
		)
	}

	relPath, err := appCtx.GetRelPath(absPath)
	if err != nil {
		return fmt.Errorf("get relative path: %w", err)
	}

	for i := range studentRepoDirs {
		destPath := filepath.Join(studentRepoDirs[i], relPath)
		err = fs.CopyFile(absPath, destPath)
	}

	return nil
}
