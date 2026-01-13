package app

import (
	"fmt"
	"path/filepath"

	"github.com/TypingHare/course-sync/internal/adapter/repo/jsonstore"
	"github.com/TypingHare/course-sync/internal/domain/model"
	"github.com/TypingHare/course-sync/internal/domain/service"
	"github.com/TypingHare/course-sync/internal/support/io"
)

func GetAssignmentService(
	assignmentDataFile string,
) *service.AssignmentService {
	return service.NewAssignmentService(
		jsonstore.NewAssignmentRepo(assignmentDataFile),
	)
}

// GetAllAssignment retrieves all assignments from the JSON store.
func GetAllAssignment(dataDir string) ([]model.Assignment, error) {
	return jsonstore.NewAssignmentRepo(GetAssignmentDataFile(dataDir)).GetAll()
}

// GetUserAssignmentDir returns the path to the user assignment directory where
// the files for the specified assignment are stored.
func GetUserAssignmentDir(
	outputMode io.OutputMode,
	srcDir string,
	assignmentName string,
) (string, error) {
	userWorkspaceDir, err := GetUserWorkspaceDir(outputMode, srcDir)
	if err != nil {
		return "", fmt.Errorf("get user assignment dir: %w", err)
	}

	return filepath.Join(userWorkspaceDir, assignmentName), nil
}
