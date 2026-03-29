package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Part 1
func parseRotateBy(input string) int {
	sdir := input[0:1] // L or R
	snum := input[1:]  // Number

	num, e := strconv.Atoi(snum)
	if e != nil {
		panic("Cannot parse number")
	}

	dir := 1

	switch sdir {
	case "L":
		dir = -1
	case "R":
		dir = 1
	}

	return num * dir
}

func rotate(num int, rotateBy int) int {
	total := (num + rotateBy) % 100

	if total < 0 {
		total = total + 100
	}

	return total
}

func getRotateValueFromFile(path string) int {
	// 1. Open the file
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // Ensure file is closed

	// 2. Initialize the scanner
	scanner := bufio.NewScanner(file)

	answer := 0
	total := 50 // given inital value

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

	return answer
}

// Part 2

// Count the number of zeroes that the rotation passes by
func rotateAndCountZeroes(num int, rotateBy int) (value int, zeroes int) {
	total := num + rotateBy
	val := total % 100
	if val < 0 {
		val = val + 100
	}

	z := 0
	if total <= 0 {
		// when the original val is positive,
		// it needs to cross 0 once to get to negative
		score := 0
		if num > 0 {
			score = 1
		}

		z = total*-1/100 + score
	} else {
		z = total / 100
	}
	return val, z
}

func countZeroesFromFile(path string) int {
	// 1. Open the file
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // Ensure file is closed

	// 2. Initialize the scanner
	scanner := bufio.NewScanner(file)

	totalZeroes := 0
	total := 50 // given inital value

	// 3. Iterate through lines
	for scanner.Scan() {
		line := scanner.Text() // Get the line as a string
		rotateBy := parseRotateBy(line)

		val, zeroes := rotateAndCountZeroes(total, rotateBy)
		totalZeroes += zeroes
		total = val

	}

	return totalZeroes
}

func main() {

	part1 := getRotateValueFromFile("input.txt")
	fmt.Printf("Part 1's Answer = %v\n", part1)

	part2 := countZeroesFromFile("input.txt")
	fmt.Printf("Part 2's Answer = %v\n", part2)

}
