// -----------------------------------------------------------------------------
// asserts for more convinient testing - documentation
//
// Copyright (C) 2024 Frank Mueller / Oldenburg / Germany / World
// -----------------------------------------------------------------------------

// Package asserts helps writing convenient and powerful unit tests. There are different
// assertions for booleans, nils, comparisons, errors, and time. Direct fails or just
// logging is possible. The package is designed to be used with the standard testing
// package. The testing argument of the functions is the *testing.T. So statements like
// asserts.True(t, myBoolean) or asserts.Less(t, expected, actual) are possible. At the
// end of a test function the number of failures can be checked with asserts.Failures(tester, 2).

package asserts // import "tideland.dev/go/asserts"

// -----------------------------------------------------------------------------
// end of file
// -----------------------------------------------------------------------------
