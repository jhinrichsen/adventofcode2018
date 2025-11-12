package adventofcode2018

import (
	"fmt"
)

// Day22Puzzle represents the cave system parameters.
type Day22Puzzle struct {
	depth   int
	targetX int
	targetY int
}

// NewDay22 parses the depth and target.
func NewDay22(data []byte) (Day22Puzzle, error) {
	n := len(data)
	i := 0

	var depth, targetX, targetY int

	// Parse "depth: "
	for i < n && data[i] != ':' {
		i++
	}
	if i < n {
		i++ // Skip ':'
	}
	for i < n && data[i] == ' ' {
		i++
	}

	// Read depth
	for i < n && data[i] >= '0' && data[i] <= '9' {
		depth = depth*10 + int(data[i]-'0')
		i++
	}

	// Skip to next line
	for i < n && (data[i] == '\n' || data[i] == '\r') {
		i++
	}

	// Parse "target: "
	for i < n && data[i] != ':' {
		i++
	}
	if i < n {
		i++ // Skip ':'
	}
	for i < n && data[i] == ' ' {
		i++
	}

	// Read targetX
	for i < n && data[i] >= '0' && data[i] <= '9' {
		targetX = targetX*10 + int(data[i]-'0')
		i++
	}

	// Skip comma
	if i < n && data[i] == ',' {
		i++
	}

	// Read targetY
	for i < n && data[i] >= '0' && data[i] <= '9' {
		targetY = targetY*10 + int(data[i]-'0')
		i++
	}

	return Day22Puzzle{depth: depth, targetX: targetX, targetY: targetY}, nil
}

// Day22 calculates the cave risk level.
// Part 1: Sum of risk levels in rectangle from (0,0) to target.
func Day22(puzzle Day22Puzzle, part1 bool) string {
	if !part1 {
		return ""
	}

	// Build erosion level cache
	erosion := make(map[pos]int)

	totalRisk := 0
	for y := 0; y <= puzzle.targetY; y++ {
		for x := 0; x <= puzzle.targetX; x++ {
			p := pos{x, y}
			el := getErosionLevel(p, puzzle, erosion)
			risk := el % 3
			totalRisk += risk
		}
	}

	return fmt.Sprintf("%d", totalRisk)
}

// getErosionLevel calculates erosion level for a position.
func getErosionLevel(p pos, puzzle Day22Puzzle, cache map[pos]int) int {
	if v, ok := cache[p]; ok {
		return v
	}

	// Calculate geologic index
	var geoIndex int
	if (p.x == 0 && p.y == 0) || (p.x == puzzle.targetX && p.y == puzzle.targetY) {
		geoIndex = 0
	} else if p.y == 0 {
		geoIndex = p.x * 16807
	} else if p.x == 0 {
		geoIndex = p.y * 48271
	} else {
		el1 := getErosionLevel(pos{p.x - 1, p.y}, puzzle, cache)
		el2 := getErosionLevel(pos{p.x, p.y - 1}, puzzle, cache)
		geoIndex = el1 * el2
	}

	// Calculate erosion level
	erosionLevel := (geoIndex + puzzle.depth) % 20183
	cache[p] = erosionLevel
	return erosionLevel
}
