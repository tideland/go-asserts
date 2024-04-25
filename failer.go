// -----------------------------------------------------------------------------
// asserts for more convinient testing - failer providing interfaces for positive
// and negative testing
//
// Copyright (C) 2024 Frank Mueller / Oldenburg / Germany / World
// -----------------------------------------------------------------------------

package asserts // import "tideland.dev/go/asserts"

import (
	"fmt"
	"os"
	"testing"
)

// Tester defines the expected functions of any testing and logging type
// needed.
type Tester interface {
	Logf(format string, args ...any)
	Errorf(format string, args ...any)
	Fail()
}

// tester defines the expected functions of a testing.T but cares about failing.
type tester struct {
	t        *testing.T
	negative bool
	failed   int
}

// Logf is used to print additional information during testing.
func (t tester) Logf(format string, args ...any) {
	t.t.Logf(format, args...)
	// fmt.Fprintf(os.Stderr, format, args...)
}

// Errorf is used to fail a test with a formatted message.
func (t tester) Errorf(format string, args ...any) {
	if !t.negative {
		fmt.Fprintf(os.Stderr, format, args...)
	}
}

// Fail is used to signal afailing test.
func (t *tester) Fail() {
	if t.negative {
		t.failed++
	} else {
		t.t.Fail()
	}
}

// Failed returns true if the given failed.
func Failed(t Tester, count int) bool {
	tt, ok := t.(*testing.T)
	if ok {
		return tt.Failed()
	}
	failed := t.(*tester).failed
	t.(*tester).failed = 0
	if failed != count {
		t.Logf("failed: %d, expected: %d", failed, count)
		t.Fail()
	}
	return failed == count
}

// MkPosNeg creates a positive and a negative tester.
func MkPosNeg(t *testing.T) (Tester, Tester) {
	pt := &tester{
		t:        t,
		negative: false,
	}
	nt := &tester{
		t:        t,
		negative: true,
	}
	return pt, nt
}

// -----------------------------------------------------------------------------
// end of file
// -----------------------------------------------------------------------------
