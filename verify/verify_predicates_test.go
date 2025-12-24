// Tideland Go Asserts - Verify - Predicate Tests
//
// Copyright (C) 2024-2025 Frank Mueller / Tideland / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package verify_test

import (
	"testing"

	"tideland.dev/go/asserts/verify"
)

// TestAll tests the All verification function.
func TestAll(t *testing.T) {
	isPositive := func(n int) bool { return n > 0 }
	isEven := func(n int) bool { return n%2 == 0 }

	// Positive: all elements pass
	verify.All(t, []int{1, 2, 3, 4, 5}, isPositive)
	verify.All(t, []int{2, 4, 6, 8}, isEven)
	verify.All(t, []int{}, isPositive) // empty slice passes

	// Create continuation testing for negative cases
	ct := verify.ContinuedTesting(t)

	// Negative: not all elements pass
	verify.All(ct, []int{1, 2, -3, 4}, isPositive)
	verify.All(ct, []int{2, 4, 5, 8}, isEven)

	// Negative: nil predicate
	verify.All(ct, []int{1, 2, 3}, nil)

	verify.FailureCount(ct, 3)
}

// TestAny tests the Any verification function.
func TestAny(t *testing.T) {
	isPositive := func(n int) bool { return n > 0 }
	isEven := func(n int) bool { return n%2 == 0 }

	// Positive: at least one element passes
	verify.Any(t, []int{-1, -2, 3, -4}, isPositive)
	verify.Any(t, []int{1, 3, 5, 6}, isEven)
	verify.Any(t, []int{5}, isPositive)

	// Create continuation testing for negative cases
	ct := verify.ContinuedTesting(t)

	// Negative: no elements pass
	verify.Any(ct, []int{-1, -2, -3}, isPositive)
	verify.Any(ct, []int{1, 3, 5}, isEven)
	verify.Any(ct, []int{}, isPositive) // empty slice fails

	// Negative: nil predicate
	verify.Any(ct, []int{1, 2, 3}, nil)

	verify.FailureCount(ct, 4)
}

// TestNone tests the None verification function.
func TestNone(t *testing.T) {
	isNegative := func(n int) bool { return n < 0 }
	isOdd := func(n int) bool { return n%2 != 0 }

	// Positive: no elements pass
	verify.None(t, []int{1, 2, 3, 4}, isNegative)
	verify.None(t, []int{2, 4, 6, 8}, isOdd)
	verify.None(t, []int{}, isNegative) // empty slice passes

	// Create continuation testing for negative cases
	ct := verify.ContinuedTesting(t)

	// Negative: at least one element passes
	verify.None(ct, []int{1, 2, -3, 4}, isNegative)
	verify.None(ct, []int{2, 4, 5, 8}, isOdd)

	// Negative: nil predicate
	verify.None(ct, []int{1, 2, 3}, nil)

	verify.FailureCount(ct, 3)
}

// TestPredicatesWithStrings tests predicate functions with string slices.
func TestPredicatesWithStrings(t *testing.T) {
	isLong := func(s string) bool { return len(s) > 5 }

	// All
	verify.All(t, []string{"hello!", "world!", "testing"}, isLong)

	// Any
	verify.Any(t, []string{"hi", "hello!", "yo"}, isLong)

	// None
	verify.None(t, []string{"hi", "yo", "ok"}, isLong)

	ct := verify.ContinuedTesting(t)

	// All should fail
	verify.All(ct, []string{"hi", "hello!", "yo"}, isLong)

	verify.FailureCount(ct, 1)
}
