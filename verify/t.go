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
)

// -----------------------------------------------------------------------------
// testing.T Replacement
// -----------------------------------------------------------------------------

// T replaces testing.T for tests. Missing methods are handled internally.
type T interface {
	Errorf(format string, args ...any)
}

// continuedTesting is a wrapper around *testing.T that
// indicates the test should continue running even after
// a verification failure
type continuedTesting struct {
	*testing.T
	failed int
	msgs   []string
}

// Ensure the wrapper implement T
var _ T = (*continuedTesting)(nil)

func (ct *continuedTesting) Errorf(format string, args ...any) {
	ct.failed++

	location := here(5)
	locatedformat := "    " + location + ": continuation " + format

	ct.msgs = append(ct.msgs, fmt.Sprintf(locatedformat, args...))

	for _, msg := range ct.msgs {
		fmt.Println(msg)
	}
	ct.msgs = nil
}

// -----------------------------------------------------------------------------
// Library API
// -----------------------------------------------------------------------------

// ContinuedTesting creates a new T instance that continues after
// testing failures.
func ContinuedTesting(t *testing.T) T {
	ct := &continuedTesting{t, 0, nil}
	return ct
}

// IsContinueT checks if a testing.T is a continueTesting type.
func IsContinueT(t T) bool {
	_, ok := t.(*continuedTesting)
	return ok
}

// FailureCount validates how many tests failed during continued
// test to verify the expected number.
func FailureCount(t T, expected int) bool {
	var ct *continuedTesting
	var ok bool

	if ct, ok = t.(*continuedTesting); !ok {
		t.Errorf("t is no continuation testing")
		return false
	}

	if ct.failed != expected {
		failed := ct.failed
		verificationFailure(t, "failure count", expected, failed)
		ct.T.Fail()
	}
	return true
}

// -----------------------------------------------------------------------------
// UTILS
// -----------------------------------------------------------------------------

type helperT interface {
	Helper()
}

// verificationFailure raises an error containing the failure message.
func verificationFailure(t T, verification string, expected, got any) {
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
		tt.Errorf("fail %q verification: want '%v', got '%v'", verification, expected, got)
		tt.FailNow()
		return
	}
	if ct, ok := t.(*continuedTesting); ok {
		ct.Errorf("fail %q verification: want '%v', got '%v'", verification, expected, got)
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
		location := fmt.Sprintf("%s:%d", file, frame.Line)
		if !more {
			return location
		}
	}
}

// -----------------------------------------------------------------------------
// EOF
// -----------------------------------------------------------------------------
