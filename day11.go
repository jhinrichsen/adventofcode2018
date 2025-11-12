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

// Day11 finds the 3x3 square with the largest total power.
// Part 1: Returns the X,Y coordinate of the top-left corner.
func Day11(puzzle Day11Puzzle, part1 bool) string {
	if !part1 {
		// Part 2 not implemented yet
		return ""
	}

	const gridSize = 300
	const squareSize = 3

	// Calculate power grid
	grid := make([][]int, gridSize+1)
	for x := range gridSize + 1 {
		grid[x] = make([]int, gridSize+1)
	}

	for x := 1; x <= gridSize; x++ {
		for y := 1; y <= gridSize; y++ {
			grid[x][y] = powerLevel(x, y, puzzle.serial)
		}
	}

	// Find 3x3 square with max power
	maxPower := int(-1 << 31) // min int
	maxX, maxY := 0, 0

	for x := 1; x <= gridSize-squareSize+1; x++ {
		for y := 1; y <= gridSize-squareSize+1; y++ {
			// Calculate sum of 3x3 square starting at (x, y)
			sum := 0
			for dx := range squareSize {
				for dy := range squareSize {
					sum += grid[x+dx][y+dy]
				}
			}

			if sum > maxPower {
				maxPower = sum
				maxX = x
				maxY = y
			}
		}
	}

	return fmt.Sprintf("%d,%d", maxX, maxY)
}
