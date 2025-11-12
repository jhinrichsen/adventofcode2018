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
	t.Skip("Skipping - answer too high, need to debug")
	testWithParserBytes(t, 15, file, true, NewDay15, Day15, "0")
}

func BenchmarkDay15Part1(b *testing.B) {
	benchWithParserBytes(b, 15, true, NewDay15, Day15)
}
