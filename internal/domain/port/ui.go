package port

// OutputMode is the minimal surface area domain logic needs to decide how noisy
// it should be.
type OutputMode interface {
	IsVerbose() bool
	IsQuiet() bool
}
