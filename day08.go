package adventofcode2018

// NewDay08 parses a whitespace-delimited sequence of positive integers from data.
// It avoids strconv/strings for speed.
func NewDay08(data []byte) ([]int, error) {
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

// Day08Part2 computes the value of the root node according to:
// - If a node has no children, its value is the sum of its metadata entries.
// - If a node has children, metadata entries are 1-based indices referencing child values; sum the referenced child values (ignore out-of-range, allow repeats, 0 references nothing).
func Day08Part2(numbers []int) int {
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
	return valueRec()
}
