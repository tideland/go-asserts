// -----------------------------------------------------------------------------
// asserts for more convinient testing - printing fail information
//
// Copyright (C) 2024 Frank Mueller / Oldenburg / Germany / World
// -----------------------------------------------------------------------------

package asserts // import "tideland.dev/go/asserts"

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
)

// logf prints a log message with the given information on stdout.
func logf(format string, args ...any) {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	// Retrieve location and function with standard offset.
	location, fun := here(4)

	// Prefix log with location and function.
	w.WriteString(fmt.Sprintf("[LOG] %s log in %s() ", location, fun))
	w.WriteString(fmt.Sprintf(format, args...))
	w.WriteString("\n")
}

// failf prints a fail message with the given information on stderr.
func failf(t Tester, assertion string, format string, args ...any) {
	w := bufio.NewWriter(os.Stderr)
	defer w.Flush()

	// Retrieve location and function with standard offset.
	location, fun := here(4)

	// Prefix with location and function.
	if assertion == "" {
		w.WriteString(fmt.Sprintf("[ERR] %s assert in %s() failed {", location, fun))
	} else {
		w.WriteString(fmt.Sprintf("[ERR] %s assert '%s' in %s() failed {", location, assertion, fun))
	}

	// Add the given information.
	w.WriteString(fmt.Sprintf(format, args...))
	w.WriteString("}\n")

	t.FailNow()
}

// here returns the filename and position based on a given offset.
// It's used by the fail function to print the location of the failure.
func here(offset int) (string, string) {
	// Retrieve program counters.
	pcs := make([]uintptr, 1)
	n := runtime.Callers(offset, pcs)
	if n == 0 {
		return "", ""
	}
	pcs = pcs[:n]
	// Build ID based on program counters.
	frames := runtime.CallersFrames(pcs)
	for {
		frame, more := frames.Next()
		_, fun := path.Split(frame.Function)
		parts := strings.Split(fun, ".")
		fun = strings.Join(parts[1:], ".")
		_, file := path.Split(frame.File)
		location := fmt.Sprintf("%s:%d:0:", file, frame.Line)
		if !more {
			return location, fun
		}
	}
}

// -----------------------------------------------------------------------------
// end of file
// -----------------------------------------------------------------------------
