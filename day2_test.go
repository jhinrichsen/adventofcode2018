package adventofcode2018

import (
	"bufio"
	"os"
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
	ids := []string{
		"abcdef",
		"bababc",
		"abbcde",
		"abcccd",
		"aabcdd",
		"abcdee",
		"ababab",
	}
	want := 12
	got := day2(ids)
	if want != got {
		t.Fatalf("want %v but got %v\n", want, got)
	}
}

func TestDay2(t *testing.T) {
	f, err := os.Open("testdata/day2")
	if err != nil {
		t.Fatal(err)
	}
	var IDs []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		IDs = append(IDs, line)
	}
	if sc.Err() != nil {
		t.Fatal(sc.Err())
	}
	want := 5434
	got := day2(IDs)
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
	want := "fgij"
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
	ids := []string{
		"abcde",
		"fghij",
		"klmno",
		"pqrst",
		"fguij",
		"axcye",
		"wvxyz",
	}
	want := "fgij"
	got := day2Part2(ids)
	if want != got {
		t.Fatalf("want %v but got %v\n", want, got)
	}
}

func TestDay2Part2(t *testing.T) {
	f, err := os.Open("testdata/day2")
	if err != nil {
		t.Fatal(err)
	}
	var IDs []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		IDs = append(IDs, line)
	}
	if sc.Err() != nil {
		t.Fatal(sc.Err())
	}
	want := "agimdjvlhedpsyoqfzuknpjwt"
	got := day2Part2(IDs)
	if want != got {
		t.Fatalf("want %v but got %v\n", want, got)
	}
}
