package main

import (
	"fmt"
	"testing"
)

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
	rangeInts, ingredients := parseInput(input)
	counts := FindFreshIngredients(rangeInts, ingredients)
	if counts != expected {
		t.Errorf("Answer: %v, Expected: %v", counts, expected)
	}

}

func TestFindRangeOfGoodIngredients(t *testing.T) {
	input := `3-5
10-14
16-20
12-18
21-25
3-5
100-120
21-25
3-5
100-120
21-25
3-5
100-120

1
5
8
11
17
32`

	var expected uint64 = 40
	rangeInts, _ := parseInput(input)
	ranges, total := FindRangesOfGoodIngredients(rangeInts)
	if total != expected {
		t.Errorf("Answer: %v, Expected: %v", total, expected)
	}

	fmt.Println("Finger")
	fmt.Println(ranges)

}

func TestMergeRanges(t *testing.T) {
	tests := []struct {
		r1   [2]uint64
		r2   [2]uint64
		want [2]uint64
	}{
		// merge r1 and r2
		{[2]uint64{10, 20}, [2]uint64{15, 20}, [2]uint64{10, 20}},
		{[2]uint64{15, 25}, [2]uint64{18, 20}, [2]uint64{15, 25}},
		{[2]uint64{90, 125}, [2]uint64{100, 120}, [2]uint64{90, 125}},
		{[2]uint64{10, 18}, [2]uint64{12, 20}, [2]uint64{10, 20}},
		{[2]uint64{10, 18}, [2]uint64{18, 20}, [2]uint64{10, 20}},

		// swap r1 and r2
		{[2]uint64{15, 20}, [2]uint64{10, 20}, [2]uint64{10, 20}},
		{[2]uint64{18, 20}, [2]uint64{15, 25}, [2]uint64{15, 25}},
		{[2]uint64{100, 120}, [2]uint64{90, 125}, [2]uint64{90, 125}},
		{[2]uint64{12, 20}, [2]uint64{10, 18}, [2]uint64{10, 20}},

		// r1 = r2
		{[2]uint64{10, 20}, [2]uint64{10, 20}, [2]uint64{10, 20}},

		// difference of 1 of either lower or upper bounds
		{[2]uint64{30, 35}, [2]uint64{36, 50}, [2]uint64{30, 50}},
		{[2]uint64{30, 35}, [2]uint64{10, 29}, [2]uint64{10, 35}},
		{[2]uint64{10, 20}, [2]uint64{21, 25}, [2]uint64{10, 25}},

		// cannot merge
		{[2]uint64{40, 60}, [2]uint64{80, 100}, [2]uint64{0, 0}},
		{[2]uint64{1, 10}, [2]uint64{12, 15}, [2]uint64{0, 0}},
	}

	for _, tt := range tests {
		t.Run("testing case", func(t *testing.T) {
			merged, _ := MergeRanges(tt.r1, tt.r2)
			if merged != tt.want {
				t.Errorf("MergeRanges(%v, %v) = %v, want %v;",
					tt.r1, tt.r2, merged, tt.want)
			}
		})
	}
}
