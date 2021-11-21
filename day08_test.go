package adventofcode2018

import "testing"

func TestDay08Part1Example(t *testing.T) {
	const want = 138
	lines, err := linesFromFilename(exampleFilename(8))
	if err != nil {
		t.Fatal(err)
	}
	numbers, err := parseDay08(lines[0])
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
	lines, err := linesFromFilename(filename(8))
	if err != nil {
		t.Fatal(err)
	}
	numbers, err := parseDay08(lines[0])
	if err != nil {
		t.Fatal(err)
	}
	got := Day08Part1(numbers)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}

}
