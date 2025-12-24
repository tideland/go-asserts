// Tideland Go Asserts - Verify - Time
//
// Copyright (C) 2024-2025 Frank Mueller / Tideland / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package verify

import (
	"fmt"
	"testing"
	"time"
)

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
// The boundaries are inclusive (gotten can equal expectedBegin or expectedEnd).
// If expectedBegin is after expectedEnd, they will be automatically swapped.
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
	if gotten >= expected {
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
	if gotten <= expected {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "duration is longer", expected, gotten, infos...)
		return false
	}
	return true
}

// ftim formats a time.Time value into RFC3339 format for consistent
// test output.
func ftim(t time.Time) string {
	return t.Format(time.RFC3339)
}
