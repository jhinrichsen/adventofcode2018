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

// pqItem represents an item in the priority queue
type pqItem struct {
	x, y, tool int
	priority   int
	index      int
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
		return day22Part1(puzzle)
	}
	return day22Part2(puzzle)
}

func day22Part1(puzzle Day22Puzzle) string {
	// Build erosion level cache using iterative approach (faster than recursive)
	erosion := make([][]int, puzzle.targetY+1)
	for y := range erosion {
		erosion[y] = make([]int, puzzle.targetX+1)
	}

	for y := 0; y <= puzzle.targetY; y++ {
		for x := 0; x <= puzzle.targetX; x++ {
			var geoIndex int
			if (x == 0 && y == 0) || (x == puzzle.targetX && y == puzzle.targetY) {
				geoIndex = 0
			} else if y == 0 {
				geoIndex = x * 16807
			} else if x == 0 {
				geoIndex = y * 48271
			} else {
				geoIndex = erosion[y][x-1] * erosion[y-1][x]
			}
			erosion[y][x] = (geoIndex + puzzle.depth) % 20183
		}
	}

	totalRisk := 0
	for y := 0; y <= puzzle.targetY; y++ {
		for x := 0; x <= puzzle.targetX; x++ {
			totalRisk += erosion[y][x] % 3
		}
	}

	return fmt.Sprintf("%d", totalRisk)
}

func day22Part2(puzzle Day22Puzzle) string {
	// Estimate bounds for search space (can go beyond target)
	maxX := puzzle.targetX * 2
	maxY := puzzle.targetY * 2
	if maxX < 100 {
		maxX = 100
	}
	if maxY < 100 {
		maxY = 100
	}

	// Pre-compute erosion levels for the entire search space
	erosion := make([][]int, maxY)
	for y := range erosion {
		erosion[y] = make([]int, maxX)
	}

	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			var geoIndex int
			if (x == 0 && y == 0) || (x == puzzle.targetX && y == puzzle.targetY) {
				geoIndex = 0
			} else if y == 0 {
				geoIndex = x * 16807
			} else if x == 0 {
				geoIndex = y * 48271
			} else {
				geoIndex = erosion[y][x-1] * erosion[y-1][x]
			}
			erosion[y][x] = (geoIndex + puzzle.depth) % 20183
		}
	}

	// Use fixed-size arrays instead of maps
	const inf = 1 << 30
	dist := make([][][3]int, maxY)
	visited := make([][][3]bool, maxY)
	for y := range dist {
		dist[y] = make([][3]int, maxX)
		visited[y] = make([][3]bool, maxX)
		for x := range dist[y] {
			dist[y][x] = [3]int{inf, inf, inf}
		}
	}

	// Start at (0,0) with torch equipped
	dist[0][0][torch] = 0

	pq := make(priorityQueue, 0, 1000)
	heap.Init(&pq)
	heap.Push(&pq, &pqItem{x: 0, y: 0, tool: torch, priority: 0})

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*pqItem)
		x, y, tool := item.x, item.y, item.tool
		currentDist := item.priority

		// Check bounds
		if x >= maxX || y >= maxY {
			continue
		}

		// Skip if already visited
		if visited[y][x][tool] {
			continue
		}
		visited[y][x][tool] = true

		// If we reached the target with torch, return the distance
		if x == puzzle.targetX && y == puzzle.targetY && tool == torch {
			return fmt.Sprintf("%d", currentDist)
		}

		// Try moving to adjacent cells with the same tool
		dirs := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

		for _, dir := range dirs {
			nx, ny := x+dir[0], y+dir[1]

			// Check bounds
			if nx < 0 || ny < 0 || nx >= maxX || ny >= maxY {
				continue
			}

			// Skip if already visited
			if visited[ny][nx][tool] {
				continue
			}

			// Check if current tool is valid in the new region
			newRegionType := erosion[ny][nx] % 3
			if !isToolValid(tool, newRegionType) {
				continue
			}

			// Moving takes 1 minute
			newDist := currentDist + 1

			if newDist < dist[ny][nx][tool] {
				dist[ny][nx][tool] = newDist
				heap.Push(&pq, &pqItem{x: nx, y: ny, tool: tool, priority: newDist})
			}
		}

		// Try switching tools at the current position
		currentRegionType := erosion[y][x] % 3
		for newTool := 0; newTool < 3; newTool++ {
			if newTool == tool {
				continue
			}

			// Skip if already visited
			if visited[y][x][newTool] {
				continue
			}

			// Check if the new tool is valid for current region
			if !isToolValid(newTool, currentRegionType) {
				continue
			}

			// Switching tools takes 7 minutes
			newDist := currentDist + 7

			if newDist < dist[y][x][newTool] {
				dist[y][x][newTool] = newDist
				heap.Push(&pq, &pqItem{x: x, y: y, tool: newTool, priority: newDist})
			}
		}
	}

	return "-1" // Should never reach here
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
