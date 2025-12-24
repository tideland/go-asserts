// Tideland Go Asserts - Verify - Predicates
//
// Copyright (C) 2024-2025 Frank Mueller / Tideland / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package verify

import (
	"fmt"
	"testing"
)

// All checks if all elements in the slice satisfy the predicate function.
// Returns true if all elements pass the predicate, false otherwise.
func All[S ~[]E, E any](t T, gotten S, predicate func(E) bool, infos ...string) bool {
	if ht, ok := t.(testing.TB); ok {
		ht.Helper()
	}

	if predicate == nil {
		verificationFailure(t, "all elements match predicate", "valid predicate", "nil predicate", infos...)
		return false
	}

	for i, elem := range gotten {
		if !predicate(elem) {
			msg := fmt.Sprintf("element at index %d failed predicate", i)
			verificationFailure(t, "all elements match predicate", "all pass", msg, infos...)
			return false
		}
	}
	return true
}

// Any checks if at least one element in the slice satisfies the predicate function.
// Returns true if any element passes the predicate, false if none do.
func Any[S ~[]E, E any](t T, gotten S, predicate func(E) bool, infos ...string) bool {
	if ht, ok := t.(testing.TB); ok {
		ht.Helper()
	}

	if predicate == nil {
		verificationFailure(t, "any element matches predicate", "valid predicate", "nil predicate", infos...)
		return false
	}

	for _, elem := range gotten {
		if predicate(elem) {
			return true
		}
	}

	verificationFailure(t, "any element matches predicate", "at least one pass", "none passed", infos...)
	return false
}

// None checks if no elements in the slice satisfy the predicate function.
// Returns true if no elements pass the predicate, false if any do.
func None[S ~[]E, E any](t T, gotten S, predicate func(E) bool, infos ...string) bool {
	if ht, ok := t.(testing.TB); ok {
		ht.Helper()
	}

	if predicate == nil {
		verificationFailure(t, "no elements match predicate", "valid predicate", "nil predicate", infos...)
		return false
	}

	for i, elem := range gotten {
		if predicate(elem) {
			msg := fmt.Sprintf("element at index %d passed predicate", i)
			verificationFailure(t, "no elements match predicate", "none pass", msg, infos...)
			return false
		}
	}
	return true
}
