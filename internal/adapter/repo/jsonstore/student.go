package jsonstore

import "github.com/TypingHare/course-sync/internal/domain/model"

type StudentRepo struct {
	store *JsonStore[[]model.Student]
}

func NewStudentRepo(filePath string) *StudentRepo {
	return &StudentRepo{
		store: NewJsonStore[[]model.Student](filePath),
	}
}

func (r *StudentRepo) GetAll() ([]model.Student, error) {
	return r.store.Read()
}

func (r *StudentRepo) SaveAll(docs []model.Student) error {
	return r.store.Write(docs)
}
