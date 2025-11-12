package adventofcode2018

import (
	"container/heap"
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

// Tool constants
const (
	torch        = 0
	climbingGear = 1
	neither      = 2
)

// state represents a position with equipped tool
type state struct {
	x, y int
	tool int
}

// pqItem represents an item in the priority queue
type pqItem struct {
	st       state
	priority int
	index    int
}

// priorityQueue implements heap.Interface
type priorityQueue []*pqItem

func (pq priorityQueue) Len() int           { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool { return pq[i].priority < pq[j].priority }
func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*pqItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// Day22 calculates the cave risk level.
// Part 1: Sum of risk levels in rectangle from (0,0) to target.
// Part 2: Shortest time to reach target with tool constraints.
func Day22(puzzle Day22Puzzle, part1 bool) string {
	if part1 {
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

	// Part 2: Find shortest path with tool switching
	minTime := findShortestPath(puzzle)
	return fmt.Sprintf("%d", minTime)
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

// getRegionType returns the region type (0=rocky, 1=wet, 2=narrow)
func getRegionType(x, y int, puzzle Day22Puzzle, erosion map[pos]int) int {
	el := getErosionLevel(pos{x, y}, puzzle, erosion)
	return el % 3
}

// isToolValid checks if a tool can be used in a region type
func isToolValid(tool, regionType int) bool {
	// Rocky (0): torch or climbing gear
	// Wet (1): climbing gear or neither
	// Narrow (2): torch or neither
	switch regionType {
	case 0: // rocky
		return tool == torch || tool == climbingGear
	case 1: // wet
		return tool == climbingGear || tool == neither
	case 2: // narrow
		return tool == torch || tool == neither
	}
	return false
}

// findShortestPath uses Dijkstra's algorithm to find the shortest path
func findShortestPath(puzzle Day22Puzzle) int {
	erosion := make(map[pos]int)
	dist := make(map[state]int)
	visited := make(map[state]bool)

	// Start at (0,0) with torch equipped
	start := state{0, 0, torch}
	dist[start] = 0

	pq := make(priorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &pqItem{st: start, priority: 0})

	// Target state: at target position with torch equipped
	target := state{puzzle.targetX, puzzle.targetY, torch}

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*pqItem)
		current := item.st
		currentDist := item.priority

		// Skip if already visited
		if visited[current] {
			continue
		}
		visited[current] = true

		// If we reached the target, return the distance
		if current == target {
			return currentDist
		}

		// Try moving to adjacent cells with the same tool
		dirs := []struct{ dx, dy int }{
			{0, 1}, {1, 0}, {0, -1}, {-1, 0},
		}

		for _, dir := range dirs {
			nx, ny := current.x+dir.dx, current.y+dir.dy

			// Check bounds (can't go to negative coordinates)
			if nx < 0 || ny < 0 {
				continue
			}

			newState := state{nx, ny, current.tool}

			// Skip if already visited
			if visited[newState] {
				continue
			}

			// Check if current tool is valid in the new region
			newRegionType := getRegionType(nx, ny, puzzle, erosion)
			if !isToolValid(current.tool, newRegionType) {
				continue
			}

			// Moving takes 1 minute
			newDist := currentDist + 1

			if d, ok := dist[newState]; !ok || newDist < d {
				dist[newState] = newDist
				heap.Push(&pq, &pqItem{st: newState, priority: newDist})
			}
		}

		// Try switching tools at the current position
		currentRegionType := getRegionType(current.x, current.y, puzzle, erosion)
		for tool := 0; tool < 3; tool++ {
			if tool == current.tool {
				continue
			}

			newState := state{current.x, current.y, tool}

			// Skip if already visited
			if visited[newState] {
				continue
			}

			// Check if the new tool is valid for current region
			if !isToolValid(tool, currentRegionType) {
				continue
			}

			// Switching tools takes 7 minutes
			newDist := currentDist + 7

			if d, ok := dist[newState]; !ok || newDist < d {
				dist[newState] = newDist
				heap.Push(&pq, &pqItem{st: newState, priority: newDist})
			}
		}
	}

	return -1 // Should never reach here
}
