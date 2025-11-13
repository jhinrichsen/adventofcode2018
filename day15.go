package adventofcode2018

import (
	"fmt"
	"sort"
)

// Day15Puzzle represents the battle map.
type Day15Puzzle struct {
	grid   [][]byte
	units  []unit
	width  int
	height int
}

type unit struct {
	x, y        int
	hp          int
	isElf       bool
	attackPower int
}

type pos struct {
	x, y int
}

// NewDay15 parses the battle map from data.
func NewDay15(data []byte) (Day15Puzzle, error) {
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
		return Day15Puzzle{}, fmt.Errorf("no map data")
	}

	// Build grid and find units
	grid := make([][]byte, len(lines))
	var units []unit
	height := len(lines)
	width := 0

	for y, line := range lines {
		if len(line) > width {
			width = len(line)
		}
		grid[y] = make([]byte, len(line))
		for x := 0; x < len(line); x++ {
			c := line[x]
			switch c {
			case 'E':
				units = append(units, unit{x: x, y: y, hp: 200, isElf: true, attackPower: 3})
				grid[y][x] = '.'
			case 'G':
				units = append(units, unit{x: x, y: y, hp: 200, isElf: false, attackPower: 3})
				grid[y][x] = '.'
			default:
				grid[y][x] = c
			}
		}
	}

	return Day15Puzzle{grid: grid, units: units, width: width, height: height}, nil
}

// Day15 simulates the battle.
// Part 1: Returns the outcome (rounds * remaining HP).
// Part 2: Returns the outcome with minimum elf attack power for no elf deaths.
func Day15(puzzle Day15Puzzle, part1 bool) uint {
	if part1 {
		outcome, _ := simulate(puzzle, 3)
		return outcome
	}

	// Part 2: Find minimum elf attack power for elves to win with no deaths
	for elfPower := 4; ; elfPower++ {
		outcome, allElvesSurvived := simulate(puzzle, elfPower)
		if allElvesSurvived {
			return outcome
		}
	}
}

// simulate runs the battle with given elf attack power.
// Returns (outcome, allElvesSurvived).
func simulate(puzzle Day15Puzzle, elfAttackPower int) (uint, bool) {
	// Make copies and set elf attack power
	units := make([]unit, len(puzzle.units))
	copy(units, puzzle.units)

	// Build occupancy grid
	occupied := make([][]bool, puzzle.height)
	for i := range occupied {
		occupied[i] = make([]bool, puzzle.width)
	}

	initialElfCount := 0
	for i := range units {
		occupied[units[i].y][units[i].x] = true
		if units[i].isElf {
			units[i].attackPower = elfAttackPower
			initialElfCount++
		}
	}

	// Pre-allocate visited array for BFS
	visited := make([][]bool, puzzle.height)
	for i := range visited {
		visited[i] = make([]bool, puzzle.width)
	}

	rounds := 0
	for {
		// Sort units by reading order
		sort.Slice(units, func(i, j int) bool {
			if units[i].y != units[j].y {
				return units[i].y < units[j].y
			}
			return units[i].x < units[j].x
		})

		// Execute each unit's turn
		roundCompleted := true
		for i := range units {
			if units[i].hp <= 0 {
				continue
			}

			// Check if there are any targets
			hasTargets := false
			for j := range units {
				if j != i && units[j].hp > 0 && units[j].isElf != units[i].isElf {
					hasTargets = true
					break
				}
			}
			if !hasTargets {
				roundCompleted = false
				break
			}

			// Move if not in range of any target
			adjIdx := adjacentEnemyIdx(units, i)
			if adjIdx == -1 {
				move(puzzle.grid, &units, i, visited, occupied)
				adjIdx = adjacentEnemyIdx(units, i)
			}

			// Attack if in range
			if adjIdx >= 0 {
				attackAt(&units, i, adjIdx, occupied)
			}
		}

		if !roundCompleted {
			break
		}

		rounds++

		// Remove dead units
		alive := make([]unit, 0, len(units))
		for _, u := range units {
			if u.hp > 0 {
				alive = append(alive, u)
			}
		}
		units = alive
	}

	// Calculate outcome and check elf survival
	totalHP := 0
	elfCount := 0
	for _, u := range units {
		if u.hp > 0 {
			totalHP += u.hp
			if u.isElf {
				elfCount++
			}
		}
	}

	allElvesSurvived := (elfCount == initialElfCount)
	return uint(rounds * totalHP), allElvesSurvived
}

// adjacentEnemyIdx returns index of adjacent enemy to attack, or -1 if none.
// Chooses target with lowest HP, breaking ties by reading order.
func adjacentEnemyIdx(units []unit, idx int) int {
	u := units[idx]
	bestIdx := -1
	bestHP := 201

	// Check in reading order: up, left, right, down
	dirs := []pos{{0, -1}, {-1, 0}, {1, 0}, {0, 1}}
	for _, d := range dirs {
		nx, ny := u.x+d.x, u.y+d.y
		for j := range units {
			if j == idx || units[j].hp <= 0 || units[j].isElf == u.isElf {
				continue
			}
			t := units[j]
			if t.x == nx && t.y == ny {
				if t.hp < bestHP || (t.hp == bestHP && (bestIdx == -1 || readingOrderLess(t.x, t.y, units[bestIdx].x, units[bestIdx].y))) {
					bestHP = t.hp
					bestIdx = j
				}
			}
		}
	}
	return bestIdx
}

func readingOrderLess(x1, y1, x2, y2 int) bool {
	if y1 != y2 {
		return y1 < y2
	}
	return x1 < x2
}

// move moves the unit toward the nearest enemy.
func move(grid [][]byte, units *[]unit, idx int, visited [][]bool, occupied [][]bool) {
	u := &(*units)[idx]

	// Find all in-range positions (adjacent to enemies)
	var inRange []pos
	for j := range *units {
		if j == idx || (*units)[j].hp <= 0 || (*units)[j].isElf == u.isElf {
			continue
		}
		t := (*units)[j]
		for _, d := range []pos{{0, -1}, {-1, 0}, {1, 0}, {0, 1}} {
			nx, ny := t.x+d.x, t.y+d.y
			if ny < 0 || ny >= len(grid) || nx < 0 || nx >= len(grid[ny]) {
				continue
			}
			if grid[ny][nx] == '.' && !occupied[ny][nx] {
				// Check if already in list
				found := false
				for _, p := range inRange {
					if p.x == nx && p.y == ny {
						found = true
						break
					}
				}
				if !found {
					inRange = append(inRange, pos{nx, ny})
				}
			}
		}
	}

	if len(inRange) == 0 {
		return
	}

	// BFS to find nearest reachable in-range position
	nearest := bfs(grid, pos{u.x, u.y}, inRange, visited, occupied)
	if nearest.x == -1 {
		return
	}

	// Update occupancy grid
	occupied[u.y][u.x] = false
	occupied[nearest.y][nearest.x] = true

	// Move one step toward nearest
	u.x = nearest.x
	u.y = nearest.y
}

// bfs finds the next step to move toward any target position.
func bfs(grid [][]byte, start pos, targets []pos, visited [][]bool, occupied [][]bool) pos {
	// Clear visited array
	for y := range visited {
		for x := range visited[y] {
			visited[y][x] = false
		}
	}

	// Check if already at target
	for _, t := range targets {
		if start.x == t.x && start.y == t.y {
			return pos{-1, -1}
		}
	}

	type state struct {
		pos       pos
		dist      int
		firstStep pos
	}

	queue := make([]state, 0, 256)
	queue = append(queue, state{pos: start, dist: 0})
	visited[start.y][start.x] = true

	var reachable []state
	minDist := -1

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		// Stop searching if we've found targets and moved past that distance
		if minDist >= 0 && curr.dist > minDist {
			break
		}

		// Check all neighbors in reading order
		for _, d := range []pos{{0, -1}, {-1, 0}, {1, 0}, {0, 1}} {
			nx, ny := curr.pos.x+d.x, curr.pos.y+d.y

			// Check bounds
			if ny < 0 || ny >= len(grid) || nx < 0 || nx >= len(grid[ny]) {
				continue
			}

			// Check if valid move
			if grid[ny][nx] != '.' || occupied[ny][nx] {
				continue
			}

			// Skip if already visited
			if visited[ny][nx] {
				continue
			}
			visited[ny][nx] = true

			np := pos{nx, ny}

			// Determine first step
			firstStep := curr.firstStep
			if curr.dist == 0 {
				firstStep = np
			}

			newDist := curr.dist + 1

			// Check if this is a target
			isTarget := false
			for _, t := range targets {
				if np.x == t.x && np.y == t.y {
					isTarget = true
					break
				}
			}

			if isTarget {
				if minDist < 0 {
					minDist = newDist
				}
				// Only record targets at minimum distance
				if newDist == minDist {
					reachable = append(reachable, state{pos: np, dist: newDist, firstStep: firstStep})
				}
			} else {
				queue = append(queue, state{pos: np, dist: newDist, firstStep: firstStep})
			}
		}
	}

	if len(reachable) == 0 {
		return pos{-1, -1}
	}

	// Sort by target position (reading order), then by first step (reading order)
	sort.Slice(reachable, func(i, j int) bool {
		ri, rj := reachable[i], reachable[j]
		// Sort by target position
		if ri.pos.y != rj.pos.y {
			return ri.pos.y < rj.pos.y
		}
		if ri.pos.x != rj.pos.x {
			return ri.pos.x < rj.pos.x
		}
		// Sort by first step
		if ri.firstStep.y != rj.firstStep.y {
			return ri.firstStep.y < rj.firstStep.y
		}
		return ri.firstStep.x < rj.firstStep.x
	})

	return reachable[0].firstStep
}

// attackAt attacks the enemy at the given index.
func attackAt(units *[]unit, attackerIdx, targetIdx int, occupied [][]bool) {
	u := (*units)[attackerIdx]
	target := &(*units)[targetIdx]

	target.hp -= u.attackPower
	if target.hp <= 0 {
		occupied[target.y][target.x] = false
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
