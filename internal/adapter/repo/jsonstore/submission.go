package jsonstore

import "github.com/TypingHare/course-sync/internal/domain/model"

type SubmissionRepo struct {
	store *JsonStore[[]model.Submission]
}

// NewSubmissionRepo constructs a SubmissionRepo backed by a JSON store.
func NewSubmissionRepo(filePath string) *SubmissionRepo {
	return &SubmissionRepo{
		store: NewJsonStore[[]model.Submission](filePath),
	}
}

// GetAll returns all submissions from the JSON store.
func (r *SubmissionRepo) GetAll() ([]model.Submission, error) {
	return r.store.Read()
}

// SaveAll persists all submissions to the JSON store.
func (r *SubmissionRepo) SaveAll(docs []model.Submission) error {
	return r.store.Write(docs)
}
