package adventofcode2018

import (
	"fmt"
	"sort"
	"testing"
)

func TestDay24Part1Example(t *testing.T) {
	testWithParserBytes(t, 24, exampleFile, true, NewDay24, Day24, 5216)
}

func TestDay24ExampleTimeline(t *testing.T) {
	puzzle, err := NewDay24(exampleFile(t, 24))
	if err != nil {
		t.Fatal(err)
	}

	// Expected state after each round from AoC description
	tests := []struct {
		round      int
		immune1    int
		immune2    int
		infection1 int
		infection2 int
	}{
		{0, 17, 989, 801, 4485}, // Initial
		{1, 0, 905, 797, 4434},  // After round 1
		{2, 0, 761, 793, 4434},  // After round 2
		{3, 0, 618, 789, 4434},  // After round 3
		{4, 0, 475, 786, 4434},  // After round 4
		{5, 0, 333, 784, 4434},  // After round 5
		{6, 0, 191, 783, 4434},  // After round 6
		{7, 0, 49, 782, 4434},   // After round 7
		{8, 0, 0, 782, 4434},    // Final: immune eliminated
	}

	immune := make([]group, len(puzzle.immuneSystem))
	copy(immune, puzzle.immuneSystem)
	infect := make([]group, len(puzzle.infection))
	copy(infect, puzzle.infection)

	for _, tt := range tests {
		// Check current state
		var immune1Units, immune2Units, infect1Units, infect2Units int
		for _, g := range immune {
			if g.units > 0 {
				if g.id == 1 {
					immune1Units = g.units
				} else if g.id == 2 {
					immune2Units = g.units
				}
			}
		}
		for _, g := range infect {
			if g.units > 0 {
				if g.id == 1 {
					infect1Units = g.units
				} else if g.id == 2 {
					infect2Units = g.units
				}
			}
		}

		if immune1Units != tt.immune1 || immune2Units != tt.immune2 ||
			infect1Units != tt.infection1 || infect2Units != tt.infection2 {
			t.Errorf("round %d: got immune=[%d,%d] infection=[%d,%d], want immune=[%d,%d] infection=[%d,%d]",
				tt.round, immune1Units, immune2Units, infect1Units, infect2Units,
				tt.immune1, tt.immune2, tt.infection1, tt.infection2)
		}

		// Stop if we've reached the end
		if len(immune) == 0 || len(infect) == 0 {
			break
		}

		// Simulate one round
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

		unitsKilledThisRound := 0
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
				unitsKilledThisRound += unitsKilled
			}
		}

		if unitsKilledThisRound == 0 {
			break
		}

		// Update armies
		immune = nil
		infect = nil
		for _, g := range groupMap {
			if g.units > 0 {
				if g.army == "immune" {
					immune = append(immune, *g)
				} else {
					infect = append(infect, *g)
				}
			}
		}
		sort.Slice(immune, func(i, j int) bool { return immune[i].id < immune[j].id })
		sort.Slice(infect, func(i, j int) bool { return infect[i].id < infect[j].id })
	}
}

func TestDay24Part1(t *testing.T) {
	testWithParserBytes(t, 24, file, true, NewDay24, Day24, 16530)
}

func TestDay24Part2ExampleWithBoost(t *testing.T) {
	puzzle, err := NewDay24(exampleFile(t, 24))
	if err != nil {
		t.Fatal(err)
	}

	// With boost of 1570, immune system should win with 51 units
	winner, units := SimulateCombat(puzzle, 1570)
	if winner != "immune" {
		t.Errorf("with boost 1570: expected immune to win, but %s won", winner)
	}
	if units != 51 {
		t.Errorf("with boost 1570: expected 51 units, got %d", units)
	}
}

func TestDay24Part2Example(t *testing.T) {
	testWithParserBytes(t, 24, exampleFile, false, NewDay24, Day24, 51)
}

func TestDay24Part2(t *testing.T) {
	testWithParserBytes(t, 24, file, false, NewDay24, Day24, 3313)
}

func BenchmarkDay24Part1(b *testing.B) {
	benchWithParserBytes(b, 24, true, NewDay24, Day24)
}

func BenchmarkDay24Part2(b *testing.B) {
	benchWithParserBytes(b, 24, false, NewDay24, Day24)
}
