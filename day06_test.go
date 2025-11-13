package adventofcode2018

import (
	"testing"
)

func TestDay06Part1(t *testing.T) {
	testWithSolverBytes(t, 6, true, Day06, 4342)
}

func BenchmarkDay06Part1(b *testing.B) {
	benchWithSolverBytes(b, 6, true, Day06)
}

func TestDay06Part2(t *testing.T) {
	testWithSolverBytes(t, 6, false, Day06, 42966)
}

func BenchmarkDay06Part2(b *testing.B) {
	benchWithSolverBytes(b, 6, false, Day06)
}
