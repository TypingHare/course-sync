package config

import "github.com/TypingHare/course-sync/internal/domain/student"

// Config represents the configuration for Course Sync.
type Config struct {
	// Roster contains a list of students in the course.
	Roster []student.Student `mapstructure:"roster"`
}
