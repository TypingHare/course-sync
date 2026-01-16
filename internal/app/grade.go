package app

import (
	"path/filepath"

	"github.com/TypingHare/course-sync/internal/adapter/repo/jsonstore"
	"github.com/TypingHare/course-sync/internal/domain/service"
)

const GradeDataFileName = "grades.json"

// GetGradeDataFile returns the grades data file path.
func GetGradeDataFile(dataDir string) string {
	return filepath.Join(dataDir, GradeDataFileName)
}

// GetGradeService constructs a GradeService backed by the data file.
func GetGradeService(gradeDataFile string) *service.GradeService {
	return service.NewGradeService(
		jsonstore.NewGradeRepo(gradeDataFile),
	)
}
