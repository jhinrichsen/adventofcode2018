package adventofcode2018

import (
	"testing"
)

func appear1(ID string, n int) bool {
	occurences := make(map[rune]int)
	for _, v := range ID {
		occurences[v]++
	}
	for _, v := range occurences {
		if v == n {
			return true
		}
	}
	return false
}

func appear(boxIDs []string, n int) int {
	count := 0
	for _, ID := range boxIDs {
		if appear1(ID, n) {
			count++
		}
	}
	return count
}

func day2(boxIDs []string) int {
	return appear(boxIDs, 2) * appear(boxIDs, 3)
}

func TestSamplePart1(t *testing.T) {
	const want = 12
	ids := []string{
		"abcdef",
		"bababc",
		"abbcde",
		"abcccd",
		"aabcdd",
		"abcdee",
		"ababab",
	}
	got := day2(ids)
	if want != got {
		t.Fatalf("want %v but got %v\n", want, got)
	}
}

func TestDay02Part1(t *testing.T) {
	const want = 5704
	lines := linesFromFilename(t, filename(2))
	got := day2(lines)
	if want != got {
		t.Fatalf("want %v but got %v\n", want, got)
	}
}

func sameRunes(s1, s2 string) string {
	var bs []byte
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			bs = append(bs, s1[i])
		}
	}
	return string(bs)
}

func TestSameRunes(t *testing.T) {
	const want = "fgij"
	got := sameRunes("fghij", "fguij")
	if want != got {
		t.Fatalf("want %v but got %v\n", want, got)
	}
}

func day2Part2(IDs []string) string {
	for i := range IDs {
		for j := i; j < len(IDs); j++ {
			s := sameRunes(IDs[i], IDs[j])
			if len(IDs[i])-1 == len(s) {
				return s
			}
		}
	}
	return ""
}

func TestSamplePart2(t *testing.T) {
	const want = "fgij"
	ids := []string{
		"abcde",
		"fghij",
		"klmno",
		"pqrst",
		"fguij",
		"axcye",
		"wvxyz",
	}
	got := day2Part2(ids)
	if want != got {
		t.Fatalf("want %v but got %v\n", want, got)
	}
}

func TestDay02Part2(t *testing.T) {
	const want = "umdryabviapkozistwcnihjqx"
	lines := linesFromFilename(t, filename(2))
	got := day2Part2(lines)
	if want != got {
		t.Fatalf("want %v but got %v\n", want, got)
	}
}
