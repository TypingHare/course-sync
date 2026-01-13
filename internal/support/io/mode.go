package io

// OutputMode captures how command output should be rendered.
type OutputMode struct {
	Verbose bool
	Quiet   bool
	Plain   bool
}

// NewOutputMode returns an OutputMode configured with the provided flags.
func NewOutputMode(verbose bool, quiet bool, plain bool) *OutputMode {
	return &OutputMode{
		Verbose: verbose,
		Quiet:   quiet,
		Plain:   plain,
	}
}

// IsVerbose reports whether verbose output is enabled.
func (o *OutputMode) IsVerbose() bool {
	return o.Verbose
}

// IsQuiet reports whether non-error output should be suppressed.
func (o *OutputMode) IsQuiet() bool {
	return o.Quiet
}

// IsPlain reports whether output should avoid styling.
func (o *OutputMode) IsPlain() bool {
	return o.Plain
}
