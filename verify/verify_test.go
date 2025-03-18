// -----------------------------------------------------------------------------
// Convenient verification of unit tests in Go libraries and applications.
// 
// Unit tests
//
// Copyright (C) 2034-2025 Frank Mueller / Oldenburg / Germany / Earth
// -----------------------------------------------------------------------------

package verify_test

import (
	"testing"

	"tideland.dev/go/assert/verify"
)

// TestTrue tests the True verification function.
func TestTrue(t *testing.T) {
	// Standard case - should pass without error
	verify.True(t, true)

	// Now test with ContinueTesting to check failures
	ct := verify.ContinueTesting(t)
	
	// This should fail
	verify.True(ct, false)
	
	// This should pass
	verify.True(ct, true)
	
	// Check that we had one failure
	if verify.ContinousFailings(ct) != 1 {
		t.Errorf("expected 1 failure, got %d", verify.ContinousFailings(ct))
	}
}

// TestFalse tests the False verification function.
func TestFalse(t *testing.T) {
	// Standard case - should pass without error
	verify.False(t, false)

	// Now test with ContinueTesting to check failures
	ct := verify.ContinueTesting(t)
	
	// This should fail
	verify.False(ct, true)
	
	// This should pass
	verify.False(ct, false)
	
	// Check that we had one failure
	if verify.ContinousFailings(ct) != 1 {
		t.Errorf("expected 1 failure, got %d", verify.ContinousFailings(ct))
	}
}

// TestNilAndNotNil tests the Nil and NotNil verification functions.
func TestNilAndNotNil(t *testing.T) {
	// Test cases for Nil
	var nilPtr *int
	var nonNilPtr = new(int)
	
	// Standard case - should pass without error
	verify.Nil(t, nil)
	verify.Nil(t, nilPtr)
	
	// Now test with ContinueTesting to check failures
	ct := verify.ContinueTesting(t)
	
	// These should fail
	verify.Nil(ct, 42)
	verify.Nil(ct, "string")
	verify.Nil(ct, nonNilPtr)
	
	// These should pass
	verify.Nil(ct, nil)
	verify.Nil(ct, nilPtr)
	
	// Check for NotNil
	verify.NotNil(ct, 42)
	verify.NotNil(ct, "string")
	verify.NotNil(ct, nonNilPtr)
	
	// These should fail
	verify.NotNil(ct, nil)
	verify.NotNil(ct, nilPtr)
	
	// Check that we had expected failures
	if verify.ContinousFailings(ct) != 5 {
		t.Errorf("expected 5 failures, got %d", verify.ContinousFailings(ct))
	}
}

// TestIsContinueT tests the IsContinueT function.
func TestIsContinueT(t *testing.T) {
	// Create a continue testing instance
	ct := verify.ContinueTesting(t)
	
	// Check that it is identified correctly
	if !verify.IsContinueT(ct) {
		t.Error("IsContinueT should recognize ContinueTesting instance")
	}
	
	// Regular testing.T should not be recognized
	if verify.IsContinueT(t) {
		t.Error("IsContinueT should not recognize regular testing.T")
	}
}

// -----------------------------------------------------------------------------
// EOF
// -----------------------------------------------------------------------------
