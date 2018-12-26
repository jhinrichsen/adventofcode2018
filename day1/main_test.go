package main

import (
	"fmt"
	"os"
	"testing"
)

func TestDay1(t *testing.T) {
	f, err := os.Open("testdata/input")
	if err != nil {
		t.Error(err)
	}
	sum, err := day1_1(f)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%d\n", sum)
}
