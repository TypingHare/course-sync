package jsonstore

import "github.com/TypingHare/course-sync/internal/domain/model"

type GradeRepo struct {
	store *JsonStore[[]model.Grade]
}

func NewGradeRepo(filePath string) *GradeRepo {
	return &GradeRepo{
		store: NewJsonStore[[]model.Grade](filePath),
	}
}

func (r *GradeRepo) GetAll() ([]model.Grade, error) {
	return r.store.Read()
}

func (r *GradeRepo) SaveAll(docs []model.Grade) error {
	return r.store.Write(docs)
}
