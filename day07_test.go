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
