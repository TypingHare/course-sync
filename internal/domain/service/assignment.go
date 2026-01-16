package service

import (
	"fmt"

	"github.com/TypingHare/course-sync/internal/domain/model"
	"github.com/TypingHare/course-sync/internal/domain/repo"
)

type AssignmentService struct {
	repo repo.IAssignmentRepo
}

// NewAssignmentService constructs an AssignmentService with the provided repo.
func NewAssignmentService(repo repo.IAssignmentRepo) *AssignmentService {
	return &AssignmentService{
		repo: repo,
	}
}

// GetAllAssignments retrieves all assignments from the repository.
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

	for i := range assignments {
		if assignments[i].Name == assignment.Name {
			return fmt.Errorf(
				"assignment with name %q already exists",
				assignment.Name,
			)
		}
	}

	return s.repo.SaveAll(append(assignments, *assignment))
}

// GetAssignmentByName finds an assignment by name.
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
