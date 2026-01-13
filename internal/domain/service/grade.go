package service

import (
	"github.com/TypingHare/course-sync/internal/domain/model"
	"github.com/TypingHare/course-sync/internal/domain/repo"
)

type GradeService struct {
	repo repo.IGradeRepo
}

func NewGradeService(r repo.IGradeRepo) *GradeService {
	return &GradeService{repo: r}
}

func (s *GradeService) GetAllGrades() ([]model.Grade, error) {
	return s.repo.GetAll()
}

func (s *GradeService) GetGradeBySubmissionHash(
	submissionHash string,
) (*model.Grade, error) {
	grades, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	for i := range grades {
		if grades[i].SubmissionHash == submissionHash {
			return &grades[i], nil
		}
	}

	return nil, nil
}

func (s *GradeService) GetLastGradeByAssignmentName(
	assignmentName string,
) (*model.Grade, error) {
	grades, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var latest *model.Grade
	for i := range grades {
		if grades[i].AssignmentName != assignmentName {
			continue
		}

		if latest == nil || grades[i].GradedAt.After(latest.GradedAt) {
			latest = &grades[i]
		}
	}

	return latest, nil
}
