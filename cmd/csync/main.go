package main

import (
	"fmt"
	"os"

	"github.com/TypingHare/course-sync/internal/cli"
)

func main() {
	err := cli.Command().Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
