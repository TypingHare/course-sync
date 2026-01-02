package shell

import (
	"fmt"
	"os"
	"sync"
	"time"

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
	message  string        // Static message shown next to the spinner
	frames   []rune        // Animation frames (Unicode characters)
	delay    time.Duration // Time between frame updates
	mutex    sync.Mutex    // Guards stop/finish operations
	stopChan chan struct{} // Signals the spinner goroutine to stop
	doneChan chan struct{} // Closed when the spinner goroutine exits
	tty      bool          // Whether stdout is a terminal (TTY)
}

func NewSpinner(message string) *Spinner {
	return &Spinner{
		message:  message,
		frames:   []rune{'⠋', '⠙', '⠹', '⠸', '⠼', '⠴', '⠦', '⠧', '⠇', '⠏'},
		delay:    90 * time.Millisecond,
		stopChan: make(chan struct{}),
		doneChan: make(chan struct{}),
		tty:      term.IsTerminal(int(os.Stdout.Fd())),
	}
}

func (spinner *Spinner) Start() {
	// If not a TTY, just print a normal line (no animation).
	if !spinner.tty {
		fmt.Println(spinner.message + "...")
		close(spinner.doneChan)
		return
	}

	go func() {
		defer close(spinner.doneChan)

		i := 0
		for {
			select {
			case <-spinner.stopChan:
				return
			default:
				frame := spinner.frames[i%len(spinner.frames)]
				// \r = carriage return, \033[2K = clear entire line
				fmt.Fprintf(os.Stdout, "\r\033[2K%c  %s", frame, spinner.message)
				i++
				time.Sleep(spinner.delay)
			}
		}
	}()
}

func (spinner *Spinner) StopWithMessage(message string) {
	spinner.mutex.Lock()
	defer spinner.mutex.Unlock()

	// Stop animation goroutine (if any) and wait for it.
	select {
	case <-spinner.stopChan:
	default:
		close(spinner.stopChan)
	}

	<-spinner.doneChan

	if spinner.tty {
		fmt.Fprintf(os.Stdout, "\r\033[2K%s\n", message)
	} else {
		fmt.Println(message)
	}
}
