// Tideland Go Asserts - Verify - Comparisons
//
// Copyright (C) 2024-2025 Frank Mueller / Tideland / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package verify

import (
	"fmt"
	"testing"

	"golang.org/x/exp/constraints"
)

// Less checks if the gotten value is less than the expected one.
// Supports integers, floats, and time.Duration.
func Less[C constraints.Integer | constraints.Float](t T, gotten, expected C, infos ...string) bool {
	if gotten >= expected {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is less", expected, gotten, infos...)
		return false
	}
	return true
}

// More checks if the gotten value is more than the expected one.
// Supports integers, floats, and time.Duration.
func More[C constraints.Integer | constraints.Float](t T, gotten, expected C, infos ...string) bool {
	if gotten <= expected {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is more", expected, gotten, infos...)
		return false
	}
	return true
}

// Zero checks if the gotten value equals zero.
func Zero[N constraints.Integer | constraints.Float](t T, gotten N, infos ...string) bool {
	if gotten != 0 {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is zero", 0, gotten, infos...)
		return false
	}
	return true
}

// NotZero checks if the gotten value is not zero.
func NotZero[N constraints.Integer | constraints.Float](t T, gotten N, infos ...string) bool {
	if gotten == 0 {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is not zero", "non-zero", gotten, infos...)
		return false
	}
	return true
}

// Positive checks if the gotten value is positive (greater than zero).
func Positive[N constraints.Integer | constraints.Float](t T, gotten N, infos ...string) bool {
	if gotten <= 0 {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is positive", "> 0", gotten, infos...)
		return false
	}
	return true
}

// Negative checks if the gotten value is negative (less than zero).
func Negative[N constraints.Integer | constraints.Float](t T, gotten N, infos ...string) bool {
	if gotten >= 0 {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is negative", "< 0", gotten, infos...)
		return false
	}
	return true
}

// Even checks if the gotten integer value is even.
func Even[I constraints.Integer](t T, gotten I, infos ...string) bool {
	if gotten%2 != 0 {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is even", "even number", gotten, infos...)
		return false
	}
	return true
}

// Odd checks if the gotten integer value is odd.
func Odd[I constraints.Integer](t T, gotten I, infos ...string) bool {
	if gotten%2 == 0 {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is odd", "odd number", gotten, infos...)
		return false
	}
	return true
}

// About checks if the gotten values equal within a expected delta. Possible
// values are integers, floats, and time.Duration.
// The tolerance must be non-negative.
func About[C constraints.Integer | constraints.Float](t T, gotten, expected, tolerance C, infos ...string) bool {
	if tolerance < 0 {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is about equal", "tolerance >= 0", tolerance, infos...)
		return false
	}
	if gotten < expected-tolerance || gotten > expected+tolerance {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		expectedDescr := fmt.Sprintf("%v' +/- '%v'", expected, tolerance)
		verificationFailure(t, "is about equal", expectedDescr, gotten, infos...)
		return false
	}
	return true
}
