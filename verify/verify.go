// Tideland Go Asserts - Verify
//
// Copyright (C) 2024-2025 Frank Mueller / Tideland / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package verify

import (
	"testing"
)

// Verifications

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
		verificationFailure(t, "is false", false, gotten, infos...)
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
