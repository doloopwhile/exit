// Package exit provides constants exit statuses defined in sysexits(3).
// It also provides wrapper functions to call os.Exit with a status code.
package exit

import (
	"fmt"
	"io"
	"os"
)

//go:generate ./gen

// Exit privides functions which print error message and exit
//
//   exit := exit.New("MyApp: ")
//
//   err := MainTask(path)
if err != nil {
	if err == ErrFileOpen {
		exit.IOErr("failed to open file: %s", path)
	} else if err == ErrInvalidFormat{
		exit.DataErr("invalid format: %s", path)
	}
}
//
type Exit struct {
	Prefix string
	Stderr io.Writer
}

// New returns a new instance of Exit with os.Stderr.
func New(prefix string) *Exit {
	return &Exit{Prefix: prefix, Stderr: os.Stderr}
}

// Exit call os.Exit with given code.
func (e *Exit) Exit(code int) {
	os.Exit(code)
}

// Exitf print error message to Stderr and call os.Exit with given code.
func (e *Exit) Exitf(code int, format string, args ...interface{}) {
	e.Stderr.Write([]byte(e.Prefix))
	fmt.Fprintf(os.Stderr, format, args...)
	e.Exit(code)
}
