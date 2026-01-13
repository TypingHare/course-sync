package service

import "github.com/TypingHare/course-sync/internal/domain/repo"

type DocService struct {
	repo repo.IDocRepo
}

func NewDocService(r repo.IDocRepo) *DocService {
	return &DocService{repo: r}
}
