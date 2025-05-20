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
	verify.Equal(t, hello, cptrd.String())
	verify.Equal(t, len(hello), cptrd.Len())
}

// TestStderr tests the capturing of writings to stderr.
func TestStderr(t *testing.T) {
	ouch := "ouch"
	cptrd := capture.Stderr(func() {
		fmt.Fprint(os.Stderr, ouch)
	})
	verify.Equal(t, ouch, cptrd.String())
	verify.Equal(t, len(ouch), cptrd.Len())
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
	verify.Equal(t, hello, cout.String())
	verify.Equal(t, len(hello), cout.Len())
	verify.Equal(t, ouch, cerr.String())
	verify.Equal(t, len(ouch), cerr.Len())
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
	verify.Equal(t, foo, cout.String())
	verify.Equal(t, len(foo), cout.Len())
	verify.Equal(t, foo, cerr.String())
	verify.Equal(t, len(foo), cerr.Len())
	verify.Equal(t, oldOut, os.Stdout)
	verify.Equal(t, oldErr, os.Stderr)
}

// -----------------------------------------------------------------------------
// EOF
// -----------------------------------------------------------------------------
