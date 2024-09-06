// -----------------------------------------------------------------------------
// asserts for more convinient testing - tests
//
// Copyright (C) 2024 Frank Mueller / Oldenburg / Germany / World
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

// TestBooleans tests the True and False assertions.
func TestBooleans(t *testing.T) {
	tester := asserts.NewTester(t, asserts.CONTINUE)

	asserts.True(tester, true)
	asserts.True(tester, false)
	asserts.False(tester, false)
	asserts.False(tester, true)

	asserts.Failures(tester, 2)
}

// TestNils tests the Nil and NotNil assertions.
func TestNils(t *testing.T) {
	tester := asserts.NewTester(t, asserts.CONTINUE)

	asserts.Nil(tester, nil)
	asserts.Nil(tester, "not nil")
	asserts.NotNil(tester, "not nil")
	asserts.NotNil(tester, nil)

	asserts.Failures(tester, 2)
}

// TestComparisons tests the Equal, Different, Less, More, and AboutEqual assertions.
func TestComparisons(t *testing.T) {
	tester := asserts.NewTester(t, asserts.CONTINUE)

	asserts.Equal(tester, 42, 42)
	asserts.Equal(tester, 42, 43)
	asserts.Equal(tester, "foo", "foo")
	asserts.Equal(tester, "foo", "bar")

	asserts.Different(tester, 42, 43)
	asserts.Different(tester, 42, 42)
	asserts.Different(tester, "foo", "bar")
	asserts.Different(tester, "foo", "foo")

	asserts.Less(tester, 5, 3)
	asserts.Less(tester, 3, 5)
	asserts.Less(tester, 5.0, 3.0)
	asserts.Less(tester, 3.0, 5.0)

	asserts.More(tester, 3, 5)
	asserts.More(tester, 5, 3)
	asserts.More(tester, 3.0, 5.0)
	asserts.More(tester, 5.0, 3.0)

	asserts.AboutEqual(tester, 45, 43, 5)
	asserts.AboutEqual(tester, 45, 43, 1)
	asserts.AboutEqual(tester, 4.5, 4.3, 0.1)
	asserts.AboutEqual(tester, 4.5, 4.3, 0.2)

	asserts.Failures(tester, 14)
}

// TestTimes tests the Time assertions.
func TestTimes(t *testing.T) {
	tester := asserts.NewTester(t, asserts.CONTINUE)
	now := time.Now()
	later := now.Add(2 * time.Hour)
	earlier := now.Add(-2 * time.Hour)

	asserts.Before(tester, later, now)
	asserts.Before(tester, earlier, now)
	asserts.After(tester, earlier, now)
	asserts.After(tester, later, now)
	asserts.Between(tester, earlier, later, now)
	asserts.Between(tester, earlier, now, later)
	asserts.Between(tester, later, earlier, now)

	asserts.Failures(tester, 4)
}

// TestDuration tests the Duration assertion.
func TestDurations(t *testing.T) {
	tester := asserts.NewTester(t, asserts.CONTINUE)

	// Define a function to measure duration.
	fnerr := func() error {
		return errors.New("ouch")
	}
	fnok := func() error {
		time.Sleep(100 * time.Millisecond)
		return nil
	}

	// Measure the duration of the function.
	duration := asserts.Duration(tester, fnerr)
	duration = asserts.Duration(tester, fnok)

	asserts.Shorter(tester, duration, 200*time.Millisecond)
	asserts.Longer(tester, duration, 50*time.Millisecond)

	// Test Shorter and Longer directly.
	asserts.Shorter(tester, 100*time.Millisecond, 200*time.Millisecond)
	asserts.Shorter(tester, 200*time.Millisecond, 100*time.Millisecond)
	asserts.Longer(tester, 200*time.Millisecond, 100*time.Millisecond)
	asserts.Longer(tester, 100*time.Millisecond, 200*time.Millisecond)

	asserts.Failures(tester, 3)
}

// TestRange tests the Ranges assertions.
func TestRange(t *testing.T) {
	tester := asserts.NewTester(t, asserts.CONTINUE)

	asserts.InRange(tester, 42, 40, 45)
	asserts.InRange(tester, 42, 45, 50)
	asserts.OutOfRange(tester, 42, 45, 50)
	asserts.OutOfRange(tester, 42, 40, 45)

	asserts.InRange(tester, 4.2, 4, 4.5)
	asserts.InRange(tester, 4.2, 4.5, 5)
	asserts.OutOfRange(tester, 4.2, 4.5, 5)
	asserts.OutOfRange(tester, 4.2, 4, 4.5)

	asserts.InRange(tester, 2*time.Second, time.Second, 3*time.Second)
	asserts.InRange(tester, 2*time.Second, 3*time.Second, 4*time.Second)
	asserts.OutOfRange(tester, 2*time.Second, 3*time.Second, 4*time.Second)
	asserts.OutOfRange(tester, 2*time.Second, time.Second, 3*time.Second)

	asserts.Failures(tester, 6)
}

// TestErrors tests the Error assertions.
func TestErrors(t *testing.T) {
	tester := asserts.NewTester(t, asserts.CONTINUE)
	testErr := errors.New("booom")

	asserts.Error(tester, testErr)
	asserts.Error(tester, nil)
	asserts.NoError(tester, nil)
	asserts.NoError(tester, testErr)
	asserts.IsError(tester, testErr, testErr)
	asserts.IsError(tester, testErr, errors.New("ouch"))
	asserts.ErrorContains(tester, testErr, "ooo")
	asserts.ErrorContains(tester, testErr, "BOOOM")
	asserts.ErrorMatches(tester, testErr, "^bo.*")
	asserts.ErrorMatches(tester, testErr, ".*ou$")

	asserts.Failures(tester, 5)
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
