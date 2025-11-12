package adventofcode2018

import (
	"fmt"
	"sort"
)

// Day15Puzzle represents the battle map.
type Day15Puzzle struct {
	grid  [][]byte
	units []unit
}

type unit struct {
	x, y     int
	hp       int
	isElf    bool
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

	for y, line := range lines {
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

	return Day15Puzzle{grid: grid, units: units}, nil
}

// Day15 simulates the battle.
// Part 1: Returns the outcome (rounds * remaining HP).
// Part 2: Returns the outcome with minimum elf attack power for no elf deaths.
func Day15(puzzle Day15Puzzle, part1 bool) string {
	if part1 {
		outcome, _ := simulate(puzzle, 3)
		return fmt.Sprintf("%d", outcome)
	}

	// Part 2: Find minimum elf attack power for elves to win with no deaths
	for elfPower := 4; ; elfPower++ {
		outcome, allElvesSurvived := simulate(puzzle, elfPower)
		if allElvesSurvived {
			return fmt.Sprintf("%d", outcome)
		}
	}
}

// simulate runs the battle with given elf attack power.
// Returns (outcome, allElvesSurvived).
func simulate(puzzle Day15Puzzle, elfAttackPower int) (int, bool) {
	// Make copies and set elf attack power
	units := make([]unit, len(puzzle.units))
	copy(units, puzzle.units)

	initialElfCount := 0
	for i := range units {
		if units[i].isElf {
			units[i].attackPower = elfAttackPower
			initialElfCount++
		}
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
			targets := findTargets(units, i)
			if len(targets) == 0 {
				roundCompleted = false
				break
			}

			// Move if not in range of any target
			adjacent := adjacentEnemies(units, i)
			if len(adjacent) == 0 {
				move(puzzle.grid, &units, i, targets)
				adjacent = adjacentEnemies(units, i)
			}

			// Attack if in range
			if len(adjacent) > 0 {
				attack(&units, i, adjacent)
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
	return rounds * totalHP, allElvesSurvived
}

// findTargets returns indices of enemy units.
func findTargets(units []unit, idx int) []int {
	var targets []int
	for j := range units {
		if j != idx && units[j].hp > 0 && units[j].isElf != units[idx].isElf {
			targets = append(targets, j)
		}
	}
	return targets
}

// adjacentEnemies returns indices of enemies adjacent to the unit.
func adjacentEnemies(units []unit, idx int) []int {
	u := units[idx]
	var adjacent []int
	for j := range units {
		if j == idx || units[j].hp <= 0 || units[j].isElf == u.isElf {
			continue
		}
		t := units[j]
		if abs(t.x-u.x)+abs(t.y-u.y) == 1 {
			adjacent = append(adjacent, j)
		}
	}
	return adjacent
}

// move moves the unit toward the nearest enemy.
func move(grid [][]byte, units *[]unit, idx int, targets []int) {
	u := &(*units)[idx]

	// Find all in-range positions (adjacent to targets)
	inRangePos := make(map[pos]bool)
	for _, ti := range targets {
		t := (*units)[ti]
		if t.hp <= 0 {
			continue
		}
		for _, d := range []pos{{0, -1}, {-1, 0}, {1, 0}, {0, 1}} {
			np := pos{t.x + d.x, t.y + d.y}
			if grid[np.y][np.x] == '.' && !occupied(*units, np.x, np.y) {
				inRangePos[np] = true
			}
		}
	}

	if len(inRangePos) == 0 {
		return
	}

	// BFS to find nearest reachable in-range position
	nearest := bfs(grid, *units, pos{u.x, u.y}, inRangePos)
	if nearest == nil {
		return
	}

	// Move one step toward nearest
	u.x = nearest.x
	u.y = nearest.y
}

// bfs finds the next step to move toward any target position.
func bfs(grid [][]byte, units []unit, start pos, targets map[pos]bool) *pos {
	if targets[start] {
		return nil // Already at target
	}

	type state struct {
		pos       pos
		dist      int
		firstStep pos
	}

	visited := make(map[pos]bool)
	queue := []state{{pos: start, dist: 0}}
	visited[start] = true

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
			np := pos{curr.pos.x + d.x, curr.pos.y + d.y}

			// Check bounds
			if np.y < 0 || np.y >= len(grid) || np.x < 0 || np.x >= len(grid[np.y]) {
				continue
			}

			// Check if valid move
			if grid[np.y][np.x] != '.' || occupied(units, np.x, np.y) {
				continue
			}

			// Skip if already visited
			if visited[np] {
				continue
			}
			visited[np] = true

			// Determine first step
			firstStep := curr.firstStep
			if curr.dist == 0 {
				firstStep = np
			}

			newDist := curr.dist + 1

			// If this is a target, record it
			if targets[np] {
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
		return nil
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

	result := reachable[0].firstStep
	return &result
}

// occupied checks if a position is occupied by a living unit.
func occupied(units []unit, x, y int) bool {
	for _, u := range units {
		if u.hp > 0 && u.x == x && u.y == y {
			return true
		}
	}
	return false
}

// attack attacks an adjacent enemy with lowest HP.
func attack(units *[]unit, idx int, adjacent []int) {
	u := (*units)[idx]

	if len(adjacent) == 0 {
		return
	}

	// Choose target with lowest HP (ties by reading order)
	sort.Slice(adjacent, func(i, j int) bool {
		ti, tj := (*units)[adjacent[i]], (*units)[adjacent[j]]
		if ti.hp != tj.hp {
			return ti.hp < tj.hp
		}
		if ti.y != tj.y {
			return ti.y < tj.y
		}
		return ti.x < tj.x
	})

	// Deal damage
	target := adjacent[0]
	(*units)[target].hp -= u.attackPower
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
