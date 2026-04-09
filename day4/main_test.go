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

	expected_marked := `..xx.xx@x.
x@@.@.@.@@
@@@@@.x.@@
@.@@@@..@.
x@.@@@@.@x
.@@@@@@@.@
.@.@.@.@@@
x.@@@.@@@@
.@@@@@@@@.
x.x.@@@.x.`

	counts, marked := FindGoodPapers(input, "x")

	expected := 13 // Given in AOC 2025 day 3 part 1

	if counts != expected {
		t.Errorf("Answer: %d; Expected: %d\n\n", counts, expected)
	}

	if marked != expected_marked {
		t.Errorf("Answer: \n%v \nExpected: \n%v", expected_marked, marked)
	}

}

func TestCountGoodPaperRollsRecurse(t *testing.T) {
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

	expected_marked := `..........
..........
..........
....@@....
...@@@@...
...@@@@@..
...@.@.@@.
...@@.@@@.
...@@@@@..
....@@@...`

	counts, marked := FindGoodPapersRecurse(input)

	expected := 43 // Given in AOC 2025 day 3 part 1

	if counts != expected {
		t.Errorf("Answer: %d; Expected: %d\n\n", counts, expected)
	}

	if marked != expected_marked {
		t.Errorf("Answer: \n%v \nExpected: \n%v", expected_marked, marked)
	}

}
