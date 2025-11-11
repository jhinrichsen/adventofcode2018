package adventofcode2018

import (
	"strings"

	"gitlab.com/jhinrichsen/aococr"
)

// Day10Puzzle represents the moving points of light.
type Day10Puzzle struct {
	points []point
}

type point struct {
	x, y   int
	vx, vy int
}

// NewDay10 parses the points from data.
// Expected format: "position=<X, Y> velocity=<VX, VY>"
func NewDay10(data []byte) (Day10Puzzle, error) {
	isdigit := func(b byte) bool {
		return b >= '0' && b <= '9'
	}

	n := len(data)
	i := 0

	// number reads the next signed integer from data starting at position i
	number := func() int {
		// Skip to next digit or sign
		for i < n && !isdigit(data[i]) && data[i] != '-' {
			i++
		}

		// Check for negative
		sign := 1
		if i < n && data[i] == '-' {
			sign = -1
			i++
		}

		// Read digits
		num := 0
		for i < n && isdigit(data[i]) {
			num = num*10 + int(data[i]-'0')
			i++
		}

		return sign * num
	}

	points := make([]point, 0, 32)

	// Parse each line
	for i < n {
		if data[i] == 'p' { // Start of "position"
			p := point{
				x:  number(),
				y:  number(),
				vx: number(),
				vy: number(),
			}
			points = append(points, p)
		} else {
			i++
		}
	}

	return Day10Puzzle{points: points}, nil
}

// Day10 simulates the moving points and finds when they align to form a message.
// Part 1: Returns the message text using OCR.
func Day10(puzzle Day10Puzzle, part1 bool) string {
	if !part1 {
		// Part 2 not implemented yet
		return ""
	}

	// Find the time when the bounding box area is minimized
	// This is when the points form the message

	// Make a copy of points to simulate
	points := make([]point, len(puzzle.points))
	copy(points, puzzle.points)

	minArea := int64(1<<63 - 1) // max int64
	minTime := uint(0)

	// Simulate for a reasonable number of steps
	// The message appears when points are closest together
	for t := uint(0); t < 20000; t++ {
		// Calculate bounding box
		minX, maxX := points[0].x, points[0].x
		minY, maxY := points[0].y, points[0].y

		for i := range len(points) {
			if points[i].x < minX {
				minX = points[i].x
			}
			if points[i].x > maxX {
				maxX = points[i].x
			}
			if points[i].y < minY {
				minY = points[i].y
			}
			if points[i].y > maxY {
				maxY = points[i].y
			}
		}

		width := int64(maxX - minX + 1)
		height := int64(maxY - minY + 1)
		area := width * height

		if area < minArea {
			minArea = area
			minTime = t
		}

		// Move all points
		for i := range len(points) {
			points[i].x += points[i].vx
			points[i].y += points[i].vy
		}
	}

	// Reset points and advance to the optimal time
	copy(points, puzzle.points)
	for range minTime {
		for i := range len(points) {
			points[i].x += points[i].vx
			points[i].y += points[i].vy
		}
	}

	// Calculate final bounding box
	minX, maxX := points[0].x, points[0].x
	minY, maxY := points[0].y, points[0].y

	for i := range len(points) {
		if points[i].x < minX {
			minX = points[i].x
		}
		if points[i].x > maxX {
			maxX = points[i].x
		}
		if points[i].y < minY {
			minY = points[i].y
		}
		if points[i].y > maxY {
			maxY = points[i].y
		}
	}

	// Render as ASCII art
	width := maxX - minX + 1
	height := maxY - minY + 1

	pointSet := make(map[[2]int]bool)
	for i := range len(points) {
		pointSet[[2]int{points[i].x - minX, points[i].y - minY}] = true
	}

	var grid strings.Builder
	for y := range height {
		for x := range width {
			if pointSet[[2]int{x, y}] {
				grid.WriteByte('#')
			} else {
				grid.WriteByte('.')
			}
		}
		grid.WriteByte('\n')
	}

	// Use aococr to parse the message
	charSet := map[rune]bool{'#': true}
	message, _ := aococr.ParseLetters(grid.String(), charSet)
	return message
}
