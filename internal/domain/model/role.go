package model

// Role represents the user role in the application.
type Role string

// Defined user roles.
const (
	RoleUnknown    Role = "unknown"
	RoleStudent    Role = "student"
	RoleInstructor Role = "instructor"
)
