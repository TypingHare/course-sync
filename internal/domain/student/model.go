package student

// Student represents a student with their basic information.
type Student struct {
	// ID is the unique identifier for the student.
	ID int `json:"id"`

	// Name is the full name of the student.
	Name string `json:"name"`

	// Email is the email address of the student.
	Email string `json:"email"`

	// RepositoryURL is the URL of the student's code repository.
	RepositoryURL string `json:"repository_url"`
}
