package app

import (
	"github.com/TypingHare/course-sync/internal/adapter/repo/jsonstore"
	"github.com/TypingHare/course-sync/internal/domain/model"
)

// GetAllAssignment retrieves all assignments from the JSON store.
func GetAllAssignment() ([]model.Assignment, error) {
	return jsonstore.NewAssignmentRepo(".csync/assignments.json").GetAll()
}
