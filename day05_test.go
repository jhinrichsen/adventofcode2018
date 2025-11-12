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

type reactFn func(polymer []byte) []byte

func reactByResclicing(polymer []byte) []byte {
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

func reactByCreatingHoles(polymer []byte) []byte {
	const hole = '.'
	l := len(polymer) - 1
	next := func(i int) int {
		if i == len(polymer) {
			return l
		}
		for i < len(polymer) && polymer[i] == hole {
			i++
		}
		return i
	}
	prev := func(i int) int {
		if i == -1 {
			return 0
		}
		for i > 0 && polymer[i] == hole {
			i--
		}
		return i
	}
	var holes int
	for i := 0; i < l; {
		if polymer[i] == hole {
			i++
			continue
		}
		j := next(i + 1)
		if i != j && reactive(polymer[i], polymer[j]) {
			polymer[i], polymer[j] = hole, hole
			holes += 2
			i = prev(i - 1)
		} else {
			i++
		}
	}
	into := make([]byte, len(polymer)-holes)
	var j int
	for i := range polymer {
		if polymer[i] != hole {
			into[j] = polymer[i]
			j++
		}
	}
	return into
}

func react(polymer []byte) []byte {
	return reactByResclicing(polymer)
	// return reactByCreatingHoles(polymer)
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

func TestDay05Part1CreatingHoles(t *testing.T) {
	day05Part1(t, reactByCreatingHoles)
}

func TestDay05Part1Reslicing(t *testing.T) {
	day05Part1(t, reactByResclicing)
}

func day05Part1(t *testing.T, f func([]byte) []byte) {
	const want = 9288
	buf, err := os.ReadFile(filename(5))
	if err != nil {
		t.Fatal(err)
	}
	got := len(f([]byte(strings.TrimSpace(string(buf)))))
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
	reacts := []struct {
		name string
		fn   reactFn
	}{
		{"reactByReslicing", reactByResclicing},
		{"reactByCreatingHoles", reactByCreatingHoles},
	}
	buf, err := os.ReadFile(filename(5))
	if err != nil {
		b.Fatal(err)
	}
	buf = []byte(strings.TrimSpace(string(buf)))
	b.ResetTimer()
	for _, r := range reacts {
		b.Run(r.name, func(b *testing.B) {
			for b.Loop() {
				_ = len(r.fn(buf))
			}
		})
	}
}
