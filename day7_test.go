package adventofcode2018

import (
	"os"
	"testing"
)

func ParseSample() (*DAC, error) {
	f, err := os.Open("testdata/day7_sample")
	if err != nil {
		return nil, err
	}
	return Parse(f)
}

func TestDay7SampleEdges(t *testing.T) {
	want := 7
	g, err := ParseSample()
	if err != nil {
		t.Fatal(err)
	}
	got := len(g.edges)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func TestDay7SampleNodes(t *testing.T) {
	want := 6
	g, err := ParseSample()
	if err != nil {
		t.Fatal(err)
	}
	got := len(g.nodes)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func TestDay7Sample(t *testing.T) {
	want := "CABDFE"
	g, err := ParseSample()
	if err != nil {
		t.Fatal(err)
	}
	got := g.Flatten()
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}
