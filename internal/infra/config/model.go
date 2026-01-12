package config

import "github.com/TypingHare/course-sync/internal/domain/student"

// Config represents the configuration for Course Sync.
type Config struct {
	// Roster contains a list of students in the course.
	Roster []student.Student `mapstructure:"roster"`
}

// GetDefault returns a Config instance with default values.
func GetDefault() *Config {
	return &Config{
		Roster: []student.Student{},
	}
}
