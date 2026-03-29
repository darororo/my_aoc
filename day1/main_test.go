package main

import (
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

func TestCountZeroes(t *testing.T) {
	tests := []struct {
		val, rotateBy, result, zeroes int
	}{
		{50, 180, 30, 2},
		{20, -20, 0, 1},
		{30, -50, 80, 1},
		{0, -200, 0, 2},
		{1, -200, 1, 2},
		{10, 200, 10, 2},
		{50, -1000, 50, 10},
		{50, 1000, 50, 10},
	}

	for _, tt := range tests {
		t.Run("testing case", func(t *testing.T) {
			val, zeroes := rotateAndCountZeroes(tt.val, tt.rotateBy)
			if val != tt.result || zeroes != tt.zeroes {
				t.Errorf("CountZeroes(%d, %d) = (%d, %d); want (%d, %d)",
					tt.val, tt.rotateBy, val, zeroes, tt.result, tt.zeroes)
			}
		})
	}
}

func TestAnswerFromTestFile(t *testing.T) {
	answer := getRotateValueFromFile("test.txt")
	expected := 3 // Given
	if answer != expected {
		t.Errorf("Answer: %d; Expected: %d", answer, expected)
	}
}

func TestCountAllCrossedZeroesFromFile(t *testing.T) {
	zeroes := countZeroesFromFile("test.txt")
	expected := 6 // Given
	if zeroes != expected {
		t.Errorf("Answer: %d; Expected: %d", zeroes, expected)
	}
}
