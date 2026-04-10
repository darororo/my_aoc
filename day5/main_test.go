package main

import "testing"

func TestFindFreshIngredients(t *testing.T) {
	input := `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

	expected := 3
	counts := FindFreshIngredients(input)
	if counts != expected {
		t.Errorf("Answer: %v, Expected: %v", counts, expected)
	}

}
