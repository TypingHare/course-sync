package service

import (
	"github.com/TypingHare/course-sync/internal/domain/model"
	"github.com/TypingHare/course-sync/internal/domain/repo"
)

type AssignmentService struct {
	repo repo.IAssignmentRepo
}

func NewAssignmentService(repo repo.IAssignmentRepo) *AssignmentService {
	return &AssignmentService{
		repo: repo,
	}
}

// AddAssignment adds a new assignment to the existing list of assignments and
// saves the updated list back to the repository.
func (s *AssignmentService) AddAssignment(
	assignment *model.Assignment,
) error {
	assignments, err := s.repo.GetAll()
	if err != nil {
		return err
	}

	return s.repo.SaveAll(append(assignments, *assignment))
}
