// -----------------------------------------------------------------------------
// Convenient verification of unit tests in Go libraries and applications.
//
// Replacement of testing.T to allow teest without immediate fail
//
// Copyright (C) 2024-2025 Frank Mueller / Oldenburg / Germany / Earth
// -----------------------------------------------------------------------------

package verify // import "tideland.dev/go/assert/verify"

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"testing"
	"time"
)

// -----------------------------------------------------------------------------
// testing.T Replacement
// -----------------------------------------------------------------------------

// T replaces testing.T for tests. Missing methods are handled internally.
type T interface {
	Errorf(format string, args ...any)
}

// continuationTesting is a wrapper around *testing.T that
// indicates the test should continue running even after
// a verification failure
type continuationTesting struct {
	*testing.T
	start  time.Time
	failed int
	msgs   []string
}

// Ensure the wrapper implement T
var _ T = (*continuationTesting)(nil)

func (ct *continuationTesting) Errorf(format string, args ...any) {
	if ct.failed == 0 {
		now := time.Now()
		ct.msgs = append(ct.msgs, fmt.Sprintf("--- FAIL: %s (%s)", ct.T.Name(), now.Sub(ct.start)))
	}

	ct.failed++

	location := here(5)
	locatedformat := "    " + location + ": " + format

	ct.msgs = append(ct.msgs, fmt.Sprintf(locatedformat, args...))

	for _, msg := range ct.msgs {
		fmt.Printf("%s\n", msg)
	}

	ct.msgs = nil
}

// -----------------------------------------------------------------------------
// Library API
// -----------------------------------------------------------------------------

// ContinuationTesting creates a new T instance that continues after
// testing failures.
func ContinuationTesting(t *testing.T) T {
	return &continuationTesting{t, time.Now(), 0, nil}
}

// IsContinueT checks if a testing.T is a continueTesting type.
func IsContinueT(t T) bool {
	_, ok := t.(*continuationTesting)
	return ok
}

// ConinuedFails validates how many tests failed during continued
// test to verify the expected number.
func ConinuedFails(t T, expected int) bool {
	if ct, ok := t.(*continuationTesting); ok {
		if ct.failed != expected {
			ct.Errorf("fail %q verification: want '%v', got '%v'", "continued fails", expected, ct.failed)
		}
		return true
	}
	t.Errorf("t is no continuation testing")
	return false
}

// -----------------------------------------------------------------------------
// UTILS
// -----------------------------------------------------------------------------

type failNowT interface {
	FailNow()
}

// verificationFailure raises an error containing the failure message.
func verificationFailure(t T, verification string, expected, got any) {
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
		tt.Errorf("fail %q verification: want '%v', got '%v'", verification, expected, got)
		return
	}
	if ct, ok := t.(*continuationTesting); ok {
		ct.Errorf("fail %q verification: want '%v', got '%v'", verification, expected, got)
		return
	}
	if ft, ok := t.(failNowT); ok {
		ft.FailNow()
		return
	}
	t.Errorf("fail %q verification: want '%v', got '%v'", verification, expected, got)
}

// here returns the location at the offseet of the caller.
func here(offset int) string {
	// Retrieve program counters
	pcs := make([]uintptr, 1)
	n := runtime.Callers(offset, pcs)
	if n == 0 {
		return ""
	}
	pcs = pcs[:n]
	// Build ID based on program counters
	frames := runtime.CallersFrames(pcs)
	for {
		frame, more := frames.Next()
		_, function := path.Split(frame.Function)
		parts := strings.Split(function, ".")
		function = strings.Join(parts[1:], ".")
		_, file := path.Split(frame.File)
		location := fmt.Sprintf("%s:%d: %s()", file, frame.Line, function)
		if !more {
			return location
		}
	}
}

// -----------------------------------------------------------------------------
// EOF
// -----------------------------------------------------------------------------
