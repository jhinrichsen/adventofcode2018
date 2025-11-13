package adventofcode2018

import (
	"fmt"
	"testing"
)

func TestDay14Part1Examples(t *testing.T) {
	tests := []struct {
		recipes int
		want    uint
	}{
		{9, 5158916779},
		{5, 124515891}, // Note: leading zero lost when converted to uint
		{18, 9251071085},
		{2018, 5941429882},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("recipes-%d", tt.recipes), func(t *testing.T) {
			puzzle := Day14Puzzle{recipes: tt.recipes}
			got := Day14(puzzle, true)
			if got != tt.want {
				t.Fatalf("want %d but got %d", tt.want, got)
			}
		})
	}
}

func TestDay14Part1(t *testing.T) {
	testWithParserBytes(t, 14, file, true, NewDay14, Day14, uint(3147574107))
}

func BenchmarkDay14Part1(b *testing.B) {
	benchWithParserBytes(b, 14, true, NewDay14, Day14)
}

func TestDay14Part2Examples(t *testing.T) {
	tests := []struct {
		recipes int
		want    uint
	}{
		{51589, 9},
		{92510, 18},
		{59414, 2018},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("recipes-%d", tt.recipes), func(t *testing.T) {
			puzzle := Day14Puzzle{recipes: tt.recipes}
			got := Day14(puzzle, false)
			if got != tt.want {
				t.Fatalf("want %d but got %d", tt.want, got)
			}
		})
	}
}

func TestDay14Part2(t *testing.T) {
	testWithParserBytes(t, 14, file, false, NewDay14, Day14, uint(20280190))
}

func BenchmarkDay14Part2(b *testing.B) {
	benchWithParserBytes(b, 14, false, NewDay14, Day14)
}
