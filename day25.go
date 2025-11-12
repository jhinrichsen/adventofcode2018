package adventofcode2018

import (
	"fmt"
)

// Day25Puzzle represents 4D points.
type Day25Puzzle struct {
	points []point4D
}

type point4D struct {
	x, y, z, w int
}

// NewDay25 parses the 4D coordinates.
func NewDay25(data []byte) (Day25Puzzle, error) {
	n := len(data)
	i := 0

	var points []point4D

	for i < n {
		if i >= n {
			break
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

		// Skip ','
		if i < n && data[i] == ',' {
			i++
		}

		// Read w
		w := 0
		negative = false
		if i < n && data[i] == '-' {
			negative = true
			i++
		}
		for i < n && data[i] >= '0' && data[i] <= '9' {
			w = w*10 + int(data[i]-'0')
			i++
		}
		if negative {
			w = -w
		}

		points = append(points, point4D{x: x, y: y, z: z, w: w})

		// Skip newline
		for i < n && (data[i] == '\n' || data[i] == '\r') {
			i++
		}
	}

	return Day25Puzzle{points: points}, nil
}

// Day25 counts constellations.
// Part 1: Returns the number of constellations.
func Day25(puzzle Day25Puzzle, part1 bool) string {
	if !part1 {
		return ""
	}

	// Union-find
	parent := make(map[int]int)
	for i := range puzzle.points {
		parent[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(x, y int) {
		px, py := find(x), find(y)
		if px != py {
			parent[px] = py
		}
	}

	// Connect points within Manhattan distance 3
	for i := range puzzle.points {
		for j := i + 1; j < len(puzzle.points); j++ {
			if manhattanDist4D(puzzle.points[i], puzzle.points[j]) <= 3 {
				union(i, j)
			}
		}
	}

	// Count distinct roots
	roots := make(map[int]bool)
	for i := range puzzle.points {
		roots[find(i)] = true
	}

	return fmt.Sprintf("%d", len(roots))
}

// manhattanDist4D calculates Manhattan distance in 4D.
func manhattanDist4D(p1, p2 point4D) int {
	return abs(p2.x-p1.x) + abs(p2.y-p1.y) + abs(p2.z-p1.z) + abs(p2.w-p1.w)
}
