package repo

import "github.com/TypingHare/course-sync/internal/domain/model"

type IStudentRepo interface {
	IBaseRepo[model.Student]
}
