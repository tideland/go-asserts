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

// -----------------------------------------------------------------------------
// Tests
// -----------------------------------------------------------------------------

// TestVerify ensures the correct error handling of the package.
func TestVerify(t *testing.T) {
	// t.Skip("Verification skip")

	t.Log("verification test started")

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
	verify.Different(t, "world", "hello")

	// Negative test cases
	verify.Equal(ct, "world", "hello")
	verify.Different(ct, "same", "same")

	verify.FailureCount(ct, 2)
}

// TestComparisons tests the Equal, Different, Less, More, and AboutEqual verification functions.
func TestComparisons(t *testing.T) {
	// Positive test cases with regular testing.T
	verify.Equal(t, 42, 42)
	verify.Equal(t, "foo", "foo")
	verify.Different(t, 43, 42)
	verify.Different(t, "bar", "foo")
	verify.Less(t, 5, 10)
	verify.Less(t, 5.0, 10.0)
	verify.More(t, 10, 5)
	verify.More(t, 10.0, 5.0)
	verify.AboutEqual(t, 45, 43, 5)
	verify.AboutEqual(t, 4.5, 4.3, 0.3)

	// Create continuation testing instance for negative test cases
	ct := verify.ContinuedTesting(t)

	// Negative test cases with continuation testing
	verify.Equal(ct, 43, 42)
	verify.Equal(ct, "bar", "foo")
	verify.Different(ct, 42, 42)
	verify.Different(ct, "foo", "foo")
	verify.Less(ct, 10, 5)
	verify.Less(ct, 10.0, 5.0)
	verify.More(ct, 5, 10)
	verify.More(ct, 5.0, 10.0)
	verify.AboutEqual(ct, 45, 43, 1)
	verify.AboutEqual(ct, 4.5, 4.3, 0.1)

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
	verify.Contains(t, "hello world", "world")
	verify.Contains(t, "hello world", "hello")
	verify.Contains(t, "hello world", "") // Empty string is contained in any string

	// Create continuation testing instance for negative test cases
	ct := verify.ContinuedTesting(t)

	// Negative test cases with continuation testing
	verify.Contains(ct, "hello world", "universe")
	verify.Contains(ct, "hello world", "HELLO") // Case-sensitive check

	verify.FailureCount(ct, 2)
}

// TestContainsAny tests the ContainsAny verification function.
func TestContainsAny(t *testing.T) {
	// Positive test cases with regular testing.T
	verify.ContainsAny(t, "hello world", []string{"world", "universe"})
	verify.ContainsAny(t, "hello world", []string{"hello", "hi"})
	verify.ContainsAny(t, "hello world", []string{"planet", "world"})

	// Create continuation testing instance for negative test cases
	ct := verify.ContinuedTesting(t)

	// Negative test cases with continuation testing
	verify.ContainsAny(ct, "hello world", []string{"universe", "planet"})
	verify.ContainsAny(ct, "hello world", []string{"HELLO", "WORLD"}) // Case-sensitive check
	verify.ContainsAny(ct, "hello world", []string{})                 // Empty slice will always fail

	verify.FailureCount(ct, 3)
}

// TestMatch tests the Match function.
func TestMatch(t *testing.T) {
	// Positive test cases with regular testing.T
	verify.Match(t, "hello world", "^hello")
	verify.Match(t, "hello world", "world$")
	verify.Match(t, "hello", "h.llo")

	// Create continuation testing instance for negative test cases
	ct := verify.ContinuedTesting(t)

	// Test cases where the regular expression should not match the actual string.
	verify.Match(ct, "hello world", "^world")
	verify.Match(ct, "hello world", "hello$")
	verify.Match(ct, "world", "h.llo")

	// Test case where the regular expression compilation should fail.
	verify.Match(ct, "hello world", "[invalid")

	verify.FailureCount(ct, 4)
}

// TestTimes tests the Time verification functions.
func TestTimes(t *testing.T) {
	now := time.Now()
	later := now.Add(2 * time.Hour)
	earlier := now.Add(-2 * time.Hour)

	// Positive test cases
	verify.Before(t, now, later)
	verify.After(t, now, earlier)
	verify.Between(t, now, earlier, later)

	// Create continuation testing instance
	ct := verify.ContinuedTesting(t)

	// Negative test cases with continuation testing
	verify.Before(ct, now, earlier)
	verify.After(ct, now, later)
	verify.Between(ct, now.Add(5*time.Hour), later, earlier)

	verify.FailureCount(ct, 3)
}

// TestDuration tests the Duration verification function.
func TestDurations(t *testing.T) {
	// Positive test cases.
	verify.Shorter(t, time.Second, 2*time.Second)
	verify.Longer(t, 2*time.Second, time.Second)

	// Create continuation testing instance
	ct := verify.ContinuedTesting(t)

	// Negative test cases with continuation testing
	verify.Shorter(ct, 2*time.Second, time.Second)
	verify.Longer(ct, time.Second, 2*time.Second)

	verify.FailureCount(ct, 2)
}

// TestRange tests the Ranges assertions.
func TestRange(t *testing.T) {
	// Positive test cases
	verify.InRange(t, 40, 30, 50)
	verify.OutOfRange(t, 50, 30, 40)
	verify.InRange(t, 4.0, 3.0, 5.0)
	verify.OutOfRange(t, 5.0, 3.0, 4.0)
	verify.InRange(t, 4*time.Second, 3*time.Second, 5*time.Second)
	verify.OutOfRange(t, 5*time.Second, 3*time.Second, 4*time.Second)

	// Create continuation testing instance for negative test cases
	ct := verify.ContinuedTesting(t)

	// Negative test cases.
	verify.InRange(ct, 50, 30, 40)
	verify.OutOfRange(ct, 40, 30, 50)
	verify.InRange(ct, 5.0, 3.0, 4.0)
	verify.OutOfRange(ct, 4.0, 3.0, 5.0)
	verify.InRange(ct, 5*time.Second, 3*time.Second, 4*time.Second)
	verify.OutOfRange(ct, 4*time.Second, 3*time.Second, 5*time.Second)

	verify.FailureCount(ct, 6)
}

// TestErrors tests the Error verification functions.
func TestErrors(t *testing.T) {
	testErr := errors.New("booom")

	// Positive test cases with regular testing.T
	verify.Error(t, testErr)
	verify.NoError(t, nil)
	verify.IsError(t, testErr, testErr)
	verify.ErrorMatch(t, testErr, "^bo.*")

	// Create continuation testing instance for negative test cases
	ct := verify.ContinuedTesting(t)

	// Negative test cases with continuation testing
	verify.Error(ct, nil)
	verify.NoError(ct, testErr)
	verify.IsError(ct, errors.New("ouch"), testErr)
	verify.ErrorMatch(ct, testErr, ".*ou$")

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
			verify.Equal(t, positive.got, positive.expect)
		}
	})

	// Negative test cases with continuation testing
	ct := verify.ContinuedTesting(t)

	t.Log("negative tests")
	t.Run("negatives", func(t *testing.T) {
		t.Log("inside negative tests")
		for _, negative := range negatives {
			verify.Equal(ct, negative.got, negative.expect)
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
