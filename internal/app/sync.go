package app

import (
	"github.com/TypingHare/course-sync/internal/domain/model"
	"github.com/TypingHare/course-sync/internal/support/exec"
	"github.com/TypingHare/course-sync/internal/support/io"
)

// StudentCommit stages student files and commits them to the repository.
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

// InstructorCommit stages instructor files and commits them to the repository.
func InstructorCommit(outputMode *io.OutputMode, dataDir string) error {
	return nil
}

// Commit commits changes for the given role.
func Commit(role model.Role, outputMode *io.OutputMode, dataDir string) error {
	switch role {
	case model.RoleStudent:
		return StudentCommit(outputMode, dataDir)
	case model.RoleInstructor:
		return InstructorCommit(outputMode, dataDir)
	}

	return nil
}

// Pull updates the local repository from the remote.
func Pull(role model.Role, outputMode *io.OutputMode, dataDir string) error {
	return exec.GitPull(outputMode, true)
}

// Push pushes local commits to the remote.
func Push(role model.Role, outputMode *io.OutputMode, dataDir string) error {
	return exec.GitPush(outputMode)
}
