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

func TestSample(t *testing.T) {
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
	want := 5434
	got := day2(IDs)
	if want != got {
		t.Fatalf("want %v but got %v\n", want, got)
	}
}
