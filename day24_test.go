package adventofcode2018

import "testing"

func TestDay24Part1Example(t *testing.T) {
	puzzle, err := NewDay24(exampleFile(t, 24))
	if err != nil {
		t.Fatal(err)
	}
	got := Day24(puzzle, true)
	const want = "5216"
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestDay24Part1(t *testing.T) {
	testWithParserBytes(t, 24, file, true, NewDay24, Day24, "16530")
}

func BenchmarkDay24Part1(b *testing.B) {
	benchWithParserBytes(b, 24, true, NewDay24, Day24)
}
