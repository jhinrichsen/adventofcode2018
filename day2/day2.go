package day2

func appear1(ID string, n int) bool {
	occurences := make(map[rune]int)
	for _, v := range ID {
		occurences[v]++
	}
	for _, v := range occurences {
		if v == n {
			return true
		}
	}
	return false
}

func appear(boxIDs []string, n int) int {
	count := 0
	for _, ID := range boxIDs {
		if appear1(ID, n) {
			count++
		}
	}
	return count
}

func day2(boxIDs []string) int {
	return appear(boxIDs, 2) * appear(boxIDs, 3)
}
