package repo

import "github.com/TypingHare/course-sync/internal/domain/model"

type ISubmissionRepo interface {
	IBaseRepo[model.Submission]
}
