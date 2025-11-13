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

// testWithParserLines is a generic test helper for day part tests using a []string parser and solver.
func testWithParserLines[P any, R comparable](
	t *testing.T,
	day uint8,
	part1 bool,
	parser func([]string) (P, error),
	solver func(P, bool) R,
	want R,
) {
	t.Helper()
	lines := linesFromFilename(t, filename(day))
	puzzle, err := parser(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := solver(puzzle, part1)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

// benchWithParserLines is a generic benchmark helper for day part benchmarks with []string parser.
// I/O is not measured, parsing and solving are measured.
// No testing/verification is performed in benchmarks.
func benchWithParserLines[P any, R comparable](
	b *testing.B,
	day uint8,
	part1 bool,
	parser func([]string) (P, error),
	solver func(P, bool) R,
) {
	b.Helper()
	lines := linesFromFilename(b, filename(day))
	for b.Loop() {
		puzzle, _ := parser(lines)
		_ = solver(puzzle, part1)
	}
}

// testWithSolverBytes is a generic test helper for solvers that take []byte directly (no parser).
func testWithSolverBytes[R comparable](
	t *testing.T,
	day uint8,
	part1 bool,
	solver func([]byte, bool) R,
	want R,
) {
	t.Helper()
	data := file(t, day)
	got := solver(data, part1)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

// benchWithSolverBytes is a generic benchmark helper for solvers that take []byte directly (no parser).
// I/O is not measured, only solving is measured.
func benchWithSolverBytes[R comparable](
	b *testing.B,
	day uint8,
	part1 bool,
	solver func([]byte, bool) R,
) {
	b.Helper()
	data := file(b, day)
	for b.Loop() {
		_ = solver(data, part1)
	}
}
