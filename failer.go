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

// Behavior is used to define the behavior for breaking a test.
type Behavior bool

const (
	FAIL     Behavior = true
	CONTINUE Behavior = false
)

// Tester is the interface for the testing.TB interface. It helps to count
// failures without breaking the test.
type Tester interface {
	testing.TB
}

// tester implements the Tester interface and overwrites the Fail and FailNow
// methods to count the number of failures.
type tester struct {
	testing.TB
	fail   bool
	failed int
}

// NewTester creates a new tester. The behavior defines if the test should
// break on the first failure.
func NewTester(tb testing.TB, behavior Behavior) Tester {
	return &tester{
		TB:   tb,
		fail: bool(behavior),
	}
}

// Error overwrites the testing.TB method to count the number of failures if
// the behavior is set to CONTINUE. Otherwise it fails the test.
func (t *tester) Error(args ...any) {
	t.failed++
	if t.fail {
		t.TB.Error(args...)
	}
}

// Errorf overwrites the testing.TB method to count the number of failures if
// the behavior is set to CONTINUE. Otherwise it fails the test.
func (t *tester) Errorf(format string, args ...any) {
	t.failed++
	if t.fail {
		t.TB.Errorf(format, args...)
	}
}

// Fail overwrites the testing.TB method to count the number of failures if
// the behavior is set to CONTINUE. Otherwise it fails the test.
func (t *tester) Fail() {
	t.failed++
	if t.fail {
		t.TB.Fail()
	}
}

// FailNow overwrites the testing.TB method to count the number of failures if
// the behavior is set to CONTINUE. Otherwise it fails the test.
func (t *tester) FailNow() {
	t.failed++
	if t.fail {
		t.TB.FailNow()
	}
}

// Failures checks the number of failures and fails the test if it is not
// the expected one.
func Failures(t Tester, expected int) {
	tt, ok := t.(*tester)
	if !ok {
		t.Fail()
		return
	}
	if tt.failed != expected {
		t.Errorf("expected %d failures, but got %d", expected, tt.failed)
	}
}

// -----------------------------------------------------------------------------
// end of file
// -----------------------------------------------------------------------------
