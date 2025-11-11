package adventofcode2018

import "testing"

// testDay is a generic test function that reads input, parses it, solves it, and checks the result.
// - day: the day number (e.g., 8)
// - example: true for example input, false for actual input
// - parser: function that takes []byte and returns (parsedData, error)
// - solver: function that takes parsedData and returns int
// - want: expected result
func testDay[T any](t *testing.T, day uint8, example bool, parser func([]byte) (T, error), solver func(T) int, want int) {
	t.Helper()
	var buf []byte
	if example {
		buf = exampleFile(t, day)
	} else {
		buf = file(t, day)
	}
	data, err := parser(buf)
	if err != nil {
		t.Fatal(err)
	}
	got := solver(data)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
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
