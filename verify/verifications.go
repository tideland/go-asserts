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
	"strings"
	"testing"
	"time"

	"golang.org/x/exp/constraints"
)

// True checks if the given value is true.
func True(t T, actual bool) bool {
	if !actual {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is true", true, actual)
		return false
	}
	return true
}

// False checks if the given value is false. It's the opposite of True.
func False(t T, actual bool) bool {
	if actual {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is false", false, actual)
		return false
	}
	return true
}

// Nil checks if the given value is nil.
func Nil(t T, actual any) bool {
	if actual != nil {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is nil", nil, actual)
		return false
	}
	return true
}

// NotNil checks if the given value is not nil. It's the opposite of Nil.
func NotNil(t T, actual any) bool {
	if actual == nil {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is not nil", nil, actual)
		return false
	}
	return true
}

// Equal checks if the given values are equal.
// It uses the == operator for comparable types and supports time.Duration.
func Equal[C comparable](t T, expected, actual C) bool {
	if expected != actual {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is equal", expected, actual)
		return false
	}
	return true
}

// Different checks if the given values are different.
// It uses the != operator for comparable types and supports time.Duration.
func Different[C comparable](t T, expected, actual C) bool {
	if expected == actual {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is different", expected, actual)
		return false
	}
	return true
}

// Length checks if the given value has the expected length. This only
// works for the according types for len(). All others fail.
func Length(t T, actual any, length int) bool {
	rv := reflect.ValueOf(actual)
	switch rv.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		actualLength := rv.Len()
		if actualLength != length {
			if ht, ok := t.(testing.TB); ok {
				ht.Helper()
			}
			verificationFailure(t, "has length", length, actualLength)
			return false
		}
	default:
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "has length", length, "not quantifiable")
		return false
	}
	return true
}

// Less checks if the actual value is less than the expected one.
// Supports integers, floats, and time.Duration.
func Less[C constraints.Integer | constraints.Float](t T, expected, actual C) bool {
	if actual >= expected {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is less", expected, actual)
		return false
	}
	return true
}

// More checks if the actual value is more than the expected one.
// Supports integers, floats, and time.Duration.
func More[C constraints.Integer | constraints.Float](t T, expected, actual C) bool {
	if actual <= expected {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is more", expected, actual)
		return false
	}
	return true
}

// AboutEqual checks if the given values are equal within a delta. Possible
// values are integers, floats, and time.Duration.
func AboutEqual[C constraints.Integer | constraints.Float](t T, expected, actual, delta C) bool {
	if expected < actual-delta || expected > actual+delta {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		expectedDescr := fmt.Sprintf("%v' +/- '%v'", expected, delta)
		verificationFailure(t, "is about equal", expectedDescr, actual)
		return false
	}
	return true
}

// Contains check if the actual string contains the expected string.
func Contains(t T, expected, actual string) bool {
	if !strings.Contains(actual, expected) {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "contains", expected, actual)
		return false
	}
	return true
}

// ContainsAny checks if the actual string contains any of the expected strings.
func ContainsAny(t T, expected []string, actual string) bool {
	for _, exp := range expected {
		if strings.Contains(actual, exp) {
			return true
		}
	}
	if ht, ok := t.(testing.TB); ok {
		ht.Helper()
	}
	expectedList := "[" + strings.Join(expected, ", ") + "]"
	verificationFailure(t, "contains any", expectedList, actual)
	return false
}

// Match checks if the actual string matches the given regular expression.
func Match(t T, expected, actual string) bool {
	re, err := regexp.Compile(expected)
	if err != nil {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "matches", expected, err.Error())
		return false
	}
	if !re.MatchString(actual) {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "matches", expected, actual)
		return false
	}
	return true
}

// Before checks if the actual time is before the expected time.
func Before(t T, expected, actual time.Time) bool {
	if !actual.Before(expected) {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is time before", ftim(expected), ftim(actual))
		return false
	}
	return true
}

// After checks if the actual time is after the expected time.
func After(t T, expected, actual time.Time) bool {
	if !actual.After(expected) {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is time after", ftim(expected), ftim(actual))
		return false
	}
	return true
}

// Between checks if the actual time is between the expected start and end times.
func Between(t T, start, end, actual time.Time) bool {
	expstr := ""
	if start.After(end) {
		start, end = end, start
	}
	if actual.Before(start) || actual.After(end) {
		expstr = fmt.Sprintf("'%s' and '%s'", ftim(start), ftim(end))
	}
	if expstr != "" {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is between", expstr, ftim(actual))
		return false
	}
	return true
}

// Shorter checks if the actual duration is shorter than the expected duration.
func Shorter(t T, expected, actual time.Duration) bool {
	if actual > expected {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "duration is shorter", expected, actual)
		return false
	}
	return true
}

// Longer checks if the actual duration is longer than the expected duration.
func Longer(t T, expected, actual time.Duration) bool {
	if actual < expected {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "duration is longer", expected, actual)
		return false
	}
	return true
}

// DurationAboutEqual checks if the given durations are equal within a delta.
func DurationAboutEqual(t T, expected, actual, delta time.Duration) bool {
	if expected < actual-delta || expected > actual+delta {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		expectedDesc := fmt.Sprintf("'%v +/- '%s'", expected, delta)
		verificationFailure(t, "duration is about equal", expectedDesc, actual)
		return false
	}
	return true
}

// InRange checks if the given value is within lower and upper bounds. Possible
// values are integers, floats, and time.Duration.
func InRange[C constraints.Integer | constraints.Float](t T, lower, upper, actual C) bool {
	if lower > upper {
		lower, upper = upper, lower
	}
	if actual <= lower || actual >= upper {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		expectedDescr := fmt.Sprintf("'%v' to '%v'", lower, upper)
		verificationFailure(t, "is in range", expectedDescr, actual)
		return false
	}
	return true
}

// OutOfRange checks if the given value is outside lower and upper bounds. It's the
// opposite of InRange.
func OutOfRange[C constraints.Integer | constraints.Float](t T, lower, upper, actual C) bool {
	if lower > upper {
		lower, upper = upper, lower
	}
	if actual >= lower && actual <= upper {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		expectedDescr := fmt.Sprintf("'%v' to '%v'", lower, upper)
		verificationFailure(t, "is out of range", expectedDescr, actual)
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
func NoError(t T, err error) bool {
	if err != nil {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is no error", nil, err)
		return false
	}
	return true
}

// IsError checks if the given error is not nil and of the expected type.
// It uses the errors.Is() function.
func IsError(t T, expected, actual error) bool {
	if !errors.Is(expected, actual) {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is expected error", expected, actual)
		return false
	}
	return true
}

// ErrorContains check if the given error is not nil and its message
// contains an expected string.
func ErrorContains(t T, expected string, actual error) bool {
	if actual == nil {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "error contains", expected, actual)
		return false
	}
	if !strings.Contains(actual.Error(), expected) {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "error contains", expected, actual.Error())
		return false
	}
	return true
}

// ErrorMatch checks if the given error is not nil and its message
// matches the expected regular expression.
func ErrorMatch(t T, expected string, actual error) bool {
	if actual == nil {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "error does match", expected, actual)
		return false
	}
	re := regexp.MustCompile(expected)
	if !re.MatchString(actual.Error()) {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "error does match", expected, actual.Error())
		return false
	}
	return true
}

// Implements checks if the given instance implements the expected interface.
// The expected parameter has to be an interface type as nil pointer like
// (*fmt.Stringer)(nil) or (*io.Reader)(nil).
func Implements(t T, expected, actual any) bool {
	if expected == nil {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "does implement", "expected instance", nil)
		return false
	}

	if actual == nil {
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

	actualType := reflect.TypeOf(actual)
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
func Assignability(t T, expected, actual any) bool {
	if expected == nil {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is assignable to", "expected type", nil)
		return false
	}

	if actual == nil {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "is assignable to", "actual type", nil)
		return false
	}

	expectedType := reflect.TypeOf(expected)
	actualType := reflect.TypeOf(actual)

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
func Panics(t T, fn func()) bool {
	if fn == nil {
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

	fn()
	return true
}

// NotPanics checks if the given functions does not panic.
func NotPanics(t T, fn func()) bool {
	if fn == nil {
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

	fn()
	return true
}

// -----------------------------------------------------------------------------
// Helper
// -----------------------------------------------------------------------------

// ftim is a short to format times in test output.
func ftim(t time.Time) string {
	return t.Format(time.RFC3339)
}

// -----------------------------------------------------------------------------
// EOF
// -----------------------------------------------------------------------------
