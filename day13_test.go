package adventofcode2018

import "testing"

func TestDay13Part1Example(t *testing.T) {
	testWithParserBytes(t, 13, exampleFile, true, NewDay13, Day13, "7,3")
}

func TestDay13Part1(t *testing.T) {
	testWithParserBytes(t, 13, file, true, NewDay13, Day13, "53,133")
}

func BenchmarkDay13Part1(b *testing.B) {
	benchWithParserBytes(b, 13, true, NewDay13, Day13)
}
