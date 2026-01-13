package repo

import "github.com/TypingHare/course-sync/internal/domain/model"

type IGradeRepo interface {
	IBaseRepo[model.Grade]
}
