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
	// marble represents a node in the circular doubly-linked list.
	type marble struct {
		value uint
		prev  *marble
		next  *marble
	}

	lastMarble := puzzle.lastMarble
	if !part1 {
		lastMarble *= 100
	}

	// Initialize scores
	scores := make([]uint, puzzle.players)

	// Create the initial marble (0) in a circular list
	current := &marble{value: 0}
	current.prev = current
	current.next = current

	// Simulate the game
	for marbleNum := uint(1); marbleNum <= lastMarble; marbleNum++ {
		if marbleNum%23 == 0 {
			// Special rule: marble is multiple of 23
			player := (marbleNum - 1) % puzzle.players
			scores[player] += marbleNum

			// Move 7 positions counter-clockwise
			for range 7 {
				current = current.prev
			}

			// Remove this marble and add to score
			scores[player] += current.value
			current.prev.next = current.next
			current.next.prev = current.prev
			current = current.next
		} else {
			// Normal rule: insert between 1 and 2 clockwise
			current = current.next
			newMarble := &marble{
				value: marbleNum,
				prev:  current,
				next:  current.next,
			}
			current.next.prev = newMarble
			current.next = newMarble
			current = newMarble
		}
	}

	return slices.Max(scores)
}
