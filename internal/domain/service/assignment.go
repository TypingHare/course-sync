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

func (s *AssignmentService) GetAllAssignments() ([]model.Assignment, error) {
	return s.repo.GetAll()
}

// AddAssignment adds a new assignment to the existing list of assignments and
// saves the updated list back to the repository.
func (s *AssignmentService) AddAssignment(assignment *model.Assignment) error {
	assignments, err := s.repo.GetAll()
	if err != nil {
		return err
	}

	return s.repo.SaveAll(append(assignments, *assignment))
}

func (s *AssignmentService) GetAssignmentByName(
	name string,
) (*model.Assignment, error) {
	assignments, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	for i := range assignments {
		if assignments[i].Name == name {
			return &assignments[i], nil
		}
	}

	return nil, nil
}
