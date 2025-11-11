package adventofcode2018

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

	var players, lastMarble uint
	n := len(data)
	i := 0

	// Parse players
	for i < n && !isdigit(data[i]) {
		i++
	}
	for i < n && isdigit(data[i]) {
		players = players*10 + uint(data[i]-'0')
		i++
	}

	// Parse last marble value
	for i < n && !isdigit(data[i]) {
		i++
	}
	for i < n && isdigit(data[i]) {
		lastMarble = lastMarble*10 + uint(data[i]-'0')
		i++
	}

	return Day09Puzzle{players: players, lastMarble: lastMarble}, nil
}

// Day09 simulates the marble game and returns the winning score.
// Part 1: Play the game as described.
func Day09(puzzle Day09Puzzle, part1 bool) uint {
	// marble represents a node in the circular doubly-linked list.
	type marble struct {
		value uint
		prev  *marble
		next  *marble
	}
	if !part1 {
		// Part 2 not implemented yet
		return 0
	}

	// Initialize scores
	scores := make([]uint, puzzle.players)

	// Create the initial marble (0) in a circular list
	current := &marble{value: 0}
	current.prev = current
	current.next = current

	// Simulate the game
	for marbleNum := uint(1); marbleNum <= puzzle.lastMarble; marbleNum++ {
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

	// Find highest score
	var maxScore uint
	for _, score := range scores {
		if score > maxScore {
			maxScore = score
		}
	}

	return maxScore
}
