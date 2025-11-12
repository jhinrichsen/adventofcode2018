package adventofcode2018

import "testing"

func TestDay17Part1Example(t *testing.T) {
	testWithParserBytes(t, 17, exampleFile, true, NewDay17, Day17, uint(57))
}

func TestDay17Part1(t *testing.T) {
	testWithParserBytes(t, 17, file, true, NewDay17, Day17, uint(37858))
}

func TestDay17Part2Example(t *testing.T) {
	testWithParserBytes(t, 17, exampleFile, false, NewDay17, Day17, uint(29))
}

func TestDay17Part2(t *testing.T) {
	testWithParserBytes(t, 17, file, false, NewDay17, Day17, uint(30410))
}

func BenchmarkDay17Part1(b *testing.B) {
	benchWithParserBytes(b, 17, true, NewDay17, Day17)
}

func BenchmarkDay17Part2(b *testing.B) {
	benchWithParserBytes(b, 17, false, NewDay17, Day17)
}
