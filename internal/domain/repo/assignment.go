package repo

import "github.com/TypingHare/course-sync/internal/domain/model"

type IAssignmentRepo interface {
	// GetAll retrieves all assignments from the data source.
	GetAll() ([]model.Assignment, error)

	// SaveAll saves multiple assignments to the data source.
	SaveAll(assignments []model.Assignment) error
}
