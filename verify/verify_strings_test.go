// Tideland Go Asserts - Verify - String Tests
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

// TestHasPrefix tests the HasPrefix verification function.
func TestHasPrefix(t *testing.T) {
	// Positive: has prefix
	verify.HasPrefix(t, "hello world", "hello")
	verify.HasPrefix(t, "testing", "test")
	verify.HasPrefix(t, "abc", "")
	verify.HasPrefix(t, "same", "same")

	// Create continuation testing for negative cases
	ct := verify.ContinuedTesting(t)

	// Negative: doesn't have prefix
	verify.HasPrefix(ct, "hello world", "world")
	verify.HasPrefix(ct, "testing", "ing")
	verify.HasPrefix(ct, "abc", "xyz")

	verify.FailureCount(ct, 3)
}

// TestHasSuffix tests the HasSuffix verification function.
func TestHasSuffix(t *testing.T) {
	// Positive: has suffix
	verify.HasSuffix(t, "hello world", "world")
	verify.HasSuffix(t, "testing", "ing")
	verify.HasSuffix(t, "abc", "")
	verify.HasSuffix(t, "same", "same")

	// Create continuation testing for negative cases
	ct := verify.ContinuedTesting(t)

	// Negative: doesn't have suffix
	verify.HasSuffix(ct, "hello world", "hello")
	verify.HasSuffix(ct, "testing", "test")
	verify.HasSuffix(ct, "abc", "xyz")

	verify.FailureCount(ct, 3)
}

// TestContainsIgnoreCase tests the ContainsIgnoreCase verification function.
func TestContainsIgnoreCase(t *testing.T) {
	// Positive: contains (case-insensitive)
	verify.ContainsIgnoreCase(t, "Hello World", "hello")
	verify.ContainsIgnoreCase(t, "Hello World", "WORLD")
	verify.ContainsIgnoreCase(t, "Hello World", "o W")
	verify.ContainsIgnoreCase(t, "TESTING", "test")
	verify.ContainsIgnoreCase(t, "testing", "TEST")

	// Create continuation testing for negative cases
	ct := verify.ContinuedTesting(t)

	// Negative: doesn't contain
	verify.ContainsIgnoreCase(ct, "Hello World", "goodbye")
	verify.ContainsIgnoreCase(ct, "testing", "xyz")

	verify.FailureCount(ct, 2)
}
