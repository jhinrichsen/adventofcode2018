package adventofcode2018

import "testing"

func TestDay07Part1Example(t *testing.T) {
	const want = "CABDFE"
	lines, err := linesFromFilename(exampleFilename(7))
	if err != nil {
		t.Fatal(err)
	}
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
	lines, err := linesFromFilename(filename(7))
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day07(lines)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}
