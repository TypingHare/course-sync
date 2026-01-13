package jsonstore

import (
	"github.com/TypingHare/course-sync/internal/domain/model"
)

type AssignmentRepo struct {
	store *JsonStore[[]model.Assignment]
}

func NewAssignmentRepo(filePath string) *AssignmentRepo {
	return &AssignmentRepo{
		store: NewJsonStore[[]model.Assignment](filePath),
	}
}

func (r *AssignmentRepo) GetAll() ([]model.Assignment, error) {
	return r.store.Read()
}

func (r *AssignmentRepo) SaveAll(assignments []model.Assignment) error {
	return r.store.Write(assignments)
}
