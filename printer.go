// -----------------------------------------------------------------------------
// Asserts for a more convenient testing in Go libraries and applications.
//
// Printer for failure details
//
// Copyright (C) 2034-2025 Frank Mueller / Oldenburg / Germany / Earth
// -----------------------------------------------------------------------------

package asserts // import "tideland.dev/go/asserts"

import (
	"bufio"
	"bytes"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"unicode"
	"unicode/utf8"
)

const indent = "    "

// logf prints a log message with the given information on stdout.
func logf(format string, args ...any) {
	location := here(0)
	format = fmt.Sprintf("%s assertion log: %s\n", location, format)
	outout := fmt.Sprintf(format, args...)

	fmt.Println(outout)
}

// failf prints a fail message with the given information on stderr.
func failf(t Tester, assertion string, format string, args ...any) {
	location := here(2)
	// testingWd
	format = fmt.Sprintf("%s assertion '%s' fail: %s\n", location, assertion, format)

	t.Errorf(format, args...)
}

// here returns the filename and position based on a given offset.
// It's used by the fail function to print the location of the failure.
func here(skip int) string {
	_, file, line, ok := runtime.Caller(skip + 1)
	if !ok {
		return "unknown file:"
	}
	filename := filepath.Base(file)
	return fmt.Sprintf("%s%s:%d:", indent, filename, line)
}

// Stolen from the `go test` tool.
// isTest tells whether name looks like a test (or benchmark, according to prefix).
// It is a Test (say) if there is a character after Test that is not a lower-case letter.
// We don't want TesticularCancer.
func isTest(name, prefix string) bool {
	if !strings.HasPrefix(name, prefix) {
		return false
	}
	if len(name) == len(prefix) { // "Test" is ok
		return true
	}
	r, _ := utf8.DecodeRuneInString(name[len(prefix):])
	return !unicode.IsLower(r)
}

func messageFromMsgAndArgs(msgAndArgs ...any) string {
	if len(msgAndArgs) == 0 || msgAndArgs == nil {
		return ""
	}
	if len(msgAndArgs) == 1 {
		msg := msgAndArgs[0]
		if msgAsStr, ok := msg.(string); ok {
			return msgAsStr
		}
		return fmt.Sprintf("%+v", msg)
	}
	if len(msgAndArgs) > 1 {
		return fmt.Sprintf(msgAndArgs[0].(string), msgAndArgs[1:]...)
	}
	return ""
}

// Aligns the provided message so that all lines after the first line start at the same location as the first line.
// Assumes that the first line starts at the correct location (after carriage return, tab, label, spacer and tab).
// The longestLabelLen parameter specifies the length of the longest label in the output (required because this is the
// basis on which the alignment occurs).
func indentMessageLines(message string, longestLabelLen int) string {
	outBuf := new(bytes.Buffer)

	for i, scanner := 0, bufio.NewScanner(strings.NewReader(message)); scanner.Scan(); i++ {
		// no need to align first line because it starts at the correct location (after the label)
		if i != 0 {
			// append alignLen+1 spaces to align with "{{longestLabel}}:" before adding tab
			outBuf.WriteString("\n\t" + strings.Repeat(" ", longestLabelLen+1) + "\t")
		}
		outBuf.WriteString(scanner.Text())
	}

	return outBuf.String()
}

// -----------------------------------------------------------------------------
// EOF
// -----------------------------------------------------------------------------
