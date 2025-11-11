package adventofcode2018

import "strings"

// NewDay08 parses a whitespace-delimited sequence of positive integers from lines.
// It avoids strconv for speed.
func NewDay08(lines []string) ([]int, error) {
	// Join all lines into single string and parse
	data := []byte(strings.Join(lines, " "))
	nums := make([]int, 0, 1024)
	n := len(data)
	i := 0
	for i < n {
		// skip whitespace
		for i < n {
			b := data[i]
			if b > ' ' { // not space/tab/newline
				break
			}
			i++
		}
		if i >= n {
			break
		}
		// read digits
		v := 0
		for i < n {
			b := data[i]
			if b < '0' || b > '9' {
				break
			}
			v = v*10 + int(b-'0')
			i++
		}
		nums = append(nums, v)
		// continue; next loop skips whitespace
	}
	return nums, nil
}

// Day08 solves day 8 for both parts.
// Part 1: Sum all metadata entries.
// Part 2: Compute the value of the root node where metadata entries reference child values.
func Day08(numbers []int, part1 bool) uint {
	if part1 {
		idx := 0
		sum := 0
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
		return uint(sum)
	}

	// Part 2
	idx := 0
	var valueRec func() int
	valueRec = func() int {
		children := numbers[idx]
		idx++
		nMeta := numbers[idx]
		idx++
		if children == 0 {
			v := 0
			for i := 0; i < nMeta; i++ {
				v += numbers[idx]
				idx++
			}
			return v
		}
		childVals := make([]int, children)
		for c := 0; c < children; c++ {
			childVals[c] = valueRec()
		}
		v := 0
		for i := 0; i < nMeta; i++ {
			ref := numbers[idx]
			idx++
			if ref >= 1 && ref <= len(childVals) {
				v += childVals[ref-1]
			}
		}
		return v
	}
	return uint(valueRec())
}
