package adventofcode2018

import (
	"testing"
)

func TestDay01Part1(t *testing.T) {
	testWithParserLines(t, 1, true, NewDay01, Day01, 454)
}

func BenchmarkDay01Part1(b *testing.B) {
	benchWithParserLines(b, 1, true, NewDay01, Day01)
}

func TestDay01Part2(t *testing.T) {
	testWithParserLines(t, 1, false, NewDay01, Day01, 566)
}

func BenchmarkDay01Part2(b *testing.B) {
	benchWithParserLines(b, 1, false, NewDay01, Day01)
}
