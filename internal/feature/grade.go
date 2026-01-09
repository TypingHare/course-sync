package feature

import (
	"fmt"
	"time"

	"github.com/TypingHare/course-sync/internal/app"
)

type Grade struct {
	AssignmentName string    // Name of the assignment.
	SubmissionHash string    // Hash of the submission.
	Score          float64   // Score received for the assignment.
	Feedback       string    // Feedback provided for the assignment.
	GradedAt       time.Time // Timestamp when the assignment was graded.
}

// GetGradeHistory retrieves the grade history from the assignments file.
func GetGradeHistory() ([]Grade, error) {
	grades, err := ReadJSONSlice[[]Grade](app.GRADE_HISTORY_FILE_NAME)
	if err != nil {
		return nil, err
	}

	return grades, nil
}

// FindLastGradeByAssignmentName searches for the most recent grade entry for a given assignment
// name.
func FindLastGradeByAssignmentName(grades []Grade, assignmentName string) *Grade {
	var latest *Grade

	for i := range grades {
		grade := &grades[i]

		if grade.AssignmentName != assignmentName {
			continue
		}

		if latest == nil || grade.GradedAt.After(latest.GradedAt) {
			latest = grade
		}
	}

	return latest
}

// FindGradeBySubmissionHash searches for a grade entry by its submission hash.
func FindGradeBySubmissionHash(grades []Grade, submissionHash string) *Grade {
	for _, grade := range grades {
		if grade.SubmissionHash == submissionHash {
			return &grade
		}
	}

	return nil
}

// AppendGradeHistory appends a new grade entry to the grade history file.
func AppendGradeHistory(grade Grade) error {
	gradeHistory, err := GetGradeHistory()
	if err != nil {
		return fmt.Errorf("get grade history: %w", err)
	}

	gradeHistory = append(gradeHistory, grade)

	return WriteJSONSlice(app.GRADE_HISTORY_FILE_NAME, gradeHistory)
}
