package app

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/TypingHare/course-sync/internal/domain/model"
	"github.com/TypingHare/course-sync/internal/support/exec"
	"github.com/TypingHare/course-sync/internal/support/io"
)

const InstructorSyncCommitMessage = "feat: synchronize files (Course Sync)"

// StudentCommit stages student files and commits them to the repository.
func StudentCommit(outputMode *io.OutputMode, dataDir string) error {
	projectDir := filepath.Dir(dataDir)
	studentFiles := GetStudentFiles(projectDir)

	for _, file := range studentFiles {
		if _, err := os.Stat(file); err != nil {
			if errors.Is(err, os.ErrNotExist) {
				continue
			}

			return fmt.Errorf("stat managed student file %q: %w", file, err)
		}

		err := exec.GitAdd(outputMode, file)
		if err != nil {
			return err
		}
	}

	return gitCommitInDir(outputMode, projectDir, "csync: update student files")
}

func getStudentsForInstructorSync(dataDir string) ([]model.Student, error) {
	studentService := GetStudentService(GetStudentDataFile(dataDir))
	students, err := studentService.GetAllStudents()
	if err != nil {
		return nil, fmt.Errorf("get students: %w", err)
	}

	return students, nil
}

func isNothingToCommit(stdout string, stderr string) bool {
	combined := strings.ToLower(stdout + "\n" + stderr)

	return strings.Contains(combined, "nothing to commit") ||
		strings.Contains(combined, "no changes added to commit")
}

func gitAddAllInDir(outputMode *io.OutputMode, repoDir string) error {
	return exec.NewCommandRunner(
		outputMode,
		[]string{"git", "add", "-A"},
		"Staging all changes for commit...",
		"Staged all changes for commit.",
		"Failed to stage all changes for commit.",
	).SetWorkingDir(repoDir).StartE()
}

func gitCommitInDir(
	outputMode *io.OutputMode,
	repoDir string,
	commitMessage string,
) error {
	runner := exec.NewCommandRunner(
		outputMode,
		[]string{"git", "commit", "-m", commitMessage},
		fmt.Sprintf("Committing changes (%s)...", commitMessage),
		fmt.Sprintf("Committed changes (%s).", commitMessage),
		fmt.Sprintf("Failed to commit changes (%s).", commitMessage),
	).SetWorkingDir(repoDir)

	result, err := runner.Start()
	if err != nil {
		if isNothingToCommit(result.Stdout, result.Stderr) {
			return nil
		}
		return err
	}

	return nil
}

func gitPushInDir(outputMode *io.OutputMode, repoDir string) error {
	return exec.NewCommandRunner(
		outputMode,
		[]string{"git", "push"},
		"Pushing changes to remote repository...",
		"Pushed changes to remote repository.",
		"Failed to push changes to remote repository.",
	).SetWorkingDir(repoDir).StartE()
}

// InstructorCommit stages instructor files and commits them to the repository.
func InstructorCommit(outputMode *io.OutputMode, dataDir string) error {
	projectDir := filepath.Dir(dataDir)

	students, err := getStudentsForInstructorSync(dataDir)
	if err != nil {
		return err
	}

	for i := range students {
		repoDir := GetStudentRepoDir(projectDir, students[i].Name)

		if err := gitAddAllInDir(outputMode, repoDir); err != nil {
			return fmt.Errorf(
				"stage student repo for %q: %w",
				students[i].Name,
				err,
			)
		}

		if err := gitCommitInDir(
			outputMode,
			repoDir,
			InstructorSyncCommitMessage,
		); err != nil {
			return fmt.Errorf(
				"commit student repo for %q: %w",
				students[i].Name,
				err,
			)
		}
	}

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
	switch role {
	case model.RoleStudent:
		return exec.GitPull(outputMode, true)
	case model.RoleInstructor:
		return exec.GitPull(outputMode, false)
	}

	return exec.GitPull(outputMode, false)
}

// Push pushes local commits to the remote.
func Push(role model.Role, outputMode *io.OutputMode, dataDir string) error {
	if role == model.RoleInstructor {
		projectDir := filepath.Dir(dataDir)

		students, err := getStudentsForInstructorSync(dataDir)
		if err != nil {
			return err
		}

		for i := range students {
			repoDir := GetStudentRepoDir(projectDir, students[i].Name)

			if err := gitPushInDir(outputMode, repoDir); err != nil {
				return fmt.Errorf(
					"push student repo for %q: %w",
					students[i].Name,
					err,
				)
			}
		}

		return nil
	}

	return exec.GitPush(outputMode)
}
