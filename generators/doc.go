// Tideland Go Asserts - Generators
//
// Copyright (C) 2024-2025 Frank Mueller / Tideland / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

/*
Package generators provides utilities for generating random test data in a
controlled and reproducible manner. All data generation is based on a
rand.Rand instance, allowing for deterministic test data when using a fixed
seed.

The package offers two convenience functions for creating random generators:

  - SimpleRand(): Creates a generator with a time-based seed (non-deterministic)
  - FixedRand(): Creates a generator with a fixed seed of 42 (deterministic)

The Generator type provides methods for generating various types of data:

Basic Types:

  - Byte, Bytes: Generate random bytes within a range
  - Int, Ints: Generate random integers within a range
  - Percent: Generate a random percentage (0-100)
  - Duration: Generate random time.Duration values
  - Time: Generate random time.Time values
  - UUID: Generate pseudo-UUIDs

Selection:

  - FlipCoin: Random boolean based on percentage
  - OneOf, OneByteOf, OneRuneOf, OneIntOf, OneStringOf, OneDurationOf: Select random elements

Text Generation:

  - Word, Words, LimitedWord: Generate random words
  - Pattern: Generate strings based on escape patterns
  - Sentence, SentenceWithNames: Generate random sentences
  - Paragraph, ParagraphWithNames: Generate random paragraphs

Identity Generation:

  - Name, MaleName, FemaleName, Names: Generate random names
  - Domain: Generate random domain names
  - URL: Generate random URLs
  - EMail: Generate random email addresses

Helper Functions:

  - ToUpperFirst: Capitalize the first rune of a string
  - BuildEMail: Construct an email address from name parts and domain
  - BuildTime: Generate formatted time strings with offsets
  - UUIDString: Convert a UUID byte array to string format

Example with fixed seed for reproducible tests:

	func TestWithFixedData(t *testing.T) {
		gen := generators.New(generators.FixedRand())

		// These will always generate the same values
		name := gen.Word()
		number := gen.Int(1, 100)
		email := gen.EMail()

		verify.NotEmpty(t, name)
		verify.InRange(t, number, 1, 100)
		verify.Match(t, email, `^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]*$`)
	}

Example with pattern-based generation:

	func TestPatternGeneration(t *testing.T) {
		gen := generators.New(generators.FixedRand())

		// Generate a phone number: (XXX) XXX-XXXX
		phone := gen.Pattern("(^1^0^0) ^1^0^0-^0^0^0^0")
		verify.Match(t, phone, `^\(\d{3}\) \d{3}-\d{4}$`)

		// Generate a product code: ABC-12345
		code := gen.Pattern("^A^A^A-^1^0^0^0^0")
		verify.Match(t, code, `^[A-Z]{3}-\d{5}$`)
	}

The Generator type is safe for concurrent use as it protects its internal
random number generator with a mutex.
*/
package generators
