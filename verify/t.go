// Convenient verification of unit tests in Go libraries and applications.
//
// Replacement of testing.T to allow teest without immediate fail
//
// Copyright (C) 2024-2025 Frank Mueller / Oldenburg / Germany / Earth

package verify

import (
	"fmt"
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
	ct.Helper()

	ct.failed++

	ct.msgs = append(ct.msgs, fmt.Sprintf(format, args...))

	for _, msg := range ct.msgs {
		ct.T.Log(msg)
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

// IsContinued checks if a testing.T is a continueTesting type.
func IsContinued(t T) bool {
	_, ok := t.(*continuedTesting)
	return ok
}

// FailureCount validates how many tests failed during continued
// test to verify the expected number.
func FailureCount(t T, expected int) bool {
	var ct *continuedTesting
	var ok bool

	if ct, ok = t.(*continuedTesting); !ok {
		t.Errorf("t is no continued testing")
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

// verificationFailure raises an error containing the failure message.
func verificationFailure(t T, verification string, expected, got any, infos ...string) {
	info := strings.Join(infos, ",")
	msg := fmt.Sprintf("fail %q verification: got '%v', expected '%v'", verification, got, expected)
	if len(info) > 0 {
		msg = msg + " (" + info + ")"
	}
	if tt, ok := t.(*testing.T); ok {
		tt.Helper()
		tt.Errorf("%s", msg)
		tt.FailNow()
		return
	}
	if ct, ok := t.(*continuedTesting); ok {
		ct.Helper()
		ct.Errorf("%s", msg)
		return
	}
	t.Errorf("%s", msg)
}

// -----------------------------------------------------------------------------
// EOF
// -----------------------------------------------------------------------------
