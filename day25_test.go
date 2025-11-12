package adventofcode2018

import "testing"

func TestDay25Part1Examples(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			"example 1",
			`0,0,0,0
3,0,0,0
0,3,0,0
0,0,3,0
0,0,0,3
0,0,0,6
9,0,0,0
12,0,0,0`,
			"2",
		},
		{
			"example 2",
			`-1,2,2,0
0,0,2,-2
0,0,0,-2
-1,2,0,0
-2,-2,-2,2
3,0,2,-1
-1,3,2,2
-1,0,-1,0
0,2,1,-2
3,0,0,0`,
			"4",
		},
		{
			"example 3",
			`1,-1,0,1
2,0,-1,0
3,2,-1,0
0,0,3,1
0,0,-1,-1
2,3,-2,0
-2,2,0,0
2,-2,0,-1
1,-1,0,-1
3,2,0,2`,
			"3",
		},
		{
			"example 4",
			`1,-1,-1,-2
-2,-2,0,1
0,2,1,3
-2,3,-2,1
0,2,3,-2
-1,-1,1,-2
0,-2,-1,0
-2,2,3,-1
1,2,2,0
-1,-2,0,-2`,
			"8",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			puzzle, err := NewDay25([]byte(tt.input))
			if err != nil {
				t.Fatal(err)
			}
			got := Day25(puzzle, true)
			if got != tt.want {
				t.Errorf("want %s but got %s", tt.want, got)
			}
		})
	}
}

func TestDay25Part1(t *testing.T) {
	testWithParserBytes(t, 25, file, true, NewDay25, Day25, "338")
}

func BenchmarkDay25Part1(b *testing.B) {
	benchWithParserBytes(b, 25, true, NewDay25, Day25)
}
