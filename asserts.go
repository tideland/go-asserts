// -----------------------------------------------------------------------------
// asserts for more convinient testing
//
// Copyright (C) 2024 Frank Mueller / Oldenburg / Germany / World
// -----------------------------------------------------------------------------

package asserts // import "tideland.dev/go/assert"

import (
	"errors"
	"regexp"
	"strings"

	"golang.org/x/exp/constraints"
)

// It's used to print additional information during testing.
// The location and function name are added automatically.
func Logf(t Tester, format string, args ...interface{}) {
	logf(t, format, args...)
}

// Failf is used to fail a test with a formatted message.
func Failf(t Tester, format string, args ...interface{}) {
	failf(t, "fail", format, args...)
}

// True checks if the given condition is true.
func True(t Tester, condition bool) {
	if !condition {
		failf(t, "true", "condition is false")
	}
}

// False checks if the given condition is false. It's the opposite of True.
func False(t Tester, condition bool) {
	if condition {
		failf(t, "false", "condition is true")
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
// It uses the == operator for comparable types.
func Equal[T comparable](t Tester, expected, actual T) {
	if expected != actual {
		failf(t, "equal", "expected is '%v', actual is '%v'", expected, actual)
	}
}

// Different checks if the given values are different.
// It uses the != operator for comparable types and is the opposite of Equal.
func Different[T comparable](t Tester, expected, actual T) {
	if expected == actual {
		failf(t, "different", "expected is '%v', actual is '%v'", expected, actual)
	}
}

// AboutEqual checks if the given values are equal within a delta. Possible
// values are integers, floats, and time.Duration.
func AboutEqual[T constraints.Integer | constraints.Float](t Tester, expected, actual, delta T) {
	if expected < actual-delta || expected > actual+delta {
		failf(t, "about equal", "expected is '%v' +/- '%v', actual is '%v'", expected, delta, actual)
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

// ErrorContains checks if the given error is not nil and its message
// contains the expected substring.
func ErrorContains(t Tester, err error, contains string) {
	if err == nil {
		failf(t, "error contains", "error is nil")
	}
	if !strings.Contains(err.Error(), contains) {
		failf(t, "error contains", "error does not contain '%s'", contains)
	}
}

// ErrorMatches checks if the given error is not nil and its message
// matches the expected regular expression.
func ErrorMatches(t Tester, err error, pattern string) {
	if err == nil {
		failf(t, "error matches", "error is nil")
	}
	if !regexp.MustCompile(pattern).MatchString(err.Error()) {
		failf(t, "error matches", "error does not match '%s'", pattern)
	}
}

// -----------------------------------------------------------------------------
// end of file
// -----------------------------------------------------------------------------
