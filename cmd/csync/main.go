package main

import (
	"fmt"
	"os"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/cli"
)

func main() {
	ctx, err := app.NewContext()
	if err != nil {
		fmt.Fprintln(
			os.Stderr,
			"Error initializing application context:", err,
		)
		os.Exit(1)
	}

	if err := cli.Cmd(ctx).Execute(); err != nil {
		fmt.Fprintln(
			os.Stderr,
			"Error:", err,
		)
		os.Exit(1)
	}
}
