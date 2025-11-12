package adventofcode2018

import "testing"

func TestDay23Part1Example(t *testing.T) {
	testWithParserBytes(t, 23, exampleFile, true, NewDay23, Day23, "7")
}

func TestDay23Part1(t *testing.T) {
	testWithParserBytes(t, 23, file, true, NewDay23, Day23, "164")
}

func BenchmarkDay23Part1(b *testing.B) {
	benchWithParserBytes(b, 23, true, NewDay23, Day23)
}
