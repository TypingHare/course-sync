package main

import (
	"fmt"
	"os"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/cli"
	"github.com/fatih/color"
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
			color.HiRedString(fmt.Sprintf("Error: %s", err)),
		)
		os.Exit(1)
	}
}
