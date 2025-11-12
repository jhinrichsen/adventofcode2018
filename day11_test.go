package adventofcode2018

import (
	"fmt"
	"testing"
)

func TestDay11Part1Examples(t *testing.T) {
	tests := []struct {
		serial int
		want   string
	}{
		{18, "33,45"},
		{42, "21,61"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("serial-%d", tt.serial), func(t *testing.T) {
			puzzle := Day11Puzzle{serial: tt.serial}
			got := Day11(puzzle, true)
			if got != tt.want {
				t.Fatalf("want %s but got %s", tt.want, got)
			}
		})
	}
}

func TestDay11Part1(t *testing.T) {
	testWithParserBytes(t, 11, file, true, NewDay11, Day11, "243,72")
}

func BenchmarkDay11Part1(b *testing.B) {
	benchWithParserBytes(b, 11, true, NewDay11, Day11)
}
