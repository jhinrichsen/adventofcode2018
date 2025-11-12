package adventofcode2018

import (
	"fmt"
	"testing"
)

func TestDay14Part1Examples(t *testing.T) {
	tests := []struct {
		recipes int
		want    string
	}{
		{9, "5158916779"},
		{5, "0124515891"},
		{18, "9251071085"},
		{2018, "5941429882"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("recipes-%d", tt.recipes), func(t *testing.T) {
			puzzle := Day14Puzzle{recipes: tt.recipes}
			got := Day14(puzzle, true)
			if got != tt.want {
				t.Fatalf("want %s but got %s", tt.want, got)
			}
		})
	}
}

func TestDay14Part1(t *testing.T) {
	testWithParserBytes(t, 14, file, true, NewDay14, Day14, "3147574107")
}

func BenchmarkDay14Part1(b *testing.B) {
	benchWithParserBytes(b, 14, true, NewDay14, Day14)
}
