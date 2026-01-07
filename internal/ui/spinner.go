package ui

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/charmbracelet/x/ansi"
	"golang.org/x/term"
)

// Spinner renders a terminal loading indicator on a single line while background work is in
// progress.
//
// It periodically updates a Unicode frame (e.g. ⠋ ⠙ ⠹) prefixed by a descriptive message. When
// stopped, the spinner clears the line and replaces it with a final status message.
//
// Spinner is safe for concurrent use. It automatically disables animation when stdout is not a
// TTY (e.g. when output is redirected or in CI).
type Spinner struct {
	tty      bool          // Whether stdout is a terminal (TTY)
	quiet    bool          // Whether to suppress output
	message  string        // Static message shown next to the spinner
	numLines int           // Number of lines the message occupies
	frames   []rune        // Animation frames (an array of Unicode characters)
	delay    time.Duration // Time between frame updates
	mutex    sync.Mutex    // Guards stop/finish operations
	stopChan chan struct{} // Signals the spinner goroutine to stop
	doneChan chan struct{} // Closed when the spinner goroutine exits
}

// NewSpinner creates a new Spinner with the given message.
func NewSpinner(message string, quiet bool) *Spinner {
	return &Spinner{
		tty:      term.IsTerminal(int(os.Stdout.Fd())),
		quiet:    quiet,
		message:  message,
		numLines: strings.Count(message, "\n") + 1,
		frames:   []rune{'⠋', '⠙', '⠹', '⠸', '⠼', '⠴', '⠦', '⠧', '⠇', '⠏'},
		delay:    100 * time.Millisecond,
		stopChan: make(chan struct{}),
		doneChan: make(chan struct{}),
	}
}

// ClearMessage clears the spinner message from the terminal. It also moves the cursor up for
// multi-line messages.
func (spinner *Spinner) ClearMessage() {
	for i := 0; i < spinner.numLines; i++ {
		fmt.Print("\r", ansi.EraseLine(2))
		if i < spinner.numLines-1 {
			fmt.Print(ansi.CursorUp(1))
		}
	}
	fmt.Print("\r")
}

func (spinner *Spinner) Start() {
	if spinner.quiet {
		return
	}

	if !spinner.tty {
		fmt.Println(spinner.message + "...")
		close(spinner.doneChan)
		return
	}

	num_frames := len(spinner.frames)
	go func() {
		defer close(spinner.doneChan)

		i := 0
		for {
			select {
			case <-spinner.stopChan:
				return
			default:
				if i > 0 {
					spinner.ClearMessage()
				}

				// Print the spinner frame and message.
				frame := spinner.frames[i%num_frames]
				fmt.Fprintf(os.Stdout, "\r\033[2K%c  %s", frame, spinner.message)
				i++

				// Wait before the next frame.
				time.Sleep(spinner.delay)
			}
		}
	}()
}

// Stop stops the spinner animation and clears the spinner message from the terminal. If the spinner
// is not running, Stop is a no-op.
func (spinner *Spinner) Stop() {
	if spinner.quiet {
		return
	}

	spinner.mutex.Lock()
	defer spinner.mutex.Unlock()

	select {
	case <-spinner.stopChan:
	default:
		close(spinner.stopChan)
		spinner.ClearMessage()
	}

	<-spinner.doneChan
}
