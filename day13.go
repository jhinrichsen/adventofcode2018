package adventofcode2018

import (
	"fmt"
	"sort"
)

// Day13Puzzle represents the mine cart track system.
type Day13Puzzle struct {
	track [][]byte
	carts []cart
}

type cart struct {
	x, y      int
	dx, dy    int // direction
	turnState int // 0=left, 1=straight, 2=right
}

// NewDay13 parses the track and carts from data.
func NewDay13(data []byte) (Day13Puzzle, error) {
	n := len(data)
	i := 0

	// Parse lines
	var lines [][]byte
	for i < n {
		start := i
		for i < n && data[i] != '\n' && data[i] != '\r' {
			i++
		}
		if i > start {
			line := make([]byte, i-start)
			copy(line, data[start:i])
			lines = append(lines, line)
		}
		for i < n && (data[i] == '\n' || data[i] == '\r') {
			i++
		}
	}

	if len(lines) == 0 {
		return Day13Puzzle{}, fmt.Errorf("no track data")
	}

	// Find max width
	maxWidth := 0
	for _, line := range lines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}

	// Build track grid and find carts
	track := make([][]byte, len(lines))
	var carts []cart

	for y, line := range lines {
		track[y] = make([]byte, maxWidth)
		for x := range maxWidth {
			if x < len(line) {
				track[y][x] = line[x]
			} else {
				track[y][x] = ' '
			}

			// Check for cart
			c := track[y][x]
			switch c {
			case '^':
				carts = append(carts, cart{x: x, y: y, dx: 0, dy: -1, turnState: 0})
				track[y][x] = '|'
			case 'v':
				carts = append(carts, cart{x: x, y: y, dx: 0, dy: 1, turnState: 0})
				track[y][x] = '|'
			case '<':
				carts = append(carts, cart{x: x, y: y, dx: -1, dy: 0, turnState: 0})
				track[y][x] = '-'
			case '>':
				carts = append(carts, cart{x: x, y: y, dx: 1, dy: 0, turnState: 0})
				track[y][x] = '-'
			}
		}
	}

	return Day13Puzzle{track: track, carts: carts}, nil
}

// Day13 simulates mine cart movement.
// Part 1: Returns X,Y coordinates of first collision.
func Day13(puzzle Day13Puzzle, part1 bool) string {
	if !part1 {
		// Part 2 not implemented yet
		return ""
	}

	// Make a copy of carts
	carts := make([]cart, len(puzzle.carts))
	copy(carts, puzzle.carts)

	for {
		// Sort carts by position (top to bottom, left to right)
		sort.Slice(carts, func(i, j int) bool {
			if carts[i].y != carts[j].y {
				return carts[i].y < carts[j].y
			}
			return carts[i].x < carts[j].x
		})

		// Move each cart
		for i := range carts {
			// Move cart
			carts[i].x += carts[i].dx
			carts[i].y += carts[i].dy

			// Check for collision
			for j := range carts {
				if i != j && carts[i].x == carts[j].x && carts[i].y == carts[j].y {
					return fmt.Sprintf("%d,%d", carts[i].x, carts[i].y)
				}
			}

			// Update direction based on track
			cell := puzzle.track[carts[i].y][carts[i].x]
			switch cell {
			case '/':
				// Swap and negate one
				if carts[i].dx != 0 {
					carts[i].dy = -carts[i].dx
					carts[i].dx = 0
				} else {
					carts[i].dx = -carts[i].dy
					carts[i].dy = 0
				}
			case '\\':
				// Swap
				if carts[i].dx != 0 {
					carts[i].dy = carts[i].dx
					carts[i].dx = 0
				} else {
					carts[i].dx = carts[i].dy
					carts[i].dy = 0
				}
			case '+':
				// Intersection - turn based on state
				switch carts[i].turnState {
				case 0: // Turn left
					carts[i].dx, carts[i].dy = carts[i].dy, -carts[i].dx
				case 2: // Turn right
					carts[i].dx, carts[i].dy = -carts[i].dy, carts[i].dx
				// case 1: straight - no change
				}
				carts[i].turnState = (carts[i].turnState + 1) % 3
			}
		}
	}
}
