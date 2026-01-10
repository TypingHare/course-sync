package student

// Student represents a student with their basic information.
type Student struct {
	// ID is the unique identifier for the student.
	ID int `mapstructure:"id"`

	// Name is the full name of the student.
	Name string `mapstructure:"name"`

	// Email is the email address of the student.
	Email string `mapstructure:"email"`

	// RepositoryUrl is the URL of the student's code repository.
	RepositoryUrl string `mapstructure:"repositoryUrl"`
}
