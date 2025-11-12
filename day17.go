package adventofcode2018

import (
	"fmt"
)

// Day17Puzzle represents the underground scan.
type Day17Puzzle struct {
	clay  map[pos]bool
	minY  int
	maxY  int
	water map[pos]byte // '|' for flowing, '~' for settled
}

// NewDay17 parses the clay positions.
func NewDay17(data []byte) (Day17Puzzle, error) {
	clay := make(map[pos]bool)
	n := len(data)
	i := 0

	minY := int(1e9)
	maxY := 0

	for i < n {
		// Parse line: "x=495, y=2..7" or "y=7, x=495..501"
		if i >= n {
			break
		}

		// Read first coordinate type
		var firstType byte
		if i < n && (data[i] == 'x' || data[i] == 'y') {
			firstType = data[i]
			i++
		}
		// Skip '='
		if i < n && data[i] == '=' {
			i++
		}

		// Read first value
		firstVal := 0
		for i < n && data[i] >= '0' && data[i] <= '9' {
			firstVal = firstVal*10 + int(data[i]-'0')
			i++
		}

		// Skip ", "
		for i < n && (data[i] == ',' || data[i] == ' ') {
			i++
		}

		// Skip second coordinate type (x or y) and '='
		if i < n && (data[i] == 'x' || data[i] == 'y') {
			i++
		}
		if i < n && data[i] == '=' {
			i++
		}

		// Read range start
		rangeStart := 0
		for i < n && data[i] >= '0' && data[i] <= '9' {
			rangeStart = rangeStart*10 + int(data[i]-'0')
			i++
		}

		// Check for ".."
		rangeEnd := rangeStart
		if i+1 < n && data[i] == '.' && data[i+1] == '.' {
			i += 2
			rangeEnd = 0
			for i < n && data[i] >= '0' && data[i] <= '9' {
				rangeEnd = rangeEnd*10 + int(data[i]-'0')
				i++
			}
		}

		// Add clay positions
		if firstType == 'x' {
			x := firstVal
			for y := rangeStart; y <= rangeEnd; y++ {
				clay[pos{x, y}] = true
				if y < minY {
					minY = y
				}
				if y > maxY {
					maxY = y
				}
			}
		} else {
			y := firstVal
			for x := rangeStart; x <= rangeEnd; x++ {
				clay[pos{x, y}] = true
			}
			if y < minY {
				minY = y
			}
			if y > maxY {
				maxY = y
			}
		}

		// Skip newline
		for i < n && (data[i] == '\n' || data[i] == '\r') {
			i++
		}
	}

	return Day17Puzzle{clay: clay, minY: minY, maxY: maxY, water: make(map[pos]byte)}, nil
}

// Day17 simulates water flow.
// Part 1: Count all tiles reached by water (settled + flowing).
// Part 2: Count only retained water (settled).
func Day17(puzzle Day17Puzzle, part1 bool) string {
	// Start flowing from (500, 0)
	flow(&puzzle, 500, 0)

	// Count water tiles within valid y range
	count := 0
	for p, w := range puzzle.water {
		if p.y >= puzzle.minY && p.y <= puzzle.maxY {
			if part1 {
				// Part 1: Count all water (flowing + settled)
				if w == '|' || w == '~' {
					count++
				}
			} else {
				// Part 2: Count only retained water (settled)
				if w == '~' {
					count++
				}
			}
		}
	}

	return fmt.Sprintf("%d", count)
}

// flow simulates water flowing from position (x, y).
func flow(p *Day17Puzzle, x, y int) bool {
	// Out of bounds
	if y > p.maxY {
		return false
	}

	// Already has water here
	if _, ok := p.water[pos{x, y}]; ok {
		return p.water[pos{x, y}] == '~'
	}

	// Hit clay
	if p.clay[pos{x, y}] {
		return true
	}

	// Mark as flowing water
	p.water[pos{x, y}] = '|'

	// Try to flow down
	below := flow(p, x, y+1)

	if !below {
		// Water continues flowing down
		return false
	}

	// Water is blocked below, try to spread horizontally
	leftBlocked := fillHorizontal(p, x, y, -1)
	rightBlocked := fillHorizontal(p, x, y, 1)

	if leftBlocked && rightBlocked {
		// Water settles - mark entire row as settled
		settleRow(p, x, y)
		return true
	}

	return false
}

// fillHorizontal fills water horizontally in direction dx (-1 for left, 1 for right).
// Returns true if blocked by clay, false if flows down.
func fillHorizontal(p *Day17Puzzle, x, y, dx int) bool {
	for {
		x += dx

		// Hit clay
		if p.clay[pos{x, y}] {
			return true
		}

		// Mark as flowing
		p.water[pos{x, y}] = '|'

		// Check below - need clay or settled water to continue spreading
		belowPos := pos{x, y + 1}
		if !p.clay[belowPos] && p.water[belowPos] != '~' {
			// Water can flow down from here
			flow(p, x, y+1)
			// After flowing, check if it's now blocked
			if p.water[belowPos] != '~' {
				return false
			}
		}
	}
}

// settleRow marks an entire row as settled water.
func settleRow(p *Day17Puzzle, x, y int) {
	// Mark center
	p.water[pos{x, y}] = '~'

	// Mark left
	for dx := x - 1; ; dx-- {
		if p.clay[pos{dx, y}] {
			break
		}
		p.water[pos{dx, y}] = '~'
	}

	// Mark right
	for dx := x + 1; ; dx++ {
		if p.clay[pos{dx, y}] {
			break
		}
		p.water[pos{dx, y}] = '~'
	}
}
