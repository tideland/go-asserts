// -----------------------------------------------------------------------------
// Asserts for a more convenient testing in Go libraries and applications.
//
// Allow to capturing of stdout and stderr
//
// Copyright (C) 2034-2025 Frank Mueller / Oldenburg / Germany / Earth
// -----------------------------------------------------------------------------

package asserts // "tideland.dev/go/asserts"

import (
	"bytes"
	"io"
	"log"
	"os"
)

// Captured provides access to the captured output in
// multiple ways.
type Captured struct {
	buffer []byte
}

// Bytes returns the captured content as bytes.
func (c Captured) Bytes() []byte {
	buf := make([]byte, c.Len())
	copy(buf, c.buffer)
	return buf
}

// String implements fmt.Stringer.
func (c Captured) String() string {
	return string(c.Bytes())
}

// Len returns the number of captured bytes.
func (c Captured) Len() int {
	return len(c.buffer)
}

// CaptureStdout allows to capture Stdout by the given function.
// The result is stored in Captured and can be retrieved as
// []byte or string for aseertions.
func CaptureStdout(f func()) Captured {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	outC := make(chan []byte)

	go func() {
		var buf bytes.Buffer
		if _, err := io.Copy(&buf, r); err != nil {
			log.Fatalf("error capturing stdout: %v", err)
		}
		outC <- buf.Bytes()
	}()

	w.Close()
	os.Stdout = old
	return Captured{
		buffer: <-outC,
	}
}

// CaptureStdout allows to capture Stderr by the given function.
// The result is stored in Captured and can be retrieved as
// []byte or string for aseertions.
func CaptureStderr(f func()) Captured {
	old := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	f()

	outC := make(chan []byte)

	go func() {
		var buf bytes.Buffer
		if _, err := io.Copy(&buf, r); err != nil {
			log.Fatalf("error capturing stderr: %v", err)
		}
		outC <- buf.Bytes()
	}()

	w.Close()
	os.Stderr = old
	return Captured{
		buffer: <-outC,
	}
}

// CaptureStdout allows to capture Stdout and Stderr by the given
// function. The result is stored in two Captureds for each and can 
// be retrieved as []byte or string for aseertions.
func CaptureBoth(f func()) (Captured, Captured) {
	var cerr Captured
	ff := func() {
		cerr = CaptureStderr(f)
	}
	cout := CaptureStdout(ff)
	return cout, cerr
}

// -----------------------------------------------------------------------------
// EOF
// -----------------------------------------------------------------------------
