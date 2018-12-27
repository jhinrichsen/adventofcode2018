package day2

import (
	"bufio"
	"log"
	"os"
	"testing"
)

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

func TestInput(t *testing.T) {
	f, err := os.Open("testdata/input")
	if err != nil {
		t.Fatal(err)
	}
	var IDs []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		IDs = append(IDs, line)
	}
	log.Printf("day 2: checksum = %v\n", day2(IDs))
}
