package adventofcode2018

import (
	"bufio"
	"os"
	"strconv"
	"testing"
)

func frequencies() ([]int, error) {
	f, err := os.Open("testdata/day1")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	var fs []int
	for sc.Scan() {
		line := sc.Text()
		n, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		fs = append(fs, n)
	}
	return fs, nil
}

func TestDay1Part1(t *testing.T) {
	fs, err := frequencies()
	if err != nil {
		t.Fatal(err)
	}
	want := 531
	got := 0
	for _, n := range fs {
		got += n
	}
	if want != got {
		t.Fatalf("want %v but got %v\n", want, got)
	}
}

func TestDay1Part2(t *testing.T) {
	fs, err := frequencies()
	if err != nil {
		t.Fatal(err)
	}
	twice := make(map[int]bool)
	want := 76787
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
