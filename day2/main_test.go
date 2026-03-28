package main

import (
	"bufio"
	"log"
	"os"
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

func TestAnswerFromTestFile(t *testing.T) {
	// 1. Open the file
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // Ensure file is closed

	// 2. Initialize the scanner
	scanner := bufio.NewScanner(file)
	scanner.Split(splitOnComma())

	total := 0

	expected := 1227775554 // Given in AOC 2025 day 2

	// 3. Iterate through lines
	for scanner.Scan() {
		text := scanner.Text() // Get the line as a string
		// At EOF, we get 0 from the scanner
		if text == "0" {
			continue
		}
		low, high := ParseRange(text)
		subTotal := AddInvalidsFromRange(low, high)
		total = total + subTotal
	}

	if total != expected {
		t.Errorf("Answer: %d; Expected: %d", total, expected)
	}
}
