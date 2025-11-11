package adventofcode2018

import "testing"

func TestDay08Part1Example(t *testing.T) {
	testDay(t, 8, true, NewDay08, Day08Part1, 138)
}

func TestDay08Part1(t *testing.T) {
	testDay(t, 8, false, NewDay08, Day08Part1, 47464)
}

func TestDay08Part2Example(t *testing.T) {
	testDay(t, 8, true, NewDay08, Day08Part2, 66)
}

func TestDay08Part2(t *testing.T) {
	testDay(t, 8, false, NewDay08, Day08Part2, 23054)
}

func BenchmarkDay08Part1(b *testing.B) {
	benchDay(b, 8, NewDay08, Day08Part1)
}

func BenchmarkDay08Part2(b *testing.B) {
	benchDay(b, 8, NewDay08, Day08Part2)
}
