package repo

import "github.com/TypingHare/course-sync/internal/domain/model"

type IDocRepo interface {
	IBaseRepo[model.Doc]
}
