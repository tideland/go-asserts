// Tideland Go Asserts - Capture
//
// Copyright (C) 2024-2025 Frank Mueller / Tideland / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

/*
Package capture provides utilities for capturing stdout and stderr output
during test execution. This is useful for testing functions that write to
standard output or error streams.

The package provides three main capture functions:

  - Stdout: Captures only stdout
  - Stderr: Captures only stderr
  - Both: Captures both stdout and stderr simultaneously

Each function accepts a function to execute and returns a Captured result
that can be inspected as bytes or string.

Example for capturing stdout:

	func TestPrintFunction(t *testing.T) {
		captured := capture.Stdout(func() {
			fmt.Println("Hello, World!")
		})

		verify.Equal(t, "Hello, World!\n", captured.String())
		verify.Equal(t, 14, captured.Len())
	}

Example for capturing both stdout and stderr:

	func TestOutputs(t *testing.T) {
		cout, cerr := capture.Both(func() {
			fmt.Fprintln(os.Stdout, "standard output")
			fmt.Fprintln(os.Stderr, "standard error")
		})

		verify.Contains(t, cout.String(), "standard output")
		verify.Contains(t, cerr.String(), "standard error")
	}

The captured output is stored in a Captured struct that provides:

  - Bytes(): Returns the captured content as a byte slice
  - String(): Returns the captured content as a string
  - Len(): Returns the number of bytes captured

All capture functions ensure that stdout/stderr are properly restored even
if the captured function panics, preventing test pollution.
*/
package capture
