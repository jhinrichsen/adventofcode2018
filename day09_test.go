package adventofcode2018

import (
	"fmt"
	"testing"
)

func TestDay09Part1Examples(t *testing.T) {
	tests := []struct {
		players    uint
		lastMarble uint
		want       uint
	}{
		{10, 1618, 8317},
		{13, 7999, 146373},
		{17, 1104, 2764},
		{21, 6111, 54718},
		{30, 5807, 37305},
		{455, 71223, 384288},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%dp-%dm", tt.players, tt.lastMarble), func(t *testing.T) {
			puzzle := Day09Puzzle{players: tt.players, lastMarble: tt.lastMarble}
			got := Day09(puzzle, true)
			if got != tt.want {
				t.Fatalf("want %d but got %d", tt.want, got)
			}
		})
	}
}

func TestDay09Part1(t *testing.T) {
	testWithParserBytes(t, 9, file, true, NewDay09, Day09, 384288)
}

func TestDay09Part2(t *testing.T) {
	testWithParserBytes(t, 9, file, false, NewDay09, Day09, 3189426841)
}

func BenchmarkDay09Part1(b *testing.B) {
	benchWithParserBytes(b, 9, true, NewDay09, Day09)
}

func BenchmarkDay09Part2(b *testing.B) {
	benchWithParserBytes(b, 9, false, NewDay09, Day09)
}
