// -----------------------------------------------------------------------------
// Convenient verification of unit tests in Go libraries and applications.
//
// A set of individual verifications
//
// Copyright (C) 2034-2025 Frank Mueller / Oldenburg / Germany / Earth
// -----------------------------------------------------------------------------

package verify // import "tideland.dev/go/assert/verify"

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"testing"
	"time"

	"golang.org/x/exp/constraints"
)

// True checks if the given value is true.
func True(t T, actual bool) bool {
	if !actual {
		failAfterVerification(t, "true", true, actual)
		return false
	}
	return true
}

// False checks if the given value is false. It's the opposite of True.
func False(t T, actual bool) bool {
	if !actual {
		failAfterVerification(t, "false", false, actual)
		return false
	}
	return true
}

// Nil checks if the given value is nil.
func Nil(t T, actual any) bool {
	if actual != nil {
		failAfterVerification(t, "nil", nil, actual)
		return false
	}
	return true
}

// NotNil checks if the given value is not nil. It's the opposite of Nil.
func NotNil(t T, actual any) bool {
	if actual == nil {
		failAfterVerification(t, "not nil", nil, actual)
		return false
	}
	return true
}

// Equal checks if the given values are equal.
// It uses the == operator for comparable types and supports time.Duration.
func Equal[C comparable](t T, expected, actual C) bool {
	if expected != actual {
		failAfterVerification(t, "equal", expected, actual)
		return false
	}
	return true
}

// Different checks if the given values are different.
// It uses the != operator for comparable types and supports time.Duration.
func Different[C comparable](t T, expected, actual C) bool {
	if expected == actual {
		failAfterVerification(t, "different", expected, actual)
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
			failAfterVerification(t, "length", length, actualLength)
			return false
		}
	default:
		failAfterVerification(t, "length", length, "not quantifiable")
		return false
	}
	return true
}

// Match checks if the actual string matches the expected regular expression.
// The regular expression is compiled from the expected string. If the compilation
// fails, the assertion fails.
func Match(t T, expected, actual string) bool {
	re, err := regexp.Compile(expected)
	if err != nil {
		failAfterVerification(t, "match", expected, err)
		return false
	}
	if !re.MatchString(actual) {
		failAfterVerification(t, "match", expected, actual)
		return false
	}
	return true
}

// Less checks if the actual value is less than the expected one.
// Supports integers, floats, and time.Duration.
func Less[C constraints.Integer | constraints.Float](t T, expected, actual C) bool {
	if actual >= expected {
		failAfterVerification(t, "less", expected, actual)
		return false
	}
	return true
}

// More checks if the actual value is more than the expected one.
// Supports integers, floats, and time.Duration.
func More[C constraints.Integer | constraints.Float](t T, expected, actual C) bool {
	if actual <= expected {
		failAfterVerification(t, "more", expected, actual)
		return false
	}
	return true
}

// AboutEqual checks if the given values are equal within a delta. Possible
// values are integers, floats, and time.Duration.
func AboutEqual[C constraints.Integer | constraints.Float](t T, expected, actual, delta C) bool {
	if expected < actual-delta || expected > actual+delta {
		expectedDescr := fmt.Sprintf("%v' +/- '%v'", expected, delta)
		failAfterVerification(t, "about equal", expectedDescr, actual)
		return false
	}
	return true
}

// Before checks if the actual time is before the expected time.
func Before(t T, expected, actual time.Time) bool {
	if !actual.Before(expected) {
		failAfterVerification(t, "time before", expected, actual)
		return false
	}
	return true
}

// After checks if the actual time is after the expected time.
func After(t T, expected, actual time.Time) bool {
	if !actual.After(expected) {
		failAfterVerification(t, "time after", expected, actual)
		return false
	}
	return true
}

// Between checks if the actual time is between the expected start and end times.
func Between(t T, start, end, actual time.Time) bool {
	if actual.Before(start) || actual.After(end) {
		expectedDescr := fmt.Sprintf("'%v'/ '%v'", start, end)
		failAfterVerification(t, "time after", expectedDescr, actual)
		return false
	}
	return true
}

// Duration calculates the duration of a function execution. If the function returns an error,
// the test fails. The returned duration can be used as expected duration in further tests.
func Duration(t T, fn func() error) (time.Duration, error) {
	start := time.Now()
	err := fn()
	if err != nil {
		failAfterVerification(t, "func exec", nil, err)
		return 0, err
	}
	return time.Since(start), nil
}

// Shorter checks if the actual duration is shorter than the expected duration.
func Shorter(t T, expected, actual time.Duration) bool {
	if actual > expected {
		failAfterVerification(t, "duration shorter", expected, actual)
		return false
	}
	return true
}

// Longer checks if the actual duration is longer than the expected duration.
func Longer(t T, expected, actual time.Duration) bool {
	if actual < expected {
		failAfterVerification(t, "duration longer", expected, actual)
		return false
	}
	return true
}

// DurationAboutEqual checks if the given durations are equal within a delta.
func DurationAboutEqual(t T, expected, actual, delta time.Duration) bool {
	if expected < actual-delta || expected > actual+delta {
		expectedDesc := fmt.Sprintf("'%v +/- '%s'", expected, delta)
		failAfterVerification(t, "duration about equal", expectedDesc, actual)
		return false
	}
	return true
}

// InRange checks if the given value is within lower and upper bounds. Possible
// values are integers, floats, and time.Duration.
func InRange[C constraints.Integer | constraints.Float](t *testing.T, lower, upper, actual C) bool {
	if lower > upper {
		lower, upper = upper, lower
	}
	if actual <= lower || actual >= upper {
		expectedDescr := fmt.Sprintf("value from '%v'to '%v' ", lower, upper)
		failAfterVerification(t, "in range", expectedDescr, actual)
		return false
	}
	return true
}

// OutOfRange checks if the given value is outside lower and upper bounds. It's the
// opposite of InRange.
func OutOfRange[C constraints.Integer | constraints.Float](t *testing.T, lower, upper, actual C) bool {
	if lower > upper {
		lower, upper = upper, lower
	}
	if actual >= lower && actual <= upper {
		expectedDescr := fmt.Sprintf("value lower than '%v' and higher than '%v'", lower, upper)
		failAfterVerification(t, "out of range", expectedDescr, actual)
		return false
	}
	return true
}

// Error checks if the given error is not nil.
func Error(t T, err error) bool {
	if err == nil {
		failAfterVerification(t, "error", "error", nil)
		return false
	}
	return true
}

// NoError checks if the given error is nil.
// It's the opposite of Error.
func NoError(t T, err error) bool {
	if err != nil {
		failAfterVerification(t, "no error", nil, err)
		return false
	}
	return true
}

// IsError checks if the given error is not nil and of the expected type.
// It uses the errors.Is() function.
func IsError(t T, expected, actual error) bool {
	if !errors.Is(expected, actual) {
		failAfterVerification(t, "is error", expected, actual)
		return false
	}
	return true
}

// ErrorMatch checks if the given error is not nil and its message
// matches the expected regular expression.
func ErrorMatch(t *testing.T, expected string, actual error) bool {
	if actual == nil {
		failAfterVerification(t, "error match", expected, actual)
		return false
	}
	re := regexp.MustCompile(expected)
	if !re.MatchString(actual.Error()) {
		failAfterVerification(t, "error match", expected, actual.Error())
		return false
	}
	return true
}

// Implements checks if the given instance implements the expected interface.
// The expected parameter has to be an interface type as nil pointer like
// (*fmt.Stringer)(nil) or (*io.Reader)(nil).
func Implements(t *testing.T, expected, actual any) bool {
	if expected == nil {
		failAfterVerification(t, "implements", "expected instance", nil)
		return false
	}

	if actual == nil {
		failAfterVerification(t, "implements", "actual instance", nil)
		return false
	}

	expectedType := reflect.TypeOf(expected).Elem()
	if expectedType.Kind() != reflect.Interface {
		failAfterVerification(t, "implements", "expected interface", nil)
		return false
	}

	actualType := reflect.TypeOf(actual)
	if !actualType.Implements(expectedType) {
		failAfterVerification(t, "implements", expectedType, actualType)
		return false
	}
	return true
}

// -----------------------------------------------------------------------------
// EOF
// -----------------------------------------------------------------------------
