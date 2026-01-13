package app

import (
	"github.com/TypingHare/course-sync/internal/domain/model"
	"github.com/TypingHare/course-sync/internal/support/exec"
	"github.com/TypingHare/course-sync/internal/support/io"
)

func StudentCommit(outputMode *io.OutputMode, dataDir string) error {
	studentFiles := GetStudentFiles(dataDir)

	for _, file := range studentFiles {
		err := exec.GitAdd(outputMode, file)
		if err != nil {
			return err
		}
	}

	return exec.GitCommit(outputMode, "csync: update student files")
}

func InstructorCommit(outputMode *io.OutputMode, dataDir string) error {
	return nil
}

func Commit(role model.Role, outputMode *io.OutputMode, dataDir string) error {
	switch role {
	case model.RoleStudent:
		return StudentCommit(outputMode, dataDir)
	case model.RoleInstructor:
		return InstructorCommit(outputMode, dataDir)
	}

	return nil
}

func Pull(role model.Role, outputMode *io.OutputMode, dataDir string) error {
	return nil
}

func Push(role model.Role, outputMode *io.OutputMode, dataDir string) error {
	return nil
}
