// Tideland Go Asserts - Verify - Unit Tests
//
// Copyright (C) 2024-2025 Frank Mueller / Tideland / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package verify_test

import (
	"bytes"
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"

	"tideland.dev/go/asserts/verify"
)

// Tests

// TestVerify ensures the correct error handling of the package.
func TestVerify(t *testing.T) {
	// t.Skip("Verification skip")

	t.Log("verification test started")

	ct := verify.ContinuedTesting(t)

	verify.True(ct, false)
	verify.False(ct, true)

	verify.FailureCount(ct, 2)
}

// TestImplements tests the Implements verification function.
func TestImplements(t *testing.T) {
	// Define test interface as nil value of interface type
	var stringer fmt.Stringer

	// Positive: type that implements interface
	var buf bytes.Buffer
	verify.Implements(t, &buf, &stringer)

	// Another type that implements Stringer (time.Time implements Stringer)
	now := time.Now()
	verify.Implements(t, now, &stringer)

	// Create continuation testing for negative cases
	ct := verify.ContinuedTesting(t)

	// Negative: type that doesn't implement interface
	verify.Implements(ct, 42, &stringer)
	verify.Implements(ct, "string", &stringer)

	// Negative: nil expected (not an interface)
	verify.Implements(ct, &buf, nil)

	// Negative: nil gotten
	verify.Implements(ct, nil, &stringer)

	// Negative: expected is not interface type pointer
	var notInterface int
	verify.Implements(ct, &buf, &notInterface)

	verify.FailureCount(ct, 5)
}

// TestErrorContains tests the ErrorContains verification function.
func TestErrorContains(t *testing.T) {
	testErr := errors.New("database connection failed: timeout after 30s")

	// Positive: error contains substring
	verify.ErrorContains(t, testErr, "database")
	verify.ErrorContains(t, testErr, "connection")
	verify.ErrorContains(t, testErr, "timeout")
	verify.ErrorContains(t, testErr, "30s")
	verify.ErrorContains(t, testErr, "database connection failed: timeout after 30s") // full match

	// Create continuation testing for negative cases
	ct := verify.ContinuedTesting(t)

	// Negative: nil error
	verify.ErrorContains(ct, nil, "anything")
	verify.ErrorContains(ct, nil, "database")

	// Negative: substring not in error
	verify.ErrorContains(ct, testErr, "network")
	verify.ErrorContains(ct, testErr, "SUCCESS")
	verify.ErrorContains(ct, testErr, "redis")

	verify.FailureCount(ct, 5)
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
	verify.About(t, 45, 43, 5)
	verify.About(t, 4.5, 4.3, 0.3)

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
	verify.About(ct, 45, 43, 1)
	verify.About(ct, 4.5, 4.3, 0.1)

	verify.FailureCount(ct, 10)
}

// TestLengths tests the Length, Empty, and NotEmpty verification functions.
func TestLengths(t *testing.T) {
	// Positive test cases with regular testing.T
	verify.Length(t, []int{1, 2, 3}, 3)
	verify.Length(t, "hello", 5)
	verify.Length(t, map[string]int{"a": 1, "b": 2}, 2)
	verify.Length(t, [2]bool{true, false}, 2)
	verify.Length(t, make(chan int, 5), 0)

	// Positive test cases for Empty and NotEmpty
	verify.Empty(t, []int{})
	verify.Empty(t, "")
	verify.Empty(t, map[string]int{})
	verify.Empty(t, make(chan int, 5))
	verify.NotEmpty(t, []int{1, 2, 3})
	verify.NotEmpty(t, "hello")
	verify.NotEmpty(t, map[string]int{"a": 1, "b": 2})

	// Create continuation testing instance for negative test cases
	ct := verify.ContinuedTesting(t)

	// Negative test cases with continuation testing
	verify.Length(ct, []int{1, 2, 3}, 4)
	verify.Length(ct, "hello", 6)
	verify.Length(ct, map[string]int{"a": 1, "b": 2}, 3)
	verify.Length(ct, [2]bool{true, false}, 1)

	// Negative test cases for Empty and NotEmpty
	verify.Empty(ct, []int{1, 2, 3})
	verify.Empty(ct, "hello")
	verify.Empty(ct, map[string]int{"a": 1})
	verify.NotEmpty(ct, []int{})
	verify.NotEmpty(ct, "")
	verify.NotEmpty(ct, map[string]int{})

	// Invalid type test case
	verify.Length(ct, 42, 0)
	verify.Empty(ct, 42)
	verify.NotEmpty(ct, 42)

	verify.FailureCount(ct, 13)
}

// TestContains tests the Contains and Substring verification functions.
func TestContains(t *testing.T) {
	// Positive test cases with regular testing.T
	verify.Contains(t, "world", []string{"hello", "world"})
	verify.Substring(t, "ello", "hello, world")
	verify.Contains(t, "hello", []string{"hello", "world"})
	verify.Substring(t, "", "hello, world")

	// Create continuation testing instance for negative test cases
	ct := verify.ContinuedTesting(t)

	// Negative test cases with continuation testing
	verify.Contains(ct, "universe", []string{"hello", "world"})
	verify.Contains(ct, "HELLO", []string{"hello", "world"})
	verify.Substring(ct, "hello, world", "HELLO")

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
	verify.Simultaneous(t, now, now)
	verify.Before(t, now, later)
	verify.After(t, now, earlier)
	verify.Between(t, now, earlier, later)

	// Create continuation testing instance
	ct := verify.ContinuedTesting(t)

	// Negative test cases with continuation testing
	verify.Simultaneous(ct, now, later)
	verify.Before(ct, now, earlier)
	verify.After(ct, now, later)
	verify.Between(ct, now.Add(5*time.Hour), later, earlier)

	verify.FailureCount(ct, 4)
}

// TestDuration tests the Duration verification function.
func TestDurations(t *testing.T) {
	// Positive test cases.
	verify.Shorter(t, time.Second, 2*time.Second)
	verify.Longer(t, 2*time.Second, time.Second)
	verify.About(t, 5*time.Second, 4*time.Second, 2*time.Second)

	// Create continuation testing instance
	ct := verify.ContinuedTesting(t)

	// Negative test cases with continuation testing
	verify.Shorter(ct, 2*time.Second, time.Second)
	verify.Longer(ct, time.Second, 2*time.Second)
	verify.About(ct, 2*time.Second, 5*time.Second, 500*time.Millisecond)

	verify.FailureCount(ct, 3)
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

// Define custom error types for AsError testing
type customError struct {
	msg string
}

func (e customError) Error() string {
	return e.msg
}

type anotherError struct{}

func (e anotherError) Error() string {
	return "another error"
}

// TestErrors tests the Error verification functions.
func TestErrors(t *testing.T) {
	testErr := errors.New("booom")

	customErr := customError{msg: "custom error"}
	wrappedErr := fmt.Errorf("wrapped: %w", customErr)
	doubleWrappedErr := fmt.Errorf("double wrapped: %w", wrappedErr)

	// Positive test cases with regular testing.T
	verify.Error(t, testErr)
	verify.NoError(t, nil)
	verify.IsError(t, testErr, testErr)
	verify.ErrorMatch(t, testErr, "^bo.*")

	// Test AsError with custom error types
	var targetCustom customError
	verify.AsError(t, customErr, &targetCustom)

	// Test UnwrapError
	verify.UnwrapError(t, wrappedErr, customErr)

	// Create continuation testing instance for negative test cases
	ct := verify.ContinuedTesting(t)

	// Negative test cases with continuation testing
	verify.Error(ct, nil)
	verify.NoError(ct, testErr)
	verify.IsError(ct, errors.New("ouch"), testErr)
	verify.ErrorMatch(ct, testErr, ".*ou$")

	// Test AsError negative cases
	var targetAnother anotherError
	verify.AsError(ct, customErr, &targetAnother) // wrong type
	verify.AsError(ct, nil, &targetCustom)        // nil error

	// Test UnwrapError negative cases
	verify.UnwrapError(ct, customErr, testErr)        // doesn't unwrap to expected
	verify.UnwrapError(ct, nil, testErr)              // nil error
	verify.UnwrapError(ct, doubleWrappedErr, testErr) // unwraps to wrong error

	verify.FailureCount(ct, 9)
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

// TestPanics tests the Panics verification function.
func TestPanics(t *testing.T) {
	// Positive: function that panics
	verify.Panics(t, func() { panic("test panic") })

	// Create continuation testing for negative cases
	ct := verify.ContinuedTesting(t)

	// Negative: function that doesn't panic
	verify.Panics(ct, func() {
		// do nothing - no panic
	})

	// Negative: nil function
	verify.Panics(ct, nil)

	verify.FailureCount(ct, 2)
}

// TestNotPanics tests the NotPanics verification function.
func TestNotPanics(t *testing.T) {
	// Positive: function that doesn't panic
	verify.NotPanics(t, func() {
		// safe code
		_ = 1 + 1
	})

	// Create continuation testing for negative cases
	ct := verify.ContinuedTesting(t)

	// Negative: function that panics
	verify.NotPanics(ct, func() {
		panic("boom")
	})

	// Negative: nil function
	verify.NotPanics(ct, nil)

	verify.FailureCount(ct, 2)
}

// TestConcurrentContinuedTesting tests thread-safety of continuedTesting.
func TestConcurrentContinuedTesting(t *testing.T) {
	ct := verify.ContinuedTesting(t)

	// Run multiple goroutines using same continued testing instance
	var wg sync.WaitGroup
	numGoroutines := 10

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			// Intentional failure - each goroutine adds one failure
			verify.Equal(ct, n, n+1)
		}(i)
	}

	wg.Wait()

	// All 10 failures should be counted correctly (tests thread safety)
	verify.FailureCount(ct, numGoroutines)
}

// TestBetweenBoundaries tests the Between function with boundary conditions.
func TestBetweenBoundaries(t *testing.T) {
	now := time.Now()
	earlier := now.Add(-1 * time.Hour)
	later := now.Add(1 * time.Hour)

	// Positive: exact boundary matches should pass (inclusive)
	verify.Between(t, earlier, earlier, later)
	verify.Between(t, later, earlier, later)
	verify.Between(t, now, earlier, later)

	// Positive: boundaries auto-swap if reversed
	verify.Between(t, now, later, earlier)

	// Create continuation testing for negative cases
	ct := verify.ContinuedTesting(t)

	// Negative: outside the range
	verify.Between(ct, earlier.Add(-1*time.Second), earlier, later)
	verify.Between(ct, later.Add(1*time.Second), earlier, later)

	verify.FailureCount(ct, 2)
}

// TestInRangeBoundaries tests InRange and OutOfRange with boundary conditions.
func TestInRangeBoundaries(t *testing.T) {
	// Positive: exact boundary matches should pass (inclusive)
	verify.InRange(t, 30, 30, 50)
	verify.InRange(t, 50, 30, 50)
	verify.InRange(t, 40, 30, 50)

	verify.InRange(t, 3.0, 3.0, 5.0)
	verify.InRange(t, 5.0, 3.0, 5.0)

	verify.InRange(t, 3*time.Second, 3*time.Second, 5*time.Second)
	verify.InRange(t, 5*time.Second, 3*time.Second, 5*time.Second)

	// Positive: boundaries auto-swap if reversed
	verify.InRange(t, 40, 50, 30)

	// Create continuation testing for negative cases
	ct := verify.ContinuedTesting(t)

	// Negative: outside range
	verify.InRange(ct, 29, 30, 50)
	verify.InRange(ct, 51, 30, 50)

	verify.FailureCount(ct, 2)
}

// TestOutOfRangeBoundaries tests OutOfRange with boundary conditions.
func TestOutOfRangeBoundaries(t *testing.T) {
	// Positive: outside the range
	verify.OutOfRange(t, 29, 30, 50)
	verify.OutOfRange(t, 51, 30, 50)

	// Create continuation testing for negative cases
	ct := verify.ContinuedTesting(t)

	// Negative: exact boundary matches should fail (boundaries are exclusive for OutOfRange)
	verify.OutOfRange(ct, 30, 30, 50)
	verify.OutOfRange(ct, 50, 30, 50)
	verify.OutOfRange(ct, 40, 30, 50)

	verify.FailureCount(ct, 3)
}

// TestAboutTolerance tests About function with tolerance validation.
func TestAboutTolerance(t *testing.T) {
	// Positive: valid tolerance
	verify.About(t, 45, 43, 5)
	verify.About(t, 4.5, 4.3, 0.3)
	verify.About(t, 5*time.Second, 4*time.Second, 2*time.Second)

	// Positive: zero tolerance
	verify.About(t, 10, 10, 0)

	// Create continuation testing for negative cases
	ct := verify.ContinuedTesting(t)

	// Negative: negative tolerance should fail
	verify.About(ct, 45, 43, -5)
	verify.About(ct, 4.5, 4.3, -0.3)
	verify.About(ct, 5*time.Second, 4*time.Second, -2*time.Second)

	verify.FailureCount(ct, 3)
}
