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

// benchDay is a generic benchmark function that reads input, parses it, and solves it.
// - day: the day number (e.g., 8)
// - parser: function that takes []byte and returns (parsedData, error)
// - solver: function that takes parsedData and returns int
func benchDay[T any](b *testing.B, day uint8, parser func([]byte) (T, error), solver func(T) int) {
	b.Helper()
	buf := file(b, day)
	b.ResetTimer()
	for b.Loop() {
		data, err := parser(buf)
		if err != nil {
			b.Fatal(err)
		}
		_ = solver(data)
	}
}
