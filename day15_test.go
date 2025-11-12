package adventofcode2018

import "testing"

func TestDay15Part1Examples(t *testing.T) {
	tests := []struct {
		name string
		map_ string
		want string
	}{
		{
			name: "main example",
			map_: `#######
#.G...#
#...EG#
#.#.#G#
#..G#E#
#.....#
#######`,
			want: "27730",
		},
		{
			name: "example 1",
			map_: `#######
#G..#E#
#E#E.E#
#G.##.#
#...#E#
#...E.#
#######`,
			want: "36334",
		},
		{
			name: "example 2",
			map_: `#######
#E..EG#
#.#G.E#
#E.##E#
#G..#.#
#..E#.#
#######`,
			want: "39514",
		},
		{
			name: "example 3",
			map_: `#######
#E.G#.#
#.#G..#
#G.#.G#
#G..#.#
#...E.#
#######`,
			want: "27755",
		},
		{
			name: "example 4",
			map_: `#######
#.E...#
#.#..G#
#.###.#
#E#G#G#
#...#G#
#######`,
			want: "28944",
		},
		{
			name: "example 5",
			map_: `#########
#G......#
#.E.#...#
#..##..G#
#...##..#
#...#...#
#.G...G.#
#.....G.#
#########`,
			want: "18740",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			puzzle, err := NewDay15([]byte(tt.map_))
			if err != nil {
				t.Fatal(err)
			}
			got := Day15(puzzle, true)
			if got != tt.want {
				t.Errorf("want %s but got %s", tt.want, got)
			}
		})
	}
}

func TestDay15Part1(t *testing.T) {
	testWithParserBytes(t, 15, file, true, NewDay15, Day15, "207542")
}

func BenchmarkDay15Part1(b *testing.B) {
	benchWithParserBytes(b, 15, true, NewDay15, Day15)
}

func TestDay15Part2Examples(t *testing.T) {
	tests := []struct {
		name string
		map_ string
		want string
	}{
		{
			name: "main example",
			map_: `#######
#.G...#
#...EG#
#.#.#G#
#..G#E#
#.....#
#######`,
			want: "4988",
		},
		{
			name: "example 1",
			map_: `#######
#E..EG#
#.#G.E#
#E.##E#
#G..#.#
#..E#.#
#######`,
			want: "31284",
		},
		{
			name: "example 2",
			map_: `#######
#E.G#.#
#.#G..#
#G.#.G#
#G..#.#
#...E.#
#######`,
			want: "3478",
		},
		{
			name: "example 3",
			map_: `#######
#.E...#
#.#..G#
#.###.#
#E#G#G#
#...#G#
#######`,
			want: "6474",
		},
		{
			name: "example 4",
			map_: `#########
#G......#
#.E.#...#
#..##..G#
#...##..#
#...#...#
#.G...G.#
#.....G.#
#########`,
			want: "1140",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			puzzle, err := NewDay15([]byte(tt.map_))
			if err != nil {
				t.Fatal(err)
			}
			got := Day15(puzzle, false)
			if got != tt.want {
				t.Errorf("want %s but got %s", tt.want, got)
			}
		})
	}
}

func TestDay15Part2(t *testing.T) {
	testWithParserBytes(t, 15, file, false, NewDay15, Day15, "64688")
}

func BenchmarkDay15Part2(b *testing.B) {
	benchWithParserBytes(b, 15, false, NewDay15, Day15)
}
