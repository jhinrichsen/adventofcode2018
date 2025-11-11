package adventofcode2018

import "testing"

func TestDay08Part1Example(t *testing.T) {
	testWithParserBytes(t, 8, exampleFile, true, NewDay08, Day08, 138)
}

func TestDay08Part1(t *testing.T) {
	testWithParserBytes(t, 8, file, true, NewDay08, Day08, 47464)
}

func TestDay08Part2Example(t *testing.T) {
	testWithParserBytes(t, 8, exampleFile, false, NewDay08, Day08, 66)
}

func TestDay08Part2(t *testing.T) {
	testWithParserBytes(t, 8, file, false, NewDay08, Day08, 23054)
}

func BenchmarkDay08Part1(b *testing.B) {
	benchWithParserBytes(b, 8, true, NewDay08, Day08)
}

func BenchmarkDay08Part2(b *testing.B) {
	benchWithParserBytes(b, 8, false, NewDay08, Day08)
}

