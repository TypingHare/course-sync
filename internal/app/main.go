package app

// RunAll runs the given functions in order. It stops at the first error and returns it.
func RunAll(fns ...func() error) error {
	for _, fn := range fns {
		if err := fn(); err != nil {
			return err
		}
	}

	return nil
}
