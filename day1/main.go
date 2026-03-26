package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

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

func main() {

	// 1. Open the file
	file, err := os.Open("input.txt")
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

	fmt.Println("Answer")
	fmt.Println(answer)

}
