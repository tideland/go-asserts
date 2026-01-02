// Tideland Go Asserts - Verify - Channel Tests
//
// Copyright (C) 2024-2025 Frank Mueller / Tideland / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package verify_test

import (
	"testing"
	"time"

	"tideland.dev/go/asserts/verify"
)

// TestChannelClosed tests the ChannelClosed verification function.
func TestChannelClosed(t *testing.T) {
	// Create a closed channel
	closedCh := make(chan any)
	close(closedCh)

	// Positive: channel is closed
	verify.ChannelClosed(t, closedCh)

	// Create continuation testing for negative cases
	ct := verify.ContinuedTesting(t)

	// Negative: channel is open with no value
	openCh := make(chan any)
	verify.ChannelClosed(ct, openCh)

	// Negative: channel is open with a value
	bufferedCh := make(chan any, 1)
	bufferedCh <- "value"
	verify.ChannelClosed(ct, bufferedCh)

	// Negative: nil channel
	var nilCh chan any
	verify.ChannelClosed(ct, nilCh)

	verify.FailureCount(ct, 3)
}

// TestChannelReceives tests the ChannelReceives verification function.
func TestChannelReceives(t *testing.T) {
	// Positive: channel receives value
	ch := make(chan int, 1)
	ch <- 42
	verify.ChannelReceives(t, ch, 100*time.Millisecond)

	// Positive: channel receives value from goroutine
	ch2 := make(chan string)
	go func() {
		time.Sleep(10 * time.Millisecond)
		ch2 <- "hello"
	}()
	verify.ChannelReceives(t, ch2, 100*time.Millisecond)

	// Create continuation testing for negative cases
	ct := verify.ContinuedTesting(t)

	// Negative: timeout (no value sent)
	ch3 := make(chan int)
	verify.ChannelReceives(ct, ch3, 10*time.Millisecond)

	// Negative: closed channel
	closedCh := make(chan int)
	close(closedCh)
	verify.ChannelReceives(ct, closedCh, 100*time.Millisecond)

	// Negative: nil channel
	var nilCh chan int
	verify.ChannelReceives(ct, nilCh, 100*time.Millisecond)

	verify.FailureCount(ct, 3)
}

// TestChannelReceivesValue tests the ChannelReceivesValue verification function.
func TestChannelReceivesValue(t *testing.T) {
	// Positive: channel receives expected value
	ch := make(chan int, 1)
	ch <- 42
	verify.ChannelReceivesValue(t, ch, 42, 100*time.Millisecond)

	// Positive: channel receives value from goroutine
	ch2 := make(chan string)
	go func() {
		time.Sleep(10 * time.Millisecond)
		ch2 <- "expected"
	}()
	verify.ChannelReceivesValue(t, ch2, "expected", 100*time.Millisecond)

	// Create continuation testing for negative cases
	ct := verify.ContinuedTesting(t)

	// Negative: wrong value received
	ch3 := make(chan int, 1)
	ch3 <- 99
	verify.ChannelReceivesValue(ct, ch3, 42, 100*time.Millisecond)

	// Negative: timeout
	ch4 := make(chan int)
	verify.ChannelReceivesValue(ct, ch4, 42, 10*time.Millisecond)

	// Negative: closed channel
	closedCh := make(chan int)
	close(closedCh)
	verify.ChannelReceivesValue(ct, closedCh, 42, 100*time.Millisecond)

	// Negative: nil channel
	var nilCh chan int
	verify.ChannelReceivesValue(ct, nilCh, 42, 100*time.Millisecond)

	verify.FailureCount(ct, 4)
}
