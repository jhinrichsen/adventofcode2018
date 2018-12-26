package main

import (
	"bufio"
	"io"
	"strconv"
)

func main() {
}

func day1_1(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r)
	sum := 0
	for sc.Scan() {
		line := sc.Text()
		n, err := strconv.Atoi(line)
		if err != nil {
			return 0, err
		}
		sum += n
	}
	return sum, nil
}
