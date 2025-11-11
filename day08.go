package adventofcode2018

// NewDay08 parses a whitespace-delimited sequence of positive integers from data.
// It avoids strconv for speed and works directly with []byte.
func NewDay08(data []byte) ([]uint, error) {
	nums := make([]uint, 0, len(data)/2) // rough estimate: each number is ~2 bytes
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
		v := uint(0)
		for i < n {
			b := data[i]
			if b < '0' || b > '9' {
				break
			}
			v = v*10 + uint(b-'0')
			i++
		}
		nums = append(nums, v)
	}
	return nums, nil
}

// Day08 solves day 8 for both parts using iterative (non-recursive) approach.
// Part 1: Sum all metadata entries.
// Part 2: Compute the value of the root node where metadata entries reference child values.
func Day08(numbers []uint, part1 bool) uint {
	idx := 0

	if part1 {
		// Part 1: sum all metadata entries
		sum := uint(0)

		type item struct {
			childrenLeft int
			nMeta        int
		}
		stack := make([]item, 0, 64)

		// Process root
		nChildren := int(numbers[idx])
		idx++
		nMeta := int(numbers[idx])
		idx++
		stack = append(stack, item{nChildren, nMeta})

		for len(stack) > 0 {
			top := &stack[len(stack)-1]

			if top.childrenLeft > 0 {
				// Process next child
				top.childrenLeft--
				nChildren := int(numbers[idx])
				idx++
				nMeta := int(numbers[idx])
				idx++
				stack = append(stack, item{nChildren, nMeta})
			} else {
				// Read metadata using Go 1.25 range over int
				for range top.nMeta {
					sum += numbers[idx]
					idx++
				}
				stack = stack[:len(stack)-1]
			}
		}

		return sum
	}

	// Part 2: compute value using stack-based approach (no recursion)
	type item struct {
		childrenLeft int
		nMeta        int
		childValues  []uint
	}
	stack := make([]item, 0, 64)

	// Process root
	nChildren := int(numbers[idx])
	idx++
	nMeta := int(numbers[idx])
	idx++
	stack = append(stack, item{nChildren, nMeta, make([]uint, 0, nChildren)})

	for len(stack) > 0 {
		top := &stack[len(stack)-1]

		if top.childrenLeft > 0 {
			// Process next child
			top.childrenLeft--
			nChildren := int(numbers[idx])
			idx++
			nMeta := int(numbers[idx])
			idx++
			stack = append(stack, item{nChildren, nMeta, make([]uint, 0, nChildren)})
		} else {
			// Compute value for this node
			var value uint

			if len(top.childValues) == 0 {
				// No children: sum metadata using Go 1.25 range over int
				for range top.nMeta {
					value += numbers[idx]
					idx++
				}
			} else {
				// Has children: use metadata as 1-based indices
				for range top.nMeta {
					ref := numbers[idx]
					idx++
					if ref >= 1 && ref <= uint(len(top.childValues)) {
						value += top.childValues[ref-1]
					}
				}
			}

			stack = stack[:len(stack)-1]

			// Add value to parent if exists
			if len(stack) > 0 {
				stack[len(stack)-1].childValues = append(stack[len(stack)-1].childValues, value)
			} else {
				// This is the root - return its value
				return value
			}
		}
	}

	return 0 // shouldn't reach here
}


