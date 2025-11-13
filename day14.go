package adventofcode2018

// Day14Puzzle represents the number of recipes to skip.
type Day14Puzzle struct {
	recipes int
}

// NewDay14 parses the number of recipes from data.
func NewDay14(data []byte) (Day14Puzzle, error) {
	n := len(data)
	i := 0

	isdigit := func(b byte) bool {
		return b >= '0' && b <= '9'
	}

	// Skip whitespace
	for i < n && !isdigit(data[i]) {
		i++
	}

	// Read number
	num := 0
	for i < n && isdigit(data[i]) {
		num = num*10 + int(data[i]-'0')
		i++
	}

	return Day14Puzzle{recipes: num}, nil
}

// Day14 generates chocolate recipe scores.
// Part 1: Returns the 10 scores after the given number of recipes (as a number).
// Part 2: Returns the number of recipes before the input sequence appears.
func Day14(puzzle Day14Puzzle, part1 bool) uint {
	if part1 {
		return day14Part1(puzzle.recipes)
	}
	return day14Part2(puzzle.recipes)
}

func day14Part1(recipes int) uint {
	// Pre-allocate with reasonable capacity
	scores := make([]byte, 2, recipes+20)
	scores[0] = 3
	scores[1] = 7
	elf1, elf2 := 0, 1

	// Generate recipes until we have enough
	target := recipes + 10
	for len(scores) < target {
		// Sum current recipes
		sum := scores[elf1] + scores[elf2]

		// Add new recipes
		if sum >= 10 {
			scores = append(scores, 1, sum-10)
		} else {
			scores = append(scores, sum)
		}

		// Move elves
		elf1 = (elf1 + 1 + int(scores[elf1])) % len(scores)
		elf2 = (elf2 + 1 + int(scores[elf2])) % len(scores)
	}

	// Build result as uint from 10 scores after recipes
	var result uint
	for i := 0; i < 10; i++ {
		result = result*10 + uint(scores[recipes+i])
	}
	return result
}

func day14Part2(recipes int) uint {
	// Convert input number to byte digits
	target := make([]byte, 0, 8)
	n := recipes
	if n == 0 {
		target = []byte{0}
	} else {
		for n > 0 {
			target = append(target, byte(n%10))
			n /= 10
		}
		// Reverse
		for i := 0; i < len(target)/2; i++ {
			target[i], target[len(target)-1-i] = target[len(target)-1-i], target[i]
		}
	}

	// Pre-allocate large capacity to avoid reallocations
	scores := make([]byte, 2, 25000000)
	scores[0] = 3
	scores[1] = 7
	elf1, elf2 := 0, 1

	targetLen := len(target)

	// Generate recipes until we find the sequence
	for {
		// Sum current recipes
		sum := scores[elf1] + scores[elf2]

		// Add new recipes
		if sum >= 10 {
			scores = append(scores, 1, sum-10)
			// Check both positions
			if len(scores) >= targetLen {
				// Check if match at position len(scores)-targetLen-1 (for the first digit added)
				if len(scores) >= targetLen+1 {
					pos := len(scores) - targetLen - 1
					match := true
					for i := 0; i < targetLen; i++ {
						if scores[pos+i] != target[i] {
							match = false
							break
						}
					}
					if match {
						return uint(pos)
					}
				}
				// Check if match at position len(scores)-targetLen (for the second digit added)
				pos := len(scores) - targetLen
				match := true
				for i := 0; i < targetLen; i++ {
					if scores[pos+i] != target[i] {
						match = false
						break
					}
				}
				if match {
					return uint(pos)
				}
			}
		} else {
			scores = append(scores, sum)
			// Check for match
			if len(scores) >= targetLen {
				pos := len(scores) - targetLen
				match := true
				for i := 0; i < targetLen; i++ {
					if scores[pos+i] != target[i] {
						match = false
						break
					}
				}
				if match {
					return uint(pos)
				}
			}
		}

		// Move elves
		elf1 = (elf1 + 1 + int(scores[elf1])) % len(scores)
		elf2 = (elf2 + 1 + int(scores[elf2])) % len(scores)
	}
}
