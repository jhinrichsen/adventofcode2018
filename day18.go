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

// countAdjacent counts how many of the 8 neighbors match the target character.
func countAdjacent(grid []byte, width, height, x, y int, target byte) uint {
	count := uint(0)
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := x+dx, y+dy
			if nx >= 0 && nx < width && ny >= 0 && ny < height {
				if grid[ny*width+nx] == target {
					count++
				}
			}
		}
	}
	return count
}

// step simulates one minute of the lumber collection area.
func step(grid []byte, width, height int) []byte {
	next := make([]byte, len(grid))
	for y := range height {
		for x := range width {
			idx := y*width + x
			current := grid[idx]

			switch current {
			case '.': // open ground
				// Becomes trees if 3+ adjacent are trees
				if countAdjacent(grid, width, height, x, y, '|') >= 3 {
					next[idx] = '|'
				} else {
					next[idx] = '.'
				}
			case '|': // trees
				// Becomes lumberyard if 3+ adjacent are lumberyards
				if countAdjacent(grid, width, height, x, y, '#') >= 3 {
					next[idx] = '#'
				} else {
					next[idx] = '|'
				}
			case '#': // lumberyard
				// Stays lumberyard if adjacent to at least 1 lumberyard AND 1 trees
				// Otherwise becomes open ground
				if countAdjacent(grid, width, height, x, y, '#') >= 1 &&
					countAdjacent(grid, width, height, x, y, '|') >= 1 {
					next[idx] = '#'
				} else {
					next[idx] = '.'
				}
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

// Day18Part1 simulates 10 minutes and returns the resource value.
func Day18Part1(p Day18Puzzle) uint {
	current := make([]byte, len(p.grid))
	copy(current, p.grid)

	for range 10 {
		current = step(current, p.width, p.height)
	}

	return resourceValue(current)
}

// Day18Part2 simulates 1000000000 minutes using cycle detection.
func Day18Part2(p Day18Puzzle) uint {
	current := make([]byte, len(p.grid))
	copy(current, p.grid)

	// Track seen states to detect cycles
	seen := make(map[string]int)
	minute := 0
	target := 1000000000

	for minute < target {
		// Convert current state to string for comparison
		state := string(current)

		// Check if we've seen this state before
		if prevMinute, found := seen[state]; found {
			// Found a cycle!
			cycleLength := minute - prevMinute
			// How many complete cycles can we skip?
			remaining := target - minute
			fullCycles := remaining / cycleLength
			minute += fullCycles * cycleLength

			// Clear the cache and continue simulating the remainder
			seen = make(map[string]int)
		}

		if minute < target {
			seen[state] = minute
			current = step(current, p.width, p.height)
			minute++
		}
	}

	return resourceValue(current)
}

// Day18 is a unified solver for both parts.
func Day18(p Day18Puzzle, part1 bool) uint {
	if part1 {
		return Day18Part1(p)
	}
	return Day18Part2(p)
}
