// -----------------------------------------------------------------------------
// asserts for more convinient testing
//
// Copyright (C) 2024 Frank Mueller / Oldenburg / Germany / World
// -----------------------------------------------------------------------------

package asserts // import "tideland.dev/go/assert"

import "errors" // Logf prints a log message with the given information on stdout.
// It's used to print additional information during testing.
// The location and function name are added automatically.
func Logf(format string, args ...interface{}) {
	logf(format, args...)
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
		failf(t, "no error", "error is %v", err)
	}
}

// IsError checks if the given error is not nil and of the expected type.
// It uses the errors.Is() function.
func IsError(t Tester, expected, actual error) {
	if !errors.Is(expected, actual) {
		failf(t, "is error", "expected is '%v', actual is '%v'", expected, actual)
	}
}

// -----------------------------------------------------------------------------
// end of file
// -----------------------------------------------------------------------------
