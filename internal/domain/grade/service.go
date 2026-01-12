package grade

import (
	"fmt"
	"path/filepath"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/infra/jsonstore"
)

// GRADES_FILE_NAME is the name of the file where grade history is stored.
const GRADES_FILE_NAME = "grades.json"

// GetGrades retrieves the list of grades from the grades JSON file in the
// application data directory.
func GetGrades(appCtx *app.Context) ([]Grade, error) {
	grades, err := jsonstore.ReadJSONFile[[]Grade](
		filepath.Join(appCtx.AppDataDir, GRADES_FILE_NAME),
	)
	if err != nil {
		return nil, fmt.Errorf("read grades file: %w", err)
	}

	return grades, nil
}

// FindLastGradeByAssignmentName searches for the most recent grade entry for a
// given assignment name.
func FindLastGradeByAssignmentName(
	grades []Grade,
	assignmentName string,
) *Grade {
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
	for i := range grades {
		if grades[i].SubmissionHash == submissionHash {
			return &grades[i]
		}
	}

	return nil
}

// AppendGradeHistory appends a new grade entry to the grades file.
func AppendGradeHistory(appCtx *app.Context, grade Grade) error {
	grades, err := GetGrades(appCtx)
	if err != nil {
		return fmt.Errorf("get grades: %w", err)
	}

	grades = append(grades, grade)

	return jsonstore.WriteJSONFile(
		filepath.Join(appCtx.AppDataDir, GRADES_FILE_NAME),
		grades,
	)
}
