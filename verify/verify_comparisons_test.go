// Tideland Go Asserts - Verify - Comparison Tests
//
// Copyright (C) 2024-2025 Frank Mueller / Tideland / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package verify_test

import (
	"testing"
	"time"

	"tideland.dev/go/asserts/verify"
)

// TestZero tests the Zero verification function.
func TestZero(t *testing.T) {
	// Positive: zero values
	verify.Zero(t, 0)
	verify.Zero(t, 0.0)
	verify.Zero(t, int64(0))

	// Create continuation testing for negative cases
	ct := verify.ContinuedTesting(t)

	// Negative: non-zero values
	verify.Zero(ct, 1)
	verify.Zero(ct, -1)
	verify.Zero(ct, 0.1)

	verify.FailureCount(ct, 3)
}

// TestNotZero tests the NotZero verification function.
func TestNotZero(t *testing.T) {
	// Positive: non-zero values
	verify.NotZero(t, 1)
	verify.NotZero(t, -1)
	verify.NotZero(t, 0.1)
	verify.NotZero(t, -0.1)

	// Create continuation testing for negative cases
	ct := verify.ContinuedTesting(t)

	// Negative: zero value
	verify.NotZero(ct, 0)
	verify.NotZero(ct, 0.0)

	verify.FailureCount(ct, 2)
}

// TestPositive tests the Positive verification function.
func TestPositive(t *testing.T) {
	// Positive: positive values
	verify.Positive(t, 1)
	verify.Positive(t, 100)
	verify.Positive(t, 0.1)

	// Create continuation testing for negative cases
	ct := verify.ContinuedTesting(t)

	// Negative: zero and negative values
	verify.Positive(ct, 0)
	verify.Positive(ct, -1)
	verify.Positive(ct, -0.1)

	verify.FailureCount(ct, 3)
}

// TestNegative tests the Negative verification function.
func TestNegative(t *testing.T) {
	// Positive: negative values
	verify.Negative(t, -1)
	verify.Negative(t, -100)
	verify.Negative(t, -0.1)

	// Create continuation testing for negative cases
	ct := verify.ContinuedTesting(t)

	// Negative: zero and positive values
	verify.Negative(ct, 0)
	verify.Negative(ct, 1)
	verify.Negative(ct, 0.1)

	verify.FailureCount(ct, 3)
}

// TestEven tests the Even verification function.
func TestEven(t *testing.T) {
	// Positive: even values
	verify.Even(t, 0)
	verify.Even(t, 2)
	verify.Even(t, -2)
	verify.Even(t, 100)

	// Create continuation testing for negative cases
	ct := verify.ContinuedTesting(t)

	// Negative: odd values
	verify.Even(ct, 1)
	verify.Even(ct, 3)
	verify.Even(ct, -1)

	verify.FailureCount(ct, 3)
}

// TestOdd tests the Odd verification function.
func TestOdd(t *testing.T) {
	// Positive: odd values
	verify.Odd(t, 1)
	verify.Odd(t, 3)
	verify.Odd(t, -1)
	verify.Odd(t, 99)

	// Create continuation testing for negative cases
	ct := verify.ContinuedTesting(t)

	// Negative: even values
	verify.Odd(ct, 0)
	verify.Odd(ct, 2)
	verify.Odd(ct, -2)

	verify.FailureCount(ct, 3)
}

// TestNumericPredicatesWithDurations tests numeric predicates work with time.Duration.
func TestNumericPredicatesWithDurations(t *testing.T) {
	// Zero
	verify.Zero(t, time.Duration(0))

	// NotZero
	verify.NotZero(t, 1*time.Second)

	// Positive
	verify.Positive(t, 1*time.Second)

	// Negative
	verify.Negative(t, -1*time.Second)

	ct := verify.ContinuedTesting(t)

	// Zero should fail for non-zero duration
	verify.Zero(ct, 1*time.Second)

	verify.FailureCount(ct, 1)
}
