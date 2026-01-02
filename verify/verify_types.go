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

// Assignability checks if a value of `gotten`'s type is assignable to a
// variable of `expected`'s type.
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

// Panics checks if calling the function `gotten` causes a panic.
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

// NotPanics checks that calling the function `gotten` does not cause a panic.
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

// DeepEqual checks if the gotten and expected values are deeply equal using reflection.
// This is useful for comparing complex structures, slices, maps, etc.
func DeepEqual(t T, gotten, expected any, infos ...string) bool {
	if ht, ok := t.(testing.TB); ok {
		ht.Helper()
	}

	if !reflect.DeepEqual(gotten, expected) {
		verificationFailure(t, "is deeply equal", expected, gotten, infos...)
		return false
	}
	return true
}

// SameType checks if the gotten and expected values have the same type.
func SameType(t T, gotten, expected any, infos ...string) bool {
	if ht, ok := t.(testing.TB); ok {
		ht.Helper()
	}

	gottenType := reflect.TypeOf(gotten)
	expectedType := reflect.TypeOf(expected)

	if gottenType != expectedType {
		verificationFailure(t, "is same type", expectedType, gottenType, infos...)
		return false
	}
	return true
}

// SamePointer checks if the gotten and expected values point to the same memory address.
// Both values must be pointers, slices, maps, channels, or functions.
func SamePointer(t T, gotten, expected any, infos ...string) bool {
	if ht, ok := t.(testing.TB); ok {
		ht.Helper()
	}

	gottenVal := reflect.ValueOf(gotten)
	expectedVal := reflect.ValueOf(expected)

	// Check if both are valid pointer-like types
	if !isPointerLike(gottenVal) {
		verificationFailure(t, "same pointer", "pointer-like type", "non-pointer type", infos...)
		return false
	}
	if !isPointerLike(expectedVal) {
		verificationFailure(t, "same pointer", "pointer-like type", "non-pointer type", infos...)
		return false
	}

	if gottenVal.Pointer() != expectedVal.Pointer() {
		verificationFailure(t, "same pointer", expectedVal.Pointer(), gottenVal.Pointer(), infos...)
		return false
	}
	return true
}

// NotSamePointer checks if the gotten and expected values point to different memory addresses.
// Both values must be pointers, slices, maps, channels, or functions.
func NotSamePointer(t T, gotten, expected any, infos ...string) bool {
	if ht, ok := t.(testing.TB); ok {
		ht.Helper()
	}

	gottenVal := reflect.ValueOf(gotten)
	expectedVal := reflect.ValueOf(expected)

	// Check if both are valid pointer-like types
	if !isPointerLike(gottenVal) {
		verificationFailure(t, "different pointer", "pointer-like type", "non-pointer type", infos...)
		return false
	}
	if !isPointerLike(expectedVal) {
		verificationFailure(t, "different pointer", "pointer-like type", "non-pointer type", infos...)
		return false
	}

	if gottenVal.Pointer() == expectedVal.Pointer() {
		verificationFailure(t, "different pointer", "different addresses", "same address", infos...)
		return false
	}
	return true
}

// isPointerLike checks if a value is a pointer-like type (pointer, slice, map, channel, or function).
func isPointerLike(v reflect.Value) bool {
	if !v.IsValid() {
		return false
	}
	kind := v.Kind()
	return kind == reflect.Ptr || kind == reflect.Slice || kind == reflect.Map ||
		kind == reflect.Chan || kind == reflect.Func
}
