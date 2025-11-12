package adventofcode2018

import (
	"fmt"
	"sort"
	"testing"
)

func TestDay24Part1Example(t *testing.T) {
	puzzle, err := NewDay24(exampleFile(t, 24))
	if err != nil {
		t.Fatal(err)
	}
	got := Day24(puzzle, true)
	const want = "5216"
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestDay24ExampleRound1(t *testing.T) {
	puzzle, err := NewDay24(exampleFile(t, 24))
	if err != nil {
		t.Fatal(err)
	}

	// Verify initial state
	if len(puzzle.immuneSystem) != 2 {
		t.Errorf("expected 2 immune groups, got %d", len(puzzle.immuneSystem))
	}
	if puzzle.immuneSystem[0].units != 17 {
		t.Errorf("immune group 1: expected 17 units, got %d", puzzle.immuneSystem[0].units)
	}
	if puzzle.immuneSystem[1].units != 989 {
		t.Errorf("immune group 2: expected 989 units, got %d", puzzle.immuneSystem[1].units)
	}
	if puzzle.infection[0].units != 801 {
		t.Errorf("infection group 1: expected 801 units, got %d", puzzle.infection[0].units)
	}
	if puzzle.infection[1].units != 4485 {
		t.Errorf("infection group 2: expected 4485 units, got %d", puzzle.infection[1].units)
	}

	// Simulate one round
	immune := make([]group, len(puzzle.immuneSystem))
	copy(immune, puzzle.immuneSystem)
	infect := make([]group, len(puzzle.infection))
	copy(infect, puzzle.infection)

	var allGroups []group
	allGroups = append(allGroups, immune...)
	allGroups = append(allGroups, infect...)

	targets := selectTargets(allGroups)

	groupMap := make(map[string]*group)
	for i := range immune {
		key := fmt.Sprintf("immune-%d", immune[i].id)
		groupMap[key] = &immune[i]
	}
	for i := range infect {
		key := fmt.Sprintf("infection-%d", infect[i].id)
		groupMap[key] = &infect[i]
	}

	sort.Slice(allGroups, func(i, j int) bool {
		return allGroups[i].initiative > allGroups[j].initiative
	})

	for _, attacker := range allGroups {
		attackerKey := fmt.Sprintf("%s-%d", attacker.army, attacker.id)
		attackerPtr := groupMap[attackerKey]
		if attackerPtr == nil || attackerPtr.units <= 0 {
			continue
		}

		targetID := targets[attackerKey]
		if targetID == "" {
			continue
		}

		defenderPtr := groupMap[targetID]
		if defenderPtr != nil && defenderPtr.units > 0 {
			damage := calculateDamage(*attackerPtr, *defenderPtr)
			unitsKilled := damage / defenderPtr.hp
			if unitsKilled > defenderPtr.units {
				unitsKilled = defenderPtr.units
			}
			defenderPtr.units -= unitsKilled
		}
	}

	// Verify state after round 1 (from AoC description)
	// Expected: Immune Group 2: 905, Infection Group 1: 797, Infection Group 2: 4434
	g := groupMap["immune-1"]
	if g.units != 0 {
		t.Errorf("after round 1: immune group 1 should be eliminated, got %d units", g.units)
	}

	g = groupMap["immune-2"]
	if g.units != 905 {
		t.Errorf("after round 1: immune group 2 expected 905 units, got %d", g.units)
	}

	g = groupMap["infection-1"]
	if g.units != 797 {
		t.Errorf("after round 1: infection group 1 expected 797 units, got %d", g.units)
	}

	g = groupMap["infection-2"]
	if g.units != 4434 {
		t.Errorf("after round 1: infection group 2 expected 4434 units, got %d", g.units)
	}
}

func TestDay24Part1(t *testing.T) {
	testWithParserBytes(t, 24, file, true, NewDay24, Day24, "16530")
}

func BenchmarkDay24Part1(b *testing.B) {
	benchWithParserBytes(b, 24, true, NewDay24, Day24)
}
