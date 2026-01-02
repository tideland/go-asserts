// Tideland Go Asserts - Verify - Type Tests
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

// TestDeepEqual tests the DeepEqual verification function.
func TestDeepEqual(t *testing.T) {
	type nested struct {
		Value int
		Name  string
	}

	// Positive: deeply equal values
	verify.DeepEqual(t, []int{1, 2, 3}, []int{1, 2, 3})
	verify.DeepEqual(t, map[string]int{"a": 1, "b": 2}, map[string]int{"a": 1, "b": 2})
	verify.DeepEqual(t, nested{Value: 42, Name: "test"}, nested{Value: 42, Name: "test"})
	verify.DeepEqual(t, []nested{{Value: 1, Name: "a"}}, []nested{{Value: 1, Name: "a"}})

	// Create continuation testing for negative cases
	ct := verify.ContinuedTesting(t)

	// Negative: not deeply equal
	verify.DeepEqual(ct, []int{1, 2, 3}, []int{1, 2, 4})
	verify.DeepEqual(ct, map[string]int{"a": 1}, map[string]int{"a": 2})
	verify.DeepEqual(ct, nested{Value: 42, Name: "test"}, nested{Value: 42, Name: "other"})

	verify.FailureCount(ct, 3)
}

// TestSameType tests the SameType verification function.
func TestSameType(t *testing.T) {
	// Positive: same types
	verify.SameType(t, 42, 100)
	verify.SameType(t, "hello", "world")
	verify.SameType(t, []int{1}, []int{2, 3})
	verify.SameType(t, map[string]int{}, map[string]int{"a": 1})

	// Create continuation testing for negative cases
	ct := verify.ContinuedTesting(t)

	// Negative: different types
	verify.SameType(ct, 42, "42")
	verify.SameType(ct, 42, int64(42))
	verify.SameType(ct, []int{1}, []string{"1"})

	verify.FailureCount(ct, 3)
}

// TestSamePointer tests the SamePointer verification function.
func TestSamePointer(t *testing.T) {
	slice := []int{1, 2, 3}
	sameSlice := slice
	differentSlice := []int{1, 2, 3}

	ptr := new(int)
	samePtr := ptr
	differentPtr := new(int)

	// Positive: same pointer
	verify.SamePointer(t, slice, sameSlice)
	verify.SamePointer(t, ptr, samePtr)

	// Create continuation testing for negative cases
	ct := verify.ContinuedTesting(t)

	// Negative: different pointers
	verify.SamePointer(ct, slice, differentSlice)
	verify.SamePointer(ct, ptr, differentPtr)

	verify.FailureCount(ct, 2)
}

// TestNotSamePointer tests the NotSamePointer verification function.
func TestNotSamePointer(t *testing.T) {
	slice := []int{1, 2, 3}
	sameSlice := slice
	differentSlice := []int{1, 2, 3}

	ptr := new(int)
	samePtr := ptr
	differentPtr := new(int)

	// Positive: different pointers
	verify.NotSamePointer(t, slice, differentSlice)
	verify.NotSamePointer(t, ptr, differentPtr)

	// Create continuation testing for negative cases
	ct := verify.ContinuedTesting(t)

	// Negative: same pointer
	verify.NotSamePointer(ct, slice, sameSlice)
	verify.NotSamePointer(ct, ptr, samePtr)

	verify.FailureCount(ct, 2)
}
