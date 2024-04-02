// -----------------------------------------------------------------------------
// asserts for more convinient testing
//
// Copyright (C) 2024 Frank Mueller / Oldenburg / Germany / World
// -----------------------------------------------------------------------------

package asserts // import "tideland.dev/go/assert"

// Logf prints a log message with the given information on stdout.
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

// False checks if the given condition is false.
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

// NotNil checks if the given value is not nil.
func NotNil(t Tester, value any) {
	if value == nil {
		failf(t, "not nil", "value is nil")
	}
}

// -----------------------------------------------------------------------------
// end of file
// -----------------------------------------------------------------------------
