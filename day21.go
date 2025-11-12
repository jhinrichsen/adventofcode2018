package adventofcode2018

import (
	"fmt"
)

// Day21Puzzle represents the device program.
type Day21Puzzle struct {
	ipReg        int
	instructions []instruction
}

// NewDay21 parses the program (same format as day 19).
func NewDay21(data []byte) (Day21Puzzle, error) {
	n := len(data)
	i := 0

	var ipReg int
	var instructions []instruction

	// Parse #ip line
	if i+4 <= n && string(data[i:i+4]) == "#ip " {
		i += 4
		ipReg = 0
		for i < n && data[i] >= '0' && data[i] <= '9' {
			ipReg = ipReg*10 + int(data[i]-'0')
			i++
		}
		// Skip newline
		for i < n && (data[i] == '\n' || data[i] == '\r') {
			i++
		}
	}

	// Parse instructions
	for i < n {
		// Read opcode name
		opcodeStart := i
		for i < n && data[i] >= 'a' && data[i] <= 'z' {
			i++
		}
		if i == opcodeStart {
			break
		}
		opcode := string(data[opcodeStart:i])

		// Skip space
		if i < n && data[i] == ' ' {
			i++
		}

		// Read A
		a := 0
		for i < n && data[i] >= '0' && data[i] <= '9' {
			a = a*10 + int(data[i]-'0')
			i++
		}
		if i < n && data[i] == ' ' {
			i++
		}

		// Read B
		b := 0
		for i < n && data[i] >= '0' && data[i] <= '9' {
			b = b*10 + int(data[i]-'0')
			i++
		}
		if i < n && data[i] == ' ' {
			i++
		}

		// Read C
		c := 0
		for i < n && data[i] >= '0' && data[i] <= '9' {
			c = c*10 + int(data[i]-'0')
			i++
		}

		instructions = append(instructions, instruction{opcode: opcode, a: a, b: b, c: c})

		// Skip newline
		for i < n && (data[i] == '\n' || data[i] == '\r') {
			i++
		}
	}

	return Day21Puzzle{ipReg: ipReg, instructions: instructions}, nil
}

// Day21 finds the value for register 0 that causes the program to halt.
// Part 1: Find the value that halts with fewest instructions (first value checked).
// Part 2: Find the value that halts with most instructions (last unique value before cycle).
func Day21(puzzle Day21Puzzle, part1 bool) string {
	opcodes := map[string]func([6]int, int, int, int) [6]int{
		"addr": addr19, "addi": addi19, "mulr": mulr19, "muli": muli19,
		"banr": banr19, "bani": bani19, "borr": borr19, "bori": bori19,
		"setr": setr19, "seti": seti19,
		"gtir": gtir19, "gtri": gtri19, "gtrr": gtrr19,
		"eqir": eqir19, "eqri": eqri19, "eqrr": eqrr19,
	}

	regs := [6]int{}
	ip := 0

	// Find the instruction that checks register 0
	// Typically it's an eqrr comparing some register with register 0
	checkIP := -1
	checkReg := -1
	for i, inst := range puzzle.instructions {
		if inst.opcode == "eqrr" && (inst.a == 0 || inst.b == 0) {
			checkIP = i
			if inst.a == 0 {
				checkReg = inst.b
			} else {
				checkReg = inst.a
			}
			break
		}
	}

	// Part 1: Return the first value
	if part1 {
		for ip >= 0 && ip < len(puzzle.instructions) {
			// If we're at the check instruction, return the value
			if ip == checkIP {
				return fmt.Sprintf("%d", regs[checkReg])
			}

			// Write IP to bound register
			regs[puzzle.ipReg] = ip

			// Execute instruction
			inst := puzzle.instructions[ip]
			regs = opcodes[inst.opcode](regs, inst.a, inst.b, inst.c)

			// Read IP from bound register and increment
			ip = regs[puzzle.ipReg]
			ip++
		}
		return "0"
	}

	// Part 2: Track all values and return the last one before cycle repeats
	seen := make(map[int]bool)
	lastValue := 0

	for ip >= 0 && ip < len(puzzle.instructions) {
		// If we're at the check instruction
		if ip == checkIP {
			val := regs[checkReg]
			if seen[val] {
				// We've seen this value before, cycle detected
				// Return the last unique value
				return fmt.Sprintf("%d", lastValue)
			}
			seen[val] = true
			lastValue = val
		}

		// Write IP to bound register
		regs[puzzle.ipReg] = ip

		// Execute instruction
		inst := puzzle.instructions[ip]
		regs = opcodes[inst.opcode](regs, inst.a, inst.b, inst.c)

		// Read IP from bound register and increment
		ip = regs[puzzle.ipReg]
		ip++
	}

	return fmt.Sprintf("%d", lastValue)
}
