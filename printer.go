// -----------------------------------------------------------------------------
// asserts for more convinient testing - printing fail information
//
// Copyright (C) 2024 Frank Mueller / Oldenburg / Germany / World
// -----------------------------------------------------------------------------

package asserts // import "tideland.dev/go/asserts"

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

// logf prints a log message with the given information on stdout.
func logf(t Tester, format string, args ...any) {
	location, fun := here(3)
	format = fmt.Sprintf("%s assertion log at %s(): %s\n", location, fun, format)

	t.Logf(format, args...)
}

// failf prints a fail message with the given information on stderr.
func failf(t Tester, assertion string, format string, args ...any) {
	location, fun := here(4)
	format = fmt.Sprintf("%s assertion '%s' fail at %s(): %s\n", location, assertion, fun, format)

	t.Errorf(format, args...)
	t.Fail()
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
