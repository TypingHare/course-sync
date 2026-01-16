package app

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/TypingHare/course-sync/internal/adapter/repo/jsonstore"
	"github.com/TypingHare/course-sync/internal/domain/model"
	"github.com/TypingHare/course-sync/internal/domain/service"
	"github.com/TypingHare/course-sync/internal/support/exec"
	"github.com/TypingHare/course-sync/internal/support/hash"
	"github.com/TypingHare/course-sync/internal/support/io"
)

const SubmissionDataFileName = "submissions.json"

// GetSubmissionDataFile returns the submissions data file path.
func GetSubmissionDataFile(dataDir string) string {
	return filepath.Join(dataDir, SubmissionDataFileName)
}

// GetSubmissionService constructs a SubmissionService backed by the data file.
func GetSubmissionService(
	submissionDataFile string,
) *service.SubmissionService {
	return service.NewSubmissionService(
		jsonstore.NewSubmissionRepo(submissionDataFile),
	)
}

// CreateHashForUserAssignmentDir generates a hash for the contents of the
// specified user assignment directory, ignoring certain files and directories.
func CreateHashForUserAssignmentDir(userAssignmentDir string) (string, error) {
	return hash.CreateHashForDir(
		userAssignmentDir,
		[]string{"__pycache__", ".DS_Store"},
	)
}

// CreateSubmission creates a new submission for the given assignment. It stages
// and commits the submission in Git. It returns the created Submission object.
func CreateSubmission(
	outputMode *io.OutputMode,
	dataDir string,
	srcDir string,
	assignmentName string,
) (*model.Submission, error) {
	userAssignmentDir, err := GetUserAssignmentDir(
		outputMode, srcDir,
		assignmentName,
	)
	if err != nil {
		return nil, fmt.Errorf("get user assignment dir path: %w", err)
	}

	submissionHash, err := CreateHashForUserAssignmentDir(userAssignmentDir)
	if err != nil {
		return nil, fmt.Errorf("create hash for submission: %w", err)
	}

	// Get submissions.
	submissionService := GetSubmissionService(
		GetSubmissionDataFile(dataDir),
	)
	submissions, err := submissionService.GetAllSubmissions()
	if err != nil {
		return nil, fmt.Errorf("get submissions: %w", err)
	}

	// Check for duplicate submission.
	for _, sub := range submissions {
		if sub.Hash == submissionHash && sub.AssignmentName == assignmentName {
			return nil, fmt.Errorf(
				"duplicate submission detected (hash: %s)",
				submissionHash,
			)
		}
	}

	if err = exec.GitAdd(outputMode, userAssignmentDir); err != nil {
		return nil, fmt.Errorf("git add submission: %w", err)
	}

	err = exec.GitCommit(outputMode, "feat: student submission "+submissionHash)
	if err != nil {
		return nil, fmt.Errorf("git commit submission: %w", err)
	}

	// Get the latest commit hash.
	gitCommitHash, err := exec.GitRevParseHead(outputMode)
	if err != nil {
		return nil, err
	}

	return &model.Submission{
		Hash:           submissionHash,
		GitCommitHash:  gitCommitHash,
		AssignmentName: assignmentName,
		SubmittedAt:    time.Now().UTC(),
	}, nil
}
