package adventofcode2018

import (
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"testing"
)

func TestDay4Sample(t *testing.T) {
	buf, err := ioutil.ReadFile("testdata/day4_sample")
	if err != nil {
		t.Fatal(err)
	}
	events := Lines(string(buf))
	want := 10 * 24
	got := day4(events)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}

func dec(s string, i int) int {
	var j int
	for j = i; isDigit(s[j]); j++ {
	}
	n, err := strconv.Atoi(s[i:j])
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func TestDec(t *testing.T) {
	want := 1317
	got := dec("1317 ", 0)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

// guard already used
func guardID(event string) int {
	return dec(event, 26)
}

// minute already used
func min(event string) int {
	return dec(event, 15)
}

// A minute and related asleeps count
type minuteReport [60]int

// a guard ID, and related minute report
type guardReport map[int]minuteReport

func newMinuteReport() minuteReport {
	var r minuteReport
	return r
}

func day4Report(events []string) guardReport {
	guardReport := make(guardReport)
	var g, asleep, awake int
	for _, event := range events {
		switch event[19] {
		case 'G': // Guard #.. begins shift
			g = guardID(event)
		case 'f': // falls asleep
			asleep = min(event)
		case 'w': // wakes up
			awake = min(event)
			mr := guardReport[g]
			for t := asleep; t < awake; t++ {
				mr[t]++
			}
			// TODO use slice to avoid copying back and forth
			guardReport[g] = mr
		}
	}
	return guardReport
}

func minutesAsleep(mr minuteReport) int {
	sum := 0
	for _, m := range mr {
		sum += m
	}
	return sum
}

func day4Strategy1(r guardReport) int {
	var max, sleepiestGuard int
	for g, mr := range r {
		n := minutesAsleep(mr)
		if n > max {
			max = n
			sleepiestGuard = g
		}
	}

	var n, sleepiestMinute int
	for k, v := range r[sleepiestGuard] {
		if v > n {
			n = v
			sleepiestMinute = k
		}
	}
	return sleepiestGuard * sleepiestMinute
}

func day4(events []string) int {
	r := day4Report(events)
	return day4Strategy1(r)
}

func TestDay4(t *testing.T) {
	buf, err := ioutil.ReadFile("testdata/day4")
	if err != nil {
		t.Fatal(err)
	}
	events := Lines(string(buf))
	sort.Strings(events)
	want := 26281
	got := day4(events)
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func Benchmark(b *testing.B) {
	buf, err := ioutil.ReadFile("testdata/day4")
	if err != nil {
		b.Fatal(err)
	}
	events := Lines(string(buf))
	sort.Strings(events)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		want := 703
		got := day4(events)
		if want != got {
			b.Fatalf("want %d but got %d\n", want, got)
		}
	}
}

func TestDay4Part2Sample(t *testing.T) {
	buf, err := ioutil.ReadFile("testdata/day4_sample")
	if err != nil {
		t.Fatal(err)
	}
	events := Lines(string(buf))
	want := 4455
	got := day4Part2(events)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func day4Part2(events []string) int {
	r := day4Report(events)
	return day4Strategy2(r)
}

func day4Strategy2(r guardReport) int {
	var guard, minute, times int
	for k, v := range r {
		for j := range v {

			if v[j] > times {
				guard = k
				minute = j
				times = v[j]
			}
		}
	}
	return guard * minute
}

func TestDay4Part2(t *testing.T) {
	buf, err := ioutil.ReadFile("testdata/day4")
	if err != nil {
		t.Fatal(err)
	}
	events := Lines(string(buf))
	sort.Strings(events)
	want := 73001
	got := day4Part2(events)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
