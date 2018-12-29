package adventofcode2018

import (
	"fmt"
	"io/ioutil"
	"log"
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

func reset(square [][]int) {
	for y := range square {
		for x := range square[y] {
			square[y][x] = 0
		}
	}
}

func dump(square [][]int) {
	fmt.Println("----------------------------------")
	for y := range square {
		for x := range square[y] {
			fmt.Printf("%d", square[y][x])
		}
		fmt.Println()
	}
	fmt.Println("----------------------------------")
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

func TestDay3Sample(t *testing.T) {
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

func TestDay3(t *testing.T) {
	buf, err := ioutil.ReadFile("testdata/day3")
	if err != nil {
		t.Fatal(err)
	}
	lines := Lines(string(buf))
	claims, err := claimsFromString(lines)
	if err != nil {
		t.Fatal(err)
	}
	want := 110827
	got := day3(claims)
	if want != got {
		t.Fatalf("want %v but got %v\n", want, got)
	}
}

func day3Part2(claims []claim) (int, error) {
	// assume all claims are disjoint and remove those that overlap
	// disjoint := !overlap
	disjoint := make(map[int]bool)
	for i := range claims {
		disjoint[i] = true
	}
	square := mkSquare(dimension(claims))
	for i := range claims {
		for j := range claims {
			if i == j {
				continue
			}
			reset(square)
			fabric(square, claims[i])
			fabric(square, claims[j])
			if overlaps(square) > 0 {
				delete(disjoint, i)
				delete(disjoint, j)
			}
		}
	}
	for i := range disjoint {
		log.Printf("index %d does not overlap with any other claim\n",
			i+1)
	}
	for i := range disjoint {
		return i + 1, nil
	}
	return -1, nil
}

func TestDay3Part2Sample(t *testing.T) {
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

func TestDay3Part2(t *testing.T) {
	// TODO requires 2.233 s on a 2,3 GHz Intel Core i7
	if testing.Short() {
		t.Skip("skipping day 3 part 2 in short mode")
	}

	buf, err := ioutil.ReadFile("testdata/day3")
	if err != nil {
		t.Fatal(err)
	}
	claims, err := claimsFromString(Lines(string(buf)))
	if err != nil {
		t.Fatal(err)
	}
	want := 116
	got, err := day3Part2(claims)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %v but got %v\n", want, got)
	}
}
