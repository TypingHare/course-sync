package jsonstore

import "github.com/TypingHare/course-sync/internal/domain/model"

type StudentRepo struct {
	store *JsonStore[[]model.Student]
}

// NewStudentRepo constructs a StudentRepo backed by a JSON store.
func NewStudentRepo(filePath string) *StudentRepo {
	return &StudentRepo{
		store: NewJsonStore[[]model.Student](filePath),
	}
}

// GetAll returns all students from the JSON store.
func (r *StudentRepo) GetAll() ([]model.Student, error) {
	return r.store.Read()
}

// SaveAll persists all students to the JSON store.
func (r *StudentRepo) SaveAll(docs []model.Student) error {
	return r.store.Write(docs)
}
