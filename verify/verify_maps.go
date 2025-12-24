// Tideland Go Asserts - Verify - Maps
//
// Copyright (C) 2024-2025 Frank Mueller / Tideland / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package verify

import (
	"fmt"
	"maps"
	"testing"
)

// MapEqual checks if two maps are equal using `maps.Equal`, which compares
// keys and values with the `==` operator. For deep equality of complex
// value types, use `DeepEqual`.
func MapEqual[M ~map[K]V, K comparable, V comparable](t T, gotten, expected M, infos ...string) bool {
	if !maps.Equal(gotten, expected) {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "maps are equal", expected, gotten, infos...)
		return false
	}
	return true
}

// MapContainsKey checks if the map contains the specified key.
func MapContainsKey[M ~map[K]V, K comparable, V any](t T, gotten M, key K, infos ...string) bool {
	if _, exists := gotten[key]; !exists {
		if ht, ok := t.(testing.TB); ok {
			ht.Helper()
		}
		verificationFailure(t, "map contains key", key, "key not found", infos...)
		return false
	}
	return true
}

// MapContainsKeys checks if the map contains all specified keys.
func MapContainsKeys[M ~map[K]V, K comparable, V any](t T, gotten M, keys []K, infos ...string) bool {
	for _, key := range keys {
		if _, exists := gotten[key]; !exists {
			if ht, ok := t.(testing.TB); ok {
				ht.Helper()
			}
			verificationFailure(t, "map contains all keys", keys, fmt.Sprintf("missing key: %v", key), infos...)
			return false
		}
	}
	return true
}

// MapContainsValue checks if any key in the map is associated with the
// specified value.
func MapContainsValue[M ~map[K]V, K comparable, V comparable](t T, gotten M, value V, infos ...string) bool {
	for _, v := range gotten {
		if v == value {
			return true
		}
	}
	if ht, ok := t.(testing.TB); ok {
		ht.Helper()
	}
	verificationFailure(t, "map contains value", value, "value not found", infos...)
	return false
}
