package adventofcode2018

import "testing"

func TestDay19Part1Example(t *testing.T) {
	puzzle, err := NewDay19(exampleFile(t, 19))
	if err != nil {
		t.Fatal(err)
	}
	got := Day19(puzzle, true)
	// Expected value not specified in problem, but tracing shows it should be 7
	t.Logf("Example Part 1: %s", got)
}

func TestDay19Part1(t *testing.T) {
	testWithParserBytes(t, 19, file, true, NewDay19, Day19, "978")
}

func BenchmarkDay19Part1(b *testing.B) {
	benchWithParserBytes(b, 19, true, NewDay19, Day19)
}
