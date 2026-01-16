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
		printErrorAndExit(
			fmt.Errorf("failed to create application context: %w", err),
		)
	}

	if err := cli.Cmd(ctx).Execute(); err != nil {
		printErrorAndExit(err)
	}
}

func printErrorAndExit(err error) {
	if err == nil {
		return
	}

	fmt.Fprintln(
		os.Stderr,
		color.HiRedString(fmt.Sprintf("Error: %s", err)),
	)

	os.Exit(1)
}
