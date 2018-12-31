package adventofcode2018

import (
	"io/ioutil"
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

func react(polymer string) string {
	for i := 0; i < len(polymer)-1; i++ {
		if reactive(polymer[i], polymer[i+1]) {
			polymer = polymer[0:i] + polymer[i+2:]
			i -= 2
			if i < 0 {
				i = -1
			}
		}
	}
	return polymer
}

func TestDay5Sample(t *testing.T) {
	want := "dabCBAcaDA"
	got := react("dabAcCaCBAcCcaDA")
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func TestDay5Sample2(t *testing.T) {
	want := "x"
	got := react("xabcCdDBADd")
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func TestDay5(t *testing.T) {
	want := 11252
	buf, err := ioutil.ReadFile("testdata/day5")
	if err != nil {
		t.Fatal(err)
	}
	got := len(react(strings.TrimSpace(string(buf))))
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}
