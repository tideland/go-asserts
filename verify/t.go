// Tideland Go Asserts - Verify
//
// Copyright (C) 2024-2025 Frank Mueller / Tideland / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package verify

import (
	"fmt"
	"strings"
	"sync"
	"testing"
)

// testing.T Replacement

// T is an interface that abstracts the standard *testing.T. It allows
// verify functions to work with both standard tests and the continued
// testing wrapper.
type T interface {
	Errorf(format string, args ...any)
}

// continuedTesting is a wrapper around *testing.T that allows test
// verifications to continue even after a failure. It collects all
// failure messages and logs them without immediately calling t.FailNow().
// The total number of failures can be asserted at the end of the test
// using FailureCount.
type continuedTesting struct {
	*testing.T
	mu     sync.Mutex
	failed int
	msgs   []string
}

// Ensure the wrapper implement T
var _ T = (*continuedTesting)(nil)

func (ct *continuedTesting) Errorf(format string, args ...any) {
	ct.Helper()

	ct.mu.Lock()
	ct.failed++
	ct.msgs = append(ct.msgs, fmt.Sprintf(format, args...))
	msgs := make([]string, len(ct.msgs))
	copy(msgs, ct.msgs)
	ct.msgs = nil
	ct.mu.Unlock()

	for _, msg := range msgs {
		ct.T.Log(msg)
	}
}

// Library API

// ContinuedTesting wraps a *testing.T to create a T instance that allows
// verifications to continue after failures. This is useful for checking
// multiple independent conditions and reporting all failures at once.
func ContinuedTesting(t *testing.T) T {
	ct := &continuedTesting{
		T:      t,
		mu:     sync.Mutex{},
		failed: 0,
		msgs:   nil,
	}
	return ct
}

// IsContinued checks if a T is a *continuedTesting instance. This can be
// useful for conditional logic in tests.
func IsContinued(t T) bool {
	_, ok := t.(*continuedTesting)
	return ok
}

// FailureCount asserts that the number of failures recorded by a
// *continuedTesting instance matches the expected count. It fails
// the test if the counts do not match. This must be called at the
// end of a test using ContinuedTesting.
func FailureCount(t T, expected int) bool {
	var ct *continuedTesting
	var ok bool

	if ct, ok = t.(*continuedTesting); !ok {
		t.Errorf("t is no continued testing")
		return false
	}

	ct.mu.Lock()
	failed := ct.failed
	ct.mu.Unlock()

	if failed != expected {
		verificationFailure(t, "failure count", expected, failed)
		ct.T.Fail()
		return false
	}
	return true
}

// UTILS

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
