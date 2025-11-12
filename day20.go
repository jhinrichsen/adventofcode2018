package adventofcode2018

import (
	"fmt"
)

// Day20Puzzle represents the facility regex.
type Day20Puzzle struct {
	regex string
}

// NewDay20 parses the regex.
func NewDay20(data []byte) (Day20Puzzle, error) {
	// Strip whitespace
	regex := ""
	for _, b := range data {
		if b != '\n' && b != '\r' && b != ' ' {
			regex += string(b)
		}
	}
	return Day20Puzzle{regex: regex}, nil
}

// Day20 finds the furthest room.
// Part 1: Returns the maximum number of doors to reach any room.
func Day20(puzzle Day20Puzzle, part1 bool) string {
	if !part1 {
		return ""
	}

	// Build the map by following all paths
	doors := make(map[pos]map[pos]bool) // doors[pos1][pos2] = true means door between pos1 and pos2

	buildMap(puzzle.regex, doors)

	// BFS to find distances
	start := pos{0, 0}
	dist := make(map[pos]int)
	dist[start] = 0
	queue := []pos{start}
	maxDist := 0

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		// Check all 4 directions
		for _, d := range []pos{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
			next := pos{curr.x + d.x, curr.y + d.y}
			// Check if there's a door
			if doors[curr][next] {
				if _, visited := dist[next]; !visited {
					dist[next] = dist[curr] + 1
					if dist[next] > maxDist {
						maxDist = dist[next]
					}
					queue = append(queue, next)
				}
			}
		}
	}

	return fmt.Sprintf("%d", maxDist)
}

// buildMap constructs the map by parsing the regex and following all paths.
func buildMap(regex string, doors map[pos]map[pos]bool) {
	// Stack for handling branches
	var stack [][]pos

	current := []pos{{0, 0}}
	i := 1 // Skip ^

	for i < len(regex)-1 { // Skip $
		c := regex[i]
		switch c {
		case 'N', 'S', 'E', 'W':
			// Move in direction
			var next []pos
			for _, p := range current {
				np := p
				switch c {
				case 'N':
					np.y--
				case 'S':
					np.y++
				case 'E':
					np.x++
				case 'W':
					np.x--
				}
				// Add door
				if doors[p] == nil {
					doors[p] = make(map[pos]bool)
				}
				if doors[np] == nil {
					doors[np] = make(map[pos]bool)
				}
				doors[p][np] = true
				doors[np][p] = true
				next = append(next, np)
			}
			current = next
		case '(':
			// Start branch - save current positions
			stack = append(stack, current)
		case '|':
			// Branch alternative - reset to saved positions
			if len(stack) > 0 {
				current = stack[len(stack)-1]
			}
		case ')':
			// End branch - pop stack
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		}
		i++
	}
}
