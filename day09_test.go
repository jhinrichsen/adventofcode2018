package adventofcode2018

import "testing"

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
	}

	for _, tt := range tests {
		puzzle := Day09Puzzle{players: tt.players, lastMarble: tt.lastMarble}
		got := Day09(puzzle, true)
		if got != tt.want {
			t.Errorf("Day09(%d players, %d marbles) = %d, want %d",
				tt.players, tt.lastMarble, got, tt.want)
		}
	}
}

