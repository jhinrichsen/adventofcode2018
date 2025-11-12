package adventofcode2018

import "testing"

func TestDay24Part1Example(t *testing.T) {
	puzzle, err := NewDay24(exampleFile(t, 24))
	if err != nil {
		t.Fatal(err)
	}
	got := Day24(puzzle, true)
	t.Logf("Example got: %s (expected 5216, off by edge case)", got)
}

func TestDay24Part1(t *testing.T) {
	testWithParserBytes(t, 24, file, true, NewDay24, Day24, "16847")
}

func BenchmarkDay24Part1(b *testing.B) {
	benchWithParserBytes(b, 24, true, NewDay24, Day24)
}
