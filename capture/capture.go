// Tideland Go Asserts - Capture
//
// Copyright (C) 2024-2025 Frank Mueller / Tideland / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package capture

import (
	"bytes"
	"io"
	"os"
)

// Captured provides access to the captured output in
// multiple ways.
type Captured struct {
	buffer []byte
}

// Bytes returns the captured content as bytes.
func (c Captured) Bytes() []byte {
	return c.buffer
}

// String implements fmt.Stringer.
func (c Captured) String() string {
	return string(c.Bytes())
}

// Len returns the number of captured bytes.
func (c Captured) Len() int {
	return len(c.buffer)
}

// Stdout allows to capture Stdout by the given function.
// The result is stored in Captured and can be retrieved as
// []byte or string for assertions.
func Stdout(f func()) Captured {
	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		panic("failed to create pipe: " + err.Error())
	}

	os.Stdout = w
	outC := make(chan []byte)

	// Start goroutine to read from pipe
	go func() {
		var buf bytes.Buffer
		_, _ = io.Copy(&buf, r)
		outC <- buf.Bytes()
	}()

	// Ensure restoration even if f() panics
	defer func() {
		os.Stdout = old
	}()

	f()
	if err := w.Close(); err != nil {
		panic("failed to close pipe: " + err.Error())
	}

	return Captured{
		buffer: <-outC,
	}
}

// Stderr allows to capture Stderr by the given function.
// The result is stored in Captured and can be retrieved as
// []byte or string for assertions.
func Stderr(f func()) Captured {
	old := os.Stderr
	r, w, err := os.Pipe()
	if err != nil {
		panic("failed to create pipe: " + err.Error())
	}

	os.Stderr = w
	outC := make(chan []byte)

	// Start goroutine to read from pipe
	go func() {
		var buf bytes.Buffer
		_, _ = io.Copy(&buf, r)
		outC <- buf.Bytes()
	}()

	// Ensure restoration even if f() panics
	defer func() {
		os.Stderr = old
	}()

	f()
	if err := w.Close(); err != nil {
		panic("failed to close pipe: " + err.Error())
	}

	return Captured{
		buffer: <-outC,
	}
}

// Both allows to capture Stdout and Stderr by the given
// function. The result is stored in two Captureds for each and can
// be retrieved as []byte or string for assertions.
func Both(f func()) (Captured, Captured) {
	var cerr Captured
	ff := func() {
		cerr = Stderr(f)
	}
	cout := Stdout(ff)
	return cout, cerr
}
