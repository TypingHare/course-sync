package jsonstore

import "github.com/TypingHare/course-sync/internal/domain/model"

type DocRepo struct {
	store *JsonStore[[]model.Doc]
}

func NewDocRepo(filePath string) *DocRepo {
	return &DocRepo{
		store: NewJsonStore[[]model.Doc](filePath),
	}
}

func (r *DocRepo) GetAll() ([]model.Doc, error) {
	return r.store.Read()
}

func (r *DocRepo) SaveAll(docs []model.Doc) error {
	return r.store.Write(docs)
}
