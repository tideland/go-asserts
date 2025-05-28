// -----------------------------------------------------------------------------
// Convenient verification of unit tests in Go libraries and applications.
//
// A set of individual verifications
//
// Copyright (C) 2024-2025 Frank Mueller / Oldenburg / Germany / Earth
// -----------------------------------------------------------------------------

package generators_test

//--------------------
// IMPORTS
//--------------------

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"tideland.dev/go/asserts/verify"

	"tideland.dev/go/asserts/generators"
)

//--------------------
// TESTS
//--------------------

// TestBuildDate tests the generation of dates.
func TestBuildDate(t *testing.T) {
	layouts := []string{
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
	}

	for _, layout := range layouts {
		ts, bt := generators.BuildTime(layout, 0)
		tsp, err := time.Parse(layout, ts)
		verify.Nil(t, err)
		verify.Equal(t, bt, tsp)

		ts, bt = generators.BuildTime(layout, -30*time.Minute)
		tsp, err = time.Parse(layout, ts)
		verify.Nil(t, err)
		verify.Equal(t, bt, tsp)

		ts, bt = generators.BuildTime(layout, time.Hour)
		tsp, err = time.Parse(layout, ts)
		verify.Nil(t, err)
		verify.Equal(t, bt, tsp)
	}
}

// TestBytes tests the generation of bytes.
func TestBytes(t *testing.T) {
	gen := generators.New(generators.FixedRand())

	// Test individual bytes.
	for range 10000 {
		lo := gen.Byte(0, 255)
		hi := gen.Byte(0, 255)
		n := gen.Byte(lo, hi)
		if hi < lo {
			lo, hi = hi, lo
		}
		verify.True(t, lo <= n && n <= hi)
	}

	// Test byte slices.
	ns := gen.Bytes(1, 200, 1000)
	verify.Length(t, ns, 1000)
	for _, n := range ns {
		verify.True(t, n >= 1 && n <= 200)
	}

	// Test UUIDs.
	for range 10000 {
		uuid := gen.UUID()
		verify.Length(t, uuid, 16)
	}
}

// TestInts tests the generation of ints.
func TestInts(t *testing.T) {
	gen := generators.New(generators.FixedRand())

	// Test individual ints.
	for range 10000 {
		lo := gen.Int(-100, 100)
		hi := gen.Int(-100, 100)
		n := gen.Int(lo, hi)
		if hi < lo {
			lo, hi = hi, lo
		}
		verify.True(t, lo <= n && n <= hi)
	}

	// Test int slices.
	ns := gen.Ints(0, 500, 10000)
	verify.Length(t, ns, 10000)
	for _, n := range ns {
		verify.True(t, n >= 0 && n <= 500)
	}

	// Test the generation of percent.
	for range 10000 {
		p := gen.Percent()
		verify.InRange(t, p, 0, 100)
	}

	// Test the flipping of coins.
	ct := 0
	cf := 0
	for range 10000 {
		c := gen.FlipCoin(50)
		if c {
			ct++
		} else {
			cf++
		}
	}
	verify.About(t, ct, 5000, 500)
	verify.About(t, cf, 5000, 500)
}

// TestOneOf tests the generation of selections.
func TestOneOf(t *testing.T) {
	gen := generators.New(generators.FixedRand())
	stuff := []any{1, true, "three", 47.11}

	for range 10000 {
		m := gen.OneOf(stuff...)
		verify.Contains(t, m, stuff)

		b := gen.OneByteOf(1, 2, 3, 4, 5)
		verify.InRange(t, b, 1, 5)

		r := gen.OneRuneOf("abcdef")
		verify.True(t, r >= 'a' && r <= 'f')

		n := gen.OneIntOf(1, 2, 3, 4, 5)
		verify.InRange(t, n, 1, 5)

		vs := []string{"one", "two", "three", "four", "five"}
		s := gen.OneStringOf(vs...)
		verify.Contains(t, s, vs)

		d := gen.OneDurationOf(1*time.Second, 2*time.Second, 3*time.Second)
		verify.InRange(t, d, 1*time.Second, 3*time.Second)
	}
}

// TestWords tests the generation of words.
func TestWords(t *testing.T) {
	gen := generators.New(generators.FixedRand())

	// Test single words.
	for range 10000 {
		w := gen.Word()
		for _, r := range w {
			verify.True(t, r >= 'a' && r <= 'z')
		}
	}

	// Test limited words.
	for range 10000 {
		lo := gen.Int(generators.MinWordLen, generators.MaxWordLen)
		hi := gen.Int(generators.MinWordLen, generators.MaxWordLen)
		w := gen.LimitedWord(lo, hi)
		wl := len(w)
		if hi < lo {
			lo, hi = hi, lo
		}
		verify.True(t, lo <= wl && wl <= hi, info("WL %d LO %d HI %d", wl, lo, hi))
	}
}

// TestPattern tests the generation based on patterns.
func TestPattern(t *testing.T) {
	gen := generators.New(generators.FixedRand())
	assertPattern := func(pattern, runes string) {
		set := make(map[rune]bool)
		for _, r := range runes {
			set[r] = true
		}
		for range 10 {
			result := gen.Pattern(pattern)
			for _, r := range result {
				verify.True(t, set[r], pattern, result, runes)
			}
		}
	}

	assertPattern("^^", "^")
	assertPattern("^0^0^0^0^0", "0123456789")
	assertPattern("^1^1^1^1^1", "123456789")
	assertPattern("^o^o^o^o^o", "01234567")
	assertPattern("^h^h^h^h^h", "0123456789abcdef")
	assertPattern("^H^H^H^H^H", "0123456789ABCDEF")
	assertPattern("^a^a^a^a^a", "abcdefghijklmnopqrstuvwxyz")
	assertPattern("^A^A^A^A^A", "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	assertPattern("^c^c^c^c^c", "bcdfghjklmnpqrstvwxyz")
	assertPattern("^C^C^C^C^C", "BCDFGHJKLMNPQRSTVWXYZ")
	assertPattern("^v^v^v^v^v", "aeiou")
	assertPattern("^V^V^V^V^V", "AEIOU")
	assertPattern("^z^z^z^z^z", "abcdefghijklmnopqrstuvwxyz0123456789")
	assertPattern("^Z^Z^Z^Z^Z", "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	assertPattern("^1^0.^0^0^0,^0^0 €", "0123456789 .,€")
}

// TestText tests the generation of text.
func TestText(t *testing.T) {
	gen := generators.New(generators.FixedRand())
	names := gen.Names(4)

	for range 10000 {
		s := gen.Sentence()
		ws := strings.Split(s, " ")
		lws := len(ws)
		verify.True(t, 2 <= lws && lws <= 15, info("S: %v SL: %d", s, lws))
		verify.True(t, 'A' <= s[0] && s[0] <= 'Z', info("SUC: %v", s[0]))
	}

	for range 10 {
		s := gen.SentenceWithNames(names)
		verify.NotEmpty(t, s)
	}

	for range 10000 {
		p := gen.Paragraph()
		ss := strings.Split(p, ". ")
		lss := len(ss)
		verify.True(t, 2 <= lss && lss <= 10, info("PL: %d", lss))
		for _, s := range ss {
			ws := strings.Split(s, " ")
			lws := len(ws)
			verify.True(t, 2 <= lws && lws <= 15, info("S: %v PSL: %d", s, lws))
			verify.True(t, 'A' <= s[0] && s[0] <= 'Z', info("PSUC: %v", s[0]))
		}
	}

	for range 10 {
		s := gen.ParagraphWithNames(names)
		verify.NotEmpty(t, s)
	}
}

// TestName tests the generation of names.
func TestName(t *testing.T) {
	gen := generators.New(generators.FixedRand())

	verify.Equal(t, generators.ToUpperFirst("yadda"), "Yadda")

	for range 10000 {
		first, middle, last := gen.Name()

		verify.Match(t, first, `[A-Z][a-z]+(-[A-Z][a-z]+)?`)
		verify.Match(t, middle, `[A-Z][a-z]+(-[A-Z][a-z]+)?`)
		verify.Match(t, last, `[A-Z]['a-zA-Z]+`)

		first, middle, last = gen.MaleName()

		verify.Match(t, first, `[A-Z][a-z]+(-[A-Z][a-z]+)?`)
		verify.Match(t, middle, `[A-Z][a-z]+(-[A-Z][a-z]+)?`)
		verify.Match(t, last, `[A-Z]['a-zA-Z]+`)

		first, middle, last = gen.FemaleName()

		verify.Match(t, first, `[A-Z][a-z]+(-[A-Z][a-z]+)?`)
		verify.Match(t, middle, `[A-Z][a-z]+(-[A-Z][a-z]+)?`)
		verify.Match(t, last, `[A-Z]['a-zA-Z]+`)

		count := gen.Int(0, 5)
		names := gen.Names(count)

		verify.Length(t, names, count)
		for _, name := range names {
			verify.Match(t, name, `[A-Z][a-z]+(-[A-Z][a-z]+)?\s([A-Z]\.\s)?[A-Z]['a-zA-Z]+`)
		}
	}
}

// TestDomain tests the generation of domains.
func TestDomain(t *testing.T) {
	gen := generators.New(generators.FixedRand())

	for range 100 {
		domain := gen.Domain()

		verify.Match(t, domain, `^[a-z0-9.-]+\.[a-z]*$`)
	}
}

// TestURL tests the generation of URLs.
func TestURL(t *testing.T) {
	gen := generators.New(generators.FixedRand())

	for range 10000 {
		url := gen.URL()

		verify.Match(t, url, `(http|ftp|https):\/\/[\w\-_]+(\.[\w\-_]+)+([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?`)
	}
}

// TestEMail tests the generation of e-mail addresses.
func TestEMail(t *testing.T) {
	gen := generators.New(generators.FixedRand())

	for range 10000 {
		addr := gen.EMail()

		verify.Match(t, addr, `^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]*$`)
	}
}

// TestTimes tests the generation of durations and times.
func TestTimes(t *testing.T) {
	gen := generators.New(generators.FixedRand())

	for range 10000 {
		// Test durations.
		lo := gen.Duration(time.Second, time.Minute)
		hi := gen.Duration(time.Second, time.Minute)
		d := gen.Duration(lo, hi)
		if hi < lo {
			lo, hi = hi, lo
		}
		verify.InRange(t, d, lo, hi)

		// Test times.
		loc := time.Local
		now := time.Now()
		dur := gen.Duration(24*time.Hour, 30*24*time.Hour)
		tim := gen.Time(loc, now, dur)
		verify.True(t, tim.Equal(now) || tim.After(now))
		verify.True(t, tim.Before(now.Add(dur)) || tim.Equal(now.Add(dur)))
	}

	sleeps := map[int]time.Duration{
		1: 1 * time.Millisecond,
		2: 2 * time.Millisecond,
		3: 3 * time.Millisecond,
		4: 4 * time.Millisecond,
		5: 5 * time.Millisecond,
	}
	for range 1000 {
		sleep := gen.SleepOneOf(sleeps[1], sleeps[2], sleeps[3], sleeps[4], sleeps[5])
		s := int(sleep) / 1000000
		_, ok := sleeps[s]
		verify.True(t, ok)
	}
}

// TestConcurrency simply produces a number of concurrent calls simply to let
// the race detection do its work.
func TestConcurrency(t *testing.T) {
	gen := generators.New(generators.FixedRand())

	run := func() {
		go gen.Byte(0, 255)
		go gen.Int(0, 9999)
		go gen.Duration(25*time.Millisecond, 75*time.Millisecond)
	}
	for range 100 {
		go run()
	}

	time.Sleep(3 * time.Second)
}

//--------------------
// HELPER
//--------------------

var info = fmt.Sprintf

// EOF
