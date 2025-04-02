// -----------------------------------------------------------------------------
// Asserts for a more convenient testing in Go libraries and applications.
//
// Unit tests
//
// Copyright (C) 2024-2025 Frank Mueller / Oldenburg / Germany / Earth
// -----------------------------------------------------------------------------

package capture_test

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"tideland.dev/go/asserts/verify"

	"tideland.dev/go/asserts/capture"
)

// TestStdout tests the capturing of writings to stdout.
func TestStdout(t *testing.T) {
	hello := "Hello, World!"
	cptrd := capture.Stdout(func() {
		fmt.Print(hello)
	})
	verify.Equal(t, cptrd.String(), hello)
	verify.Equal(t, cptrd.Len(), len(hello))
}

// TestStderr tests the capturing of writings to stderr.
func TestStderr(t *testing.T) {
	ouch := "ouch"
	cptrd := capture.Stderr(func() {
		fmt.Fprint(os.Stderr, ouch)
	})
	verify.Equal(t, cptrd.String(), ouch)
	verify.Equal(t, cptrd.Len(), len(ouch))
}

// TestBoth tests the capturing of writings to stdout
// and stderr.
func TestBoth(t *testing.T) {
	hello := "Hello, World!"
	ouch := "ouch"
	cout, cerr := capture.Both(func() {
		fmt.Fprint(os.Stdout, hello)
		fmt.Fprint(os.Stderr, ouch)
	})
	verify.Equal(t, cout.String(), hello)
	verify.Equal(t, cout.Len(), len(hello))
	verify.Equal(t, cerr.String(), ouch)
	verify.Equal(t, cerr.Len(), len(ouch))
}

// TestBytes tests the retrieving of captures as bytes.
func TestBytes(t *testing.T) {
	foo := "foo"
	boo := []byte(foo)
	cout, cerr := capture.Both(func() {
		fmt.Fprint(os.Stdout, foo)
		fmt.Fprint(os.Stderr, foo)
	})
	verify.True(t, bytes.Equal(cout.Bytes(), boo))
	verify.True(t, bytes.Equal(cerr.Bytes(), boo))
}

// TestRestore tests the restoring of os.Stdout
// and os.Stderr after capturing.
func TestRestore(t *testing.T) {
	foo := "foo"
	oldOut := os.Stdout
	oldErr := os.Stderr
	cout, cerr := capture.Both(func() {
		fmt.Fprint(os.Stdout, foo)
		fmt.Fprint(os.Stderr, foo)
	})
	verify.Equal(t, cout.String(), foo)
	verify.Equal(t, cout.Len(), len(foo))
	verify.Equal(t, cerr.String(), foo)
	verify.Equal(t, cerr.Len(), len(foo))
	verify.Equal(t, os.Stdout, oldOut)
	verify.Equal(t, os.Stderr, oldErr)
}

// -----------------------------------------------------------------------------
// EOF
// -----------------------------------------------------------------------------
