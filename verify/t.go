// -----------------------------------------------------------------------------
// Convenient verification of unit tests in Go libraries and applications.
//
// Replacement of testing.T to allow teest without immediate fail
//
// Copyright (C) 2034-2025 Frank Mueller / Oldenburg / Germany / Earth
// -----------------------------------------------------------------------------

package verify // import "tideland.dev/go/assert/verify"

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// T is an interface that matches the common methods of testing.T
type T interface {
	Fail()
	FailNow()
	Log(args ...any)
	Logf(format string, args ...any)
	Errorf(format string, args ...any)
}

// continueTesting is a wrapper around *testing.T that
// indicates the test should continue running even after
// a verification failure
type continueTesting struct {
	*testing.T
	failed int
}

// Ensure the wrapper implement T
var _ T = (*continueTesting)(nil)

// ContinueTesting creates a new T instance that continues after
// failures.
func ContinueTesting(t *testing.T) T {
	return &continueTesting{t, 0}
}

// IsContinueT checks if a testing.T is a continueTesting type.
func IsContinueT(t T) bool {
	_, ok := t.(*continueTesting)
	return ok
}

// ContinousFailings returns the number of failures if t is a
// continued testing.
func ContinousFailings(t T) int {
	ct, ok := t.(*continueTesting)
	if !ok {
		t.Errorf("not a continuing testing")
	}
	return ct.failed
}

// Logf is used to print additional information during testing.
// The location and function name are added automatically.
func Logf(t *testing.T, format string, args ...any) {
	t.Helper()
	output := fmt.Sprintf(format, args...)
	t.Logf(output)
}

// Failf is a helper function that fails a test with a formatted
// message. If t is a continueTesting, it only calls Errorf, otherwise
// it also calls FailNow.
func Failf(t T, verification string, format string, args ...any) {
	tPtr, _ := t.(*testing.T)
	if tPtr != nil {
		tPtr.Helper()
	}

	output := failureInformation(tPtr, verification, format, args...)

	t.Errorf(output)

	// If it's not a continueTesting wrapper, also call FailNow
	if _, ok := t.(*continueTesting); !ok {
		t.FailNow()
	}

	t.(*continueTesting).failed++
}

// failAfterVerification creates a simple Failf() in default format for
// a failed verification.
func failAfterVerification(t T, verification string, expected, got any) {
	output := fmt.Sprintf("expected %v, got %v", expected, got)
	Failf(t, verification, output)
}

// failureInformation retrieves information and provides a string
// describing location and details about failure.
func failureInformation(t *testing.T, verification string, format string, args ...any) string {
	failmsg := fmt.Sprintf("fail '%s':%s ", verification, fmt.Sprintf(format, args...))

	if t == nil {
		return failmsg
	}

	// Retrieve detailed info
	callInfo := getCallInfo(t)
	indent := strings.Repeat("    ", callInfo.level)
	location := fmt.Sprintf("%s:%d", callInfo.filename, callInfo.line)

	return fmt.Sprintf("%s%s: %s", indent, location, failmsg)
}

// getCallInfo analysiert den Call Stack und gibt die relevante Aufrufposition zurück
func getCallInfo(t *testing.T) callInfo {
	var info callInfo

	info.testname = t.Name()
	parts := strings.Split(info.testname, "/")

	info.level = len(parts) - 1

	// Search in call stack
	for i := 1; i < 10; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}

		// Take none of the test library files
		if strings.Contains(file, "testing.go") ||
			strings.Contains(file, "verify/t.go") ||
			strings.Contains(file, "asserts/") {
			continue
		}

		// Get the function name
		fn := runtime.FuncForPC(pc)
		funcName := fn.Name()

		// Test-Funktionen erkennen (beginnen mit "Test")
		if !strings.Contains(funcName, ".Test") {
			// Found it
			info.filename = filepath.Base(file)
			info.line = line
			info.function = funcName
			break
		}
	}

	return info
}

// callInfo contains information about the failure location
// and description.
type callInfo struct {
	filename string // Dateiname
	line     int    // Zeilennummer
	function string // Funktionsname
	testname string // Name des Tests/Subtests
	level    int    // Verschachtelungsebene des Tests
}

// -----------------------------------------------------------------------------
// EOF
// -----------------------------------------------------------------------------
