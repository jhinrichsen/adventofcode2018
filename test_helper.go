package adventofcode2018

import (
	"os"
	"testing"
)

// testWithParser is a generic test helper for day part tests using a parser and solver.
func testWithParser[P any, R comparable](
	t *testing.T,
	day uint8,
	filenameFunc func(uint8) string,
	part1 bool,
	parser func([]string) (P, error),
	solver func(P, bool) R,
	want R,
) {
	t.Helper()
	lines := linesFromFilename(t, filenameFunc(day))
	puzzle, err := parser(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := solver(puzzle, part1)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

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

// testSolver is a generic test helper for day part tests that work directly with []byte.
func testSolver[R comparable](
	t *testing.T,
	day uint8,
	filenameFunc func(uint8) string,
	part1 bool,
	solver func([]byte, bool) (R, error),
	want R,
) {
	t.Helper()
	buf := fileFromFilename(t, filenameFunc, day)
	got, err := solver(buf, part1)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

// fileFromFilename reads file bytes using a filename function and day number.
func fileFromFilename(tb testing.TB, filenameFunc func(uint8) string, day uint8) []byte {
	tb.Helper()
	buf, err := os.ReadFile(filenameFunc(day))
	if err != nil {
		tb.Fatal(err)
	}
	return buf
}

// benchWithParser is a generic benchmark helper for day part benchmarks using a parser.
// I/O is not measured, parsing and solving are measured.
// No testing/verification is performed in benchmarks.
func benchWithParser[P any, R comparable](
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

// benchSolver is a generic benchmark helper for solvers that work directly with []byte.
// I/O is not measured, solving is measured.
// No testing/verification is performed in benchmarks.
func benchSolver[R comparable](
	b *testing.B,
	day uint8,
	part1 bool,
	solver func([]byte, bool) (R, error),
) {
	b.Helper()
	data := file(b, day)
	for b.Loop() {
		_, _ = solver(data, part1)
	}
}

