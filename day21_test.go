package adventofcode2018

import "testing"

func TestDay21Part1(t *testing.T) {
	testWithParserBytes(t, 21, file, true, NewDay21, Day21, "7216956")
}

func TestDay21Part2(t *testing.T) {
	testWithParserBytes(t, 21, file, false, NewDay21, Day21, "14596916")
}

func BenchmarkDay21Part1(b *testing.B) {
	benchWithParserBytes(b, 21, true, NewDay21, Day21)
}

func BenchmarkDay21Part2(b *testing.B) {
	benchWithParserBytes(b, 21, false, NewDay21, Day21)
}
