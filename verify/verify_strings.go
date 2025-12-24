// Tideland Go Asserts - Verify - Strings
//
// Copyright (C) 2024-2025 Frank Mueller / Tideland / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package verify

import (
	"reflect"
	"regexp"
	"strings"
	"testing"
)

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

// flexlen returns the length of types available to return their length.
func flexlen(in any) int {
	// Check for possible existing methods
	switch in := in.(type) {
	case lenner:
		return in.Len()
	case lengthier:
		return in.Length()
	default:
		// Use reflection
		rv := reflect.ValueOf(in)
		switch rv.Kind() {
		case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
			return rv.Len()
		default:
			// Good old -1 is enough here, verification is above
			return -1
		}
	}
}

type lenner interface {
	Len() int
}

type lengthier interface {
	Length() int
}

// HasPrefix checks if the gotten string has the expected prefix.
func HasPrefix(t T, gotten, expected string, infos ...string) bool {
	if !strings.HasPrefix(gotten, expected) {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "has prefix", expected, gotten, infos...)
		return false
	}
	return true
}

// HasSuffix checks if the gotten string has the expected suffix.
func HasSuffix(t T, gotten, expected string, infos ...string) bool {
	if !strings.HasSuffix(gotten, expected) {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "has suffix", expected, gotten, infos...)
		return false
	}
	return true
}

// ContainsIgnoreCase checks if the haystack string contains the needle string (case-insensitive).
func ContainsIgnoreCase(t T, haystack, needle string, infos ...string) bool {
	if !strings.Contains(strings.ToLower(haystack), strings.ToLower(needle)) {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "contains (case-insensitive)", needle, haystack, infos...)
		return false
	}
	return true
}
