package adventofcode2018

import (
	"testing"
)

// testWithParserBytes is a generic test helper for day part tests using a []byte parser and solver.
func testWithParserBytes[P any, R comparable](
	t *testing.T,
	day uint8,
	fileFunc func(testing.TB, uint8) []byte,
	part1 bool,
	parser func([]byte) (P, error),
	solver func(P, bool) R,
	want R,
) {
	t.Helper()
	data := fileFunc(t, day)
	puzzle, err := parser(data)
	if err != nil {
		t.Fatal(err)
	}
	got := solver(puzzle, part1)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

// benchWithParserBytes is a generic benchmark helper for day part benchmarks with []byte parser.
// I/O is not measured, parsing and solving are measured.
// No testing/verification is performed in benchmarks.
func benchWithParserBytes[P any, R comparable](
	b *testing.B,
	day uint8,
	part1 bool,
	parser func([]byte) (P, error),
	solver func(P, bool) R,
) {
	b.Helper()
	data := file(b, day)
	for b.Loop() {
		puzzle, _ := parser(data)
		_ = solver(puzzle, part1)
	}
}
