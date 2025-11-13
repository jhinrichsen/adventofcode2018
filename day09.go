package adventofcode2018

import "slices"

// Day09Puzzle represents the marble game configuration.
type Day09Puzzle struct {
	players    uint
	lastMarble uint
}

// NewDay09 parses the marble game configuration from data.
// Expected format: "N players; last marble is worth M points"
func NewDay09(data []byte) (Day09Puzzle, error) {
	isdigit := func(b byte) bool {
		return b >= '0' && b <= '9'
	}

	n := len(data)
	i := 0

	// number reads the next number from data starting at position i
	number := func() uint {
		// Skip non-digits
		for i < n && !isdigit(data[i]) {
			i++
		}
		// Read digits
		num := uint(0)
		for i < n && isdigit(data[i]) {
			num = num*10 + uint(data[i]-'0')
			i++
		}
		return num
	}

	return Day09Puzzle{
		players:    number(),
		lastMarble: number(),
	}, nil
}

// Day09 simulates the marble game and returns the winning score.
// Part 1: Play the game as described.
// Part 2: Last marble is 100 times larger.
func Day09(puzzle Day09Puzzle, part1 bool) uint {
	lastMarble := puzzle.lastMarble
	if !part1 {
		lastMarble *= 100
	}

	// Initialize scores
	scores := make([]uint, puzzle.players)

	// Pre-allocate marble array
	marbles := make([]marble, lastMarble+1)

	// Initialize marble 0 in a circular list
	marbles[0] = marble{value: 0, prev: 0, next: 0}
	currentIdx := uint(0)
	nextIdx := uint(1)

	// Simulate the game
	for marbleNum := uint(1); marbleNum <= lastMarble; marbleNum++ {
		if marbleNum%23 == 0 {
			// Special rule: marble is multiple of 23
			player := (marbleNum - 1) % puzzle.players
			scores[player] += marbleNum

			// Move 7 positions counter-clockwise
			for range 7 {
				currentIdx = marbles[currentIdx].prev
			}

			// Remove this marble and add to score
			scores[player] += marbles[currentIdx].value
			prevIdx := marbles[currentIdx].prev
			nextIdxTemp := marbles[currentIdx].next
			marbles[prevIdx].next = nextIdxTemp
			marbles[nextIdxTemp].prev = prevIdx
			currentIdx = nextIdxTemp
		} else {
			// Normal rule: insert between 1 and 2 clockwise
			currentIdx = marbles[currentIdx].next

			newIdx := nextIdx
			nextIdx++

			marbles[newIdx] = marble{
				value: marbleNum,
				prev:  currentIdx,
				next:  marbles[currentIdx].next,
			}

			marbles[marbles[currentIdx].next].prev = newIdx
			marbles[currentIdx].next = newIdx
			currentIdx = newIdx
		}
	}

	return slices.Max(scores)
}

// marble represents a node in the circular doubly-linked list using indices.
type marble struct {
	value uint
	prev  uint
	next  uint
}
