package main

import (
	"bufio"
	"log"
	"os"
	"testing"
)

func TestRotateLeft(t *testing.T) {
	rotateBy := parseRotateBy("L50")
	expected := -50

	if rotateBy != expected {
		t.Errorf("Result: %d; Expected: %d", rotateBy, expected)
	}
}

func TestRotateRight(t *testing.T) {
	rotateBy := parseRotateBy("R10")
	expected := 10

	if rotateBy != expected {
		t.Errorf("Result: %d; Expected: %d", rotateBy, expected)
	}
}

func TestRotateCalc(t *testing.T) {
	tests := []struct {
		val, rotateBy, want int
	}{
		{50, 180, 30},
		{20, -20, 0},
		{30, -50, 80},
		{0, -200, 0},
		{10, -200, 10},
	}

	for _, tt := range tests {
		t.Run("testing case", func(t *testing.T) {
			if got := rotate(tt.val, tt.rotateBy); got != tt.want {
				t.Errorf("Add(%d, %d) = %d, want %d", tt.val, tt.rotateBy, got, tt.want)
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

	answer := 0
	total := 50 // given inital value

	expected := 3 // Given

	// 3. Iterate through lines
	for scanner.Scan() {
		line := scanner.Text() // Get the line as a string
		rotateBy := parseRotateBy(line)

		result := rotate(total, rotateBy)
		if result == 0 {
			answer = answer + 1
		}
		total = result

	}

	if answer != expected {
		t.Errorf("Answer: %d; Expected: %d", answer, expected)
	}
}
