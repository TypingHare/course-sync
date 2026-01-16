package jsonstore

import "github.com/TypingHare/course-sync/internal/domain/model"

type GradeRepo struct {
	store *JsonStore[[]model.Grade]
}

// NewGradeRepo constructs a GradeRepo backed by a JSON store.
func NewGradeRepo(filePath string) *GradeRepo {
	return &GradeRepo{
		store: NewJsonStore[[]model.Grade](filePath),
	}
}

// GetAll returns all grades from the JSON store.
func (r *GradeRepo) GetAll() ([]model.Grade, error) {
	return r.store.Read()
}

// SaveAll persists all grades to the JSON store.
func (r *GradeRepo) SaveAll(docs []model.Grade) error {
	return r.store.Write(docs)
}
