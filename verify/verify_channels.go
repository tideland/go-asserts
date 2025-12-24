// Tideland Go Asserts - Verify - Channels
//
// Copyright (C) 2024-2025 Frank Mueller / Tideland / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package verify

import (
	"testing"
	"time"
)

// ChannelClosed checks if the given channel is closed.
// It attempts to receive from the channel with a zero timeout.
// If the channel is closed, the receive will succeed immediately with the zero value and ok=false.
func ChannelClosed(t T, gotten <-chan any, infos ...string) bool {
	if ht, ok := t.(testing.TB); ok {
		ht.Helper()
	}

	if gotten == nil {
		verificationFailure(t, "channel is closed", "valid channel", "nil channel", infos...)
		return false
	}

	// Try to receive from the channel with a select and default
	select {
	case _, ok := <-gotten:
		if ok {
			// Channel is open and received a value
			verificationFailure(t, "channel is closed", "closed channel", "open channel with value", infos...)
			return false
		}
		// Channel is closed (ok == false)
		return true
	default:
		// Channel is open but has no value ready
		verificationFailure(t, "channel is closed", "closed channel", "open channel", infos...)
		return false
	}
}

// ChannelReceives checks if the given channel receives a value within the timeout.
// Returns true if a value is received, false if timeout occurs or channel is closed.
func ChannelReceives[E any](t T, gotten <-chan E, timeout time.Duration, infos ...string) bool {
	if ht, ok := t.(testing.TB); ok {
		ht.Helper()
	}

	if gotten == nil {
		verificationFailure(t, "channel receives", "valid channel", "nil channel", infos...)
		return false
	}

	select {
	case _, ok := <-gotten:
		if !ok {
			verificationFailure(t, "channel receives", "open channel", "closed channel", infos...)
			return false
		}
		return true
	case <-time.After(timeout):
		verificationFailure(t, "channel receives", "value within timeout", "timeout exceeded", infos...)
		return false
	}
}

// ChannelReceivesValue checks if the given channel receives the expected value within the timeout.
func ChannelReceivesValue[E comparable](t T, gotten <-chan E, expected E, timeout time.Duration, infos ...string) bool {
	if ht, ok := t.(testing.TB); ok {
		ht.Helper()
	}

	if gotten == nil {
		verificationFailure(t, "channel receives value", "valid channel", "nil channel", infos...)
		return false
	}

	select {
	case val, ok := <-gotten:
		if !ok {
			verificationFailure(t, "channel receives value", expected, "closed channel", infos...)
			return false
		}
		if val != expected {
			verificationFailure(t, "channel receives value", expected, val, infos...)
			return false
		}
		return true
	case <-time.After(timeout):
		verificationFailure(t, "channel receives value", expected, "timeout exceeded", infos...)
		return false
	}
}
