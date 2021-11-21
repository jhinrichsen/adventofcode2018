package adventofcode2018

import (
	"strconv"
	"strings"
)

func parseDay08(line string) ([]int, error) {
	parts := strings.Fields(line)
	is := make([]int, len(parts))
	for i := range parts {
		var err error
		is[i], err = strconv.Atoi(parts[i])
		if err != nil {
			return is, err
		}
	}
	return is, nil
}

func Day08Part1(numbers []int) (sum int) {
	idx := 0
	var readRec func()
	readRec = func() {
		children := numbers[idx]
		idx++
		nMeta := numbers[idx]
		idx++
		for children > 0 {
			readRec()
			children--
		}
		metaIdxStart := idx
		metaIdxEnd := metaIdxStart + nMeta
		for i := metaIdxStart; idx < metaIdxEnd; i++ {
			sum += numbers[i]
			idx++
		}
	}
	readRec()
	return
}
