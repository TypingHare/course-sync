package io

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/charmbracelet/x/ansi"
	"golang.org/x/term"
)

// Spinner renders a terminal-based loading indicator while background work is
// in progress.
//
// A Spinner periodically displays an animated sequence of Unicode frames
// prefixed by a static message. When output is not a TTY, the spinner animation
// is disabled and the message is printed once instead.
//
// Spinner is safe for concurrent use. Calling Stop will terminate the animation
// and block until all background spinner activity has completed.
type Spinner struct {
	out io.Writer // Output writer (e.g., os.Stdout).
	tty bool      // Whether stdout is a terminal (TTY).

	message  string        // Static message shown next to the spinner.
	numLines int           // Number of lines the message occupies.
	frames   []rune        // Animation frames (an array of Unicode characters).
	delay    time.Duration // Time between frame updates.

	stopChan chan struct{} // Signals the spinner goroutine to stop.
	doneChan chan struct{} // Closed when the spinner goroutine exits.

	startOnce sync.Once
	stopOnce  sync.Once
	doneOnce  sync.Once
}

// NewSpinner creates a new Spinner with the given output writer and message.
func NewSpinner(out io.Writer, message string) *Spinner {
	var tty bool
	if f, ok := out.(*os.File); ok {
		tty = term.IsTerminal(int(f.Fd()))
	}

	return &Spinner{
		out:      out,
		tty:      tty,
		message:  message,
		numLines: strings.Count(message, "\n") + 1,
		frames:   []rune{'⠋', '⠙', '⠹', '⠸', '⠼', '⠴', '⠦', '⠧', '⠇', '⠏'},
		delay:    60 * time.Millisecond,
		stopChan: make(chan struct{}),
		doneChan: make(chan struct{}),
	}
}

// ClearMessage clears the spinner message from the terminal. It also moves the
// cursor up for multi-line messages.
func (s *Spinner) ClearMessage() {
	for i := 0; i < s.numLines; i++ {
		fmt.Fprint(s.out, "\r", ansi.EraseLine(2))
		if i < s.numLines-1 {
			fmt.Fprint(s.out, ansi.CursorUp(1))
		}
	}
	fmt.Fprint(s.out, "\r")
}

// Start begins rendering the spinner animation.
//
// If the output is not a TTY, Start prints the spinner message once (followed
// by "...") and returns immediately without starting any background goroutine.
//
// When running in a TTY, Start launches a goroutine that periodically updates
// the spinner frame until Stop is called. Start is non-blocking and may be
// called at most once per Spinner instance.
func (s *Spinner) Start() {
	s.startOnce.Do(func() {
		if !s.tty {
			fmt.Fprintln(s.out, s.message+"...")
			s.doneOnce.Do(func() { close(s.doneChan) })
			return
		}

		go func() {
			defer s.doneOnce.Do(func() { close(s.doneChan) })

			ticker := time.NewTicker(s.delay)
			defer ticker.Stop()

			s.PrintMessage(s.frames[0])

			numFrames := len(s.frames)
			i := 1

			for {
				select {
				case <-s.stopChan:
					return
				case <-ticker.C:
					s.ClearMessage()
					s.PrintMessage(s.frames[i%numFrames])
					i++
				}
			}
		}()
	})
}

// PrintMessage prints the spinner message prefixed by the given frame.
func (s *Spinner) PrintMessage(frame rune) {
	fmt.Fprintf(s.out, "%c  %s", frame, s.message)
}

// Stop stops the spinner animation and waits for it to terminate.
//
// Stop is safe to call multiple times; subsequent calls are no-ops. If Start
// was never called, Stop returns immediately. When Stop returns, all background
// spinner activity has completed.
func (s *Spinner) Stop() {
	s.stopOnce.Do(func() { close(s.stopChan) })

	// If Start was never called, there is no goroutine to close doneChan.
	s.doneOnce.Do(func() { close(s.doneChan) })

	<-s.doneChan
}
