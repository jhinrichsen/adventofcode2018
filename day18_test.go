package adventofcode2018

import "testing"

func TestDay18Part1Example(t *testing.T) {
	testWithParserBytes(t, 18, exampleFile, true, NewDay18, Day18, "1147")
}

func TestDay18Part1(t *testing.T) {
	testWithParserBytes(t, 18, file, true, NewDay18, Day18, "637550")
}

func BenchmarkDay18Part1(b *testing.B) {
	benchWithParserBytes(b, 18, true, NewDay18, Day18)
}
