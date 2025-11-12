package adventofcode2018

import (
	"fmt"
	"sort"
	"strings"
)

// Day24Puzzle represents the two armies.
type Day24Puzzle struct {
	immuneSystem []group
	infection    []group
}

type group struct {
	id         int
	army       string // "immune" or "infection"
	units      int
	hp         int
	weaknesses []string
	immunities []string
	attackDmg  int
	attackType string
	initiative int
}

// NewDay24 parses the army groups.
func NewDay24(data []byte) (Day24Puzzle, error) {
	lines := strings.Split(string(data), "\n")

	var immuneSystem []group
	var infection []group

	currentArmy := ""
	groupID := 1

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if line == "Immune System:" {
			currentArmy = "immune"
			continue
		}
		if line == "Infection:" {
			currentArmy = "infection"
			groupID = 1
			continue
		}

		g := parseGroup(line, groupID, currentArmy)
		if currentArmy == "immune" {
			immuneSystem = append(immuneSystem, g)
		} else {
			infection = append(infection, g)
		}
		groupID++
	}

	return Day24Puzzle{immuneSystem: immuneSystem, infection: infection}, nil
}

func parseGroup(line string, id int, army string) group {
	g := group{id: id, army: army}

	// Parse: "17 units each with 5390 hit points (weak to radiation, bludgeoning) with an attack that does 4507 fire damage at initiative 2"

	parts := strings.Split(line, " ")
	i := 0

	// Parse units
	_, _ = fmt.Sscanf(parts[i], "%d", &g.units)

	// Find "with" to get HP
	for i < len(parts) {
		if parts[i] == "with" {
			i++
			break
		}
		i++
	}
	_, _ = fmt.Sscanf(parts[i], "%d", &g.hp)
	i += 3 // skip "hit points" (the number itself, "hit", "points")

	// Check for weaknesses/immunities in parentheses
	if i < len(parts) && strings.HasPrefix(parts[i], "(") {
		// Collect everything in parentheses
		var modifiers string
		for i < len(parts) {
			modifiers += parts[i] + " "
			if strings.HasSuffix(parts[i], ")") {
				i++
				break
			}
			i++
		}

		// Parse modifiers
		modifiers = strings.Trim(modifiers, "() ")
		sections := strings.Split(modifiers, ";")
		for _, section := range sections {
			section = strings.TrimSpace(section)
			if strings.HasPrefix(section, "weak to ") {
				types := strings.TrimPrefix(section, "weak to ")
				g.weaknesses = strings.Split(strings.ReplaceAll(types, ", ", ","), ",")
			} else if strings.HasPrefix(section, "immune to ") {
				types := strings.TrimPrefix(section, "immune to ")
				g.immunities = strings.Split(strings.ReplaceAll(types, ", ", ","), ",")
			}
		}
	}

	// Find "does" for attack damage
	for i < len(parts) {
		if parts[i] == "does" {
			i++
			break
		}
		i++
	}
	_, _ = fmt.Sscanf(parts[i], "%d", &g.attackDmg)
	i++
	g.attackType = parts[i]

	// Find "initiative" for initiative value
	for i < len(parts) {
		if parts[i] == "initiative" {
			i++
			break
		}
		i++
	}
	_, _ = fmt.Sscanf(parts[i], "%d", &g.initiative)

	return g
}

// Day24 simulates the combat.
// Part 1: Returns the number of units in the winning army.
// Part 2: Returns the number of units the immune system has after winning with the smallest boost.
func Day24(puzzle Day24Puzzle, part1 bool) uint {
	if part1 {
		_, units := SimulateCombat(puzzle, 0)
		return uint(units)
	}

	// Part 2: Find smallest boost for immune system to win
	for boost := 1; ; boost++ {
		winner, units := SimulateCombat(puzzle, boost)
		if winner == "immune" {
			return uint(units)
		}
	}
}

// SimulateCombat runs the combat simulation with the given boost.
// Returns the winner ("immune" or "infection") and the number of units remaining.
func SimulateCombat(puzzle Day24Puzzle, boost int) (string, int) {
	// Make copies
	immune := make([]group, len(puzzle.immuneSystem))
	copy(immune, puzzle.immuneSystem)
	infect := make([]group, len(puzzle.infection))
	copy(infect, puzzle.infection)

	// Apply boost to immune system
	for i := range immune {
		immune[i].attackDmg += boost
	}

	// Simulate combat
	for len(immune) > 0 && len(infect) > 0 {
		// Get all groups for target selection
		var allGroups []group
		allGroups = append(allGroups, immune...)
		allGroups = append(allGroups, infect...)

		// Target selection phase
		targets := selectTargets(allGroups)

		// Create group map for easy lookup during attack
		groupMap := make(map[string]*group)
		for i := range immune {
			key := fmt.Sprintf("immune-%d", immune[i].id)
			groupMap[key] = &immune[i]
		}
		for i := range infect {
			key := fmt.Sprintf("infection-%d", infect[i].id)
			groupMap[key] = &infect[i]
		}

		// Track if any units died this round
		unitsKilledThisRound := 0

		// Attack phase (sorted by initiative, descending)
		sort.Slice(allGroups, func(i, j int) bool {
			return allGroups[i].initiative > allGroups[j].initiative
		})

		for _, attacker := range allGroups {
			// Get fresh attacker state
			attackerKey := fmt.Sprintf("%s-%d", attacker.army, attacker.id)
			attackerPtr := groupMap[attackerKey]
			if attackerPtr == nil || attackerPtr.units <= 0 {
				continue
			}

			targetID := targets[attackerKey]
			if targetID == "" {
				continue
			}

			// Find target and deal damage
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

		// If no units died, it's a stalemate
		if unitsKilledThisRound == 0 {
			break
		}

		// Update armies (remove dead groups)
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
	}

	// Determine winner and count remaining units
	if len(immune) > 0 && len(infect) > 0 {
		// Stalemate: both armies survive
		return "stalemate", 0
	}

	if len(immune) > 0 {
		total := 0
		for _, g := range immune {
			total += g.units
		}
		return "immune", total
	}

	total := 0
	for _, g := range infect {
		total += g.units
	}
	return "infection", total
}

func effectivePower(g group) int {
	return g.units * g.attackDmg
}

func calculateDamage(attacker, defender group) int {
	damage := effectivePower(attacker)

	// Check immunities
	for _, imm := range defender.immunities {
		if imm == attacker.attackType {
			return 0
		}
	}

	// Check weaknesses
	for _, weak := range defender.weaknesses {
		if weak == attacker.attackType {
			return damage * 2
		}
	}

	return damage
}

func selectTargets(groups []group) map[string]string {
	targets := make(map[string]string)
	targeted := make(map[string]bool)

	// Sort by effective power (descending), then initiative (descending)
	sort.Slice(groups, func(i, j int) bool {
		pi, pj := effectivePower(groups[i]), effectivePower(groups[j])
		if pi != pj {
			return pi > pj
		}
		return groups[i].initiative > groups[j].initiative
	})

	for _, attacker := range groups {
		if attacker.units <= 0 {
			continue
		}

		var bestTarget *group
		bestDamage := 0

		for i := range groups {
			defender := &groups[i]
			if defender.units <= 0 {
				continue
			}
			if defender.army == attacker.army {
				continue
			}

			defenderID := fmt.Sprintf("%s-%d", defender.army, defender.id)
			if targeted[defenderID] {
				continue
			}

			damage := calculateDamage(attacker, *defender)
			if damage == 0 {
				continue
			}

			if damage > bestDamage ||
				(damage == bestDamage && bestTarget != nil && effectivePower(*defender) > effectivePower(*bestTarget)) ||
				(damage == bestDamage && bestTarget != nil && effectivePower(*defender) == effectivePower(*bestTarget) && defender.initiative > bestTarget.initiative) {
				bestDamage = damage
				bestTarget = defender
			}
		}

		if bestTarget != nil {
			attackerID := fmt.Sprintf("%s-%d", attacker.army, attacker.id)
			targetID := fmt.Sprintf("%s-%d", bestTarget.army, bestTarget.id)
			targets[attackerID] = targetID
			targeted[targetID] = true
		}
	}

	return targets
}
