// -----------------------------------------------------------------------------
// asserts for more convinient testing - tests
//
// Copyright (C) 2025 Frank Mueller / Oldenburg / Germany / Earth
// -----------------------------------------------------------------------------

package asserts_test

import (
	"errors"
	"testing"
	"time"

	"tideland.dev/go/asserts"
)

// TestLogf tests the log function.
func TestLogf(t *testing.T) {
	asserts.Logf(t, "Hello, %s!", "World")
}

// TestFail tests the fail behavior of the tester.
func TestFail(t *testing.T) {
	tester := asserts.NewTester(t, asserts.FAIL)

	// fail set to true will cause the test to fail to
	// verify that the tester is working correctly. After
	// the test is run, the fail variable should be set to
	// false to prevent the test from failing.
	// fail := true
	fail := false

	if fail {
		asserts.False(tester, fail)
	}
}

// TestBooleans tests the True and False assertions.
func TestBooleans(t *testing.T) {
	tester := asserts.NewTester(t, asserts.CONTINUE)

	// Positive test cases.
	asserts.True(tester, true)
	asserts.False(tester, false)

	// Negative test cases.
	asserts.True(tester, false)
	asserts.False(tester, true)

	// Check the number of failures and if messages are correct.
	asserts.Failures(tester, 2)
	asserts.FailureMatch(tester, 0, "true.*false")
	asserts.FailureMatch(tester, 1, "false.*true")
}

// TestNils tests the Nil and NotNil assertions.
func TestNils(t *testing.T) {
	tester := asserts.NewTester(t, asserts.CONTINUE)

	// Positive test cases.
	asserts.Nil(tester, nil)
	asserts.NotNil(tester, "not nil")

	// Negative test cases.
	asserts.Nil(tester, "not nil")
	asserts.NotNil(tester, nil)

	// Check the number of failures.
	asserts.Failures(tester, 2)
}

// TestComparisons tests the Equal, Different, Less, More, and AboutEqual assertions.
func TestComparisons(t *testing.T) {
	tester := asserts.NewTester(t, asserts.CONTINUE)

	// Positive test cases.
	asserts.Equal(tester, 42, 42)
	asserts.Equal(tester, "foo", "foo")
	asserts.Different(tester, 42, 43)
	asserts.Different(tester, "foo", "bar")
	asserts.Less(tester, 5, 3)
	asserts.Less(tester, 5.0, 3.0)
	asserts.More(tester, 3, 5)
	asserts.More(tester, 3.0, 5.0)
	asserts.AboutEqual(tester, 45, 43, 5)
	asserts.AboutEqual(tester, 4.5, 4.3, 0.1)

	// Negative test cases.
	asserts.Equal(tester, 42, 43)
	asserts.Equal(tester, "foo", "bar")
	asserts.Different(tester, 42, 42)
	asserts.Different(tester, "foo", "foo")
	asserts.Less(tester, 3, 5)
	asserts.Less(tester, 3.0, 5.0)
	asserts.More(tester, 5, 3)
	asserts.More(tester, 5.0, 3.0)
	asserts.AboutEqual(tester, 45, 43, 1)
	asserts.AboutEqual(tester, 4.5, 4.3, 0.2)

	// Check the number of failures.
	asserts.Failures(tester, 10)
}

// TestMatch tests the Match function.
func TestMatch(t *testing.T) {
	tester := asserts.NewTester(t, asserts.CONTINUE)

	// Test cases where the regular expression should match the actual string.
	asserts.Match(tester, "^hello", "hello world")
	asserts.Match(tester, "world$", "hello world")
	asserts.Match(tester, "h.llo", "hello")

	// Test cases where the regular expression should not match the actual string.
	asserts.Match(tester, "^world", "hello world")
	asserts.Match(tester, "hello$", "hello world")
	asserts.Match(tester, "h.llo", "world")

	// Test case where the regular expression compilation should fail.
	asserts.Match(tester, "[invalid", "hello world")

	// Check the number of failures.
	asserts.Failures(tester, 4)
}

// TestTimes tests the Time assertions.
func TestTimes(t *testing.T) {
	tester := asserts.NewTester(t, asserts.CONTINUE)
	now := time.Now()
	later := now.Add(2 * time.Hour)
	earlier := now.Add(-2 * time.Hour)

	// Positive test cases.
	asserts.Before(tester, later, now)
	asserts.After(tester, earlier, now)
	asserts.Between(tester, earlier, later, now)

	// Negative test cases.
	asserts.Before(tester, earlier, now)
	asserts.After(tester, later, now)
	asserts.Between(tester, earlier, now, later)
	asserts.Between(tester, later, earlier, now)

	// Check the number of failures.
	asserts.Failures(tester, 4)
}

// TestDuration tests the Duration assertion.
func TestDurations(t *testing.T) {
	tester := asserts.NewTester(t, asserts.CONTINUE)

	// Define functions to measure duration. One will return an error
	// for testing the negative case.
	fnok := func() error {
		time.Sleep(100 * time.Millisecond)
		return nil
	}
	fnerr := func() error {
		return errors.New("ouch")
	}

	// Positive test cases.
	duration := asserts.Duration(tester, fnok)
	
	asserts.Shorter(tester, duration, 200*time.Millisecond)
	asserts.Shorter(tester, 100*time.Millisecond, 200*time.Millisecond)
	asserts.Longer(tester, 200*time.Millisecond, 100*time.Millisecond)

	// Negative test cases.
	_ = asserts.Duration(tester, fnerr)
	
	asserts.Longer(tester, duration, 50*time.Millisecond)
	asserts.Shorter(tester, 200*time.Millisecond, 100*time.Millisecond)
	asserts.Longer(tester, 100*time.Millisecond, 200*time.Millisecond)

	// Check the number of failures.
	asserts.Failures(tester, 3)
}

// TestRange tests the Ranges assertions.
func TestRange(t *testing.T) {
	tester := asserts.NewTester(t, asserts.CONTINUE)

	// Positive test cases.
	asserts.InRange(tester, 42, 40, 45)
	asserts.OutOfRange(tester, 42, 45, 50)
	asserts.InRange(tester, 4.2, 4, 4.5)
	asserts.OutOfRange(tester, 4.2, 4.5, 5)
	asserts.InRange(tester, 2*time.Second, time.Second, 3*time.Second)
	asserts.OutOfRange(tester, 2*time.Second, 3*time.Second, 4*time.Second)

	// Negative test cases.
	asserts.InRange(tester, 42, 45, 50)
	asserts.OutOfRange(tester, 42, 40, 45)
	asserts.InRange(tester, 4.2, 4.5, 5)
	asserts.OutOfRange(tester, 4.2, 4, 4.5)
	asserts.InRange(tester, 2*time.Second, 3*time.Second, 4*time.Second)
	asserts.OutOfRange(tester, 2*time.Second, time.Second, 3*time.Second)

	// Check the number of failures.
	asserts.Failures(tester, 6)
}

// TestErrors tests the Error assertions.
func TestErrors(t *testing.T) {
	tester := asserts.NewTester(t, asserts.CONTINUE)
	testErr := errors.New("booom")

	// Positive test cases.
	asserts.Error(tester, testErr)
	asserts.NoError(tester, nil)
	asserts.IsError(tester, testErr, testErr)
	asserts.ErrorMatch(tester, testErr, "^bo.*")

	// Negative test cases.
	asserts.Error(tester, nil)
	asserts.NoError(tester, testErr)
	asserts.IsError(tester, testErr, errors.New("ouch"))
	asserts.ErrorMatch(tester, testErr, ".*ou$")

	// Check the number of failures.
	asserts.Failures(tester, 4)
}

// TestRun tests the Run function.
func TestRun(t *testing.T) {
	tester := asserts.NewTester(t, asserts.CONTINUE)
	subtests := []struct {
		name     string
		negative bool
		expected any
		actual   any
	}{
		{"int positive", false, 42, 42},
		{"float positive", true, 4.2, 4.2},
		{"string positive", false, "foo", "foo"},
		{"int negative", true, 42, 43},
		{"float negative", true, 4.2, 4.3},
		{"string negative", true, "foo", "bar"},
	}

	for _, test := range subtests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			if test.negative {
				asserts.Different(tester, test.expected, test.actual)
			} else {
				asserts.Equal(tester, test.expected, test.actual)
			}
		})
	}
}

// -----------------------------------------------------------------------------
// end of file
// -----------------------------------------------------------------------------
