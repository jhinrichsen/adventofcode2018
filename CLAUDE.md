# Claude Code Instructions for Advent of Code 2018

## 🚨 CRITICAL: Solution Confidentiality

**Solutions MUST NEVER appear anywhere except in unit test `want` values.**

Prohibited locations:
- ❌ Git commit messages
- ❌ Pull request titles/descriptions
- ❌ Code comments
- ❌ Documentation files
- ❌ Console output or logs

**ONLY** permitted location:
- ✅ Unit test files: `TestDayNNPart1` and `TestDayNNPart2` (the `want` parameter)

**Example commit messages:**
```
✅ Good: feat(day17): implement water retention logic for part 2
❌ Bad:  feat(day17): solve part 2 - answer is 30410
```

---

## Project Overview

Advent of Code is a yearly programming contest (Dec 1-25) with daily two-part puzzles.
- Each puzzle has examples (same for all users) and unique puzzle input (user-specific)
- Part 2 unlocks only after Part 1 is verified correct
- Priority: **correctness first, then performance**

---

## Implementation Workflow

### Step 1: Implement Part 1 Examples

1. Copy `.day.go` template to `day{NN}.go`
2. Copy `.day_test.go` template to `day{NN}_test.go`
3. Remove all comments from both files
4. Rename all identifiers from `00` to `{NN}` (zero-padded, e.g., `Day08`)
5. Implement all examples from puzzle description
6. **Definition of Done**: All example tests pass
7. **Proceed automatically** to Part 1 (no user confirmation needed)

### Step 2: Implement Part 1

1. Create test case using puzzle input from `testdata/day{NN}.txt`
2. First run will fail with unknown answer: `want=?, got=X`
3. Update test with `want=X` for second run
4. **Definition of Done**: Test passes
5. **STOP HERE** - wait for user to verify answer on website before Part 2

### Step 3: Implement Part 2 Examples

1. User will provide Part 2 description (requires login, so you cannot fetch it)
2. Implement all Part 2 examples
3. **Definition of Done**: All example tests pass
4. **Proceed automatically** to Part 2 implementation

### Step 4: Implement Part 2

1. Follow same procedure as Part 1
2. **STOP** after test passes, wait for user verification

---

## Code Structure & Naming

**File organization:**
```
dayXX.go                    # Implementation
dayXX_test.go              # Tests and benchmarks
testdata/dayXX.txt         # Puzzle input
testdata/dayXX_example.txt # Example input (optional)
```

**Naming conventions:**
- Use two-digit zero-padded day numbers: `Day08` (not `Day8`)
- All code in root package (no subpackages)
- Files: `day08.go`, `day08_test.go`
- Functions: `NewDay08()`, `Day08()`
- Tests: `TestDay08Part1()`, `TestDay08Part2()`
- Benchmarks: `BenchmarkDay08Part1()`, `BenchmarkDay08Part2()`

---

## Parser & Solver Patterns

Choose the pattern that best fits the input format:

### Pattern 1: Line-Oriented Parser (Preferred for structured line-based input)

**When to use:** Input has distinct line structure (headers, groups, structured lines)

**Parser signature:**
```go
func NewDayXX(lines []string) (DayXXPuzzle, error)
```

**Test pattern:**
```go
func TestDayXXPart1(t *testing.T) {
    lines := linesFromFilename(t, filename(XX))
    puzzle, err := NewDayXX(lines)
    if err != nil {
        t.Fatal(err)
    }
    got := DayXX(puzzle, true)
    const want = 12345
    if got != want {
        t.Errorf("want %d but got %d", want, got)
    }
}

func BenchmarkDayXXPart1(b *testing.B) {
    lines := linesFromFilename(b, filename(XX))
    for i := 0; i < b.N; i++ {
        puzzle, _ := NewDayXX(lines)
        DayXX(puzzle, true)
    }
}
```

**Examples:** Day 3, 4, 6, 7

### Pattern 2: Byte-Level Parser (For compact formats)

**When to use:** Parsing coordinates, numbers, or formats requiring byte iteration

**Parser signature:**
```go
func NewDayXX(data []byte) (DayXXPuzzle, error)
```

**Test pattern:**
```go
func TestDayXXPart1(t *testing.T) {
    testWithParserBytes(t, XX, filename, true, NewDayXX, DayXX, 12345)
}

func BenchmarkDayXXPart1(b *testing.B) {
    benchWithParserBytes(b, XX, true, NewDayXX, DayXX)
}
```

**Examples:** Day 16, 17, 18, 19, 20, 21, 22, 23, 24, 25

### Pattern 3: Direct Solver (Rare, for simple problems)

**When to use:** Parsing and solving are tightly coupled

**Solver signature:**
```go
func DayXX(data []byte, part1 bool) (string, error)
```

---

## Solver Function Implementation

**Signature:**
```go
// DayXX solves the puzzle.
// Part 1: [Description]
// Part 2: [Description]
func DayXX(puzzle DayXXPuzzle, part1 bool) uint {
    if !part1 {
        return 0  // or implement Part 2
    }

    // Solution logic
    return result
}
```

**Return type:**
- Use `uint` for numeric answers (counts, sums, coordinates)
- Use `string` only for non-numeric answers
- Never return negative values for counts

---

## Parsing Best Practices

### 1. Custom Parsers (No Heavy Libraries)

- Write purpose-built parsers tailored to input format
- Avoid `strings.Fields`, `strings.Split`, or regex overhead
- Parse in single pass, building data on-the-fly
- Example: For CSV, iterate char-by-char instead of splitting into substrings

### 2. Minimize Allocations

- Pre-allocate slices when size is known: `make([]T, 0, capacity)`
- Parse in-place where possible
- Skip `fmt.Sscanf` or regex for structured text
- Use byte/rune inspection for manual parsing

### 3. Optimal Data Structures

- Fixed dimensions → Use arrays or flat 1D slices instead of maps
- Small key range (0-9) → Use slice indexed by key instead of `map[int]int`
- Example: `[]int` of length 9 is ~96% faster than `map[uint8]int` with keys 0-8

### 4. Reuse Buffers

- Allocate temporary slices outside loops
- Reset with `temp = temp[:0]` instead of reallocating
- Reuse `bytes.Buffer` for string building
- Avoid constant allocate/free cycles

### 5. Error Handling

- Handle file I/O and parsing errors properly
- Return errors from `NewDayNN`, don't use `log.Fatal` or `os.Exit`
- Can panic for truly unreachable conditions
- Remove all logging/debug output before completion

---

## Performance Requirements

**Target:** Run in under 1 second, minimal memory usage

### Memory Optimization

1. **Minimize Allocations**
   - Strive for zero allocations in hot loops
   - Use `go test -benchmem` to identify allocation sources
   - Reuse slices by resetting length instead of creating new ones

2. **Prefer Arrays/Slices over Maps**
   - Arrays are cache-friendly and have no hashing overhead
   - Use maps only for truly dynamic/unbounded data
   - For moderate data, sorted slice + binary search can outperform maps

3. **In-Place Operations**
   - Update data in-place when correctness allows
   - Use double-buffering (toggle between two arrays) for simulations
   - Copy to reuse old slice memory when filtering

4. **Algorithmic Efficiency**
   - Aim for O(n) or O(n log n) solutions
   - Use Go's optimized `sort.Slice` or `slices.Sort`
   - Prefer iterative over recursive approaches (avoid stack overhead)

### Coding Practices

1. **Simple Loops**
   - Use basic loops and conditionals (compile to efficient code)
   - Avoid goroutines/channels (overhead outweighs benefit)
   - Avoid reflective or interface-heavy constructs in hot paths
   - Use `break`/`continue` to skip unnecessary work

2. **Leverage Compiler Optimizations**
   - Write clear code that compiler can optimize
   - Use built-in functions: `copy()`, `bytes.Equal()` (assembly-optimized)
   - Modern Go inlines functions and eliminates bounds checks
   - Simple generic functions will be inlined

3. **Standard Library Only**
   - Use Go 1.24 features (check `go.mod`)
   - New packages: `slices`, `maps`, `min()`, `max()`
   - No third-party dependencies
   - No recursion (risk of stack overflow)

---

## Testing Requirements

### Example Tests

```go
func TestDayNNPart1Example(t *testing.T) {
    lines := linesFromFilename(t, exampleFilename(NN))
    data := NewDayNN(lines)
    got := DayNN(data, true)
    const want = 26  // From puzzle description
    if got != want {
        t.Errorf("Part1 example: got %d, want %d", got, want)
    }
}
```

**Requirements:**
- Test ALL examples from puzzle description
- Use separate test functions or sub-tests for multiple examples
- Examples can use `testdata/dayNN_example.txt` or inline strings

### Puzzle Input Tests

```go
func TestDayNNPart1(t *testing.T) {
    lines := linesFromFilename(t, filename(NN))
    data := NewDayNN(lines)
    got := DayNN(data, true)
    const want = 1673  // Verified answer
    if got != want {
        t.Errorf("Part1: got %d, want %d", got, want)
    }
}
```

**Requirements:**
- Include verified answers as `want` values
- No printing on success
- Use `t.Fatalf` or `t.Errorf` for failures
- Run `go test ./...` to verify all pass

---

## Benchmarking

```go
func BenchmarkDayNNPart1(b *testing.B) {
    raw := linesFromFilename(b, filename(NN))
    b.ResetTimer()
    for range b.N {
        data := NewDayNN(raw)
        _ = DayNN(data, true)
    }
}
```

**Requirements:**
- Include parsing in benchmark (it's part of performance)
- Call `b.ResetTimer()` after loading input
- Run `go test -bench . -benchmem` to check allocations
- Target: Low B/op (bytes/op) and minimal allocs/op

---

## Git Commit Guidelines

**Format:** `<type>(<scope>): <description>`

**Types:** feat, fix, refactor, test, docs, perf

**Scope:** dayXX (e.g., day01, day25)

**Good examples:**
```
feat(day13): add part 2 cart collision handling
refactor(day16): optimize opcode matching logic
fix(day17): correct water settling check after flow
perf(day21): reduce allocations in VM loop
```

**Bad examples:**
```
❌ feat(day17): solve part 2 - answer is 30410
❌ Add solution for day 13 part 1 (42)
```

**IMPORTANT: Before committing, ALWAYS run `go fmt ./...` to ensure all code is properly formatted.**

---

## Common Pitfalls to Avoid

### Parsing
- ❌ Using `strings.Split(string(data), "\n")` in byte parser (use `[]string` parameter)
- ❌ Not handling negative numbers in byte-level parsing
- ❌ Forgetting to skip delimiters
- ❌ Off-by-one errors in index iteration
- ❌ Converting `[]byte` to `string` unnecessarily

### Algorithm
- ❌ Not handling edge cases (empty input, single item)
- ❌ Infinite loops in simulation
- ❌ Incorrect tie-breaking in sorting
- ❌ Not checking state after recursive/iterative calls

### Testing
- ❌ Hardcoding wrong answers
- ❌ Not testing all examples before puzzle input
- ❌ **Including solutions in commit messages**
- ❌ Not updating tests after bug fixes

### Code Quality
- ❌ Using recursion (stack overflow risk)
- ❌ Using `panic()` instead of error handling
- ❌ Looking up solutions online
- ❌ Hardcoding results: `return 5347`
- ❌ Global state or mutable package variables
- ❌ Creating helper packages (keep each day self-contained)

---

## Common Algorithm Patterns

| Problem Type | Algorithm | Example Days |
|--------------|-----------|--------------|
| Pathfinding | BFS, Dijkstra | Day 15, 20, 22 |
| Graph connectivity | Union-Find, DFS | Day 25 |
| Simulation | State machine, cellular automaton | Day 13, 15, 17, 18, 24 |
| Pattern matching | Manual parsing | Day 20 |
| Virtual machine | Interpreter, instruction set | Day 16, 19, 21 |
| Geometry | Manhattan distance, grid ops | Day 6, 10, 17, 22, 23 |

---

## Common Helper Patterns

### Position Types
```go
type pos struct{ x, y int }
type pos3D struct{ x, y, z int }
type point4D struct{ x, y, z, w int }
```

### Helper Functions
```go
func abs(x int) int {
    if x < 0 { return -x }
    return x
}

func manhattanDist(x1, y1, x2, y2 int) int {
    return abs(x2-x1) + abs(y2-y1)
}
```

### Grid Operations
```go
func copyGrid(grid [][]byte) [][]byte {
    copy := make([][]byte, len(grid))
    for y := range grid {
        copy[y] = make([]byte, len(grid[y]))
        copy(copy[y], grid[y])
    }
    return copy
}
```

---

## Environment

- **Go version:** 1.24 (NOT 1.25)
- **Platform:** Linux
- **Test helpers:** See `test_helper.go` and `input_test.go`
- **Working directory:** `/home/user/adventofcode2018`
- **Branch:** `claude/contributi-feature-011CV47oeas4QzzfzchWHmy5`

---

## Final Checklist

Before considering implementation complete:

- [ ] All example tests pass
- [ ] Puzzle input test passes with verified answer
- [ ] No debug prints or log statements
- [ ] Code is `go fmt` formatted
- [ ] `go vet ./...` passes with no warnings
- [ ] Benchmarks show reasonable performance (<1s)
- [ ] No `panic()`, proper error handling
- [ ] Commit message follows format (no solutions!)
- [ ] Code is self-contained for the day (no shared utilities)
- [ ] Return type is `uint` for numeric answers
- [ ] No recursion used
