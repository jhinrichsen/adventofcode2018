package adventofcode2018

import "testing"

func TestDay16Part1(t *testing.T) {
	buf := file(t, 16)
	data, err := NewDay16(buf)
	if err != nil {
		t.Fatal(err)
	}
	got := Day16Part1(data)
	t.Logf("Day 16 Part 1: %d", got)
}

func TestDay16Part2(t *testing.T) {
	buf := file(t, 16)
	data, err := NewDay16(buf)
	if err != nil {
		t.Fatal(err)
	}
	got := Day16Part2(data)
	t.Logf("Day 16 Part 2: %d", got)
}

func BenchmarkDay16Part1(b *testing.B) {
	buf := file(b, 16)
	b.ResetTimer()
	for b.Loop() {
		data, err := NewDay16(buf)
		if err != nil {
			b.Fatal(err)
		}
		_ = Day16Part1(data)
	}
}

func BenchmarkDay16Part2(b *testing.B) {
	buf := file(b, 16)
	b.ResetTimer()
	for b.Loop() {
		data, err := NewDay16(buf)
		if err != nil {
			b.Fatal(err)
		}
		_ = Day16Part2(data)
	}
}
