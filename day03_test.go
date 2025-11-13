package adventofcode2018

import (
	"strconv"
	"strings"
	"testing"
)

type claim struct {
	x, y, w, h int
}

// #1 @ 1,3: 4x5
func newClaim(s string) (claim, error) {
	var c claim
	s1 := strings.Split(s, ":")
	s2 := strings.Split(s1[0], "@")
	s3 := strings.Split(strings.TrimSpace(s2[1]), ",")
	x, err := strconv.Atoi(s3[0])
	if err != nil {
		return c, nil
	}
	y, err := strconv.Atoi(s3[1])
	if err != nil {
		return c, nil
	}
	s4 := strings.Split(strings.TrimSpace(s1[1]), "x")
	w, err := strconv.Atoi(s4[0])
	if err != nil {
		return c, nil
	}
	h, err := strconv.Atoi(s4[1])
	if err != nil {
		return c, nil
	}
	return claim{x, y, w, h}, nil
}

func TestNewClaim(t *testing.T) {
	want := claim{1, 3, 4, 5}
	got, err := newClaim("#1 @ 1,3: 4x5")
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %v but got %v\n", want, got)
	}
}

func dimension(claims []claim) (int, int) {
	mx, my := 0, 0
	for _, c := range claims {
		x, y := c.x+c.w, c.y+c.h
		if x > mx {
			mx = x
		}
		if y > my {
			my = y
		}
	}
	return mx, my
}

func mkSquare(x, y int) [][]int {
	square := make([][]int, y)
	for j := range square {
		square[j] = make([]int, x)
	}
	return square
}

func TestMkSquare(t *testing.T) {
	sq := mkSquare(19, 13)
	y := len(sq)
	x := len(sq[0])
	if x != 19 {
		t.Fatalf("want 19 but got %d\n", x)
	}
	if y != 13 {
		t.Fatalf("want 13 but got %d\n", x)
	}
	sq[12][18] = 0
}

func fabric(square [][]int, c claim) {
	for y := c.y; y < c.y+c.h; y++ {
		for x := c.x; x < c.x+c.w; x++ {
			square[y][x]++
		}
	}
}

func overlaps(square [][]int) int {
	n := 0
	for y := range square {
		for x := range square[y] {
			if square[y][x] > 1 {
				n++
			}
		}
	}
	return n
}

func day3(claims []claim) int {
	x, y := dimension(claims)
	square := mkSquare(x, y)
	for _, c := range claims {
		fabric(square, c)
	}
	return overlaps(square)
}

func claimsFromString(rep []string) ([]claim, error) {
	var claims []claim
	for _, v := range rep {
		c, err := newClaim(v)
		if err != nil {
			return nil, err
		}
		claims = append(claims, c)
	}
	return claims, nil
}

func TestDay03Example(t *testing.T) {
	cs := []string{
		"#1 @ 1,3: 4x4",
		"#2 @ 3,1: 4x4",
		"#3 @ 5,5: 2x2",
	}
	claims, err := claimsFromString(cs)
	if err != nil {
		t.Fatal(err)
	}
	want := 4
	got := day3(claims)
	if want != got {
		t.Fatalf("want %v but got %v\n", want, got)
	}
}

func TestDay03Part1(t *testing.T) {
	const want = 118322
	lines := linesFromFilename(t, filename(3))
	claims, err := claimsFromString(lines)
	if err != nil {
		t.Fatal(err)
	}
	got := day3(claims)
	if want != got {
		t.Fatalf("want %v but got %v\n", want, got)
	}
}

func (a *claim) abs() (Point, Point) {
	return Point{a.x, a.y}, Point{a.x + a.w, a.y + a.h}
}

func day3Part2(claims []claim) (int, error) {
	// assume all claims are disjoint and remove those that overlap
	// disjoint := !overlap
	disjoint := make(map[int]bool)
	for i := range claims {
		disjoint[i] = true
	}
	for i := range claims {
		for j := range claims {
			if i == j {
				continue
			}
			l1, r1 := claims[i].abs()
			l2, r2 := claims[j].abs()
			if Overlap(l1, r1, l2, r2) {
				delete(disjoint, i)
				delete(disjoint, j)
			}
		}
	}
	for i := range disjoint {
		return i + 1, nil
	}
	return -1, nil
}

func TestDay03Part2Example(t *testing.T) {
	cs := []string{
		"#1 @ 1,3: 4x4",
		"#2 @ 3,1: 4x4",
		"#3 @ 5,5: 2x2",
	}
	claims, err := claimsFromString(cs)
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got, err := day3Part2(claims)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %v but got %v\n", want, got)
	}
}

func TestDay03Part2(t *testing.T) {
	const want = 1178
	lines := linesFromFilename(t, filename(3))
	claims, err := claimsFromString(lines)
	if err != nil {
		t.Fatal(err)
	}
	got, err := day3Part2(claims)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %v but got %v\n", want, got)
	}
}

func BenchmarkDay03Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(3))
	b.ResetTimer()
	for range b.N {
		claims, _ := claimsFromString(lines)
		_ = day3(claims)
	}
}

func BenchmarkDay03Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(3))
	b.ResetTimer()
	for range b.N {
		claims, _ := claimsFromString(lines)
		_, _ = day3Part2(claims)
	}
}
