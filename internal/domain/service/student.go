package service

import (
	"fmt"

	"github.com/TypingHare/course-sync/internal/domain/model"
	"github.com/TypingHare/course-sync/internal/domain/repo"
)

type StudentService struct {
	repo repo.IStudentRepo
}

// NewStudentService constructs a StudentService with the provided repo.
func NewStudentService(studentRepo repo.IStudentRepo) *StudentService {
	return &StudentService{
		repo: studentRepo,
	}
}

// GetAllStudents retrieves all students from the repository.
func (s *StudentService) GetAllStudents() ([]model.Student, error) {
	return s.repo.GetAll()
}

// AddStudent appends a new student to the repository.
func (s *StudentService) AddStudent(
	student *model.Student,
) error {
	students, err := s.repo.GetAll()
	if err != nil {
		return fmt.Errorf("get all students: %w", err)
	}

	return s.repo.SaveAll(append(students, *student))
}

// GetNextStudentID returns the next available student ID.
func (s *StudentService) GetNextStudentID() (int, error) {
	students, err := s.repo.GetAll()
	if err != nil {
		return 0, fmt.Errorf("get all students: %w", err)
	}

	largestID := 0
	for _, student := range students {
		if student.ID > largestID {
			largestID = student.ID
		}
	}

	return largestID + 1, nil
}
