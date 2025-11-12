package adventofcode2018

import "testing"

func TestDay16Part1(t *testing.T) {
	const want uint = 607
	lines := linesFromFilename(t, filename(16))
	puzzle, err := NewDay16(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := Day16Part1(puzzle)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay16Part2(t *testing.T) {
	const want = 577
	lines := linesFromFilename(t, filename(16))
	puzzle, err := NewDay16(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := Day16Part2(puzzle)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay16Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(16))
	for b.Loop() {
		puzzle, _ := NewDay16(lines)
		_ = Day16Part1(puzzle)
	}
}

func BenchmarkDay16Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(16))
	for b.Loop() {
		puzzle, _ := NewDay16(lines)
		_ = Day16Part2(puzzle)
	}
}
