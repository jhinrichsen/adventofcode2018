package adventofcode2018

// Day19Puzzle represents the device program.
type Day19Puzzle struct {
	ipReg        int
	instructions []instruction
}

type instruction struct {
	opcode  string
	a, b, c int
}

// NewDay19 parses the program.
func NewDay19(data []byte) (Day19Puzzle, error) {
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

	return Day19Puzzle{ipReg: ipReg, instructions: instructions}, nil
}

// Day19 executes the program.
// Part 1: Returns the value in register 0 when the program halts.
// Part 2: Same but with register 0 starting at 1; uses optimization to avoid slow simulation.
func Day19(puzzle Day19Puzzle, part1 bool) uint {
	regs := [6]int{}
	if !part1 {
		// Part 2: Start with register 0 = 1
		regs[0] = 1
	}

	opcodes := map[string]func([6]int, int, int, int) [6]int{
		"addr": addr19, "addi": addi19, "mulr": mulr19, "muli": muli19,
		"banr": banr19, "bani": bani19, "borr": borr19, "bori": bori19,
		"setr": setr19, "seti": seti19,
		"gtir": gtir19, "gtri": gtri19, "gtrr": gtrr19,
		"eqir": eqir19, "eqri": eqri19, "eqrr": eqrr19,
	}

	ip := 0

	// For Part 2, the program computes sum of divisors of a large number.
	// We need to run the initialization phase, then optimize.
	if !part1 {
		// Run initialization phase (first ~50 instructions or until we enter the main loop)
		// The main loop starts when ip is small and we've initialized register 2
		instructionCount := 0
		maxInitInstructions := 100 // Safety limit

		for ip >= 0 && ip < len(puzzle.instructions) && instructionCount < maxInitInstructions {
			// Write IP to bound register
			regs[puzzle.ipReg] = ip

			// Execute instruction
			inst := puzzle.instructions[ip]
			regs = opcodes[inst.opcode](regs, inst.a, inst.b, inst.c)

			// Read IP from bound register and increment
			ip = regs[puzzle.ipReg]
			ip++
			instructionCount++

			// Detect when we've entered the main loop
			// The initialization usually ends when IP jumps to a low number (like 1 or 2)
			// and register 2 has been set to a large value
			if ip <= 3 && regs[2] > 1000 {
				// We've completed initialization; register 2 contains the target number
				// The program computes sum of divisors of register 2
				return sumOfDivisors(uint(regs[2]))
			}
		}

		// If we didn't detect the pattern, fall back to full simulation
		// (shouldn't happen with typical AoC 2018 Day 19 inputs)
	}

	// Part 1 or fallback: Run full simulation
	for ip >= 0 && ip < len(puzzle.instructions) {
		// Write IP to bound register
		regs[puzzle.ipReg] = ip

		// Execute instruction
		inst := puzzle.instructions[ip]
		regs = opcodes[inst.opcode](regs, inst.a, inst.b, inst.c)

		// Read IP from bound register and increment
		ip = regs[puzzle.ipReg]
		ip++
	}

	return uint(regs[0])
}

// Opcode implementations for 6 registers

func addr19(regs [6]int, a, b, c int) [6]int {
	regs[c] = regs[a] + regs[b]
	return regs
}

func addi19(regs [6]int, a, b, c int) [6]int {
	regs[c] = regs[a] + b
	return regs
}

func mulr19(regs [6]int, a, b, c int) [6]int {
	regs[c] = regs[a] * regs[b]
	return regs
}

func muli19(regs [6]int, a, b, c int) [6]int {
	regs[c] = regs[a] * b
	return regs
}

func banr19(regs [6]int, a, b, c int) [6]int {
	regs[c] = regs[a] & regs[b]
	return regs
}

func bani19(regs [6]int, a, b, c int) [6]int {
	regs[c] = regs[a] & b
	return regs
}

func borr19(regs [6]int, a, b, c int) [6]int {
	regs[c] = regs[a] | regs[b]
	return regs
}

func bori19(regs [6]int, a, b, c int) [6]int {
	regs[c] = regs[a] | b
	return regs
}

func setr19(regs [6]int, a, b, c int) [6]int {
	regs[c] = regs[a]
	return regs
}

func seti19(regs [6]int, a, b, c int) [6]int {
	regs[c] = a
	return regs
}

func gtir19(regs [6]int, a, b, c int) [6]int {
	if a > regs[b] {
		regs[c] = 1
	} else {
		regs[c] = 0
	}
	return regs
}

func gtri19(regs [6]int, a, b, c int) [6]int {
	if regs[a] > b {
		regs[c] = 1
	} else {
		regs[c] = 0
	}
	return regs
}

func gtrr19(regs [6]int, a, b, c int) [6]int {
	if regs[a] > regs[b] {
		regs[c] = 1
	} else {
		regs[c] = 0
	}
	return regs
}

func eqir19(regs [6]int, a, b, c int) [6]int {
	if a == regs[b] {
		regs[c] = 1
	} else {
		regs[c] = 0
	}
	return regs
}

func eqri19(regs [6]int, a, b, c int) [6]int {
	if regs[a] == b {
		regs[c] = 1
	} else {
		regs[c] = 0
	}
	return regs
}

func eqrr19(regs [6]int, a, b, c int) [6]int {
	if regs[a] == regs[b] {
		regs[c] = 1
	} else {
		regs[c] = 0
	}
	return regs
}

// sumOfDivisors calculates the sum of all divisors of n (including 1 and n).
// Uses an O(√n) algorithm instead of brute force O(n).
func sumOfDivisors(n uint) uint {
	var sum uint
	// Only iterate up to sqrt(n)
	for i := uint(1); i*i <= n; i++ {
		if n%i == 0 {
			sum += i
			// Add the paired divisor if it's different
			if i != n/i {
				sum += n / i
			}
		}
	}
	return sum
}
