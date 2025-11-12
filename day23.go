package adventofcode2018

import (
	"container/heap"
	"math"
	"math/bits"
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
// Part 2: Find the coordinate in range of the most nanobots, closest to origin.
func Day23(puzzle Day23Puzzle, part1 bool) uint {
	if part1 {
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
			if dist <= uint(s.r) {
				count++
			}
		}

		return uint(count)
	}

	// Part 2: Find the coordinate in range of most nanobots
	return findBestPosition(puzzle.bots)
}

// manhattanDist3D calculates Manhattan distance in 3D.
func manhattanDist3D(x1, y1, z1, x2, y2, z2 int) uint {
	return uint(abs(x2-x1) + abs(y2-y1) + abs(z2-z1))
}

// region represents a 3D cube region.
type region struct {
	x, y, z      int  // bottom-left-front corner
	size         uint // length of each side
	inRange      uint // number of bots that can reach this region
	distToOrigin uint // minimum distance from this region to origin
}

// regionQueue implements heap.Interface for priority queue of regions.
type regionQueue []region

func (pq regionQueue) Len() int { return len(pq) }

func (pq regionQueue) Less(i, j int) bool {
	// Priority: more bots in range, then closer to origin, then smaller size
	if pq[i].inRange != pq[j].inRange {
		return pq[i].inRange > pq[j].inRange
	}
	if pq[i].distToOrigin != pq[j].distToOrigin {
		return pq[i].distToOrigin < pq[j].distToOrigin
	}
	return pq[i].size < pq[j].size
}

func (pq regionQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *regionQueue) Push(x interface{}) {
	*pq = append(*pq, x.(region))
}

func (pq *regionQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// distToRegion calculates the minimum Manhattan distance from a nanobot to a region.
func distToRegion(bot nanobot, r region) uint {
	maxX := r.x + int(r.size) - 1
	maxY := r.y + int(r.size) - 1
	maxZ := r.z + int(r.size) - 1

	dx := 0
	if bot.x < r.x {
		dx = r.x - bot.x
	} else if bot.x > maxX {
		dx = bot.x - maxX
	}

	dy := 0
	if bot.y < r.y {
		dy = r.y - bot.y
	} else if bot.y > maxY {
		dy = bot.y - maxY
	}

	dz := 0
	if bot.z < r.z {
		dz = r.z - bot.z
	} else if bot.z > maxZ {
		dz = bot.z - maxZ
	}

	return uint(dx + dy + dz)
}

// countBotsInRange counts how many bots can reach this region.
func countBotsInRange(bots []nanobot, r region) uint {
	count := uint(0)
	for _, bot := range bots {
		if distToRegion(bot, r) <= uint(bot.r) {
			count++
		}
	}
	return count
}

// minDistToOrigin calculates the minimum Manhattan distance from region to origin.
func minDistToOrigin(r region) uint {
	maxX := r.x + int(r.size) - 1
	maxY := r.y + int(r.size) - 1
	maxZ := r.z + int(r.size) - 1

	dx := 0
	if r.x > 0 {
		dx = r.x
	} else if maxX < 0 {
		dx = -maxX
	}

	dy := 0
	if r.y > 0 {
		dy = r.y
	} else if maxY < 0 {
		dy = -maxY
	}

	dz := 0
	if r.z > 0 {
		dz = r.z
	} else if maxZ < 0 {
		dz = -maxZ
	}

	return uint(dx + dy + dz)
}

// findBestPosition finds the position in range of most nanobots, closest to origin.
func findBestPosition(bots []nanobot) uint {
	// Find the bounding box of all nanobots
	minX, maxX := math.MaxInt32, math.MinInt32
	minY, maxY := math.MaxInt32, math.MinInt32
	minZ, maxZ := math.MaxInt32, math.MinInt32

	for _, bot := range bots {
		if bot.x-bot.r < minX {
			minX = bot.x - bot.r
		}
		if bot.x+bot.r > maxX {
			maxX = bot.x + bot.r
		}
		if bot.y-bot.r < minY {
			minY = bot.y - bot.r
		}
		if bot.y+bot.r > maxY {
			maxY = bot.y + bot.r
		}
		if bot.z-bot.r < minZ {
			minZ = bot.z - bot.r
		}
		if bot.z+bot.r > maxZ {
			maxZ = bot.z + bot.r
		}
	}

	// Find the smallest power of 2 that contains the bounding box
	maxDim := maxX - minX
	if d := maxY - minY; d > maxDim {
		maxDim = d
	}
	if d := maxZ - minZ; d > maxDim {
		maxDim = d
	}
	size := uint(1 << bits.Len(uint(maxDim)))

	// Start with the initial region
	initial := region{
		x:    minX,
		y:    minY,
		z:    minZ,
		size: size,
	}
	initial.inRange = countBotsInRange(bots, initial)
	initial.distToOrigin = minDistToOrigin(initial)

	pq := &regionQueue{initial}
	heap.Init(pq)

	for pq.Len() > 0 {
		r := heap.Pop(pq).(region)

		// If we've reached a single point, return it
		if r.size == 1 {
			return r.distToOrigin
		}

		// Split into 8 octants
		newSize := r.size / 2
		offset := int(newSize)
		for dx := 0; dx < 2; dx++ {
			for dy := 0; dy < 2; dy++ {
				for dz := 0; dz < 2; dz++ {
					newRegion := region{
						x:    r.x + dx*offset,
						y:    r.y + dy*offset,
						z:    r.z + dz*offset,
						size: newSize,
					}
					newRegion.inRange = countBotsInRange(bots, newRegion)
					newRegion.distToOrigin = minDistToOrigin(newRegion)
					heap.Push(pq, newRegion)
				}
			}
		}
	}

	return 0
}
