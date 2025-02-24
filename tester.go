// -----------------------------------------------------------------------------
// asserts for more convinient testing - failer providing interfaces for positive
// and negative testing
//
// Copyright (C) 2024-2025 Frank Mueller / Oldenburg / Germany / Earth
// -----------------------------------------------------------------------------

package asserts // import "tideland.dev/go/asserts"

import (
	"fmt"
	"regexp"
	"sync"
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
	mux      sync.Mutex
	fail     bool
	failures []string
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
	t.mux.Lock()
	defer t.mux.Unlock()
	t.failures = append(t.failures, fmt.Sprint(args...))
	if t.fail {
		t.TB.Error(args...)
	}
}

// Errorf overwrites the testing.TB method to count the number of failures if
// the behavior is set to CONTINUE. Otherwise it fails the test.
func (t *tester) Errorf(format string, args ...any) {
	t.mux.Lock()
	defer t.mux.Unlock()
	t.failures = append(t.failures, fmt.Sprintf(format, args...))
	if t.fail {
		t.TB.Errorf(format, args...)
	}
}

// Fail overwrites the testing.TB method to count the number of failures if
// the behavior is set to CONTINUE. Otherwise it fails the test.
func (t *tester) Fail() {
	t.mux.Lock()
	defer t.mux.Unlock()
	t.failures = append(t.failures, "fail")
	if t.fail {
		t.TB.Fail()
	}
}

// FailNow overwrites the testing.TB method to count the number of failures if
// the behavior is set to CONTINUE. Otherwise it fails the test.
func (t *tester) FailNow() {
	t.mux.Lock()
	defer t.mux.Unlock()
	t.failures = append(t.failures, "fail now")
	if t.fail {
		t.TB.FailNow()
	}
}

// Failures checks the number of failures and fails the test if it is not
// the expected one.
func Failures(t Tester, expected int) {
	tt, ok := t.(*tester)
	if !ok {
		ot, ok := t.(testing.TB)
		if !ok {
			panic("not a tester and not a testing.TB")
		}
		ot.Fatalf("not a tester")
	}
	if len(tt.failures) != expected {
		for _, f := range tt.failures {
			tt.TB.Logf("failure: %v", f)
		}
		tt.TB.Errorf("expected %d failures, but got %d", expected, len(tt.failures))
	}
}

// FailureMatch checks if a failure matches a regular expression pattern.
func FailureMatch(t Tester, number int, pattern string) {
	tt, ok := t.(*tester)
	if !ok {
		ot, ok := t.(testing.TB)
		if !ok {
			panic("not a tester and not a testing.TB")
		}
		ot.Fatalf("not a tester")
	}
	if number < 0 || number >= len(tt.failures) {
		tt.TB.Errorf("failure number %d out of range", number)
	}
	re := regexp.MustCompile(pattern)
	if !re.MatchString(tt.failures[number]) {
		for _, f := range tt.failures {
			tt.TB.Logf("failure: %v", f)
		}
		tt.TB.Errorf("failure number %d does not match pattern'%s': %v", number, pattern, tt.failures[number])
	}
}

// -----------------------------------------------------------------------------
// end of file
// -----------------------------------------------------------------------------
