package adventofcode2018

import "testing"

func TestDay10Part1Example(t *testing.T) {
	testWithParserBytes(t, 10, exampleFile, true, NewDay10, Day10, "HI")
}

func TestDay10Part1(t *testing.T) {
	testWithParserBytes(t, 10, file, true, NewDay10, Day10, "RPNNXFZR")
}

func BenchmarkDay10Part1(b *testing.B) {
	benchWithParserBytes(b, 10, true, NewDay10, Day10)
}
