package adventofcode2018

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
)

func linesFromFilename(tb testing.TB, filename string) []string {
	tb.Helper()
	f, err := os.Open(filename)
	if err != nil {
		tb.Fatal(err)
	}
	return linesFromReader(tb, f)
}

func linesFromReader(tb testing.TB, r io.Reader) []string {
	tb.Helper()
	var lines []string
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := sc.Text()
		lines = append(lines, line)
	}
	if err := sc.Err(); err != nil {
		tb.Fatal(err)
	}
	return lines
}

func exampleFilename(day uint8) string {
	return fmt.Sprintf("testdata/day%02d_example.txt", day)
}

func filename(day uint8) string {
	return fmt.Sprintf("testdata/day%02d.txt", day)
}

// file reads the main input file bytes for day N (zero-padded).
func file(tb testing.TB, day uint8) []byte {
	tb.Helper()
	buf, err := os.ReadFile(filename(day))
	if err != nil {
		tb.Fatal(err)
	}
	return buf
}

// exampleFile reads the example input file bytes for day N (zero-padded).
func exampleFile(tb testing.TB, day uint8) []byte {
	tb.Helper()
	buf, err := os.ReadFile(exampleFilename(day))
	if err != nil {
		tb.Fatal(err)
	}
	return buf
}

// example2File reads the second example input file bytes for day N (zero-padded).
func example2File(tb testing.TB, day uint8) []byte {
	tb.Helper()
	buf, err := os.ReadFile(fmt.Sprintf("testdata/day%02d_example2.txt", day))
	if err != nil {
		tb.Fatal(err)
	}
	return buf
}
