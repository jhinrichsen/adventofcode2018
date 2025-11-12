package adventofcode2018

import "strconv"

type Day01Puzzle struct {
	frequencies []int
}

func NewDay01(lines []string) (Day01Puzzle, error) {
	frequencies := make([]int, 0, len(lines))
	for _, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			return Day01Puzzle{}, err
		}
		frequencies = append(frequencies, n)
	}
	return Day01Puzzle{frequencies: frequencies}, nil
}

func Day01(puzzle Day01Puzzle, part1 bool) int {
	if part1 {
		sum := 0
		for _, n := range puzzle.frequencies {
			sum += n
		}
		return sum
	}

	seen := make(map[int]bool)
	frequency := 0
	seen[0] = true
	i := 0
	for {
		frequency += puzzle.frequencies[i]
		if seen[frequency] {
			return frequency
		}
		seen[frequency] = true
		i++
		if i == len(puzzle.frequencies) {
			i = 0
		}
	}
}
