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

func TestDay13Part2Example(t *testing.T) {
	testWithParserBytes(t, 13, example2File, false, NewDay13, Day13, "6,4")
}

func TestDay13Part2(t *testing.T) {
	testWithParserBytes(t, 13, file, false, NewDay13, Day13, "111,68")
}

func BenchmarkDay13Part2(b *testing.B) {
	benchWithParserBytes(b, 13, false, NewDay13, Day13)
}
