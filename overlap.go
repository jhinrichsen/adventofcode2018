package adventofcode2018

// Point has two coordinates.
type Point struct {
	x, y int
}

// Overlap returns true if two rectanges (l1, r1) and (l2, r2) overlap.
func Overlap(l1, r1, l2, r2 Point) bool {
	// If one rectangle is on left side of other
	if r1.x <= l2.x || r2.x <= l1.x {
		return false
	}

	// If one rectangle is above other
	if r1.y <= l2.y || r2.y <= l1.y {
		return false
	}

	return true
}
