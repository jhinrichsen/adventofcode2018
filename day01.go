package adventofcode2018

import "strconv"

func NewDay01(lines []string) ([]int, error) {
	frequencies := make([]int, 0, len(lines))
	for _, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		frequencies = append(frequencies, n)
	}
	return frequencies, nil
}

func Day01(frequencies []int, part1 bool) uint {
	if part1 {
		sum := 0
		for _, n := range frequencies {
			sum += n
		}
		return uint(sum)
	}

	seen := make(map[int]bool)
	frequency := 0
	seen[0] = true
	i := 0
	for {
		frequency += frequencies[i]
		if seen[frequency] {
			return uint(frequency)
		}
		seen[frequency] = true
		i++
		if i == len(frequencies) {
			i = 0
		}
	}
}
