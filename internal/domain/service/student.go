package service

import "github.com/TypingHare/course-sync/internal/domain/repo"

type StudentService struct {
	repo repo.IStudentRepo
}

func NewStudentService(studentRepo repo.IStudentRepo) *StudentService {
	return &StudentService{
		repo: studentRepo,
	}
}
