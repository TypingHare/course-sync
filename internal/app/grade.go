package app

import (
	"path/filepath"

	"github.com/TypingHare/course-sync/internal/adapter/repo/jsonstore"
	"github.com/TypingHare/course-sync/internal/domain/service"
)

const GradeDataFileName = "grades.json"

func GetGradeDataFile(dataDir string) string {
	return filepath.Join(dataDir, GradeDataFileName)
}

func GetGradeService(gradeDataFile string) *service.GradeService {
	return service.NewGradeService(
		jsonstore.NewGradeRepo(gradeDataFile),
	)
}
