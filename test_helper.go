package adventofcode2018

import "testing"

// testDayPart is a generic test helper for standard day part tests with []string parser.
func testDayPart[P any, R comparable](
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

// testDayPartBytes is a generic test helper for day part tests with []byte parser.
func testDayPartBytes[P any, R comparable](
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

// benchDayPart is a generic benchmark helper for standard day part benchmarks with []string parser.
// I/O is not measured, parsing and solving are measured.
// No testing/verification is performed in benchmarks.
func benchDayPart[P any, R comparable](
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

// benchDayPartBytes is a generic benchmark helper for day part benchmarks with []byte parser.
// I/O is not measured, parsing and solving are measured.
// No testing/verification is performed in benchmarks.
func benchDayPartBytes[P any, R comparable](
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
