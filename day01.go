package adventofcode2018

func Day01(data []byte, part1 bool) uint {
	frequencies := make([]int, 0, 1024)

	i := 0
	for i < len(data) {
		negative := false
		if data[i] == '+' {
			i++
		} else if data[i] == '-' {
			negative = true
			i++
		}

		num := 0
		for i < len(data) && data[i] >= '0' && data[i] <= '9' {
			num = num*10 + int(data[i]-'0')
			i++
		}

		if negative {
			num = -num
		}
		frequencies = append(frequencies, num)

		if i < len(data) && data[i] == '\n' {
			i++
		}
	}

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
	idx := 0
	for {
		frequency += frequencies[idx]
		if seen[frequency] {
			return uint(frequency)
		}
		seen[frequency] = true
		idx++
		if idx == len(frequencies) {
			idx = 0
		}
	}
}
