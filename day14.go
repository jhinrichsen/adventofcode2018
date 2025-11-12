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
// Part 2: Returns the number of recipes before the input sequence appears.
func Day14(puzzle Day14Puzzle, part1 bool) string {
	// Start with recipes [3, 7]
	scores := []int{3, 7}
	elf1, elf2 := 0, 1

	if part1 {
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

	// Part 2: Find when the input sequence appears
	// Convert input number to digits
	target := make([]int, 0)
	n := puzzle.recipes
	if n == 0 {
		target = []int{0}
	} else {
		for n > 0 {
			target = append([]int{n % 10}, target...)
			n /= 10
		}
	}

	// Generate recipes until we find the sequence
	for {
		// Sum current recipes
		sum := scores[elf1] + scores[elf2]

		// Add new recipes and check for match
		if sum >= 10 {
			scores = append(scores, 1)
			if matchesEnd(scores, target) {
				return fmt.Sprintf("%d", len(scores)-len(target))
			}
			scores = append(scores, sum%10)
			if matchesEnd(scores, target) {
				return fmt.Sprintf("%d", len(scores)-len(target))
			}
		} else {
			scores = append(scores, sum)
			if matchesEnd(scores, target) {
				return fmt.Sprintf("%d", len(scores)-len(target))
			}
		}

		// Move elves
		elf1 = (elf1 + 1 + scores[elf1]) % len(scores)
		elf2 = (elf2 + 1 + scores[elf2]) % len(scores)
	}
}

// matchesEnd checks if the end of scores matches target.
func matchesEnd(scores, target []int) bool {
	if len(scores) < len(target) {
		return false
	}
	start := len(scores) - len(target)
	for i := range target {
		if scores[start+i] != target[i] {
			return false
		}
	}
	return true
}
