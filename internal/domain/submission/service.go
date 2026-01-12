package submission

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/domain/assignment"
	"github.com/TypingHare/course-sync/internal/infra/git"
	"github.com/TypingHare/course-sync/internal/infra/hash"
	"github.com/TypingHare/course-sync/internal/infra/jsonstore"
)

// Submission represents a user's submission data.
const SubmissionsFileName = "submissions.json"

// GetSubmissions retrieves the list of submissions from the submissions JSON
// file in the application data directory.
func GetSubmissions(appCtx *app.Context) ([]Submission, error) {
	submissions, err := jsonstore.ReadJSONFile[[]Submission](
		filepath.Join(appCtx.AppDataDir, SubmissionsFileName),
	)
	if err != nil {
		return nil, err
	}

	return submissions, nil
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
	appCtx *app.Context,
	assignmentName string,
) (*Submission, error) {
	userAssignmentDir, err := assignment.GetUserAssignmentDir(
		appCtx,
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
	submissions, err := GetSubmissions(appCtx)
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

	if err = git.Add(appCtx, userAssignmentDir); err != nil {
		return nil, fmt.Errorf("git add submission: %w", err)
	}

	err = git.Commit(appCtx, "feat: student submission "+submissionHash)
	if err != nil {
		return nil, fmt.Errorf("git commit submission: %w", err)
	}

	// Get the latest commit hash.
	gitCommitHash, err := git.RevParseHead(appCtx)
	if err != nil {
		return nil, err
	}

	return &Submission{
		Hash:           submissionHash,
		GitCommitHash:  gitCommitHash,
		AssignmentName: assignmentName,
		SubmittedAt:    time.Now().UTC(),
	}, nil
}

// AddSubmissionToFile adds the given submission to the submissions JSON file.
func AddSubmissionToFile(appCtx *app.Context, submission Submission) error {
	submissions, err := GetSubmissions(appCtx)
	if err != nil {
		return fmt.Errorf("get submissions %w", err)
	}

	submissions = append(submissions, submission)

	return jsonstore.WriteJSONFile(
		filepath.Join(appCtx.AppDataDir, SubmissionsFileName),
		submissions,
	)
}
