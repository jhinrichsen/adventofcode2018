package adventofcode2018

// NewDay18 parses the lumber collection area grid.
func NewDay18(data []byte) ([]byte, int, int, error) {
	// Count dimensions
	width := 0
	for i := 0; i < len(data); i++ {
		if data[i] == '\n' {
			width = i
			break
		}
	}
	if width == 0 {
		width = len(data)
	}

	// Create grid without newlines
	grid := make([]byte, 0, len(data))
	for i := 0; i < len(data); i++ {
		if data[i] != '\n' {
			grid = append(grid, data[i])
		}
	}

	height := len(grid) / width
	return grid, width, height, nil
}

// countAdjacent counts how many of the 8 neighbors match the target character.
func countAdjacent(grid []byte, width, height, x, y int, target byte) int {
	count := 0
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
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
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
func resourceValue(grid []byte) int {
	trees := 0
	lumberyards := 0
	for i := 0; i < len(grid); i++ {
		if grid[i] == '|' {
			trees++
		} else if grid[i] == '#' {
			lumberyards++
		}
	}
	return trees * lumberyards
}

// Day18Part1 simulates 10 minutes and returns the resource value.
func Day18Part1(grid []byte, width, height int) int {
	current := make([]byte, len(grid))
	copy(current, grid)

	for minute := 0; minute < 10; minute++ {
		current = step(current, width, height)
	}

	return resourceValue(current)
}

// Day18Part2 simulates 1000000000 minutes using cycle detection.
func Day18Part2(grid []byte, width, height int) int {
	current := make([]byte, len(grid))
	copy(current, grid)

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
			current = step(current, width, height)
			minute++
		}
	}

	return resourceValue(current)
}
