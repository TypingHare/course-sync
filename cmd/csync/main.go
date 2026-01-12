package main

import (
	"fmt"
	"os"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/cli"
	"github.com/TypingHare/course-sync/internal/ui"
)

func main() {
	// Build the application context.
	appCtx, err := app.BuildContext()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Execute the CLI command.
	if err = cli.Cmd(appCtx).Execute(); err != nil {
		fmt.Fprintln(
			os.Stderr,
			ui.MakeError("Fatal error: failed to "+err.Error()),
		)
		os.Exit(1)
	}
}
