package adventofcode2018

import (
	"fmt"
	"sort"
	"strings"
)

// Day07 returns sorted characters on input lines in format
// "Step C must be finished before step A can begin.".
func Day07(lines []string) (string, error) {
	// "Each step is designated by a single letter."
	type step = byte
	var result string

	var steps = make(map[step]bool)

	// Instead of linked lists for the graph, which are hard to implement in
	// a borrow checker, use a map. Yes, Go 3 will have a borrow checker.
	var lefts = make(map[step]map[step]bool)
	var rights = make(map[step]map[step]bool)

	// keep track of what we're doing
	dones := make(map[step]bool)
	isDone := func(st step) bool {
		_, ok := dones[st]
		return ok
	}

	prospects := make(map[step]bool)
	isReady := func(st step) bool {
		befores, ok := lefts[st]
		if !ok {
			// no predecessors
			return true
		}
		for k := range befores {
			if !isDone(k) {
				// predecessor not ready
				return false
			}
		}
		return true
	}

	// return before and after step
	order := func(line string) (step, step) {
		parts := strings.Fields(line)
		// die hard 2018
		return parts[1][0], parts[7][0]
	}

	for _, line := range lines {
		l, r := order(line)
		steps[l] = true
		steps[r] = true

		m, ok := lefts[r]
		if !ok {
			// prepare initial map
			m = make(map[step]bool)
		}
		m[l] = true
		lefts[r] = m

		m, ok = rights[l]
		if !ok {
			// prepare initial map
			m = make(map[step]bool)
		}
		m[r] = true
		rights[l] = m
	}

	// populate with beginning steps (no left nodes)
	for st := range steps {
		if len(lefts[st]) == 0 {
			prospects[st] = true
		}
	}

	for len(dones) < len(steps) {
		// filter prospect list where all predecessors have run
		readies := make(map[step]bool)
		for st := range prospects {
			if isReady(st) {
				readies[st] = true
			}
		}

		// pick next step
		var srt []step
		for st := range readies {
			srt = append(srt, st)
		}
		sort.Slice(srt, func(i, j int) bool { return srt[i] < srt[j] })

		// move step from ready to done
		st := srt[0]
		delete(prospects, st)
		dones[st] = true
		result += string(st)

		// pin all successors as prospects, unless already run
		for r := range rights[st] {
			if !isDone(r) {
				prospects[r] = true
			}
		}
	}
	if len(steps) != len(result) {
		return result, fmt.Errorf("internal error: want %d steps but only got %d", len(steps), len(result))
	}
	return result, nil
}

// Day07Part2 computes the total time to complete all steps with the given
// number of workers and a base duration where A=1+base, B=2+base, ... Z=26+base.
func Day07Part2(lines []string, workers, base int) (int, error) {
    type step = byte

    duration := func(st step) int {
        return int(st-'A'+1) + base
    }

    // Build graph: lefts (prereqs), rights (successors), and set of steps
    steps := make(map[step]bool)
    lefts := make(map[step]map[step]bool)
    rights := make(map[step]map[step]bool)

    order := func(line string) (step, step) {
        parts := strings.Fields(line)
        return parts[1][0], parts[7][0]
    }

    for _, line := range lines {
        l, r := order(line)
        steps[l] = true
        steps[r] = true
        if _, ok := lefts[r]; !ok {
            lefts[r] = make(map[step]bool)
        }
        lefts[r][l] = true
        if _, ok := rights[l]; !ok {
            rights[l] = make(map[step]bool)
        }
        rights[l][r] = true
    }

    // Helper to check if all prereqs are done
    dones := make(map[step]bool)
    isReady := func(st step) bool {
        befores, ok := lefts[st]
        if !ok || len(befores) == 0 {
            return true
        }
        for k := range befores {
            if !dones[k] {
                return false
            }
        }
        return true
    }

    // Priority queue of available steps (as map, we will select smallest each tick)
    available := make(map[step]bool)
    for st := range steps {
        if len(lefts[st]) == 0 {
            available[st] = true
        }
    }

    // Workers state
    type worker struct {
        busy bool
        st   step
        rem  int
    }
    ws := make([]worker, workers)

    time := 0
    doneCount := 0
    total := len(steps)

    for doneCount < total {
        // Assign available tasks to idle workers in alphabetical order
        // Gather ready tasks from available where prereqs satisfied and not currently in progress
        inProgress := make(map[step]bool)
        for _, w := range ws {
            if w.busy {
                inProgress[w.st] = true
            }
        }
        var ready []step
        for st := range available {
            if !inProgress[st] && isReady(st) {
                ready = append(ready, st)
            }
        }
        sort.Slice(ready, func(i, j int) bool { return ready[i] < ready[j] })
        // Fill idle workers
        ridx := 0
        for i := range ws {
            if !ws[i].busy && ridx < len(ready) {
                st := ready[ridx]
                ridx++
                ws[i] = worker{busy: true, st: st, rem: duration(st)}
                delete(available, st) // taken
            }
        }

        // Advance time by 1 second
        time++
        // Tick workers and complete tasks finishing now
        for i := range ws {
            if ws[i].busy {
                ws[i].rem--
                if ws[i].rem == 0 {
                    finished := ws[i].st
                    ws[i] = worker{} // now idle
                    // mark done
                    dones[finished] = true
                    doneCount++
                    // add successors to available
                    for succ := range rights[finished] {
                        if !dones[succ] {
                            available[succ] = true
                        }
                    }
                }
            }
        }
    }

    return time, nil
}
