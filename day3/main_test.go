package main

import "testing"

func TestGetBestPair(t *testing.T) {
	tests := []struct {
		input string
		left  int
		right int
	}{
		{"12345", 4, 5},
		{"455453", 5, 5},
		{"455463", 6, 3},
		{"9234591", 9, 9},
		{"345443", 5, 4},
		{"987654321111111", 9, 8},
		{"811111111111119", 8, 9},
		{"234234234234278", 7, 8},
		{"818181911112111", 9, 2},
	}

	for _, tt := range tests {
		t.Run("testing case", func(t *testing.T) {
			l, r := GetBestPair(tt.input)
			if l != tt.left || r != tt.right {
				t.Errorf("GetBestPair(%v) = (%v, %v), want (%v, %v);",
					tt.input, l, r, tt.left, tt.right)
			}
		})
	}
}

func TestSumPairsFromFile(t *testing.T) {
	expected := 357 // AOC 2025 day 3
	got := SumPairsFromFile("test.txt")

	if got != expected {
		t.Errorf("Got: %v; Expected: %v", got, expected)
	}
}

func TestGetBest12(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"987654321111111", 987654321111},
		{"811111111111119", 811111111119},
		{"234234234234278", 434234234278},
		{"818181911112111", 888911112111},
	}

	for _, tt := range tests {
		t.Run("testing case", func(t *testing.T) {
			got := GetBest12(tt.input)
			if got != tt.want {
				t.Errorf("GetBest12(%v) = %v, want %v;",
					tt.input, got, tt.want)
			}
		})
	}
}

func TestSumBest12FromFile(t *testing.T) {
	var expected int = 3121910778619 // AOC 2025 day 3
	got := SumBest12FromFile("test.txt")
	if got != expected {
		t.Errorf("Got: %v; Expected: %v", got, expected)
	}
}
