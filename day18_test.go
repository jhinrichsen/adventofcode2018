package adventofcode2018

import "testing"

func TestDay18Part1Example(t *testing.T) {
	// "After 10 minutes, there are 37 wooded acres and 31 lumberyards.
	// Multiplying the number of wooded acres by the number of lumberyards gives
	// the total resource value after ten minutes: 37 * 31 = 1147."
	const want = 1147
	lines := linesFromFilename(t, exampleFilename(18))
	p, err := NewDay18(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := Day18Part1(p)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay18Part1(t *testing.T) {
	const want = 637550
	lines := linesFromFilename(t, filename(18))
	p, err := NewDay18(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := Day18Part1(p)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay18Part2(t *testing.T) {
	const want = 201465
	lines := linesFromFilename(t, filename(18))
	p, err := NewDay18(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := Day18Part2(p)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay18Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(18))
	for b.Loop() {
		p, _ := NewDay18(lines)
		_ = Day18Part1(p)
	}
}

func BenchmarkDay18Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(18))
	for b.Loop() {
		p, _ := NewDay18(lines)
		_ = Day18Part2(p)
	}
}
