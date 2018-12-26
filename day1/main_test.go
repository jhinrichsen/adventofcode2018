package day1

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"testing"
)

func frequencies() ([]int, error) {
	f, err := os.Open("testdata/input")
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

func TestPart1(t *testing.T) {
	fs, err := frequencies()
	if err != nil {
		t.Fatal(err)
	}
	sum := 0
	for _, n := range fs {
		sum += n
	}
	log.Printf("Sum: %d\n", sum)
}

func TestPart2(t *testing.T) {
	fs, err := frequencies()
	if err != nil {
		t.Fatal(err)
	}
	twice := make(map[int]bool)
	sum := 0
	i := 0
	for {
		sum += fs[i]
		if twice[sum] {
			break
		}
		twice[sum] = true

		i++
		if i == len(fs) {
			i = 0
		}
	}
	log.Printf("Twice: %d\n", sum)
}
