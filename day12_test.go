package adventofcode2018

import "testing"

func TestDay12Part1Example(t *testing.T) {
	testWithParserBytes(t, 12, exampleFile, true, NewDay12, Day12, "325")
}

func TestDay12Part1(t *testing.T) {
	testWithParserBytes(t, 12, file, true, NewDay12, Day12, "1184")
}

func BenchmarkDay12Part1(b *testing.B) {
	benchWithParserBytes(b, 12, true, NewDay12, Day12)
}
