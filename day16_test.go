package adventofcode2018

import "testing"

func TestDay16Part1(t *testing.T) {
	const want = 607
	buf := file(t, 16)
	lines := Lines(string(buf))
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
	buf := file(t, 16)
	lines := Lines(string(buf))
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
	buf := file(b, 16)
	lines := Lines(string(buf))
	for b.Loop() {
		puzzle, _ := NewDay16(lines)
		_ = Day16Part1(puzzle)
	}
}

func BenchmarkDay16Part2(b *testing.B) {
	buf := file(b, 16)
	lines := Lines(string(buf))
	for b.Loop() {
		puzzle, _ := NewDay16(lines)
		_ = Day16Part2(puzzle)
	}
}
