// Tideland Go Asserts - Verify - Slice Tests
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

// TestSliceEqual tests the SliceEqual verification function.
func TestSliceEqual(t *testing.T) {
	// Positive: equal slices
	verify.SliceEqual(t, []int{1, 2, 3}, []int{1, 2, 3})
	verify.SliceEqual(t, []string{"a", "b", "c"}, []string{"a", "b", "c"})
	verify.SliceEqual(t, []int{}, []int{})

	// Create continuation testing for negative cases
	ct := verify.ContinuedTesting(t)

	// Negative: different slices
	verify.SliceEqual(ct, []int{1, 2, 3}, []int{1, 2, 4})
	verify.SliceEqual(ct, []int{1, 2, 3}, []int{1, 2})
	verify.SliceEqual(ct, []string{"a", "b"}, []string{"a", "c"})

	verify.FailureCount(ct, 3)
}

// TestSliceContainsAll tests the SliceContainsAll verification function.
func TestSliceContainsAll(t *testing.T) {
	// Positive: haystack contains all needles
	verify.SliceContainsAll(t, []int{1, 2, 3, 4, 5}, []int{2, 4})
	verify.SliceContainsAll(t, []string{"a", "b", "c"}, []string{"a", "c"})
	verify.SliceContainsAll(t, []int{1, 2, 3}, []int{})

	// Create continuation testing for negative cases
	ct := verify.ContinuedTesting(t)

	// Negative: missing elements
	verify.SliceContainsAll(ct, []int{1, 2, 3}, []int{4, 5})
	verify.SliceContainsAll(ct, []string{"a", "b"}, []string{"c"})

	verify.FailureCount(ct, 2)
}

// TestSliceContainsAny tests the SliceContainsAny verification function.
func TestSliceContainsAny(t *testing.T) {
	// Positive: haystack contains at least one needle
	verify.SliceContainsAny(t, []int{1, 2, 3}, []int{2, 5})
	verify.SliceContainsAny(t, []string{"a", "b", "c"}, []string{"x", "c"})

	// Create continuation testing for negative cases
	ct := verify.ContinuedTesting(t)

	// Negative: no elements found
	verify.SliceContainsAny(ct, []int{1, 2, 3}, []int{4, 5, 6})
	verify.SliceContainsAny(ct, []string{"a", "b"}, []string{"x", "y"})
	verify.SliceContainsAny(ct, []int{1, 2, 3}, []int{})

	verify.FailureCount(ct, 3)
}

// TestSliceSorted tests the SliceSorted verification function.
func TestSliceSorted(t *testing.T) {
	// Positive: sorted slices
	verify.SliceSorted(t, []int{1, 2, 3, 4, 5})
	verify.SliceSorted(t, []string{"a", "b", "c"})
	verify.SliceSorted(t, []int{})
	verify.SliceSorted(t, []int{1})

	// Create continuation testing for negative cases
	ct := verify.ContinuedTesting(t)

	// Negative: unsorted slices
	verify.SliceSorted(ct, []int{3, 1, 2})
	verify.SliceSorted(ct, []string{"c", "a", "b"})

	verify.FailureCount(ct, 2)
}

// TestSliceSortedFunc tests the SliceSortedFunc verification function.
func TestSliceSortedFunc(t *testing.T) {
	// Descending comparison function
	descending := func(a, b int) int {
		if a > b {
			return -1
		}
		if a < b {
			return 1
		}
		return 0
	}

	// Positive: sorted with custom function
	verify.SliceSortedFunc(t, []int{5, 4, 3, 2, 1}, descending)

	// Create continuation testing for negative cases
	ct := verify.ContinuedTesting(t)

	// Negative: not sorted according to custom function
	verify.SliceSortedFunc(ct, []int{1, 2, 3, 4, 5}, descending)

	verify.FailureCount(ct, 1)
}
