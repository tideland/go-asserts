// -----------------------------------------------------------------------------
// asserts for more convinient testing - failer providing interfaces for positive
// and negative testing
//
// Copyright (C) 2024 Frank Mueller / Oldenburg / Germany / World
// -----------------------------------------------------------------------------

package asserts // import "tideland.dev/go/asserts"

import (
	"testing"
)

// Tester surrounds the testing.T to provide a more convinient way of testing.
type Tester interface {
	FailNow()
}

// tester defines the expected functions of a testing.T but cares about failing.
type tester struct {
	t        *testing.T
	negative bool
	failed   int
}

// FailNow in case of a negative test is not allowed to be called.
func (t *tester) FailNow() {
	if t.negative {
		t.failed++
	} else {
		t.t.FailNow()
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
