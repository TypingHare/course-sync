package service

import (
	"fmt"

	"github.com/TypingHare/course-sync/internal/domain/model"
	"github.com/TypingHare/course-sync/internal/domain/repo"
)

type StudentService struct {
	repo repo.IStudentRepo
}

func NewStudentService(studentRepo repo.IStudentRepo) *StudentService {
	return &StudentService{
		repo: studentRepo,
	}
}

func (s *StudentService) GetAllStudents() ([]model.Student, error) {
	return s.repo.GetAll()
}

func (s *StudentService) AddStudent(
	student *model.Student,
) error {
	students, err := s.repo.GetAll()
	if err != nil {
		return fmt.Errorf("get all students: %w", err)
	}

	return s.repo.SaveAll(append(students, *student))
}

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
