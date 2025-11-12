package adventofcode2018

import "testing"

func TestDay19Part1Example(t *testing.T) {
	puzzle, err := NewDay19(exampleFile(t, 19))
	if err != nil {
		t.Fatal(err)
	}
	got := Day19(puzzle, true)
	// Expected value not specified in problem, but tracing shows it should be 6
	t.Logf("Example Part 1: %d", got)
}

func TestDay19Part1(t *testing.T) {
	testWithParserBytes(t, 19, file, true, NewDay19, Day19, 978)
}

func TestDay19Part2(t *testing.T) {
	testWithParserBytes(t, 19, file, false, NewDay19, Day19, 10996992)
}

func BenchmarkDay19Part1(b *testing.B) {
	benchWithParserBytes(b, 19, true, NewDay19, Day19)
}

func BenchmarkDay19Part2(b *testing.B) {
	benchWithParserBytes(b, 19, false, NewDay19, Day19)
}
