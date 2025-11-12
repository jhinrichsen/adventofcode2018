package adventofcode2018

import "testing"

func TestDay22Part1Example(t *testing.T) {
	testWithParserBytes(t, 22, exampleFile, true, NewDay22, Day22, "114")
}

func TestDay22Part1(t *testing.T) {
	testWithParserBytes(t, 22, file, true, NewDay22, Day22, "11810")
}

func BenchmarkDay22Part1(b *testing.B) {
	benchWithParserBytes(b, 22, true, NewDay22, Day22)
}
