package student

// Student represents a student in the system.
const STUDENTS_FILE_NAME = "students.json"

// NewStudent creates a new Student instance.
func NewStudent(
	ID int,
	name string,
	email string,
	repositoryURL string,
) *Student {
	return &Student{
		ID:            ID,
		Name:          name,
		Email:         email,
		RepositoryURL: repositoryURL,
	}
}
