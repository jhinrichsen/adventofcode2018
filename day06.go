package adventofcode2018

func Day06(data []byte, part1 bool) uint {
	// Parse coordinates inline from bytes
	coords := make([][2]int, 0, 64)

	i := 0
	for i < len(data) {
		// Parse x coordinate
		x := 0
		for i < len(data) && data[i] >= '0' && data[i] <= '9' {
			x = x*10 + int(data[i]-'0')
			i++
		}

		// Skip ", "
		if i < len(data) && data[i] == ',' {
			i++
		}
		if i < len(data) && data[i] == ' ' {
			i++
		}

		// Parse y coordinate
		y := 0
		for i < len(data) && data[i] >= '0' && data[i] <= '9' {
			y = y*10 + int(data[i]-'0')
			i++
		}

		coords = append(coords, [2]int{x, y})

		// Skip newline
		if i < len(data) && data[i] == '\n' {
			i++
		}
	}

	if part1 {
		return part1Solution(coords)
	}
	return part2Solution(coords, 10000)
}

func part1Solution(coords [][2]int) uint {
	// Find grid dimensions
	maxX, maxY := 0, 0
	for _, c := range coords {
		if c[0] > maxX {
			maxX = c[0]
		}
		if c[1] > maxY {
			maxY = c[1]
		}
	}
	maxX++
	maxY++

	// Create grid
	grid := make([]int, maxX*maxY)
	for i := range grid {
		grid[i] = -1
	}

	// Reusable buffers for distance calculations
	distances := make([]int, len(coords))

	// Fill grid with closest coordinate index
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			minDist := 1<<31 - 1
			minCount := 0
			closestIdx := -1

			// Calculate distances to all coordinates
			for i, c := range coords {
				dx := c[0] - x
				if dx < 0 {
					dx = -dx
				}
				dy := c[1] - y
				if dy < 0 {
					dy = -dy
				}
				dist := dx + dy
				distances[i] = dist

				if dist < minDist {
					minDist = dist
					minCount = 1
					closestIdx = i
				} else if dist == minDist {
					minCount++
				}
			}

			// Only set if single closest point
			if minCount == 1 {
				grid[y*maxX+x] = closestIdx
			} else {
				grid[y*maxX+x] = -2 // tie
			}
		}
	}

	// Find infinite areas (those touching borders)
	infinite := make([]bool, len(coords))
	for x := 0; x < maxX; x++ {
		idx := grid[x]
		if idx >= 0 {
			infinite[idx] = true
		}
		idx = grid[(maxY-1)*maxX+x]
		if idx >= 0 {
			infinite[idx] = true
		}
	}
	for y := 0; y < maxY; y++ {
		idx := grid[y*maxX]
		if idx >= 0 {
			infinite[idx] = true
		}
		idx = grid[y*maxX+maxX-1]
		if idx >= 0 {
			infinite[idx] = true
		}
	}

	// Count areas for each coordinate
	areas := make([]int, len(coords))
	for _, idx := range grid {
		if idx >= 0 && !infinite[idx] {
			areas[idx]++
		}
	}

	// Find maximum area
	maxArea := 0
	for _, area := range areas {
		if area > maxArea {
			maxArea = area
		}
	}

	return uint(maxArea)
}

func part2Solution(coords [][2]int, limit int) uint {
	// Find grid dimensions
	maxX, maxY := 0, 0
	for _, c := range coords {
		if c[0] > maxX {
			maxX = c[0]
		}
		if c[1] > maxY {
			maxY = c[1]
		}
	}
	maxX++
	maxY++

	count := 0
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			totalDist := 0
			for _, c := range coords {
				dx := c[0] - x
				if dx < 0 {
					dx = -dx
				}
				dy := c[1] - y
				if dy < 0 {
					dy = -dy
				}
				totalDist += dx + dy
				if totalDist >= limit {
					break
				}
			}
			if totalDist < limit {
				count++
			}
		}
	}

	return uint(count)
}
