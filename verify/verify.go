// -----------------------------------------------------------------------------
// Convenient verification of unit tests in Go libraries and applications.
//
// A set of individual verifications
//
// Copyright (C) 2024-2025 Frank Mueller / Oldenburg / Germany / Earth
// -----------------------------------------------------------------------------

package verify // import "tideland.dev/go/assert/verify"

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"slices"
	"strings"
	"testing"
	"time"

	"golang.org/x/exp/constraints"
)

// -----------------------------------------------------------------------------
// Verifications
// -----------------------------------------------------------------------------

// True checks if the given value is true.
func True(t T, gotten bool, infos ...string) bool {
	if !gotten {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is true", true, gotten, infos...)
		return false
	}
	return true
}

// False checks if the given value is false. It's the opposite of True.
func False(t T, gotten bool, infos ...string) bool {
	if gotten {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is false", false, gotten)
		return false
	}
	return true
}

// Nil checks if the given value is nil.
func Nil(t T, gotten any, infos ...string) bool {
	if gotten != nil {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is nil", nil, gotten, infos...)
		return false
	}
	return true
}

// NotNil checks if the given value is not nil. It's the opposite of Nil.
func NotNil(t T, gotten any, infos ...string) bool {
	if gotten == nil {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is not nil", nil, gotten, infos...)
		return false
	}
	return true
}

// Equal checks if the gotten and expected values are equal.
// It uses the == operator for comparable types and supports time.Duration.
func Equal[C comparable](t T, gotten, expected C, infos ...string) bool {
	if expected != gotten {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is equal", expected, gotten, infos...)
		return false
	}
	return true
}

// Different checks if the given values are different.
// It uses the != operator for comparable types and supports time.Duration.
func Different[C comparable](t T, gotten, expected C, infos ...string) bool {
	if expected == gotten {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is different", expected, gotten, infos...)
		return false
	}
	return true
}

// Length checks if the given value has the expected length. This only
// works for the according types for len(). All others fail.
func Length(t T, gotten any, expected int, infos ...string) bool {
	if expected < 0 {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "has length", expected, "not quantifiable", infos...)
		return false
	}
	gottenLen := flexlen(gotten)
	if gottenLen < 0 {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "has length", expected, "gotten not quantifiable", infos...)
		return false
	}
	if gottenLen != expected {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "has length", expected, gottenLen, infos...)
		return false
	}
	return true
}

// Empty checks if the given value is empty. This only works for the according types for len().
// All others fail.
func Empty(t T, gotten any, infos ...string) bool {
	gottenLen := flexlen(gotten)
	if gottenLen < 0 {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "empty", 0, "gotten not quantifiable", infos...)
		return false
	}
	if gottenLen != 0 {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "empty", 0, gottenLen, infos...)
		return false
	}
	return true
}

// NotEmpty checks if the given value is not empty. This only works for the according types for len().
// All others fail.
func NotEmpty(t T, gotten any, infos ...string) bool {
	gottenLen := flexlen(gotten)
	if gottenLen < 0 {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "not empty", 0, "gotten not quantifiable", infos...)
		return false
	}
	if gottenLen == 0 {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "not empty", "> 0", gottenLen, infos...)
		return false
	}
	return true
}

// Less checks if the gotten value is less than the expected one.
// Supports integers, floats, and time.Duration.
func Less[C constraints.Integer | constraints.Float](t T, gotten, expected C, infos ...string) bool {
	if gotten >= expected {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is less", expected, gotten, infos...)
		return false
	}
	return true
}

// More checks if the gotten value is more than the expected one.
// Supports integers, floats, and time.Duration.
func More[C constraints.Integer | constraints.Float](t T, gotten, expected C, infos ...string) bool {
	if gotten <= expected {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is more", expected, gotten, infos...)
		return false
	}
	return true
}

// About checks if the gotten values equal within a expected delta. Possible
// values are integers, floats, and time.Duration.
func About[C constraints.Integer | constraints.Float](t T, gotten, expected, tolerance C, infos ...string) bool {
	if gotten < expected-tolerance || gotten > expected+tolerance {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		expectedDescr := fmt.Sprintf("%v' +/- '%v'", expected, tolerance)
		verificationFailure(t, "is about equal", expectedDescr, gotten, infos...)
		return false
	}
	return true
}

// Substring checks if the gotten string is a substring of the expected string.
func Substring(t T, gotten, expected string, infos ...string) bool {
	if !strings.Contains(expected, gotten) {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "substring", expected, gotten, infos...)
		return false
	}
	return true
}

// Contains checks if the slice contains the expected element.
func Contains[S ~[]E, E comparable](t T, gotten E, expected S, infos ...string) bool {
	if !slices.Contains(expected, gotten) {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "contains", expected, gotten, infos...)
		return false
	}
	return true
}

// Match checks if the gotten string matches the expected regular expression.
func Match(t T, gotten, expected string, infos ...string) bool {
	re, err := regexp.Compile(expected)
	if err != nil {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "matches", expected, err.Error(), infos...)
		return false
	}
	if !re.MatchString(gotten) {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "matches", expected, gotten, infos...)
		return false
	}
	return true
}

// Simultaneous checks if the gotten time is simultaneous with the expected time.
func Simultaneous(t T, gotten, expected time.Time, infos ...string) bool {
	if !gotten.Equal(expected) {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is time simultaneous", ftim(expected), ftim(gotten), infos...)
		return false
	}
	return true
}

// Before checks if the gotten time is before the expected time.
func Before(t T, gotten, expected time.Time, infos ...string) bool {
	if !gotten.Before(expected) {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is time before", ftim(expected), ftim(gotten), infos...)
		return false
	}
	return true
}

// After checks if the gotten time is after the expected time.
func After(t T, gotten, expected time.Time, infos ...string) bool {
	if !gotten.After(expected) {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is time after", ftim(expected), ftim(gotten), infos...)
		return false
	}
	return true
}

// Between checks if the gotten time is between the expected start and end times.
func Between(t T, gotten, expectedBegin, expectedEnd time.Time, infos ...string) bool {
	expstr := ""
	if expectedBegin.After(expectedEnd) {
		expectedBegin, expectedEnd = expectedEnd, expectedBegin
	}
	if gotten.Before(expectedBegin) || gotten.After(expectedEnd) {
		expstr = fmt.Sprintf("'%s' and '%s'", ftim(expectedBegin), ftim(expectedEnd))
	}
	if expstr != "" {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is between", expstr, ftim(gotten), infos...)
		return false
	}
	return true
}

// Shorter checks if the gotten duration is shorter than the expected duration.
func Shorter(t T, gotten, expected time.Duration, infos ...string) bool {
	if gotten > expected {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "duration is shorter", expected, gotten, infos...)
		return false
	}
	return true
}

// Longer checks if the gotten duration is longer than the expected duration.
func Longer(t T, gotten, expected time.Duration, infos ...string) bool {
	if gotten < expected {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "duration is longer", expected, gotten, infos...)
		return false
	}
	return true
}

// InRange checks if the given value is within lower and upper bounds. Possible
// values are integers, floats, and time.Duration.
func InRange[C constraints.Integer | constraints.Float](t T, gotten, expectedLower, expectedUpper C, infos ...string) bool {
	if expectedLower > expectedUpper {
		expectedLower, expectedUpper = expectedUpper, expectedLower
	}
	if gotten < expectedLower || gotten > expectedUpper {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		expectedDescr := fmt.Sprintf("'%v' to '%v'", expectedLower, expectedUpper)
		verificationFailure(t, "is in range", expectedDescr, gotten, infos...)
		return false
	}
	return true
}

// OutOfRange checks if the given value is outside lower and upper bounds. It's the
// opposite of InRange.
func OutOfRange[C constraints.Integer | constraints.Float](t T, gotten, expectedLower, expectedUpper C, infos ...string) bool {
	if expectedLower > expectedUpper {
		expectedLower, expectedUpper = expectedUpper, expectedLower
	}
	if gotten >= expectedLower && gotten <= expectedUpper {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		expectedDescr := fmt.Sprintf("'%v' to '%v'", expectedLower, expectedUpper)
		verificationFailure(t, "is out of range", expectedDescr, gotten, infos...)
		return false
	}
	return true
}

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

// IsError checks if the given error is not nil and of the expected type.
// It uses the errors.Is() function.
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

// AsError checks if the given error can be unwrapped to the expected error type.
// It uses the errors.As() function. The expected parameter should be a pointer
// to the error type you want to check for.
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

// UnwrapError checks if the given error unwraps to the expected error.
// It uses the errors.Unwrap() function.
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

// Implements checks if the gotten instance implements the expected interface.
// The expected parameter has to be an interface type as nil pointer. Hier e.g.
// var stringer fmt.Stinger and then verify.Implements(t, myVar, &fmtStringer).
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

	defer func() {
		if r := recover(); r == nil {
			if ht, ok := t.(testing.TB); ok {
				ht.Helper()
			}
			verificationFailure(t, "panics", "expected function", "actual function")
		}
	}()

	gotten()
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

	defer func() {
		if r := recover(); r != nil {
			if ht, ok := t.(testing.TB); ok {
				ht.Helper()
			}
			verificationFailure(t, "not panics", "expected function", "actual function")
		}
	}()

	gotten()
	return true
}

// -----------------------------------------------------------------------------
// Helper
// -----------------------------------------------------------------------------

// ftim is a short to format times in test output.
func ftim(t time.Time) string {
	return t.Format(time.RFC3339)
}

type lenner interface {
	Len() int
}

type lengthier interface {
	Length() int
}

// flexlen retruns the length of types avaialbe to return their length.
func flexlen(in any) int {
	// Check for possible existing methods
	switch in.(type) {
	case lenner:
		return in.(lenner).Len()
	case lengthier:
		return in.(lengthier).Length()
	}
	// Use reflection
	rv := reflect.ValueOf(in)
	switch rv.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		return rv.Len()
	}
	// Good old -1 is enough here, verification is above
	return -1
}

// -----------------------------------------------------------------------------
// EOF
// -----------------------------------------------------------------------------
