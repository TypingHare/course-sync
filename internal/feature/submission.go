package feature

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/TypingHare/course-sync/internal/app"
)

type Submission struct {
	Hash           string
	AssignmentName string
	SubmittedAt    time.Time
}

// GetSubmissionHistory retrieves the submission history from the assignments file.
func GetSubmissionHistory() ([]Submission, error) {
	submissions, err := ReadJSONSlice[[]Submission](app.GRADE_HISTORY_FILE_NAME)
	if err != nil {
		return nil, err
	}

	return submissions, nil
}

// CollectSubmittedFilePaths collects the files in a user assignment directory.
func CollectSubmittedFilePaths(userAssignmentDirPath string) ([]string, error) {
	filePaths, err := CollectFiles(userAssignmentDirPath)
	if err != nil {
		return nil, fmt.Errorf("collect files: %w", err)
	}

	return filePaths, nil
}

// CreateHashForSubmission creates a hash for the submission based on the submitted files. The
// returned hash string is the last 8 characters of the SHA-256 hash.
func CreateHashForSubmission(assignmentName string, userAssignmentDirPath string) (string, error) {
	submittedFilePaths, err := CollectSubmittedFilePaths(userAssignmentDirPath)
	if err != nil {
		return "", fmt.Errorf("collect submitted file paths: %w", err)
	}

	// Make deterministic by sorting file paths.
	sort.Strings(submittedFilePaths)

	_hash := sha256.New()

	for _, submittedFilePath := range submittedFilePaths {
		// Only hash regular files (skip symlinks, devices, etc.)
		info, err := os.Lstat(submittedFilePath)
		if err != nil {
			return "", fmt.Errorf("stat %q: %w", submittedFilePath, err)
		}
		if !info.Mode().IsRegular() {
			continue
		}

		relativePath, err := filepath.Rel(userAssignmentDirPath, submittedFilePath)
		if err != nil {
			return "", fmt.Errorf("relative path for %q: %w", submittedFilePath, err)
		}

		// Include path in hash so renames/moves change the submission hash
		if err := writeString(_hash, relativePath); err != nil {
			return "", fmt.Errorf("hash relative path %q: %w", relativePath, err)
		}

		// Include contents
		if err := hashFileContents(_hash, submittedFilePath); err != nil {
			return "", fmt.Errorf("hash file %q: %w", submittedFilePath, err)
		}
	}

	sum := _hash.Sum(nil)
	hashString := hex.EncodeToString(sum)

	return hashString[len(hashString)-8:], nil
}

// writeString writes a string to the hash with a separator.
func writeString(_hash hash.Hash, str string) error {
	_, err := io.WriteString(_hash, str)
	if err != nil {
		return err
	}
	_, err = io.WriteString(_hash, "\n")

	return err
}

// hashFileContents reads the file at the given path and writes its contents to the hash.
func hashFileContents(_hash hash.Hash, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(_hash, file)
	if err == nil {
		_, err = io.WriteString(_hash, "\n")
	}

	return err
}

// CreateSubmission creates a new submission for the given assignment. It stages and commits the
// submission in Git. It returns the created Submission object.
func CreateSubmission(assignmentName string) (*Submission, error) {
	userAssignmentDirPath, err := GetUserAssignmentDirPath(assignmentName)
	if err != nil {
		return nil, fmt.Errorf("get user assignment dir path: %w", err)
	}

	submissionHash, err := CreateHashForSubmission(assignmentName, userAssignmentDirPath)
	if err != nil {
		return nil, fmt.Errorf("create hash for submission: %w", err)
	}

	err = app.RunAll(
		func() error { return GitAdd(userAssignmentDirPath) },
		func() error { return GitCommit("feat: student submission " + submissionHash) },
	)
	if err != nil {
		return nil, fmt.Errorf("git add and commit: %w", err)
	}

	return &Submission{
		Hash:           submissionHash,
		AssignmentName: assignmentName,
		SubmittedAt:    time.Now(),
	}, nil
}

// AppendSubmissionToHistory appends a submission to the submission history file.
func AppendSubmissionToHistory(submission Submission) error {
	submissionHistory, err := GetSubmissionHistory()
	if err != nil {
		return fmt.Errorf("get submission history: %w", err)
	}

	submissionHistory = append(submissionHistory, submission)

	return WriteJSONSlice(app.GRADE_HISTORY_FILE_NAME, submissionHistory)
}
