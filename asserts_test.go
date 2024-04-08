// -----------------------------------------------------------------------------
// asserts for more convinient testing - tests
//
// Copyright (C) 2024 Frank Mueller / Oldenburg / Germany / World
// -----------------------------------------------------------------------------

package asserts_test

import (
	"errors"
	"testing"

	"tideland.dev/go/asserts"
)

// TestLogf tests the log function.
func TestLogf(t *testing.T) {
	asserts.Logf(t, "Hello, %s!", "World")
}

// TestBooleans tests the True and False assertions.
func TestBooleans(t *testing.T) {
	pos, neg := asserts.MkPosNeg(t)

	asserts.True(pos, true)
	asserts.True(neg, false)
	asserts.False(pos, false)
	asserts.False(neg, true)

	asserts.True(pos, asserts.Failed(neg, 2))
}

// TestNils tests the Nil and NotNil assertions.
func TestNils(t *testing.T) {
	pos, neg := asserts.MkPosNeg(t)

	asserts.Nil(pos, nil)
	asserts.Nil(neg, "not nil")
	asserts.NotNil(pos, "not nil")
	asserts.NotNil(neg, nil)

	asserts.True(pos, asserts.Failed(neg, 2))
}

// TestComparisons tests the Equal and Different assertions.
func TestComparisons(t *testing.T) {
	pos, neg := asserts.MkPosNeg(t)

	asserts.Equal(pos, 42, 42)
	asserts.Equal(neg, 42, 43)
	asserts.Equal(pos, "foo", "foo")
	asserts.Equal(neg, "foo", "bar")

	asserts.Different(pos, 42, 43)
	asserts.Different(neg, 42, 42)
	asserts.Different(pos, "foo", "bar")
	asserts.Different(neg, "foo", "foo")

	asserts.AboutEqual(pos, 45, 43, 5)
	asserts.AboutEqual(neg, 4.5, 4.3, 0.1)

	asserts.True(pos, asserts.Failed(neg, 5))
}

// TestErrors tests the Error assertions.
func TestErrors(t *testing.T) {
	pos, neg := asserts.MkPosNeg(t)

	testErr := errors.New("booom")

	asserts.Error(pos, testErr)
	asserts.Error(neg, nil)
	asserts.NoError(pos, nil)
	asserts.NoError(neg, testErr)
	asserts.IsError(pos, testErr, testErr)
	asserts.IsError(neg, testErr, errors.New("ouch"))
	asserts.ErrorContains(pos, testErr, "ooo")
	asserts.ErrorContains(neg, testErr, "BOOOM")
	asserts.ErrorMatches(pos, testErr, "^bo.*")
	asserts.ErrorMatches(neg, testErr, ".*ou$")

	asserts.True(pos, asserts.Failed(neg, 5))
}

// -----------------------------------------------------------------------------
// end of file
// -----------------------------------------------------------------------------
