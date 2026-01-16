package jsonstore

import "github.com/TypingHare/course-sync/internal/domain/model"

type DocRepo struct {
	store *JsonStore[[]model.Doc]
}

// NewDocRepo constructs a DocRepo backed by a JSON store.
func NewDocRepo(filePath string) *DocRepo {
	return &DocRepo{
		store: NewJsonStore[[]model.Doc](filePath),
	}
}

// GetAll returns all docs from the JSON store.
func (r *DocRepo) GetAll() ([]model.Doc, error) {
	return r.store.Read()
}

// SaveAll persists all docs to the JSON store.
func (r *DocRepo) SaveAll(docs []model.Doc) error {
	return r.store.Write(docs)
}
