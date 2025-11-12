package adventofcode2018

import "testing"

func TestDay16Part1Example(t *testing.T) {
	input := `Before: [3, 2, 1, 1]
9 2 1 2
After:  [3, 2, 2, 1]
`
	puzzle, err := NewDay16([]byte(input))
	if err != nil {
		t.Fatal(err)
	}

	// This sample should match 3 opcodes (mulr, addi, seti)
	if len(puzzle.samples) != 1 {
		t.Fatalf("expected 1 sample, got %d", len(puzzle.samples))
	}

	matches := countMatches(puzzle.samples[0])
	if matches != 3 {
		t.Errorf("expected 3 matches, got %d", matches)
	}
}

func TestDay16Part1(t *testing.T) {
	testWithParserBytes(t, 16, file, true, NewDay16, Day16, "607")
}

func BenchmarkDay16Part1(b *testing.B) {
	benchWithParserBytes(b, 16, true, NewDay16, Day16)
}
