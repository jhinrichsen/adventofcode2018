package adventofcode2018

// Day18Puzzle represents the lumber collection area grid.
type Day18Puzzle struct {
	grid   []byte
	width  int
	height int
}

// NewDay18 parses the lumber collection area grid from raw bytes.
func NewDay18(data []byte) (Day18Puzzle, error) {
	if len(data) == 0 {
		return Day18Puzzle{}, nil
	}

	// Find dimensions by counting first line and total lines
	width := 0
	for i := 0; i < len(data); i++ {
		if data[i] == '\n' || data[i] == '\r' {
			break
		}
		width++
	}

	if width == 0 {
		return Day18Puzzle{}, nil
	}

	// Count lines and build grid without newlines
	grid := make([]byte, 0, len(data))
	height := 0
	for i := 0; i < len(data); i++ {
		if data[i] == '\n' || data[i] == '\r' {
			if i+1 < len(data) && data[i] == '\r' && data[i+1] == '\n' {
				i++ // Skip \r\n
			}
			if len(grid) > height*width {
				height++
			}
		} else {
			grid = append(grid, data[i])
		}
	}
	// Count the last line if it doesn't end with newline
	if len(grid) > height*width {
		height++
	}

	return Day18Puzzle{
		grid:   grid,
		width:  width,
		height: height,
	}, nil
}

// step simulates one minute using C8Indices for neighbor iteration.
func step(grid []byte, width, height int) []byte {
	next := make([]byte, len(grid))
	g := Grid{W: width, H: height}

	for idx, nbrs := range g.C8Indices() {
		current := grid[idx]

		// Count trees and lumberyards in single pass
		var trees, lumberyards uint
		for nidx := range nbrs {
			switch grid[nidx] {
			case '|':
				trees++
			case '#':
				lumberyards++
			}
		}

		switch current {
		case '.': // open ground -> trees if 3+ adjacent trees
			if trees >= 3 {
				next[idx] = '|'
			} else {
				next[idx] = '.'
			}
		case '|': // trees -> lumberyard if 3+ adjacent lumberyards
			if lumberyards >= 3 {
				next[idx] = '#'
			} else {
				next[idx] = '|'
			}
		case '#': // lumberyard -> stays if 1+ lumberyard AND 1+ trees
			if lumberyards >= 1 && trees >= 1 {
				next[idx] = '#'
			} else {
				next[idx] = '.'
			}
		}
	}
	return next
}

// resourceValue calculates the total resource value (wooded acres * lumberyards).
func resourceValue(grid []byte) uint {
	var trees, lumberyards uint
	for _, cell := range grid {
		if cell == '|' {
			trees++
		} else if cell == '#' {
			lumberyards++
		}
	}
	return trees * lumberyards
}

// Day18 is a unified solver for both parts.
// Part 1: simulates 10 minutes.
// Part 2: simulates 1000000000 minutes using cycle detection.
func Day18(p Day18Puzzle, part1 bool) uint {
	current := make([]byte, len(p.grid))
	copy(current, p.grid)

	target := 10
	if !part1 {
		target = 1000000000
	}

	// For part 2, track seen states to detect cycles
	var seen map[string]int
	if !part1 {
		seen = make(map[string]int)
	}

	minute := 0
	for minute < target {
		// For part 2, check for cycles
		if !part1 {
			state := string(current)
			if prevMinute, found := seen[state]; found {
				// Found a cycle!
				cycleLength := minute - prevMinute
				remaining := target - minute
				fullCycles := remaining / cycleLength
				minute += fullCycles * cycleLength
				// Clear the cache and continue simulating the remainder
				seen = make(map[string]int)
			}
			if minute < target {
				seen[state] = minute
			}
		}

		if minute < target {
			current = step(current, p.width, p.height)
			minute++
		}
	}

	return resourceValue(current)
}
