// Tideland Go Asserts - Verify - Slices
//
// Copyright (C) 2024-2025 Frank Mueller / Tideland / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package verify

import (
	"fmt"
	"slices"
	"testing"

	"golang.org/x/exp/constraints"
)

// Contains checks if the slice contains the expected element.
func Contains[S ~[]E, E comparable](t T, gotten E, expected S, infos ...string) bool {
	if !slices.Contains(expected, gotten) {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "contains", expected, gotten, infos...)
		return false
	}
	return true
}

// SliceEqual checks if two slices are deeply equal.
func SliceEqual[S ~[]E, E comparable](t T, gotten, expected S, infos ...string) bool {
	if !slices.Equal(gotten, expected) {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "slices are equal", expected, gotten, infos...)
		return false
	}
	return true
}

// SliceContainsAll checks if the haystack slice contains all elements from the needles slice.
func SliceContainsAll[S ~[]E, E comparable](t T, haystack S, needles S, infos ...string) bool {
	for _, needle := range needles {
		if !slices.Contains(haystack, needle) {
			if ht, ok := t.(testing.TB); ok {
				ht.Helper()
			}
			verificationFailure(t, "slice contains all", needles, fmt.Sprintf("missing: %v", needle), infos...)
			return false
		}
	}
	return true
}

// SliceContainsAny checks if the haystack slice contains any element from the needles slice.
func SliceContainsAny[S ~[]E, E comparable](t T, haystack S, needles S, infos ...string) bool {
	for _, needle := range needles {
		if slices.Contains(haystack, needle) {
			return true
		}
	}
	if ht, ok := t.(testing.TB); ok {
		ht.Helper()
	}
	verificationFailure(t, "slice contains any", needles, "none found", infos...)
	return false
}

// SliceSorted checks if the slice is sorted in ascending order.
func SliceSorted[S ~[]E, E constraints.Ordered](t T, gotten S, infos ...string) bool {
	if !slices.IsSorted(gotten) {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "slice is sorted", "sorted slice", gotten, infos...)
		return false
	}
	return true
}

// SliceSortedFunc checks if the slice is sorted according to a comparison function.
// The comparison function should return -1 if a < b, 0 if a == b, and 1 if a > b.
func SliceSortedFunc[S ~[]E, E any](t T, gotten S, cmp func(E, E) int, infos ...string) bool {
	if !slices.IsSortedFunc(gotten, cmp) {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "slice is sorted (custom)", "sorted slice", gotten, infos...)
		return false
	}
	return true
}
