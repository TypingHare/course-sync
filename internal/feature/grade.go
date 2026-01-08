package feature

import (
	"time"

	"github.com/TypingHare/course-sync/internal/app"
)

type Grade struct {
	AssignmentName string
	Score          float64
	Feedback       string
	GradedAt       time.Time
}

// GetGradeHistory retrieves the grade history from the assignments file.
func GetGradeHistory() ([]Grade, error) {
	grades, err := ReadJSONSlice[[]Grade](app.GRADE_HISTORY_FILE_NAME)
	if err != nil {
		return nil, err
	}

	return grades, nil
}
