package adventofcode2018

import "testing"

func TestDay07Part1Example(t *testing.T) {
	const want = "CABDFE"
	lines := linesFromFilename(t, exampleFilename(7))
	got, err := Day07(lines)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func TestDay07Part1(t *testing.T) {
	const want = "FMOXCDGJRAUIHKNYZTESWLPBQV"
	lines := linesFromFilename(t, filename(7))
	got, err := Day07(lines)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func TestDay07Part2Example(t *testing.T) {
	const want = 15
	// Example with 2 workers and base 0 should take 15 seconds
	got, err := Day07Part2(linesFromFilename(t, exampleFilename(7)), 2, 0)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay07Part2(t *testing.T) {
	const want = 1053
	got, err := Day07Part2(linesFromFilename(t, filename(7)), 5, 60)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Fatalf("want %d but got %d", want, got)
	}
}
