package adventofcode2018

import (
	"strconv"
	"strings"
)

// Sample represents a before/after observation with an instruction
type Sample struct {
	Before [4]int
	After  [4]int
	Op     [4]int // [opcode, A, B, C]
}

// Instruction represents a single instruction in the test program
type Instruction struct {
	Op [4]int // [opcode, A, B, C]
}

// Day16Puzzle holds parsed samples and test program
type Day16Puzzle struct {
	Samples []Sample
	Program []Instruction
}

// OpFunc is a function that executes an operation
type OpFunc func(reg [4]int, a, b, c int) [4]int

// All 16 operations
var operations = map[string]OpFunc{
	"addr": func(reg [4]int, a, b, c int) [4]int { reg[c] = reg[a] + reg[b]; return reg },
	"addi": func(reg [4]int, a, b, c int) [4]int { reg[c] = reg[a] + b; return reg },
	"mulr": func(reg [4]int, a, b, c int) [4]int { reg[c] = reg[a] * reg[b]; return reg },
	"muli": func(reg [4]int, a, b, c int) [4]int { reg[c] = reg[a] * b; return reg },
	"banr": func(reg [4]int, a, b, c int) [4]int { reg[c] = reg[a] & reg[b]; return reg },
	"bani": func(reg [4]int, a, b, c int) [4]int { reg[c] = reg[a] & b; return reg },
	"borr": func(reg [4]int, a, b, c int) [4]int { reg[c] = reg[a] | reg[b]; return reg },
	"bori": func(reg [4]int, a, b, c int) [4]int { reg[c] = reg[a] | b; return reg },
	"setr": func(reg [4]int, a, b, c int) [4]int { reg[c] = reg[a]; return reg },
	"seti": func(reg [4]int, a, b, c int) [4]int { reg[c] = a; return reg },
	"gtir": func(reg [4]int, a, b, c int) [4]int {
		if a > reg[b] {
			reg[c] = 1
		} else {
			reg[c] = 0
		}
		return reg
	},
	"gtri": func(reg [4]int, a, b, c int) [4]int {
		if reg[a] > b {
			reg[c] = 1
		} else {
			reg[c] = 0
		}
		return reg
	},
	"gtrr": func(reg [4]int, a, b, c int) [4]int {
		if reg[a] > reg[b] {
			reg[c] = 1
		} else {
			reg[c] = 0
		}
		return reg
	},
	"eqir": func(reg [4]int, a, b, c int) [4]int {
		if a == reg[b] {
			reg[c] = 1
		} else {
			reg[c] = 0
		}
		return reg
	},
	"eqri": func(reg [4]int, a, b, c int) [4]int {
		if reg[a] == b {
			reg[c] = 1
		} else {
			reg[c] = 0
		}
		return reg
	},
	"eqrr": func(reg [4]int, a, b, c int) [4]int {
		if reg[a] == reg[b] {
			reg[c] = 1
		} else {
			reg[c] = 0
		}
		return reg
	},
}

// parseInts parses a line containing space-separated integers
func parseInts(line string) ([4]int, error) {
	var result [4]int
	// Extract numbers from line (handles "Before: [1, 2, 3, 4]" and "1 2 3 4" formats)
	fields := strings.FieldsFunc(line, func(r rune) bool {
		return r == ' ' || r == '[' || r == ']' || r == ',' || r == ':'
	})

	numIdx := 0
	for _, field := range fields {
		if field == "" || field == "Before" || field == "After" {
			continue
		}
		num, err := strconv.Atoi(field)
		if err != nil {
			continue
		}
		if numIdx < 4 {
			result[numIdx] = num
			numIdx++
		}
	}

	return result, nil
}

// NewDay16 parses the input lines into samples and test program
func NewDay16(lines []string) (*Day16Puzzle, error) {
	const beforePrefix = "Before:"

	samples := []Sample{}
	var programStart int

	i := 0
	for i < len(lines) {
		line := lines[i]

		// Check if this is a "Before:" line
		if strings.HasPrefix(line, beforePrefix) {
			var s Sample
			var err error

			// Parse Before
			s.Before, err = parseInts(line)
			if err != nil {
				return nil, err
			}

			// Parse instruction (next line)
			i++
			s.Op, err = parseInts(lines[i])
			if err != nil {
				return nil, err
			}

			// Parse After (next line)
			i++
			s.After, err = parseInts(lines[i])
			if err != nil {
				return nil, err
			}

			samples = append(samples, s)
			i++
			continue
		}

		// Check for double blank line (separator between samples and program)
		if len(line) == 0 && i+1 < len(lines) && len(lines[i+1]) == 0 {
			programStart = i + 2
			break
		}

		i++
	}

	// Parse test program
	program := []Instruction{}
	for i := programStart; i < len(lines); i++ {
		line := lines[i]
		if len(line) == 0 {
			continue
		}

		inst := Instruction{}
		var err error
		inst.Op, err = parseInts(line)
		if err != nil {
			return nil, err
		}
		program = append(program, inst)
	}

	return &Day16Puzzle{
		Samples: samples,
		Program: program,
	}, nil
}

// matchesOp checks if a sample matches a given operation
func matchesOp(s Sample, opFunc OpFunc) bool {
	result := opFunc(s.Before, s.Op[1], s.Op[2], s.Op[3])
	return result == s.After
}

// Day16Part1 counts samples that behave like 3 or more opcodes
func Day16Part1(puzzle *Day16Puzzle) uint {
	var count uint

	for _, sample := range puzzle.Samples {
		matches := 0
		for _, opFunc := range operations {
			if matchesOp(sample, opFunc) {
				matches++
			}
		}
		if matches >= 3 {
			count++
		}
	}

	return count
}

// deduceOpcodes determines which opcode number maps to which operation
func deduceOpcodes(samples []Sample) map[int]string {
	// For each opcode number, track which operations it could be
	possible := make(map[int]map[string]bool)
	for i := 0; i < 16; i++ {
		possible[i] = make(map[string]bool)
		for name := range operations {
			possible[i][name] = true
		}
	}

	// Eliminate impossible mappings based on samples
	for _, sample := range samples {
		opcode := sample.Op[0]
		for name, opFunc := range operations {
			if !matchesOp(sample, opFunc) {
				delete(possible[opcode], name)
			}
		}
	}

	// Deduce the mapping by repeatedly finding opcodes with only one possibility
	mapping := make(map[int]string)
	used := make(map[string]bool)

	for len(mapping) < 16 {
		// Find an opcode with only one possibility
		for opcode := 0; opcode < 16; opcode++ {
			if _, ok := mapping[opcode]; ok {
				continue // Already mapped
			}

			// Count possibilities
			var onlyOption string
			count := 0
			for name := range possible[opcode] {
				if !used[name] {
					onlyOption = name
					count++
				}
			}

			if count == 1 {
				mapping[opcode] = onlyOption
				used[onlyOption] = true
			}
		}
	}

	return mapping
}

// Day16Part2 deduces opcodes and executes the test program
func Day16Part2(puzzle *Day16Puzzle) int {
	// Deduce the opcode mapping
	mapping := deduceOpcodes(puzzle.Samples)

	// Execute the test program
	reg := [4]int{0, 0, 0, 0}

	for _, inst := range puzzle.Program {
		opcode := inst.Op[0]
		opName := mapping[opcode]
		opFunc := operations[opName]
		reg = opFunc(reg, inst.Op[1], inst.Op[2], inst.Op[3])
	}

	return reg[0]
}
