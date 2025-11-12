package adventofcode2018

import (
	"fmt"
)

// Day16Puzzle represents the samples and test program.
type Day16Puzzle struct {
	samples []sample
}

type sample struct {
	before      [4]int
	instruction [4]int
	after       [4]int
}

// NewDay16 parses the input data.
func NewDay16(data []byte) (Day16Puzzle, error) {
	n := len(data)
	i := 0

	var samples []sample

	// Parse samples (Before/After format)
	for i < n {
		// Look for "Before: "
		if i+8 <= n && string(data[i:i+8]) == "Before: " {
			// Parse Before line
			var before [4]int
			i += 8
			if i >= n || data[i] != '[' {
				return Day16Puzzle{}, fmt.Errorf("expected '[' after 'Before: '")
			}
			i++
			for j := 0; j < 4; j++ {
				num := 0
				for i < n && data[i] >= '0' && data[i] <= '9' {
					num = num*10 + int(data[i]-'0')
					i++
				}
				before[j] = num
				// Skip comma and space
				if j < 3 {
					if i < n && data[i] == ',' {
						i++
					}
					if i < n && data[i] == ' ' {
						i++
					}
				}
			}
			// Skip closing bracket and newline
			for i < n && (data[i] == ']' || data[i] == '\n' || data[i] == '\r') {
				i++
			}

			// Parse instruction line
			var instruction [4]int
			for j := 0; j < 4; j++ {
				num := 0
				for i < n && data[i] >= '0' && data[i] <= '9' {
					num = num*10 + int(data[i]-'0')
					i++
				}
				instruction[j] = num
				// Skip space
				if i < n && data[i] == ' ' {
					i++
				}
			}
			// Skip newline
			for i < n && (data[i] == '\n' || data[i] == '\r') {
				i++
			}

			// Parse After line
			var after [4]int
			// Skip "After:  ["
			if i+9 <= n && string(data[i:i+9]) == "After:  [" {
				i += 9
			} else if i+8 <= n && string(data[i:i+8]) == "After: [" {
				i += 8
			}
			for j := 0; j < 4; j++ {
				num := 0
				for i < n && data[i] >= '0' && data[i] <= '9' {
					num = num*10 + int(data[i]-'0')
					i++
				}
				after[j] = num
				// Skip comma and space
				if j < 3 {
					if i < n && data[i] == ',' {
						i++
					}
					if i < n && data[i] == ' ' {
						i++
					}
				}
			}
			// Skip closing bracket and newlines
			for i < n && (data[i] == ']' || data[i] == '\n' || data[i] == '\r') {
				i++
			}

			samples = append(samples, sample{before: before, instruction: instruction, after: after})
		} else {
			// Skip other lines (empty lines or test program)
			for i < n && data[i] != '\n' && data[i] != '\r' {
				i++
			}
			for i < n && (data[i] == '\n' || data[i] == '\r') {
				i++
			}
		}
	}

	return Day16Puzzle{samples: samples}, nil
}

// Day16 solves the puzzle.
// Part 1: Count samples that behave like 3+ opcodes.
func Day16(puzzle Day16Puzzle, part1 bool) string {
	if !part1 {
		return ""
	}

	count := 0
	for _, s := range puzzle.samples {
		matches := countMatches(s)
		if matches >= 3 {
			count++
		}
	}

	return fmt.Sprintf("%d", count)
}

// countMatches returns how many opcodes produce the expected result.
func countMatches(s sample) int {
	count := 0
	a := s.instruction[1]
	b := s.instruction[2]
	c := s.instruction[3]

	// Test all 16 opcodes
	opcodes := []func([4]int, int, int, int) [4]int{
		addr, addi, mulr, muli, banr, bani, borr, bori,
		setr, seti, gtir, gtri, gtrr, eqir, eqri, eqrr,
	}

	for _, op := range opcodes {
		result := op(s.before, a, b, c)
		if result == s.after {
			count++
		}
	}

	return count
}

// Opcode implementations

func addr(regs [4]int, a, b, c int) [4]int {
	regs[c] = regs[a] + regs[b]
	return regs
}

func addi(regs [4]int, a, b, c int) [4]int {
	regs[c] = regs[a] + b
	return regs
}

func mulr(regs [4]int, a, b, c int) [4]int {
	regs[c] = regs[a] * regs[b]
	return regs
}

func muli(regs [4]int, a, b, c int) [4]int {
	regs[c] = regs[a] * b
	return regs
}

func banr(regs [4]int, a, b, c int) [4]int {
	regs[c] = regs[a] & regs[b]
	return regs
}

func bani(regs [4]int, a, b, c int) [4]int {
	regs[c] = regs[a] & b
	return regs
}

func borr(regs [4]int, a, b, c int) [4]int {
	regs[c] = regs[a] | regs[b]
	return regs
}

func bori(regs [4]int, a, b, c int) [4]int {
	regs[c] = regs[a] | b
	return regs
}

func setr(regs [4]int, a, b, c int) [4]int {
	regs[c] = regs[a]
	return regs
}

func seti(regs [4]int, a, b, c int) [4]int {
	regs[c] = a
	return regs
}

func gtir(regs [4]int, a, b, c int) [4]int {
	if a > regs[b] {
		regs[c] = 1
	} else {
		regs[c] = 0
	}
	return regs
}

func gtri(regs [4]int, a, b, c int) [4]int {
	if regs[a] > b {
		regs[c] = 1
	} else {
		regs[c] = 0
	}
	return regs
}

func gtrr(regs [4]int, a, b, c int) [4]int {
	if regs[a] > regs[b] {
		regs[c] = 1
	} else {
		regs[c] = 0
	}
	return regs
}

func eqir(regs [4]int, a, b, c int) [4]int {
	if a == regs[b] {
		regs[c] = 1
	} else {
		regs[c] = 0
	}
	return regs
}

func eqri(regs [4]int, a, b, c int) [4]int {
	if regs[a] == b {
		regs[c] = 1
	} else {
		regs[c] = 0
	}
	return regs
}

func eqrr(regs [4]int, a, b, c int) [4]int {
	if regs[a] == regs[b] {
		regs[c] = 1
	} else {
		regs[c] = 0
	}
	return regs
}
