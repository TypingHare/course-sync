package model

// Role represents the user role in the application.
type Role string

// Defined user roles.
const (
	RoleStudent    Role = "student"
	RoleInstructor Role = "instructor"
)
