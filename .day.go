package adventofcode2018

// The key words “MUST”, “MUST NOT”, “REQUIRED”, “SHALL”, “SHALL NOT”,
// “SHOULD”, “SHOULD NOT”, “RECOMMENDED”, “NOT RECOMMENDED”, “MAY”, and
// “OPTIONAL” in this document are to be interpreted as described in BCP
// 14 RFC2119 RFC8174 when, and only when, they appear in all capitals,
// as shown here.
//
// This file contains placeholders in the form {{placeholder}}.
//
// This file `day00.go` acts as a template for day NN to aid AI agents in implementing day NN.
// Any implementation for day NN MUST use this template and create a new file `day{{NN}}.go`.
//

// Puzzle acts as a generic name, implementations MUST adjust to concrete puzzle.
// All puzzles have a unique name that MUST be used instead of `Puzzle`, e.g. `NotQuiteLisp', the first puzzle ever (Day 1, 2015).
type Puzzle struct {
}

// NewPuzzle parses the puzzle input in text form into an efficient representation for day N.
// The parser function CAN be left out if the solver (`Day00`) is implemented using puzzle input []byte for maximum performance.
// Implementations CAN use puzzle input as []byte instead of []string for best performance, or if problem solution is grid based instead of line based..
// When using []byte input, parser CAN rely on input in format end_of_line = lf, insert_final_newline = true, trim_trailing_whitespace = true.
// Parser DO NOT HAVE TO cater for malicious input, but use regular error handling e.g. when parsing numbers using strconv.Atoi().
// Parser MUST NOT suppress or ignore errors.
func NewPuzzle(lines []string) (Puzzle, error) {
	return Puzzle{}, nil
}

// Day00 solves the puzzle for day NN, and provides a result.
// The name `Day00` MUST be adjusted to `Day{{NN}}` for the given day NN.
// The identifier for day NN MUST be two digits even for single digit days (`Day07` not `Day7`).
// The name MUST NOT be changed to the puzzle name (`NotQuiteLisp`).
// The priority for implementations is maximum performance in runtime, and minimal garbage collection.
// Common practices for maintainability, ease of reading, clean code, are NOT REQUIRED.
// Implementations SHOULD use high performance parser using puzzle input as []byte and leave out the corresponding parser `NewPuzzle`.
// Implementations SHOULD limit the result types to `uint` (counts, sums), `string` (text) and `int` (solutions that include negative numbers, which is very rare).
func Day00(s Puzzle, part1 bool) (uint, error) {
	return 0, nil
}
