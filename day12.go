package adventofcode2018

import (
	"fmt"
)

// Day12Puzzle represents the plant growth simulation.
type Day12Puzzle struct {
	initial string
	rules   map[string]bool
}

// NewDay12 parses the initial state and rules from data.
func NewDay12(data []byte) (Day12Puzzle, error) {
	n := len(data)
	i := 0

	// Skip to initial state value (after "initial state: ")
	for i < n && data[i] != ':' {
		i++
	}
	i++ // skip ':'
	for i < n && data[i] == ' ' {
		i++
	}

	// Read initial state
	start := i
	for i < n && data[i] != '\n' && data[i] != '\r' {
		i++
	}
	initial := string(data[start:i])

	// Skip blank line(s)
	for i < n && (data[i] == '\n' || data[i] == '\r') {
		i++
	}

	// Parse rules
	rules := make(map[string]bool)
	for i < n {
		if i+9 >= n {
			break
		}

		// Read pattern (5 chars)
		pattern := string(data[i : i+5])
		i += 5

		// Skip " => "
		for i < n && data[i] == ' ' {
			i++
		}
		if i < n && data[i] == '=' {
			i++
		}
		if i < n && data[i] == '>' {
			i++
		}
		for i < n && data[i] == ' ' {
			i++
		}

		// Read result (1 char)
		if i < n {
			result := data[i] == '#'
			rules[pattern] = result
			i++
		}

		// Skip to next line
		for i < n && (data[i] == '\n' || data[i] == '\r') {
			i++
		}
	}

	return Day12Puzzle{initial: initial, rules: rules}, nil
}

// Day12 simulates plant growth.
// Part 1: Returns the sum of pot numbers after 20 generations.
// Part 2: Returns the sum after 50 billion generations by detecting pattern.
func Day12(puzzle Day12Puzzle, part1 bool) string {
	generations := 20
	if !part1 {
		generations = 50000000000
	}

	// Initialize plants map
	plants := make(map[int]bool)
	for i, c := range puzzle.initial {
		if c == '#' {
			plants[i] = true
		}
	}

	// For part 2, detect when pattern stabilizes
	var prevSum, prevDiff int
	stableCount := 0

	for gen := 0; gen < generations; gen++ {
		// Find bounds
		minPot := 0
		maxPot := 0
		first := true
		for pot := range plants {
			if first {
				minPot = pot
				maxPot = pot
				first = false
			}
			if pot < minPot {
				minPot = pot
			}
			if pot > maxPot {
				maxPot = pot
			}
		}

		// Expand bounds for pattern matching
		minPot -= 2
		maxPot += 2

		// Calculate next generation
		nextPlants := make(map[int]bool)
		for pot := minPot; pot <= maxPot; pot++ {
			// Build pattern for this pot
			pattern := ""
			for offset := -2; offset <= 2; offset++ {
				if plants[pot+offset] {
					pattern += "#"
				} else {
					pattern += "."
				}
			}

			// Check if this pattern produces a plant
			if puzzle.rules[pattern] {
				nextPlants[pot] = true
			}
		}

		plants = nextPlants

		// For part 2, detect stable pattern
		if !part1 && gen > 100 {
			sum := 0
			for pot := range plants {
				sum += pot
			}

			diff := sum - prevSum
			if diff == prevDiff {
				stableCount++
				if stableCount > 10 {
					// Pattern is stable, extrapolate
					remaining := generations - gen - 1
					finalSum := sum + (remaining * diff)
					return fmt.Sprintf("%d", finalSum)
				}
			} else {
				stableCount = 0
			}

			prevSum = sum
			prevDiff = diff
		}
	}

	// Sum pot numbers with plants
	sum := 0
	for pot := range plants {
		sum += pot
	}

	return fmt.Sprintf("%d", sum)
}
