package adventofcode2018

import "testing"

func TestDay18Part1Example(t *testing.T) {
	// "After 10 minutes, there are 37 wooded acres and 31 lumberyards.
	// Multiplying the number of wooded acres by the number of lumberyards gives
	// the total resource value after ten minutes: 37 * 31 = 1147."
	const want = 1147
	buf := exampleFile(t, 18)
	grid, width, height, err := NewDay18(buf)
	if err != nil {
		t.Fatal(err)
	}
	got := Day18Part1(grid, width, height)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay18Part1(t *testing.T) {
	const want = 637550
	buf := file(t, 18)
	grid, width, height, err := NewDay18(buf)
	if err != nil {
		t.Fatal(err)
	}
	got := Day18Part1(grid, width, height)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay18Part2(t *testing.T) {
	const want = 201465
	buf := file(t, 18)
	grid, width, height, err := NewDay18(buf)
	if err != nil {
		t.Fatal(err)
	}
	got := Day18Part2(grid, width, height)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay18Part1(b *testing.B) {
	buf := file(b, 18)
	b.ResetTimer()
	for b.Loop() {
		grid, width, height, err := NewDay18(buf)
		if err != nil {
			b.Fatal(err)
		}
		_ = Day18Part1(grid, width, height)
	}
}

func BenchmarkDay18Part2(b *testing.B) {
	buf := file(b, 18)
	b.ResetTimer()
	for b.Loop() {
		grid, width, height, err := NewDay18(buf)
		if err != nil {
			b.Fatal(err)
		}
		_ = Day18Part2(grid, width, height)
	}
}
