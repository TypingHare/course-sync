package jsonstore

import (
	"github.com/TypingHare/course-sync/internal/domain/model"
)

type AssignmentRepo struct {
	store *JsonStore[[]model.Assignment]
}

// NewAssignmentRepo constructs an AssignmentRepo backed by a JSON store.
func NewAssignmentRepo(filePath string) *AssignmentRepo {
	return &AssignmentRepo{
		store: NewJsonStore[[]model.Assignment](filePath),
	}
}

// GetAll returns all assignments from the JSON store.
func (r *AssignmentRepo) GetAll() ([]model.Assignment, error) {
	return r.store.Read()
}

// SaveAll persists all assignments to the JSON store.
func (r *AssignmentRepo) SaveAll(assignments []model.Assignment) error {
	return r.store.Write(assignments)
}
