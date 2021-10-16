package adventofcode2018

import "fmt"

func exampleFilename(day int) string {
	return fmt.Sprintf("testdata/day%02d_example.txt", day)
}

func filename(day int) string {
	return fmt.Sprintf("testdata/day%02d.txt", day)
}
