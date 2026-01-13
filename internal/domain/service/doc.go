package service

import (
	"github.com/TypingHare/course-sync/internal/domain/model"
	"github.com/TypingHare/course-sync/internal/domain/repo"
)

type DocService struct {
	repo repo.IDocRepo
}

func NewDocService(r repo.IDocRepo) *DocService {
	return &DocService{repo: r}
}

func (s *DocService) GetAllDocs() ([]model.Doc, error) {
	return s.repo.GetAll()
}

func (s *DocService) GetDocByName(name string) (*model.Doc, error) {
	docs, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	if docs == nil {
		return nil, nil
	}

	for i := range docs {
		if docs[i].Name == name {
			return &docs[i], nil
		}
	}

	return nil, nil
}

func (s *DocService) GetDefaultDoc() (*model.Doc, error) {
	docs, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	if docs == nil {
		return nil, nil
	}

	for i := range docs {
		if docs[i].IsDefault {
			return &docs[i], nil
		}
	}
	return nil, nil
}
