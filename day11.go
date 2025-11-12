package adventofcode2018

import "fmt"

// Day11Puzzle represents the grid serial number.
type Day11Puzzle struct {
	serial int
}

// NewDay11 parses the grid serial number from data.
func NewDay11(data []byte) (Day11Puzzle, error) {
	isdigit := func(b byte) bool {
		return b >= '0' && b <= '9'
	}

	n := len(data)
	i := 0

	// number reads the next integer from data
	number := func() int {
		// Skip to next digit
		for i < n && !isdigit(data[i]) {
			i++
		}

		// Read digits
		num := 0
		for i < n && isdigit(data[i]) {
			num = num*10 + int(data[i]-'0')
			i++
		}

		return num
	}

	return Day11Puzzle{serial: number()}, nil
}

// powerLevel calculates the power level for a fuel cell at (x, y).
func powerLevel(x, y, serial int) int {
	rackID := x + 10
	power := rackID * y
	power += serial
	power *= rackID
	power = (power / 100) % 10 // extract hundreds digit
	power -= 5
	return power
}

// Day11 finds the square with the largest total power.
// Part 1: Returns the X,Y coordinate of the 3x3 square's top-left corner.
// Part 2: Returns the X,Y,size identifier of any size square with largest power.
func Day11(puzzle Day11Puzzle, part1 bool) string {
	const gridSize = 300

	// Build summed-area table for O(1) rectangle sum queries
	// SAT[x][y] = sum of all cells from (1,1) to (x,y)
	sat := make([][]int, gridSize+1)
	for x := range gridSize + 1 {
		sat[x] = make([]int, gridSize+1)
	}

	for x := 1; x <= gridSize; x++ {
		for y := 1; y <= gridSize; y++ {
			power := powerLevel(x, y, puzzle.serial)
			sat[x][y] = power + sat[x-1][y] + sat[x][y-1] - sat[x-1][y-1]
		}
	}

	// squareSum calculates sum of square of given size starting at (x, y)
	squareSum := func(x, y, size int) int {
		x2 := x + size - 1
		y2 := y + size - 1
		return sat[x2][y2] - sat[x-1][y2] - sat[x2][y-1] + sat[x-1][y-1]
	}

	maxPower := int(-1 << 31) // min int
	maxX, maxY, maxSize := 0, 0, 0

	if part1 {
		// Part 1: Only check 3x3 squares
		const squareSize = 3
		for x := 1; x <= gridSize-squareSize+1; x++ {
			for y := 1; y <= gridSize-squareSize+1; y++ {
				sum := squareSum(x, y, squareSize)
				if sum > maxPower {
					maxPower = sum
					maxX = x
					maxY = y
				}
			}
		}
		return fmt.Sprintf("%d,%d", maxX, maxY)
	}

	// Part 2: Check all sizes from 1 to 300
	for size := 1; size <= gridSize; size++ {
		for x := 1; x <= gridSize-size+1; x++ {
			for y := 1; y <= gridSize-size+1; y++ {
				sum := squareSum(x, y, size)
				if sum > maxPower {
					maxPower = sum
					maxX = x
					maxY = y
					maxSize = size
				}
			}
		}
	}

	return fmt.Sprintf("%d,%d,%d", maxX, maxY, maxSize)
}
