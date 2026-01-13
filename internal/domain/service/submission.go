package service

import "github.com/TypingHare/course-sync/internal/domain/repo"

type SubmissionService struct {
	repo repo.ISubmissionRepo
}

func NewSubmissionService(r repo.ISubmissionRepo) *SubmissionService {
	return &SubmissionService{repo: r}
}
