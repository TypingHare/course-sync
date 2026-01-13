package service

import (
	"fmt"

	"github.com/TypingHare/course-sync/internal/domain/model"
	"github.com/TypingHare/course-sync/internal/domain/repo"
)

type SubmissionService struct {
	repo repo.ISubmissionRepo
}

func NewSubmissionService(r repo.ISubmissionRepo) *SubmissionService {
	return &SubmissionService{repo: r}
}

func (s *SubmissionService) GetAllSubmissions() ([]model.Submission, error) {
	return s.repo.GetAll()
}

func (s *SubmissionService) AddSubmission(
	submission *model.Submission,
) error {
	submissions, err := s.repo.GetAll()
	if err != nil {
		return fmt.Errorf("get submissions: %w", err)
	}

	return s.repo.SaveAll(append(submissions, *submission))
}
