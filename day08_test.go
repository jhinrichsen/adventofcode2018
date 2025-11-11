package adventofcode2018

import "testing"

func TestDay08Part1Example(t *testing.T) {
	testDayPart(t, 8, exampleFilename, true, NewDay08, Day08, uint(138))
}

func TestDay08Part1(t *testing.T) {
	testDayPart(t, 8, filename, true, NewDay08, Day08, uint(47464))
}

func TestDay08Part2Example(t *testing.T) {
	testDayPart(t, 8, exampleFilename, false, NewDay08, Day08, uint(66))
}

func TestDay08Part2(t *testing.T) {
	testDayPart(t, 8, filename, false, NewDay08, Day08, uint(23054))
}
