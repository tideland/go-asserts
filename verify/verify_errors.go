// Tideland Go Asserts - Verify - Errors
//
// Copyright (C) 2024-2025 Frank Mueller / Tideland / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package verify

import (
	"errors"
	"regexp"
	"strings"
	"testing"
)

// Error checks if the given error is not nil.
func Error(t T, err error) bool {
	if err == nil {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is error", "error", nil)
		return false
	}
	return true
}

// NoError checks if the given error is nil.
// It's the opposite of Error.
func NoError(t T, gotten error) bool {
	if gotten != nil {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is no error", nil, gotten)
		return false
	}
	return true
}

// IsError checks if the gotten error is equivalent to the expected error
// using `errors.Is`. This is useful for checking for sentinel errors.
func IsError(t T, gotten, expected error) bool {
	if !errors.Is(gotten, expected) {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is expected error", expected, gotten)
		return false
	}
	return true
}

// AsError checks if the gotten error can be unwrapped to a specific error
// type using `errors.As`. The `expected` parameter must be a pointer to a
// variable of the target error type (e.g., `var perr *os.PathError`).
func AsError(t T, gotten error, expected any) bool {
	if gotten == nil {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "error as type", expected, gotten)
		return false
	}
	if !errors.As(gotten, expected) {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "error as type", expected, gotten)
		return false
	}
	return true
}

// UnwrapError checks if unwrapping `gotten` error once results in an
// error that is equivalent to `expected` using `errors.Is`.
func UnwrapError(t T, gotten, expected error) bool {
	if gotten == nil {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "error unwraps to", expected, gotten)
		return false
	}
	unwrapped := errors.Unwrap(gotten)
	if !errors.Is(unwrapped, expected) {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "error unwraps to", expected, unwrapped)
		return false
	}
	return true
}

// ErrorContains check if the given error is not nil and its message
// contains an expected string.
func ErrorContains(t T, gotten error, expected string) bool {
	if gotten == nil {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "error contains", expected, gotten)
		return false
	}
	if !strings.Contains(gotten.Error(), expected) {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "error contains", expected, gotten.Error())
		return false
	}
	return true
}

// ErrorMatch checks if the gotten error is not nil and its message
// matches the expected regular expression.
func ErrorMatch(t T, gotten error, expected string) bool {
	if gotten == nil {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "error does match", expected, gotten)
		return false
	}
	re := regexp.MustCompile(expected)
	if !re.MatchString(gotten.Error()) {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "error does match", expected, gotten.Error())
		return false
	}
	return true
}
