package adventofcode2018

// Day19Puzzle represents the device program.
type Day19Puzzle struct {
	ipReg        int
	instructions []instruction19
}

type instruction19 struct {
	opcode  uint8 // Opcode ID (0-15)
	a, b, c int
}

// Opcode constants
const (
	opAddr uint8 = iota
	opAddi
	opMulr
	opMuli
	opBanr
	opBani
	opBorr
	opBori
	opSetr
	opSeti
	opGtir
	opGtri
	opGtrr
	opEqir
	opEqri
	opEqrr
)

// NewDay19 parses the program.
func NewDay19(data []byte) (Day19Puzzle, error) {
	n := len(data)
	i := 0

	var ipReg int
	var instructions []instruction19

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

		// Convert opcode string to ID
		var opcodeID uint8
		opcodeLen := i - opcodeStart
		if opcodeLen >= 4 {
			// Check first 4 chars to determine opcode
			switch data[opcodeStart] {
			case 'a':
				if data[opcodeStart+1] == 'd' && data[opcodeStart+2] == 'd' {
					if data[opcodeStart+3] == 'r' {
						opcodeID = opAddr
					} else {
						opcodeID = opAddi
					}
				}
			case 'm':
				if data[opcodeStart+1] == 'u' && data[opcodeStart+2] == 'l' {
					if data[opcodeStart+3] == 'r' {
						opcodeID = opMulr
					} else {
						opcodeID = opMuli
					}
				}
			case 'b':
				if data[opcodeStart+1] == 'a' && data[opcodeStart+2] == 'n' {
					if data[opcodeStart+3] == 'r' {
						opcodeID = opBanr
					} else {
						opcodeID = opBani
					}
				} else if data[opcodeStart+1] == 'o' && data[opcodeStart+2] == 'r' {
					if data[opcodeStart+3] == 'r' {
						opcodeID = opBorr
					} else {
						opcodeID = opBori
					}
				}
			case 's':
				if data[opcodeStart+1] == 'e' && data[opcodeStart+2] == 't' {
					if data[opcodeStart+3] == 'r' {
						opcodeID = opSetr
					} else {
						opcodeID = opSeti
					}
				}
			case 'g':
				if data[opcodeStart+1] == 't' {
					if data[opcodeStart+2] == 'i' {
						opcodeID = opGtir
					} else if data[opcodeStart+2] == 'r' {
						if data[opcodeStart+3] == 'i' {
							opcodeID = opGtri
						} else {
							opcodeID = opGtrr
						}
					}
				}
			case 'e':
				if data[opcodeStart+1] == 'q' {
					if data[opcodeStart+2] == 'i' {
						opcodeID = opEqir
					} else if data[opcodeStart+2] == 'r' {
						if data[opcodeStart+3] == 'i' {
							opcodeID = opEqri
						} else {
							opcodeID = opEqrr
						}
					}
				}
			}
		}

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

		instructions = append(instructions, instruction19{opcode: opcodeID, a: a, b: b, c: c})

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
			executeOpcode(inst.opcode, &regs, inst.a, inst.b, inst.c)

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
		executeOpcode(inst.opcode, &regs, inst.a, inst.b, inst.c)

		// Read IP from bound register and increment
		ip = regs[puzzle.ipReg]
		ip++
	}

	return uint(regs[0])
}

// executeOpcode executes an opcode with inline implementation
func executeOpcode(op uint8, regs *[6]int, a, b, c int) {
	switch op {
	case opAddr:
		regs[c] = regs[a] + regs[b]
	case opAddi:
		regs[c] = regs[a] + b
	case opMulr:
		regs[c] = regs[a] * regs[b]
	case opMuli:
		regs[c] = regs[a] * b
	case opBanr:
		regs[c] = regs[a] & regs[b]
	case opBani:
		regs[c] = regs[a] & b
	case opBorr:
		regs[c] = regs[a] | regs[b]
	case opBori:
		regs[c] = regs[a] | b
	case opSetr:
		regs[c] = regs[a]
	case opSeti:
		regs[c] = a
	case opGtir:
		if a > regs[b] {
			regs[c] = 1
		} else {
			regs[c] = 0
		}
	case opGtri:
		if regs[a] > b {
			regs[c] = 1
		} else {
			regs[c] = 0
		}
	case opGtrr:
		if regs[a] > regs[b] {
			regs[c] = 1
		} else {
			regs[c] = 0
		}
	case opEqir:
		if a == regs[b] {
			regs[c] = 1
		} else {
			regs[c] = 0
		}
	case opEqri:
		if regs[a] == b {
			regs[c] = 1
		} else {
			regs[c] = 0
		}
	case opEqrr:
		if regs[a] == regs[b] {
			regs[c] = 1
		} else {
			regs[c] = 0
		}
	}
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
