package adventofcode2018

// The key words “MUST”, “MUST NOT”, “REQUIRED”, “SHALL”, “SHALL NOT”,
// “SHOULD”, “SHOULD NOT”, “RECOMMENDED”, “NOT RECOMMENDED”, “MAY”, and
// “OPTIONAL” in this document are to be interpreted as described in BCP
// 14 RFC2119 RFC8174 when, and only when, they appear in all capitals,
// as shown here.
//
// This file contains placeholders in the form {{placeholder}}.
//
// This file `day00_test.go` acts as a template for day NN to aid AI agents in implementing test cases for day NN.
// Any implementation for day NN MUST use this template and create a new file `day{{NN}}_test.go`.
// Test cases are not performance critical so implementations SHOULD follow best practices such as Clean Code.
// Implementations CAN introduce helper functions for don't-repeat-yourself (DRY).

import (
	"testing"
)

// TestDay00Part1Example tests all examples given in the puzzle description.
// Getting the solution to the puzzle right in the first shot is the number one priority in this project.
// Implementations MUST test all given examples 100%.
// Implementations MUST include an abbreviated, original quote from the website for each example ("[…], so nirvana is 42").
// Implementations SHOULD use table tests for multiple inputs.
// Implementations SHOULD create external files instead of hardcoded (multiline) strings to avoid parser errors for real puzzle inputs.
// Implementations CAN use the provided functions `example{{N}}Filename()` to read example input from `testdata/day00_example{{N}}.txt`).
// Implementations CAN skip parsing from external files if the solver (`Day00()`) uses []byte for maximum performance.
func TestDay00Part1Example(t *testing.T) {
	const want = 0
	lines := linesFromFilename(t, exampleFilename(0))
	p, err := NewPuzzle(lines)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day00(p, true)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay00Part1(t *testing.T) {
	const want = 0
	lines := linesFromFilename(t, filename(0))
	p, err := NewPuzzle(lines)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day00(p, false)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

// BenchmarkDay00Part1 is used for testing the performance of day N part 1.
// Benchmarks MUST NOT include measuring file I/O (reading the file from disk into memory).
// Benchmarks MUST include parsing the puzzle input.
func BenchmarkDay00Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(0))
	for b.Loop() {
		p, _ := NewPuzzle(lines)
		_, _ = Day00(p, true)
	}
}

func TestDay00Part2Example(t *testing.T) {
	const want = 0
	lines := linesFromFilename(t, exampleFilename(0))
	p, err := NewPuzzle(lines)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day00(p, false)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay00Part2(t *testing.T) {
	const want = 0
	lines := linesFromFilename(t, filename(0))
	p, err := NewPuzzle(lines)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day00(p, false)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay00Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(0))
	for b.Loop() {
		p, _ := NewPuzzle(lines)
		_, _ = Day00(p, false)
	}
}
