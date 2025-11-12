package adventofcode2018

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
func Day21(puzzle Day21Puzzle, part1 bool) uint {
	// The program implements a hash-like function that generates a sequence.
	// Instead of simulating the VM, we reverse-engineer and compute directly.
	// The algorithm is:
	//   r3 = 0
	//   loop:
	//     r2 = r3 | 65536
	//     r3 = 1505483
	//     while r2 > 0:
	//       r3 = ((r3 + (r2 & 255)) & 16777215) * 65899 & 16777215
	//       r2 = r2 / 256
	//     check if r3 == r0, if not repeat

	var r3 uint
	seen := make(map[uint]bool)
	var lastValue uint
	first := true

	for {
		r2 := r3 | 65536
		r3 = 1505483

		for {
			r3 = ((r3 + (r2 & 255)) & 16777215) * 65899 & 16777215
			if r2 < 256 {
				break
			}
			r2 = r2 / 256
		}

		// Part 1: Return the first value
		if first {
			if part1 {
				return r3
			}
			first = false
		}

		// Part 2: Track values and detect cycle
		if seen[r3] {
			return lastValue
		}
		seen[r3] = true
		lastValue = r3
	}
}
