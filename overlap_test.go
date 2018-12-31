package adventofcode2018

import "testing"

func TestOverlap(t *testing.T) {
	var tests = []struct {
		rect1, rect2 []int
		want         bool
	}{
		{[]int{3, 3, 5, 5}, []int{6, 6, 8, 8}, false},
		{[]int{3, 3, 7, 7}, []int{6, 6, 8, 8}, true},
		{[]int{3, 3, 5, 5}, []int{5, 5, 7, 7}, false},
		{[]int{0, 0, 5, 5}, []int{2, 2, 4, 4}, true},
	}

	for i, tt := range tests {
		l1 := Point{tt.rect1[0], tt.rect1[1]}
		r1 := Point{tt.rect1[2], tt.rect1[3]}
		l2 := Point{tt.rect2[0], tt.rect2[1]}
		r2 := Point{tt.rect2[2], tt.rect2[3]}

		got := Overlap(l1, r1, l2, r2)
		if tt.want != got {
			t.Fatalf("TestOverlap #%d: want %v but got %v",
				i, tt.want, got)
		}
	}
}
