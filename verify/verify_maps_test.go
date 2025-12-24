// Tideland Go Asserts - Verify - Map Tests
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

// TestMapEqual tests the MapEqual verification function.
func TestMapEqual(t *testing.T) {
	// Positive: equal maps
	verify.MapEqual(t, map[string]int{"a": 1, "b": 2}, map[string]int{"a": 1, "b": 2})
	verify.MapEqual(t, map[int]string{1: "a", 2: "b"}, map[int]string{1: "a", 2: "b"})
	verify.MapEqual(t, map[string]int{}, map[string]int{})

	// Create continuation testing for negative cases
	ct := verify.ContinuedTesting(t)

	// Negative: different maps
	verify.MapEqual(ct, map[string]int{"a": 1, "b": 2}, map[string]int{"a": 1, "b": 3})
	verify.MapEqual(ct, map[string]int{"a": 1}, map[string]int{"a": 1, "b": 2})
	verify.MapEqual(ct, map[string]int{"a": 1}, map[string]int{"b": 1})

	verify.FailureCount(ct, 3)
}

// TestMapContainsKey tests the MapContainsKey verification function.
func TestMapContainsKey(t *testing.T) {
	testMap := map[string]int{"a": 1, "b": 2, "c": 3}

	// Positive: key exists
	verify.MapContainsKey(t, testMap, "a")
	verify.MapContainsKey(t, testMap, "b")
	verify.MapContainsKey(t, testMap, "c")

	// Create continuation testing for negative cases
	ct := verify.ContinuedTesting(t)

	// Negative: key doesn't exist
	verify.MapContainsKey(ct, testMap, "d")
	verify.MapContainsKey(ct, testMap, "x")

	verify.FailureCount(ct, 2)
}

// TestMapContainsKeys tests the MapContainsKeys verification function.
func TestMapContainsKeys(t *testing.T) {
	testMap := map[string]int{"a": 1, "b": 2, "c": 3}

	// Positive: all keys exist
	verify.MapContainsKeys(t, testMap, []string{"a", "b"})
	verify.MapContainsKeys(t, testMap, []string{"a", "b", "c"})
	verify.MapContainsKeys(t, testMap, []string{})

	// Create continuation testing for negative cases
	ct := verify.ContinuedTesting(t)

	// Negative: some keys missing
	verify.MapContainsKeys(ct, testMap, []string{"a", "d"})
	verify.MapContainsKeys(ct, testMap, []string{"x", "y"})

	verify.FailureCount(ct, 2)
}

// TestMapContainsValue tests the MapContainsValue verification function.
func TestMapContainsValue(t *testing.T) {
	testMap := map[string]int{"a": 1, "b": 2, "c": 3}

	// Positive: value exists
	verify.MapContainsValue(t, testMap, 1)
	verify.MapContainsValue(t, testMap, 2)
	verify.MapContainsValue(t, testMap, 3)

	// Create continuation testing for negative cases
	ct := verify.ContinuedTesting(t)

	// Negative: value doesn't exist
	verify.MapContainsValue(ct, testMap, 4)
	verify.MapContainsValue(ct, testMap, 0)

	verify.FailureCount(ct, 2)
}
