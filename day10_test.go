package adventofcode2018

import "testing"

func TestDay10Part1Example(t *testing.T) {
	testWithParserBytes(t, 10, exampleFile, true, NewDay10, Day10, "HI")
}

func TestDay10Part1(t *testing.T) {
	testWithParserBytes(t, 10, file, true, NewDay10, Day10, "RPNNXFZR")
}

func TestDay10Part2Example(t *testing.T) {
	testWithParserBytes(t, 10, exampleFile, false, NewDay10, Day10, "3")
}

func TestDay10Part2(t *testing.T) {
	testWithParserBytes(t, 10, file, false, NewDay10, Day10, "10946")
}

func BenchmarkDay10Part1(b *testing.B) {
	benchWithParserBytes(b, 10, true, NewDay10, Day10)
}

func BenchmarkDay10Part2(b *testing.B) {
	benchWithParserBytes(b, 10, false, NewDay10, Day10)
}
