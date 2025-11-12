package adventofcode2018

import "testing"

func TestDay17Part1Example(t *testing.T) {
	testWithParserBytes(t, 17, exampleFile, true, NewDay17, Day17, "57")
}

func TestDay17Part1(t *testing.T) {
	testWithParserBytes(t, 17, file, true, NewDay17, Day17, "37858")
}

func BenchmarkDay17Part1(b *testing.B) {
	benchWithParserBytes(b, 17, true, NewDay17, Day17)
}
