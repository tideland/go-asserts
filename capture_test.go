// -----------------------------------------------------------------------------
// Asserts for a more convenient testing in Go libraries and applications.
//
// Capturing unit tests
//
// Copyright (C) 2024-2025 Frank Mueller / Oldenburg / Germany / Earth
// -----------------------------------------------------------------------------

package asserts_test

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"tideland.dev/go/asserts"
)

// TestStdout tests the capturing of writings to stdout.
func TestStdout(t *testing.T) {
	tester := asserts.NewTester(t, asserts.FAIL)

	hello := "Hello, World!"
	cptrd := asserts.CaptureStdout(func() {
		fmt.Print(hello)
	})
	asserts.Equal(tester, cptrd.String(), hello)
	asserts.Equal(tester, cptrd.Len(), len(hello))
}

// TestStderr tests the capturing of writings to stderr.
func TestStderr(t *testing.T) {
	tester := asserts.NewTester(t, asserts.FAIL)

	ouch := "ouch"
	cptrd := asserts.CaptureStderr(func() {
		fmt.Fprint(os.Stderr, ouch)
	})
	asserts.Equal(tester, cptrd.String(), ouch)
	asserts.Equal(tester, cptrd.Len(), len(ouch))
}

// TestBoth tests the capturing of writings to stdout
// and stderr.
func TestBoth(t *testing.T) {
	tester := asserts.NewTester(t, asserts.FAIL)

	hello := "Hello, World!"
	ouch := "ouch"
	cout, cerr := asserts.CaptureBoth(func() {
		fmt.Fprint(os.Stdout, hello)
		fmt.Fprint(os.Stderr, ouch)
	})
	asserts.Equal(tester, cout.String(), hello)
	asserts.Equal(tester, cout.Len(), len(hello))
	asserts.Equal(tester, cerr.String(), ouch)
	asserts.Equal(tester, cerr.Len(), len(ouch))
}

// TestBytes tests the retrieving of captures as bytes.
func TestBytes(t *testing.T) {
	tester := asserts.NewTester(t, asserts.FAIL)

	foo := "foo"
	boo := []byte(foo)
	cout, cerr := asserts.CaptureBoth(func() {
		fmt.Fprint(os.Stdout, foo)
		fmt.Fprint(os.Stderr, foo)
	})
	asserts.True(tester, bytes.Equal(cout.Bytes(), boo))
	asserts.True(tester, bytes.Equal(cerr.Bytes(), boo))
}

// TestRestore tests the restoring of os.Stdout
// and os.Stderr after capturing.
func TestRestore(t *testing.T) {
	tester := asserts.NewTester(t, asserts.FAIL)

	foo := "foo"
	oldOut := os.Stdout
	oldErr := os.Stderr
	cout, cerr := asserts.CaptureBoth(func() {
		fmt.Fprint(os.Stdout, foo)
		fmt.Fprint(os.Stderr, foo)
	})
	asserts.Equal(tester, cout.String(), foo)
	asserts.Equal(tester, cout.Len(), len(foo))
	asserts.Equal(tester, cerr.String(), foo)
	asserts.Equal(tester, cerr.Len(), len(foo))
	asserts.Equal(tester, os.Stdout, oldOut)
	asserts.Equal(tester, os.Stderr, oldErr)
}

// -----------------------------------------------------------------------------
// EOF
// -----------------------------------------------------------------------------
