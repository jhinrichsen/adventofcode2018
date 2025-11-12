package adventofcode2018

import "testing"

func TestDay22Part1Example(t *testing.T) {
	testWithParserBytes(t, 22, exampleFile, true, NewDay22, Day22, "114")
}

func TestDay22Part1(t *testing.T) {
	testWithParserBytes(t, 22, file, true, NewDay22, Day22, "11810")
}

func TestDay22Part2Example(t *testing.T) {
	testWithParserBytes(t, 22, exampleFile, false, NewDay22, Day22, "45")
}

func TestDay22Part2(t *testing.T) {
	testWithParserBytes(t, 22, file, false, NewDay22, Day22, "1015")
}

func BenchmarkDay22Part1(b *testing.B) {
	benchWithParserBytes(b, 22, true, NewDay22, Day22)
}

func BenchmarkDay22Part2(b *testing.B) {
	benchWithParserBytes(b, 22, false, NewDay22, Day22)
}
