package adventofcode2018

import (
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"testing"
)

func coordinates(ss []string) ([]Point, error) {
	var ps []Point
	for _, s := range ss {
		parts := strings.Split(s, ", ")
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			return ps, err
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			return ps, err
		}
		p := Point{x, y}
		ps = append(ps, p)
	}
	return ps, nil
}

func size(ps []Point) (int, int) {
	mx, my := 0, 0
	for _, p := range ps {
		if p.x > mx {
			mx = p.x
		}
		if p.y > my {
			my = p.y
		}
	}
	return mx + 1, my + 1
}

const empty = -1
const draw = -2

func mkEmptySquare(mx, my int) [][]int {
	square := mkSquare(mx, my)
	for x := 0; x < mx; x++ {
		for y := 0; y < my; y++ {
			square[y][x] = empty
		}
	}
	return square
}

func setup(square [][]int, ps []Point) {
	for i, p := range ps {
		square[p.y][p.x] = i
	}
}

func manhattanDistance(p1, p2 Point) int {
	dx := p2.x - p1.x
	if dx < 0 {
		dx = -dx
	}
	dy := p2.y - p1.y
	if dy < 0 {
		dy = -dy
	}
	return dx + dy
}

// returns indices into ps of nearest Point to (x/y)
func minDist(ps []Point, p0 Point) []int {
	// phase 1: collect distances
	min := math.MaxInt32
	distances := make([]int, len(ps))
	for i := range ps {
		distances[i] = manhattanDistance(p0, ps[i])
		if distances[i] < min {
			min = distances[i]
		}
	}
	// phase 2: collect all minimal distances
	var mins []int
	for i := range distances {
		if distances[i] == min {
			mins = append(mins, i)
		}
	}
	return mins
}

func fillManhattan(square [][]int, ps []Point) {
	for y := 0; y < len(square); y++ {
		for x := 0; x < len(square[0]); x++ {
			if square[y][x] == empty {
				mins := minDist(ps, Point{x, y})
				if len(mins) == 1 {
					square[y][x] = mins[0]
				} else {
					square[y][x] = draw
				}
			}
		}
	}
}

func borders(square [][]int) map[int]bool {
	bs := make(map[int]bool)
	my := len(square)
	mx := len(square[0])
	for y := 0; y < my; y++ {
		bs[square[y][0]] = true
		bs[square[y][mx-1]] = true
	}
	for x := 0; x < mx; x++ {
		bs[square[0][x]] = true
		bs[square[my-1][x]] = true
	}
	return bs
}

func maxFinite(square [][]int, m map[int]bool) (int, int) {
	var maxIndex, max int
	for i := range m {
		n := 0
		for y := range square {
			for x := range square[0] {
				if square[y][x] == i {
					n++
				}
			}
		}
		if n > max {
			max = n
			maxIndex = i
		}
	}
	return maxIndex, max
}

func day6(ps []Point) int {
	square := mkEmptySquare(size(ps))
	setup(square, ps)
	fillManhattan(square, ps)
	all := make(map[int]bool)
	for i := range ps {
		all[i] = true
	}
	infinite := borders(square)
	for i := range infinite {
		delete(all, i)
	}
	_, n := maxFinite(square, all)
	return n
}

func TestDay6Sample(t *testing.T) {
	cs := []string{
		"1, 1",
		"1, 6",
		"8, 3",
		"3, 4",
		"5, 5",
		"8, 9",
	}

	want := 17
	ps, err := coordinates(cs)
	if err != nil {
		t.Fatal(err)
	}
	got := day6(ps)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func Test_manhattanDistance(t *testing.T) {
	type args struct {
		p1 Point
		p2 Point
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := manhattanDistance(tt.args.p1, tt.args.p2); got != tt.want {
				t.Errorf("manhattanDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay6(t *testing.T) {
	buf, err := ioutil.ReadFile("testdata/day6")
	if err != nil {
		t.Fatal(err)
	}
	want := 3276
	ps, err := coordinates(Lines(string(buf)))
	if err != nil {
		t.Fatal(err)
	}
	got := day6(ps)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}
