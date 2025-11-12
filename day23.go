package adventofcode2018

import (
	"fmt"
)

// Day23Puzzle represents the nanobots.
type Day23Puzzle struct {
	bots []nanobot
}

type nanobot struct {
	x, y, z int
	r       int
}

// NewDay23 parses the nanobots.
func NewDay23(data []byte) (Day23Puzzle, error) {
	n := len(data)
	i := 0

	var bots []nanobot

	for i < n {
		// Parse "pos=<x,y,z>, r=radius"
		if i+4 > n || string(data[i:i+4]) != "pos=" {
			break
		}
		i += 4

		// Skip '<'
		if i < n && data[i] == '<' {
			i++
		}

		// Read x
		x := 0
		negative := false
		if i < n && data[i] == '-' {
			negative = true
			i++
		}
		for i < n && data[i] >= '0' && data[i] <= '9' {
			x = x*10 + int(data[i]-'0')
			i++
		}
		if negative {
			x = -x
		}

		// Skip ','
		if i < n && data[i] == ',' {
			i++
		}

		// Read y
		y := 0
		negative = false
		if i < n && data[i] == '-' {
			negative = true
			i++
		}
		for i < n && data[i] >= '0' && data[i] <= '9' {
			y = y*10 + int(data[i]-'0')
			i++
		}
		if negative {
			y = -y
		}

		// Skip ','
		if i < n && data[i] == ',' {
			i++
		}

		// Read z
		z := 0
		negative = false
		if i < n && data[i] == '-' {
			negative = true
			i++
		}
		for i < n && data[i] >= '0' && data[i] <= '9' {
			z = z*10 + int(data[i]-'0')
			i++
		}
		if negative {
			z = -z
		}

		// Skip ">, r="
		for i < n && data[i] != 'r' {
			i++
		}
		if i < n && data[i] == 'r' {
			i++
		}
		if i < n && data[i] == '=' {
			i++
		}

		// Read radius
		r := 0
		for i < n && data[i] >= '0' && data[i] <= '9' {
			r = r*10 + int(data[i]-'0')
			i++
		}

		bots = append(bots, nanobot{x: x, y: y, z: z, r: r})

		// Skip newline
		for i < n && (data[i] == '\n' || data[i] == '\r') {
			i++
		}
	}

	return Day23Puzzle{bots: bots}, nil
}

// Day23 finds nanobots in range of strongest.
// Part 1: Count nanobots in range of the strongest nanobot.
func Day23(puzzle Day23Puzzle, part1 bool) string {
	if !part1 {
		return ""
	}

	// Find strongest nanobot
	strongest := 0
	for i := range puzzle.bots {
		if puzzle.bots[i].r > puzzle.bots[strongest].r {
			strongest = i
		}
	}

	// Count bots in range
	count := 0
	s := puzzle.bots[strongest]
	for _, b := range puzzle.bots {
		dist := manhattanDist3D(s.x, s.y, s.z, b.x, b.y, b.z)
		if dist <= s.r {
			count++
		}
	}

	return fmt.Sprintf("%d", count)
}

// manhattanDist3D calculates Manhattan distance in 3D.
func manhattanDist3D(x1, y1, z1, x2, y2, z2 int) int {
	return abs(x2-x1) + abs(y2-y1) + abs(z2-z1)
}
