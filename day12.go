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
func Day12(puzzle Day12Puzzle, part1 bool) string {
	if !part1 {
		// Part 2 not implemented yet
		return ""
	}

	// Initialize plants map
	plants := make(map[int]bool)
	for i, c := range puzzle.initial {
		if c == '#' {
			plants[i] = true
		}
	}

	// Simulate 20 generations
	for gen := 0; gen < 20; gen++ {
		// Find bounds
		minPot := 0
		maxPot := 0
		for pot := range plants {
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
	}

	// Sum pot numbers with plants
	sum := 0
	for pot := range plants {
		sum += pot
	}

	return fmt.Sprintf("%d", sum)
}
