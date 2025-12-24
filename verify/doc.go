// Tideland Go Asserts - Verify
//
// Copyright (C) 2024-2025 Frank Mueller / Tideland / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

/*
Package verify provides a comprehensive and flexible set of assertion functions
for Go's standard testing package. It is designed to write clear, readable,
and expressive tests.

A key feature is the ability to perform continued testing. By wrapping the
standard *testing.T with verify.ContinuedTesting(), assertions will report
failures without halting the test execution. This allows for multiple
independent checks within a single test function, reporting all failures
at the end. The number of expected failures can be asserted using
verify.FailureCount().

Example for a test using continued testing:

	func TestMyFunction(t *testing.T) {
		// Wrap testing.T for continued testing.
		ct := verify.ContinuedTesting(t)

		result, err := myFunction("input")
		verify.NoError(ct, err, "myFunction should not produce an error")
		verify.Length(ct, result, 7, "result should have a length of 7")
		verify.Match(ct, result, "^success:")

		value := 42
		verify.Positive(ct, value)
		verify.Even(ct, value)

		// At the end check the number of failures. Here we expect zero,
		// but if any of the above checks failed, this will fail and
		// report the number of found failures.
		verify.FailureCount(ct, 0)
	}
*/
package verify
