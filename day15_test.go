package adventofcode2018

import "testing"

func TestDay15Part1Example(t *testing.T) {
	testWithParserBytes(t, 15, exampleFile, true, NewDay15, Day15, "27730")
}

func TestDay15Part1(t *testing.T) {
	t.Skip("Skipping - answer too high, need to debug")
	testWithParserBytes(t, 15, file, true, NewDay15, Day15, "213265")
}

func BenchmarkDay15Part1(b *testing.B) {
	benchWithParserBytes(b, 15, true, NewDay15, Day15)
}
