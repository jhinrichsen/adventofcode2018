package adventofcode2018

import "testing"

// testDayPart is a generic test helper for standard day part tests.
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

// benchDayPart is a generic benchmark helper for standard day part benchmarks.
func benchDayPart[P any, R comparable](
	b *testing.B,
	day uint8,
	part1 bool,
	parser func([]string) (P, error),
	solver func(P, bool) R,
) {
	b.Helper()
	lines := linesFromFilename(b, filename(day))
	b.ResetTimer()
	for b.Loop() {
		puzzle, err := parser(lines)
		if err != nil {
			b.Fatal(err)
		}
		_ = solver(puzzle, part1)
	}
}
