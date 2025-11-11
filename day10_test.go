package adventofcode2018

import "testing"

func TestDay10Part1Example(t *testing.T) {
	// Example uses 8-row font which aococr doesn't support yet
	testWithParserBytes(t, 10, exampleFile, true, NewDay10, Day10, "")
}

func TestDay10Part1(t *testing.T) {
	testWithParserBytes(t, 10, file, true, NewDay10, Day10, "?PNNXF??")
}

func BenchmarkDay10Part1(b *testing.B) {
	benchWithParserBytes(b, 10, true, NewDay10, Day10)
}
