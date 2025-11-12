package adventofcode2018

import (
	"fmt"
)

// Day18Puzzle represents the lumber collection area.
type Day18Puzzle struct {
	grid [][]byte
}

// NewDay18 parses the initial grid state.
func NewDay18(data []byte) (Day18Puzzle, error) {
	n := len(data)
	i := 0

	var grid [][]byte

	for i < n {
		start := i
		// Read until newline
		for i < n && data[i] != '\n' && data[i] != '\r' {
			i++
		}
		if i > start {
			line := make([]byte, i-start)
			copy(line, data[start:i])
			grid = append(grid, line)
		}
		// Skip newlines
		for i < n && (data[i] == '\n' || data[i] == '\r') {
			i++
		}
	}

	return Day18Puzzle{grid: grid}, nil
}

// Day18 simulates the lumber collection.
// Part 1: Resource value after 10 minutes.
func Day18(puzzle Day18Puzzle, part1 bool) string {
	if !part1 {
		return ""
	}

	grid := copyGrid(puzzle.grid)

	// Simulate 10 minutes
	for minute := 0; minute < 10; minute++ {
		grid = step(grid)
	}

	// Count wooded acres and lumberyards
	trees := 0
	lumberyards := 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == '|' {
				trees++
			} else if grid[y][x] == '#' {
				lumberyards++
			}
		}
	}

	return fmt.Sprintf("%d", trees*lumberyards)
}

// step simulates one minute of transformation.
func step(grid [][]byte) [][]byte {
	height := len(grid)
	width := len(grid[0])
	next := make([][]byte, height)
	for y := range next {
		next[y] = make([]byte, width)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			neighbors := countNeighbors(grid, x, y)
			current := grid[y][x]

			switch current {
			case '.': // Open ground
				// Becomes trees if 3+ adjacent trees
				if neighbors['|'] >= 3 {
					next[y][x] = '|'
				} else {
					next[y][x] = '.'
				}
			case '|': // Trees
				// Becomes lumberyard if 3+ adjacent lumberyards
				if neighbors['#'] >= 3 {
					next[y][x] = '#'
				} else {
					next[y][x] = '|'
				}
			case '#': // Lumberyard
				// Remains lumberyard if adjacent to 1+ lumberyard AND 1+ trees
				if neighbors['#'] >= 1 && neighbors['|'] >= 1 {
					next[y][x] = '#'
				} else {
					next[y][x] = '.'
				}
			}
		}
	}

	return next
}

// countNeighbors counts the types of adjacent acres.
func countNeighbors(grid [][]byte, x, y int) map[byte]int {
	counts := make(map[byte]int)
	height := len(grid)
	width := len(grid[0])

	// Check all 8 adjacent positions
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			ny := y + dy
			nx := x + dx
			if ny >= 0 && ny < height && nx >= 0 && nx < width {
				counts[grid[ny][nx]]++
			}
		}
	}

	return counts
}

// copyGrid creates a deep copy of the grid.
func copyGrid(grid [][]byte) [][]byte {
	copy := make([][]byte, len(grid))
	for y := range grid {
		copy[y] = make([]byte, len(grid[y]))
		for x := range grid[y] {
			copy[y][x] = grid[y][x]
		}
	}
	return copy
}
