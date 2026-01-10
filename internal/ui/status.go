package ui

import (
	"fmt"

	"github.com/fatih/color"
)

// MarkSuccess prepends a success message with a green square emoji.
func MarkSuccess(message string) string {
	return fmt.Sprintf("🟩 %s", message)
}

// MarkWarning prepends a warning message with an orange square emoji.
func MarkWarning(message string) string {
	return fmt.Sprintf("🟧 %s", message)
}

// MarkError prepends an error message with a red square emoji.
func MarkError(message string) string {
	return fmt.Sprintf("🟥 %s", message)
}

// MakeOngoing styles an ongoing message with bright cyan color.
func MakeOngoing(message string) string {
	return color.New(color.FgHiCyan).SprintFunc()(message)
}

// MakeDone styles a done message with bright black (gray) color.
func MakeDone(message string) string {
	return color.New(color.FgHiBlack).SprintFunc()(message)
}

// MakeError styles an error message with bright red color.
func MakeError(message string) string {
	return color.New(color.FgHiRed).SprintFunc()(message)
}
