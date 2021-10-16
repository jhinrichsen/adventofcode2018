package adventofcode2018

import (
	"strconv"
	"testing"
)

func frequencies() ([]int, error) {
	var fs []int
	lines, err := linesFromFilename(filename(1))
	if err != nil {
		return fs, err
	}
	for _, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			return fs, err
		}
		fs = append(fs, n)
	}
	return fs, nil
}

func TestDay01Part1(t *testing.T) {
	fs, err := frequencies()
	if err != nil {
		t.Fatal(err)
	}
	want := 454
	got := 0
	for _, n := range fs {
		got += n
	}
	if want != got {
		t.Fatalf("want %v but got %v\n", want, got)
	}
}

func TestDay01Part2(t *testing.T) {
	fs, err := frequencies()
	if err != nil {
		t.Fatal(err)
	}
	twice := make(map[int]bool)
	want := 566
	got := 0
	i := 0
	for {
		got += fs[i]
		if twice[got] {
			break
		}
		twice[got] = true

		i++
		if i == len(fs) {
			i = 0
		}
	}
	if want != got {
		t.Fatalf("want %v but got %v\n", want, got)
	}
}
