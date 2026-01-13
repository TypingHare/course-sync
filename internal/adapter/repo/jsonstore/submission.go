package jsonstore

import "github.com/TypingHare/course-sync/internal/domain/model"

type SubmissionRepo struct {
	store *JsonStore[[]model.Submission]
}

func NewSubmissionRepo(filePath string) *SubmissionRepo {
	return &SubmissionRepo{
		store: NewJsonStore[[]model.Submission](filePath),
	}
}

func (r *SubmissionRepo) GetAll() ([]model.Submission, error) {
	return r.store.Read()
}

func (r *SubmissionRepo) SaveAll(docs []model.Submission) error {
	return r.store.Write(docs)
}
