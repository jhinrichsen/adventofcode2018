package adventofcode2018

import (
	"fmt"
)

// Day14Puzzle represents the number of recipes to skip.
type Day14Puzzle struct {
	recipes int
}

// NewDay14 parses the number of recipes from data.
func NewDay14(data []byte) (Day14Puzzle, error) {
	n := len(data)
	i := 0

	isdigit := func(b byte) bool {
		return b >= '0' && b <= '9'
	}

	// Skip whitespace
	for i < n && !isdigit(data[i]) {
		i++
	}

	// Read number
	num := 0
	for i < n && isdigit(data[i]) {
		num = num*10 + int(data[i]-'0')
		i++
	}

	return Day14Puzzle{recipes: num}, nil
}

// Day14 generates chocolate recipe scores.
// Part 1: Returns the 10 scores after the given number of recipes.
func Day14(puzzle Day14Puzzle, part1 bool) string {
	if !part1 {
		// Part 2 not implemented yet
		return ""
	}

	// Start with recipes [3, 7]
	scores := []int{3, 7}
	elf1, elf2 := 0, 1

	// Generate recipes until we have enough
	target := puzzle.recipes + 10
	for len(scores) < target {
		// Sum current recipes
		sum := scores[elf1] + scores[elf2]

		// Add new recipes
		if sum >= 10 {
			scores = append(scores, 1, sum%10)
		} else {
			scores = append(scores, sum)
		}

		// Move elves
		elf1 = (elf1 + 1 + scores[elf1]) % len(scores)
		elf2 = (elf2 + 1 + scores[elf2]) % len(scores)
	}

	// Build result string from 10 scores after puzzle.recipes
	result := ""
	for i := 0; i < 10; i++ {
		result += fmt.Sprintf("%d", scores[puzzle.recipes+i])
	}
	return result
}
