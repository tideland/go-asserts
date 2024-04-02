// -----------------------------------------------------------------------------
// asserts for more convinient testing - tests
//
// Copyright (C) 2024 Frank Mueller / Oldenburg / Germany / World
// -----------------------------------------------------------------------------

package asserts_test

import (
	"testing"

	"tideland.dev/go/asserts"
)

// TestLogf tests the log function.
func TestLogf(t *testing.T) {
	asserts.Logf("Hello, %s!", "World")
}

// TestTrueFalse tests the True and False assertions.
func TestTrueFalse(t *testing.T) {
	pt, nt := asserts.MkPosNeg(t)

	asserts.True(pt, true)
	asserts.True(nt, false)
	asserts.True(pt, asserts.Failed(nt))

	asserts.False(pt, false)
	asserts.False(nt, true)
	asserts.True(pt, asserts.Failed(nt))
}

// TestNilNotNil tests the Nil and NotNil assertions.
func TestNilNotNil(t *testing.T) {
	pt, nt := asserts.MkPosNeg(t)

	asserts.Nil(pt, nil)
	asserts.Nil(nt, "not nil")
	asserts.True(pt, asserts.Failed(nt))

	asserts.NotNil(pt, "not nil")
	asserts.NotNil(nt, nil)
	asserts.True(pt, asserts.Failed(nt))
}

// -----------------------------------------------------------------------------
// end of file
// -----------------------------------------------------------------------------
