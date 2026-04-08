package main

import (
	"testing"
)

func TestCountGoodPaperRolls(t *testing.T) {
	input := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

	// 	marked := `..xx.xx@x.
	// x@@.@.@.@@
	// @@@@@.x.@@
	// @.@@@@..@.
	// x@.@@@@.@x
	// .@@@@@@@.@
	// .@.@.@.@@@
	// x.@@@.@@@@
	// .@@@@@@@@.
	// x.x.@@@.x.`

	counts, _ := findGoodPapers(input)

	expected := 13 // Given in AOC 2025 day 3 part 1

	if counts != expected {
		t.Errorf("Answer: %d; Expected: %d\n\n", counts, expected)
	}

	// if output != marked {
	// 	t.Errorf("Answer: \n%v \nExpected: \n%v", output, marked)
	// }

}
