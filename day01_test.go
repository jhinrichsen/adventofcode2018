package adventofcode2018

import (
	"testing"
)

func TestDay01Part1(t *testing.T) {
	testWithSolverBytes(t, 1, true, Day01, 454)
}

func BenchmarkDay01Part1(b *testing.B) {
	benchWithSolverBytes(b, 1, true, Day01)
}

func TestDay01Part2(t *testing.T) {
	testWithSolverBytes(t, 1, false, Day01, 566)
}

func BenchmarkDay01Part2(b *testing.B) {
	benchWithSolverBytes(b, 1, false, Day01)
}
