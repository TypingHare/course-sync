package exec

import (
	"strings"
	"unicode"
)

// JoinCommand returns a shell-style command line string.
// It quotes any argument that contains whitespace or quotes/backslashes.
func JoinCommand(command string, args []string) string {
	parts := make([]string, 0, 1+len(args))
	parts = append(parts, command)

	for _, a := range args {
		parts = append(parts, shellQuote(a))
	}

	return strings.Join(parts, " ")
}

// shellQuote quotes a string for safe inclusion in a shell command line.
// It uses double quotes and escapes only quotes and backslashes.
func shellQuote(s string) string {
	if s == "" {
		return `""`
	}

	needQuotes := false
	for _, r := range s {
		if unicode.IsSpace(r) || r == '"' || r == '\\' {
			needQuotes = true
			break
		}
	}
	if !needQuotes {
		return s
	}

	// Minimal double-quote escaping: \ and " become escaped.
	var b strings.Builder
	b.Grow(len(s) + 2)
	b.WriteByte('"')
	for _, r := range s {
		if r == '\\' || r == '"' {
			b.WriteByte('\\')
		}
		b.WriteRune(r)
	}
	b.WriteByte('"')
	return b.String()
}
