// -----------------------------------------------------------------------------
// Convenient verification of unit tests in Go libraries and applications.
//
// Unit tests
//
// Copyright (C) 2024-2025 Frank Mueller / Oldenburg / Germany / Earth
// -----------------------------------------------------------------------------

package verify_test

import (
	"errors"
	"testing"
	"time"

	"tideland.dev/go/asserts/verify"
)

// TestVerify ensures the correct error handling of the package.
func TestVerify(t *testing.T) {
	// t.Skip("Verification skip")

	t.Log("verification test started")
	// verify.True(t, false)

	ct := verify.ContinuedTesting(t)

	verify.True(ct, false)
	verify.False(ct, true)

	verify.FailureCount(ct, 2)
}

// TestBoolean tests the True and False verification functions.
func TestBoolean(t *testing.T) {
	// Standard testing
	verify.True(t, true)
	verify.False(t, false)

	// Continued testing
	ct := verify.ContinuedTesting(t)

	verify.True(ct, false)
	verify.True(ct, true)
	verify.False(ct, true)
	verify.False(ct, false)

	verify.FailureCount(ct, 2)
}

// TestNils tests the Nil and NotNil verification functions.
func TestNils(t *testing.T) {
	// Create continuation testing instances
	ct := verify.ContinuedTesting(t)

	// Positive test cases.
	verify.Nil(t, nil)
	verify.NotNil(t, "not nil")

	// Negative test cases.
	verify.Nil(ct, "not nil")
	verify.NotNil(ct, nil)

	verify.FailureCount(ct, 2)
}

// TestStrings tests the Equal and Different verification functions for strings.
func TestStrings(t *testing.T) {
	// Create continuation testing instance
	ct := verify.ContinuedTesting(t)

	// Positive test cases
	verify.Equal(t, "hello", "hello")
	verify.Different(t, "hello", "world")

	// Negative test cases
	verify.Equal(ct, "hello", "world")
	verify.Different(ct, "same", "same")

	verify.FailureCount(ct, 2)
}

// TestComparisons tests the Equal, Different, Less, More, and AboutEqual verification functions.
func TestComparisons(t *testing.T) {
	// Positive test cases with regular testing.T
	verify.Equal(t, 42, 42)
	verify.Equal(t, "foo", "foo")
	verify.Different(t, 42, 43)
	verify.Different(t, "foo", "bar")
	verify.Less(t, 10, 5)
	verify.Less(t, 10.0, 5.0)
	verify.More(t, 5, 10)
	verify.More(t, 5.0, 10.0)
	verify.AboutEqual(t, 43, 45, 5)
	verify.AboutEqual(t, 4.3, 4.5, 0.3)

	// Create continuation testing instance for negative test cases
	ct := verify.ContinuedTesting(t)

	// Negative test cases with continuation testing
	verify.Equal(ct, 42, 43)
	verify.Equal(ct, "foo", "bar")
	verify.Different(ct, 42, 42)
	verify.Different(ct, "foo", "foo")
	verify.Less(ct, 5, 10)
	verify.Less(ct, 5.0, 10.0)
	verify.More(ct, 10, 5)
	verify.More(ct, 10.0, 5.0)
	verify.AboutEqual(ct, 43, 45, 1)
	verify.AboutEqual(ct, 4.3, 4.5, 0.1)

	verify.FailureCount(ct, 10)
}

// TestLengths tests the Length verification function.
func TestLengths(t *testing.T) {
	// Positive test cases with regular testing.T
	verify.Length(t, []int{1, 2, 3}, 3)
	verify.Length(t, "hello", 5)
	verify.Length(t, map[string]int{"a": 1, "b": 2}, 2)
	verify.Length(t, [2]bool{true, false}, 2)
	verify.Length(t, make(chan int, 5), 0)

	// Create continuation testing instance for negative test cases
	ct := verify.ContinuedTesting(t)

	// Negative test cases with continuation testing
	verify.Length(ct, []int{1, 2, 3}, 4)
	verify.Length(ct, "hello", 6)
	verify.Length(ct, map[string]int{"a": 1, "b": 2}, 3)
	verify.Length(ct, [2]bool{true, false}, 1)

	// Invalid type test case
	verify.Length(ct, 42, 0)

	verify.FailureCount(ct, 5)
}

// TestContains tests the Contains verification function.
func TestContains(t *testing.T) {
	// Positive test cases with regular testing.T
	verify.Contains(t, "world", "hello world")
	verify.Contains(t, "hello", "hello world")
	verify.Contains(t, "", "hello world") // Empty string is contained in any string

	// Create continuation testing instance for negative test cases
	ct := verify.ContinuedTesting(t)

	// Negative test cases with continuation testing
	verify.Contains(ct, "universe", "hello world")
	verify.Contains(ct, "HELLO", "hello world") // Case-sensitive check

	verify.FailureCount(ct, 2)
}

// TestContainsAny tests the ContainsAny verification function.
func TestContainsAny(t *testing.T) {
	// Positive test cases with regular testing.T
	verify.ContainsAny(t, []string{"world", "universe"}, "hello world")
	verify.ContainsAny(t, []string{"hello", "hi"}, "hello world")
	verify.ContainsAny(t, []string{"planet", "world"}, "hello world")

	// Create continuation testing instance for negative test cases
	ct := verify.ContinuedTesting(t)

	// Negative test cases with continuation testing
	verify.ContainsAny(ct, []string{"universe", "planet"}, "hello world")
	verify.ContainsAny(ct, []string{"HELLO", "WORLD"}, "hello world") // Case-sensitive check
	verify.ContainsAny(ct, []string{}, "hello world") // Empty slice will always fail

	verify.FailureCount(ct, 3)
}

// TestMatch tests the Match function.
func TestMatch(t *testing.T) {
	// Positive test cases with regular testing.T
	verify.Match(t, "^hello", "hello world")
	verify.Match(t, "world$", "hello world")
	verify.Match(t, "h.llo", "hello")

	// Create continuation testing instance for negative test cases
	ct := verify.ContinuedTesting(t)

	// Test cases where the regular expression should not match the actual string.
	verify.Match(ct, "^world", "hello world")
	verify.Match(ct, "hello$", "hello world")
	verify.Match(ct, "h.llo", "world")

	// Test case where the regular expression compilation should fail.
	verify.Match(ct, "[invalid", "hello world")

	verify.FailureCount(ct, 4)
}

// TestTimes tests the Time verification functions.
func TestTimes(t *testing.T) {
	now := time.Now()
	later := now.Add(2 * time.Hour)
	earlier := now.Add(-2 * time.Hour)

	// Positive test cases
	verify.Before(t, later, now)
	verify.After(t, earlier, now)
	verify.Between(t, earlier, later, now)

	// Create continuation testing instance
	ct := verify.ContinuedTesting(t)

	// Negative test cases with continuation testing
	verify.Before(ct, earlier, now)
	verify.After(ct, later, now)
	verify.Between(ct, later, earlier, now.Add(5*time.Hour))

	verify.FailureCount(ct, 3)
}

// TestDuration tests the Duration verification function.
func TestDurations(t *testing.T) {
	// Positive test cases.
	verify.Shorter(t, 2*time.Second, time.Second)
	verify.Longer(t, time.Second, 2*time.Second)

	// Create continuation testing instance
	ct := verify.ContinuedTesting(t)

	// Negative test cases with continuation testing
	verify.Shorter(ct, time.Second, 2*time.Second)
	verify.Longer(ct, 2*time.Second, time.Second)

	verify.FailureCount(ct, 2)
}

// TestRange tests the Ranges assertions.
func TestRange(t *testing.T) {
	// Positive test cases
	verify.InRange(t, 30, 50, 40)
	verify.OutOfRange(t, 30, 40, 50)
	verify.InRange(t, 3.0, 5.0, 4.0)
	verify.OutOfRange(t, 3.0, 4.0, 5.0)
	verify.InRange(t, 3*time.Second, 5*time.Second, 4*time.Second)
	verify.OutOfRange(t, 3*time.Second, 4*time.Second, 5*time.Second)

	// Create continuation testing instance for negative test cases
	ct := verify.ContinuedTesting(t)

	// Negative test cases.
	verify.InRange(ct, 30, 40, 50)
	verify.OutOfRange(ct, 30, 50, 40)
	verify.InRange(ct, 3.0, 4.0, 5.0)
	verify.OutOfRange(ct, 3.0, 5.0, 4.0)
	verify.InRange(ct, 3*time.Second, 4*time.Second, 5*time.Second)
	verify.OutOfRange(ct, 3*time.Second, 5*time.Second, 4*time.Second)

	verify.FailureCount(ct, 6)
}

// TestErrors tests the Error verification functions.
func TestErrors(t *testing.T) {
	testErr := errors.New("booom")

	// Positive test cases with regular testing.T
	verify.Error(t, testErr)
	verify.NoError(t, nil)
	verify.IsError(t, testErr, testErr)
	verify.ErrorMatch(t, "^bo.*", testErr)

	// Create continuation testing instance for negative test cases
	ct := verify.ContinuedTesting(t)

	// Negative test cases with continuation testing
	verify.Error(ct, nil)
	verify.NoError(ct, testErr)
	verify.IsError(ct, testErr, errors.New("ouch"))
	verify.ErrorMatch(ct, ".*ou$", testErr)

	verify.FailureCount(ct, 4)
}

// TestRun tests the Run function.
func TestRun(t *testing.T) {
	positives := []struct {
		name   string
		expect int
		got    int
	}{
		{"a", 1, 1},
		{"b", 2, 2},
		{"c", 3, 3},
	}

	negatives := []struct {
		name   string
		expect int
		got    int
	}{
		{"d", 4, 5},
		{"e", 6, 7},
		{"f", 8, 9},
	}

	// Positive test cases with continuation testing
	t.Log("positive tests")
	t.Run("positives", func(t *testing.T) {
		t.Log("inside positive tests")
		for _, positive := range positives {
			verify.Equal(t, positive.expect, positive.got)
		}
	})

	// Negative test cases with continuation testing
	ct := verify.ContinuedTesting(t)

	t.Log("negative tests")
	t.Run("negatives", func(t *testing.T) {
		t.Log("inside negative tests")
		for _, negative := range negatives {
			verify.Equal(ct, negative.expect, negative.got)
		}
	})

	verify.FailureCount(ct, 3)
}

// TestAssignability tests the Assignability verification function.
func TestAssignability(t *testing.T) {
	type tt struct {
		A int
		B string
	}

	var x int
	var y string
	var z tt

	// Positive test cases with regular testing.T
	verify.Assignability(t, x, 2)
	verify.Assignability(t, y, "Hello, World!")
	verify.Assignability(t, z, tt{1, "test"})

	// Create continuation testing instance for negative test cases
	ct := verify.ContinuedTesting(t)

	// Negative test cases with continuation testing
	verify.Assignability(ct, x, "Hello, World!")
	verify.Assignability(ct, y, tt{0, "no"})
	verify.Assignability(ct, z, "done")

	verify.FailureCount(ct, 3)
}

// TestIsContinue tests the IsContinueT function.
func TestIsContinue(t *testing.T) {
	// Create a continue testing instance
	ct := verify.ContinuedTesting(t)

	// Check that it is identified correctly
	if !verify.IsContinued(ct) {
		t.Error("IsContinueT should recognize ContinueTesting instance")
	}

	// Regular testing.T should not be recognized
	if verify.IsContinued(t) {
		t.Error("IsContinueT should not recognize regular testing.T")
	}

	// Check that we had expected failures
	verify.FailureCount(ct, 0)
}

// -----------------------------------------------------------------------------
// EOF
// -----------------------------------------------------------------------------
