package main

import (
	"fmt"
	"os"

	"github.com/TypingHare/course-sync/internal/app"
	"github.com/TypingHare/course-sync/internal/cli"
)

func main() {
	ctx := app.Context{
		Role: "student",
	}

	if err := cli.Cmd(ctx).Execute(); err != nil {
		fmt.Fprintln(
			os.Stderr,
			"Error:", err,
		)
		os.Exit(1)
	}
}
