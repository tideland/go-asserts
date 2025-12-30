# Tideland Go Asserts

[![GitHub release](https://img.shields.io/github/release/tideland/go-asserts.svg)](https://github.com/tideland/go-asserts)
[![GitHub license](https://img.shields.io/badge/license-New%20BSD-blue.svg)](https://raw.githubusercontent.com/tideland/go-asserts/main/LICENSE)
[![Go Module](https://img.shields.io/github/go-mod/go-version/tideland/go-asserts)](https://github.com/tideland/go-asserts/blob/main/go.mod)
[![GoDoc](https://godoc.org/tideland.dev/go/asserts?status.svg)](https://pkg.go.dev/tideland.dev/go/asserts)

## Description

**Tideland Go Asserts** provides comprehensive testing utilities for Go, consisting of three specialized packages that work together to make testing easier, more readable, and more maintainable.

## Packages

### verify - Human-Readable Assertions

The `verify` package provides a rich set of assertion functions for Go's standard testing package. It supports continued testing mode where assertions report failures without halting execution, allowing multiple checks in a single test.

**Key Features:**
- Type assertions: `Nil`, `NotNil`, `Equal`, `NotEqual`, etc.
- Comparison assertions: `Less`, `Greater`, `InRange`, `About`, etc.
- String assertions: `Contains`, `HasPrefix`, `Match`, etc.
- Slice and map assertions: `Length`, `Empty`, `NotEmpty`, etc.
- Error assertions: `NoError`, `ErrorContains`, `ErrorMatches`, etc.
- Channel assertions: `Readable`, `NotReadable`, `Closed`, etc.
- Predicate assertions: `True`, `False`, `Predicate`, etc.
- Continued testing mode with failure count tracking

**Example:**

```go
import "tideland.dev/go/asserts/verify"

func TestMyFunction(t *testing.T) {
    ct := verify.ContinuedTesting(t)

    result, err := myFunction("input")
    verify.NoError(ct, err)
    verify.Length(ct, result, 7)
    verify.Match(ct, result, "^success:")

    verify.FailureCount(ct, 0) // Verify no failures occurred
}
```

### capture - Output Capture for Testing

The `capture` package provides utilities for capturing stdout and stderr during test execution, enabling tests for functions that write to standard streams.

**Key Features:**
- Capture stdout only
- Capture stderr only
- Capture both stdout and stderr simultaneously
- Panic-safe restoration of streams
- Simple API with byte and string access

**Example:**

```go
import "tideland.dev/go/asserts/capture"

func TestPrintFunction(t *testing.T) {
    captured := capture.Stdout(func() {
        fmt.Println("Hello, World!")
    })

    verify.Equal(t, "Hello, World!\n", captured.String())
    verify.Equal(t, 14, captured.Len())
}
```

### generators - Random Test Data Generation

The `generators` package provides utilities for generating random test data in a controlled and reproducible manner. All generation is based on a `rand.Rand` instance, allowing deterministic test data when using a fixed seed.

**Key Features:**
- Basic types: bytes, ints, durations, times, UUIDs
- Selection helpers: `OneOf`, `FlipCoin`, etc.
- Text generation: words, sentences, paragraphs
- Pattern-based generation with escape sequences
- Identity generation: names, emails, URLs, domains
- Concurrent-safe with internal mutex
- Fixed and simple random generators

**Example:**

```go
import "tideland.dev/go/asserts/generators"

func TestWithRandomData(t *testing.T) {
    gen := generators.New(generators.FixedRand())

    email := gen.EMail()
    verify.Match(t, email, `^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]*$`)

    // Pattern-based generation: (XXX) XXX-XXXX
    phone := gen.Pattern("(^1^0^0) ^1^0^0-^0^0^0^0")
    verify.Match(t, phone, `^\(\d{3}\) \d{3}-\d{4}$`)
}
```

## Installation

```bash
go get tideland.dev/go/asserts
```

## Usage

Import the packages you need:

```go
import (
    "tideland.dev/go/asserts/verify"
    "tideland.dev/go/asserts/capture"
    "tideland.dev/go/asserts/generators"
)
```

The packages work well together:

```go
func TestComplexScenario(t *testing.T) {
    gen := generators.New(generators.FixedRand())
    ct := verify.ContinuedTesting(t)

    // Generate test data
    testName := gen.Word()
    testEmail := gen.EMail()

    // Capture output
    output := capture.Stdout(func() {
        processUser(testName, testEmail)
    })

    // Verify results
    verify.Contains(ct, output.String(), testName)
    verify.Contains(ct, output.String(), testEmail)
    verify.FailureCount(ct, 0)
}
```

## Development

Each package includes a Makefile for common development tasks:

```bash
# Run all checks for a specific package
cd verify && make all

# Run tests with race detection
cd capture && make test

# Generate coverage report
cd generators && make coverage

# Or use the root Makefile to target specific packages
make test --package=verify
make all --package=capture
make lint --package=generators

# Run for all packages
make all
```

## Contributors

- Frank Mueller (https://github.com/themue / https://github.com/tideland / https://tideland.dev)
