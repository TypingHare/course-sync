package app

type Context struct {
	Verbose bool
	Quiet   bool
	Plain   bool

	WorkingDir string
	ProjectDir string

	Role string
}
