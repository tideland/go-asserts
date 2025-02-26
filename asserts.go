// -----------------------------------------------------------------------------
// asserts for more convinient testing
//
// Copyright (C) 2024-2025 Frank Mueller / Oldenburg / Germany / Earth
// -----------------------------------------------------------------------------

package asserts // import "tideland.dev/go/assert"

import (
	"errors"
	"regexp"
	"time"

	"golang.org/x/exp/constraints"
)

// It's used to print additional information during testing.
// The location and function name are added automatically.
func Logf(t Tester, format string, args ...interface{}) {
	logf(format, args...)
}

// Failf is used to fail a test with a formatted message.
func Failf(t Tester, format string, args ...interface{}) {
	failf(t, "fail", format, args...)
}

// True checks if the given condition is true.
func True(t Tester, condition bool) {
	if !condition {
		failf(t, "true", "expected condition 'true' is 'false'")
	}
}

// False checks if the given condition is false. It's the opposite of True.
func False(t Tester, condition bool) {
	if condition {
		failf(t, "false", "expected condition 'false' is 'true'")
	}
}

// Nil checks if the given value is nil.
func Nil(t Tester, value any) {
	if value != nil {
		failf(t, "nil", "value is not nil")
	}
}

// NotNil checks if the given value is not nil. It's the opposite of Nil.
func NotNil(t Tester, value any) {
	if value == nil {
		failf(t, "not nil", "value is nil")
	}
}

// Equal checks if the given values are equal.
// It uses the == operator for comparable types and supports time.Duration.
func Equal[T comparable](t Tester, expected, actual T) {
	if expected != actual {
		failf(t, "equal", "expected is '%v', actual is '%v'", expected, actual)
	}
}

// Different checks if the given values are different.
// It uses the != operator for comparable types and supports time.Duration.
func Different[T comparable](t Tester, expected, actual T) {
	if expected == actual {
		failf(t, "different", "expected is '%v', actual is '%v'", expected, actual)
	}
}

// Match checks if the actual string matches the expected regular expression.
// The regular expression is compiled from the expected string. If the compilation
// fails, the assertion fails.
func Match(t Tester, expected, actual string) {
	re, err := regexp.Compile(expected)
	if err != nil {
		failf(t, "match", "failed to compile regexp from '%s': %v", expected, err)
		return
	}
	if !re.MatchString(actual) {
		failf(t, "match", "actual '%s' doesn't match '%s'", actual, expected)
	}
}

// Less checks if the actual value is less than the expected one.
// Supports integers, floats, and time.Duration.
func Less[T constraints.Integer | constraints.Float](t Tester, expected, actual T) {
	if actual >= expected {
		failf(t, "less", "actual '%v' is more than '%v'", actual, expected)
	}
}

// More checks if the actual value is more than the expected one.
// Supports integers, floats, and time.Duration.
func More[T constraints.Integer | constraints.Float](t Tester, expected, actual T) {
	if actual <= expected {
		failf(t, "more", "actual '%v' is less than expected '%v'", actual, expected)
	}
}

// AboutEqual checks if the given values are equal within a delta. Possible
// values are integers, floats, and time.Duration.
func AboutEqual[T constraints.Integer | constraints.Float](t Tester, expected, actual, delta T) {
	if expected < actual-delta || expected > actual+delta {
		failf(t, "about equal", "expected is '%v' +/- '%v', actual is '%v'", expected, delta, actual)
	}
}

// Before checks if the actual time is before the expected time.
func Before(t Tester, expected, actual time.Time) {
	if !actual.Before(expected) {
		failf(t, "time before", "actual time '%v' is not before expected time '%v'", actual, expected)
	}
}

// After checks if the actual time is after the expected time.
func After(t Tester, expected, actual time.Time) {
	if !actual.After(expected) {
		failf(t, "time after", "actual time '%v' is not after expected time '%v'", actual, expected)
	}
}

// Between checks if the actual time is between the expected start and end times.
func Between(t Tester, start, end, actual time.Time) {
	if actual.Before(start) || actual.After(end) {
		failf(t, "time between", "actual time '%v' is not between start time '%v' and end time '%v'", actual, start, end)
	}
}

// Duration calculates the duration of a function execution. If the function returns an error,
// the test fails. The returned duration can be used as expected duration in further tests.
func Duration(t Tester, fn func() error) time.Duration {
	start := time.Now()
	err := fn()
	if err != nil {
		failf(t, "measure duration", "function returned an error: '%v'", err)
	}
	return time.Since(start)
}

// Shorter checks if the actual duration is shorter than the expected duration.
func Shorter(t Tester, actual, expected time.Duration) {
	if actual > expected {
		failf(t, "earlier", "actual duration '%v' is not shorter than expected duration '%v'", expected, actual)
	}
}

// Longer checks if the actual duration is longer than the expected duration.
func Longer(t Tester, actual, expected time.Duration) {
	if actual < expected {
		failf(t, "later", "actual duration '%v' is not longer than expected duration '%v'", expected, actual)
	}
}

// DurationAboutEqual checks if the given durations are equal within a delta.
func DurationAboutEqual(t Tester, expected, actual, delta time.Duration) {
	if expected < actual-delta || expected > actual+delta {
		failf(t, "duration about equal", "expected duration is '%v' +/- '%v', actual duration is '%v'", expected, delta, actual)
	}
}

// InRange checks if the given value is within lower and upper bounds. Possible
// values are integers, floats, and time.Duration.
func InRange[T constraints.Integer | constraints.Float](t Tester, expected, lower, upper T) {
	if lower > upper {
		lower, upper = upper, lower
	}
	if expected <= lower || expected >= upper {
		failf(t, "range", "value is '%v', not in range '%v' to '%v'", expected, lower, upper)
	}
}

// OutOfRange checks if the given value is outside lower and upper bounds. It's the
// opposite of InRange.
func OutOfRange[T constraints.Integer | constraints.Float](t Tester, expected, lower, upper T) {
	if lower > upper {
		lower, upper = upper, lower
	}
	if expected >= lower && expected <= upper {
		failf(t, "out of range", "value is '%v', not out of range '%v' to '%v'", expected, lower, upper)
	}
}

// Error checks if the given error is not nil.
func Error(t Tester, err error) {
	if err == nil {
		failf(t, "error", "error is nil")
	}
}

// NoError checks if the given error is nil.
// It's the opposite of Error.
func NoError(t Tester, err error) {
	if err != nil {
		failf(t, "no error", "expected no error, got '%v'", err)
	}
}

// IsError checks if the given error is not nil and of the expected type.
// It uses the errors.Is() function.
func IsError(t Tester, expected, actual error) {
	if !errors.Is(expected, actual) {
		failf(t, "is error", "expected is '%v', actual is '%v'", expected, actual)
	}
}

// ErrorMatch checks if the given error is not nil and its message
// matches the expected regular expression.
func ErrorMatch(t Tester, err error, expected string) {
	if err == nil {
		failf(t, "error match", "error is nil")
		return
	}
	re := regexp.MustCompile(expected)
	if !re.MatchString(err.Error()) {
		failf(t, "error match", "error '%v' does not match '%s'", err.Error(), expected)
	}
}

// -----------------------------------------------------------------------------
// end of file
// -----------------------------------------------------------------------------
