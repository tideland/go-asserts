// Tideland Go Asserts - Verify - Types
//
// Copyright (C) 2024-2025 Frank Mueller / Tideland / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package verify

import (
	"reflect"
	"testing"
)

// Implements checks if the gotten instance implements the expected interface.
// The expected parameter has to be an interface type as nil pointer. Here e.g.
// var stringer fmt.Stringer and then verify.Implements(t, myVar, &stringer).
func Implements(t T, gotten, expected any) bool {
	if expected == nil {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "does implement", "expected instance", nil)
		return false
	}

	if gotten == nil {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "does implement", "actual instance", nil)
		return false
	}

	expectedType := reflect.TypeOf(expected).Elem()
	if expectedType.Kind() != reflect.Interface {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "does implement", "expected interface", nil)
		return false
	}

	actualType := reflect.TypeOf(gotten)
	if !actualType.Implements(expectedType) {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "does implement", expectedType, actualType)
		return false
	}
	return true
}

// Assignability checks if the actual value can be assigned to the type of the
// expected type.
func Assignability(t T, gotten, expected any) bool {
	if expected == nil {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is assignable to", "expected type", nil)
		return false
	}

	if gotten == nil {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is assignable to", "actual type", nil)
		return false
	}

	expectedType := reflect.TypeOf(expected)
	actualType := reflect.TypeOf(gotten)

	if !actualType.AssignableTo(expectedType) {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is assignable to", expectedType, actualType)
		return false
	}
	return true
}

// Panics checks if the given functions panics.
func Panics(t T, gotten func()) bool {
	if gotten == nil {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "panics", "expected function", nil)
		return false
	}

	if ht, ok := t.(testing.TB); ok {
		ht.Helper()
	}

	panicked := false
	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()

	gotten()

	if !panicked {
		verificationFailure(t, "panics", "function to panic", "function did not panic")
		return false
	}
	return true
}

// NotPanics checks if the given functions does not panic.
func NotPanics(t T, gotten func()) bool {
	if gotten == nil {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "not panics", "expected function", nil)
		return false
	}

	if ht, ok := t.(testing.TB); ok {
		ht.Helper()
	}

	panicked := false
	defer func() {
		if r := recover(); r != nil {
			panicked = true
			verificationFailure(t, "not panics", "function not to panic", "function panicked")
		}
	}()

	gotten()

	return !panicked
}
