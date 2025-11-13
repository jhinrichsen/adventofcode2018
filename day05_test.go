package adventofcode2018

import (
	"bytes"
	"math"
	"os"
	"strings"
	"testing"
)

func reactive(p1, p2 byte) bool {
	if p1 > p2 {
		p2, p1 = p1, p2
	}
	if p2-p1 == 32 {
		return true
	}
	return false
}

func react(polymer []byte) []byte {
	for i := 0; i < len(polymer)-1; i++ {
		if reactive(polymer[i], polymer[i+1]) {
			polymer = append(polymer[0:i], polymer[i+2:]...)
			i -= 2
			if i < 0 {
				i = -1
			}
		}
	}
	return polymer
}

func TestDay05Example1(t *testing.T) {
	want := []byte("dabCBAcaDA")
	got := react([]byte("dabAcCaCBAcCcaDA"))
	if !bytes.Equal(want, got) {
		t.Fatalf("want %v but got %v", string(want), string(got))
	}
}

func TestDay05Example2(t *testing.T) {
	want := []byte("x")
	got := react([]byte("xabcCdDBADd"))
	if !bytes.Equal(want, got) {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func TestDay05Part1(t *testing.T) {
	const want = 9288
	buf, err := os.ReadFile(filename(5))
	if err != nil {
		t.Fatal(err)
	}
	got := len(react([]byte(strings.TrimSpace(string(buf)))))
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func toLower(r rune) rune {
	if 'A' <= r && r <= 'Z' {
		r += 32
	}
	return r
}

func toUpper(r rune) rune {
	if 'a' <= r && r <= 'z' {
		r -= 32
	}
	return r
}

func unitTypes(polymer string) map[rune]bool {
	m := make(map[rune]bool)
	for _, v := range polymer {
		m[toLower(v)] = true
	}
	return m
}

func reduce(polymer string, unitType rune) string {
	polymer = strings.Replace(polymer, string(toLower(unitType)), "", -1)
	polymer = strings.Replace(polymer, string(toUpper(unitType)), "", -1)
	return polymer
}

func day05Part2(polymer string) int {
	min := math.MaxInt32
	for ut := range unitTypes(polymer) {
		l := len(react([]byte(reduce(polymer, ut))))
		if l < min {
			min = l
		}
	}
	return min
}

func TestDay05Part2(t *testing.T) {
	const want = 5844
	buf, err := os.ReadFile(filename(5))
	if err != nil {
		t.Fatal(err)
	}
	got := day05Part2(strings.TrimSpace(string(buf)))
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func BenchmarkDay05Part1(b *testing.B) {
	buf, err := os.ReadFile(filename(5))
	if err != nil {
		b.Fatal(err)
	}
	buf = []byte(strings.TrimSpace(string(buf)))
	b.ResetTimer()
	for range b.N {
		_ = len(react(buf))
	}
}

func BenchmarkDay05Part2(b *testing.B) {
	buf, err := os.ReadFile(filename(5))
	if err != nil {
		b.Fatal(err)
	}
	polymer := strings.TrimSpace(string(buf))
	b.ResetTimer()
	for range b.N {
		_ = day05Part2(polymer)
	}
}
