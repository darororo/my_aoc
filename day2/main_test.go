package main

import (
	"testing"
)

func TestValidateID(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		// repeating sequence
		{"123123", false},
		{"1010", false},
		{"1188511885", false},
		{"38593859", false},
		// leading zeros
		{"0123", false},
		{"000", false},
		// valid IDs
		{"101", true},
		{"111", true},
		{"212", true},
		{"999", true},
		{"1234321", true},
	}

	for _, tt := range tests {
		t.Run("testing case", func(t *testing.T) {
			if got := ValidateID(tt.input); got != tt.want {
				t.Errorf("ValidateID(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestValidateIDv2(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		// repeating sequence
		{"11", false},
		{"22", false},
		{"999", false},
		{"1010", false},
		{"38593859", false},
		{"2121212121", false},
		{"321123", true},
		{"998", true},
		{"1012", true},
		{"38593862", true},
	}

	for _, tt := range tests {
		t.Run("testing case", func(t *testing.T) {
			if got := ValidateIDv2(tt.input); got != tt.want {
				t.Errorf("ValidateIDv2(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestSumRangePart1(t *testing.T) {

	total := SumInvalidsFromFile("test.txt", ValidateID)
	expected := 1227775554 // Given in AOC 2025 day 2

	if total != expected {
		t.Errorf("Answer: %d; Expected: %d", total, expected)
	}
}

func TestSumRangePart2(t *testing.T) {

	total := SumInvalidsFromFile("test.txt", ValidateIDv2)
	expected := 4174379265 // Given in AOC 2025 day 2 part 2

	if total != expected {
		t.Errorf("Answer: %d; Expected: %d", total, expected)
	}
}
