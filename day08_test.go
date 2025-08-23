package adventofcode2018

import "testing"

func TestDay08Part1Example(t *testing.T) {
	const want = 138
	buf := exampleFile(t, 8)
	numbers, err := NewDay08(buf)
	if err != nil {
		t.Fatal(err)
	}
	got := Day08Part1(numbers)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay08Part1(t *testing.T) {
	const want = 47464
	buf := file(t, 8)
	numbers, err := NewDay08(buf)
	if err != nil {
		t.Fatal(err)
	}
	got := Day08Part1(numbers)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}

}

func TestDay08Part2Example(t *testing.T) {
	const want = 66
	buf := exampleFile(t, 8)
	numbers, err := NewDay08(buf)
	if err != nil {
		t.Fatal(err)
	}
	got := Day08Part2(numbers)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay08Part2(t *testing.T) {
	const want = 23054
	buf := file(t, 8)
	numbers, err := NewDay08(buf)
	if err != nil {
		t.Fatal(err)
	}
	got := Day08Part2(numbers)
	if got != want {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay08Part1(b *testing.B) {
	buf := file(b, 8)
	b.ResetTimer()
	for b.Loop() {
		numbers, err := NewDay08(buf)
		if err != nil {
			b.Fatal(err)
		}
		_ = Day08Part1(numbers)
	}
}

func BenchmarkDay08Part2(b *testing.B) {
	buf := file(b, 8)
	b.ResetTimer()
	for b.Loop() {
		numbers, err := NewDay08(buf)
		if err != nil {
			b.Fatal(err)
		}
		_ = Day08Part2(numbers)
	}
}
