package service

import "github.com/TypingHare/course-sync/internal/domain/repo"

type GradeService struct {
	repo repo.IGradeRepo
}

func NewGradeService(r repo.IGradeRepo) *GradeService {
	return &GradeService{repo: r}
}
