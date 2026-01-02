// Tideland Go Asserts - Verify - Ranges
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

// InRange checks if the given value is within lower and upper bounds.
// The boundaries are inclusive (gotten can equal expectedLower or expectedUpper).
// If expectedLower is greater than expectedUpper, they will be automatically swapped.
// Possible values are integers, floats, and time.Duration.
func InRange[C constraints.Integer | constraints.Float](t T, gotten, expectedLower, expectedUpper C, infos ...string) bool {
	if expectedLower > expectedUpper {
		expectedLower, expectedUpper = expectedUpper, expectedLower
	}
	if gotten < expectedLower || gotten > expectedUpper {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		expectedDescr := fmt.Sprintf("'%v' to '%v'", expectedLower, expectedUpper)
		verificationFailure(t, "is in range", expectedDescr, gotten, infos...)
		return false
	}
	return true
}

// OutOfRange checks if the given value is outside lower and upper bounds.
// It's the opposite of InRange. The boundaries are exclusive (gotten cannot equal
// expectedLower or expectedUpper to be considered out of range).
// If expectedLower is greater than expectedUpper, they will be automatically swapped.
func OutOfRange[C constraints.Integer | constraints.Float](t T, gotten, expectedLower, expectedUpper C, infos ...string) bool {
	if expectedLower > expectedUpper {
		expectedLower, expectedUpper = expectedUpper, expectedLower
	}
	if gotten >= expectedLower && gotten <= expectedUpper {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		expectedDescr := fmt.Sprintf("'%v' to '%v'", expectedLower, expectedUpper)
		verificationFailure(t, "is out of range", expectedDescr, gotten, infos...)
		return false
	}
	return true
}
