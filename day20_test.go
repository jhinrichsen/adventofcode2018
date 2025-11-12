package adventofcode2018

import "testing"

func TestDay20Part1Examples(t *testing.T) {
	tests := []struct {
		name  string
		regex string
		want  uint
	}{
		{"example 1", "^WNE$", 3},
		{"example 2", "^ENWWW(NEEE|SSE(EE|N))$", 10},
		{"example 3", "^ENNWSWW(NEWS|)SSSEEN(WNSE|)EE(SWEN|)NNN$", 18},
		{"example 4", "^ESSWWN(E|NNENN(EESS(WNSE|)SSS|WWWSSSSE(SW|NNNE)))$", 23},
		{"example 5", "^WSSEESWWWNW(S|NENNEEEENN(ESSSSW(NWSW|SSEN)|WSWWN(E|WWS(E|SS))))$", 31},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			puzzle, err := NewDay20([]byte(tt.regex))
			if err != nil {
				t.Fatal(err)
			}
			got := Day20(puzzle, true)
			if got != tt.want {
				t.Errorf("want %d but got %d", tt.want, got)
			}
		})
	}
}

func TestDay20Part1(t *testing.T) {
	testWithParserBytes(t, 20, file, true, NewDay20, Day20, uint(3755))
}

func TestDay20Part2(t *testing.T) {
	testWithParserBytes(t, 20, file, false, NewDay20, Day20, uint(8627))
}

func BenchmarkDay20Part1(b *testing.B) {
	benchWithParserBytes(b, 20, true, NewDay20, Day20)
}

func BenchmarkDay20Part2(b *testing.B) {
	benchWithParserBytes(b, 20, false, NewDay20, Day20)
}
