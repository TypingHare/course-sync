package service

import (
	"github.com/TypingHare/course-sync/internal/domain/model"
	"github.com/TypingHare/course-sync/internal/domain/repo"
)

type DocService struct {
	repo repo.IDocRepo
}

// NewDocService constructs a DocService with the provided repo.
func NewDocService(r repo.IDocRepo) *DocService {
	return &DocService{repo: r}
}

// GetAllDocs retrieves all docs from the repository.
func (s *DocService) GetAllDocs() ([]model.Doc, error) {
	return s.repo.GetAll()
}

// GetDocByName finds a doc by name.
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

// GetDefaultDoc returns the doc marked as default.
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

// AddDoc adds a new doc to the repository. If the new doc is marked as default,
// the existing default doc will be unset.
func (s *DocService) AddDoc(doc *model.Doc) error {
	docs, err := s.repo.GetAll()
	if err != nil {
		return err
	}
	if docs == nil {
		docs = []model.Doc{}
	}

	if doc.IsDefault {
		for i := range docs {
			docs[i].IsDefault = false
		}
	}

	docs = append(docs, *doc)
	return s.repo.SaveAll(docs)
}

// UpsertDoc adds a new doc or replaces an existing one with the same name.
// If the incoming doc is marked as default, all other docs are unset.
func (s *DocService) UpsertDoc(doc *model.Doc) error {
	docs, err := s.repo.GetAll()
	if err != nil {
		return err
	}
	if docs == nil {
		docs = []model.Doc{}
	}

	if doc.IsDefault {
		for i := range docs {
			docs[i].IsDefault = false
		}
	}

	existingIndex := -1
	for i := range docs {
		if docs[i].Name == doc.Name {
			existingIndex = i
			break
		}
	}

	if existingIndex >= 0 {
		docs[existingIndex] = *doc
	} else {
		docs = append(docs, *doc)
	}

	return s.repo.SaveAll(docs)
}
