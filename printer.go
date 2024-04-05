// -----------------------------------------------------------------------------
// asserts for more convinient testing - printing fail information
//
// Copyright (C) 2024 Frank Mueller / Oldenburg / Germany / World
// -----------------------------------------------------------------------------

package asserts // import "tideland.dev/go/asserts"

import (
	"fmt"
	"log/slog"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"
)

// setup initializes the asserts package only one time.
func setup() {
	var once sync.Once
	onceBody := func() {
		logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
		slog.SetDefault(logger)
	}
	// Run the setup only once.
	done := make(chan struct{})
	go func() {
		once.Do(onceBody)
		close(done)
	}()
	<-done
}

// logf prints a log message with the given information on stdout.
func logf(format string, args ...any) {
	setup()
	location, fun := here(4)

	slog.Info("asserts log", "location", location, "function", fun, "message", fmt.Sprintf(format, args...))
}

// failf prints a fail message with the given information on stderr.
func failf(t Tester, assertion string, format string, args ...any) {
	setup()
	location, fun := here(4)

	if assertion == "" {
		slog.Error("asserts fail", "location", location, "function", fun, "message", fmt.Sprintf(format, args...))
	} else {
		slog.Error("asserts fail", "location", location, "function", fun, "assertion", assertion, "message", fmt.Sprintf(format, args...))
	}

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
