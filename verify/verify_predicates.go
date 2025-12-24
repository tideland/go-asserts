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

// All checks if every element in a slice satisfies the provided predicate
// function. It passes for an empty slice.
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

// Any checks if at least one element in a slice satisfies the provided
// predicate function. It fails for an empty slice.
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

// None checks if no element in a slice satisfies the provided predicate
// function. It passes for an empty slice.
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
