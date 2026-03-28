package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"strconv"
	"strings"
)

func ValidateID(input string) bool {

	// contain a leading zero
	if input[0:1] == "0" {
		return false
	}

	// is Odd
	if len(input)%2 == 1 {
		return true
	}

	// Check repeating pattern
	length := len(input)
	leftHalf := input[0 : length/2]
	rightHalf := input[length/2:]
	if leftHalf == rightHalf {
		return false
	}

	return true
}

func AddInvalidsFromRange(low, high int) int {
	sum := 0
	for i := low; i <= high; i++ {
		s := strconv.Itoa(i)
		if ok := ValidateID(s); !ok {
			sum = sum + i
		}
	}

	return sum
}

func ParseRange(input string) (low, high int) {
	tokens := strings.Split(input, "-")
	l, err := strconv.Atoi(tokens[0])
	if err != nil {
		panic(err)
	}
	h, err := strconv.Atoi(tokens[1])
	if err != nil {
		panic(err)
	}
	return l, h
}

func splitOnComma() func([]byte, bool) (int, []byte, error) {
	// Define a custom split function for commas
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		// Search for the first comma
		if i := bytes.IndexByte(data, ','); i >= 0 {
			// Found a comma: return everything up to it,
			// and advance the scanner past it (+1)
			return i + 1, data[:i], nil
		}
		// If we're at EOF and have remaining data, return it
		if atEOF {
			return len(data), data, nil
		}
		// Request more data from the reader
		return 0, nil, nil
	}
}

func main() {
	// 1. Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // Ensure file is closed

	scanner := bufio.NewScanner(file)
	scanner.Split(splitOnComma())

	invalidTotal := 0
	for scanner.Scan() {
		text := scanner.Text()
		// At EOF, we get 0 from the scanner
		if text == "0" {
			continue
		}

		low, high := ParseRange(text)

		invalidCounts := AddInvalidsFromRange(low, high)

		invalidTotal = invalidTotal + invalidCounts
	}

	println(invalidTotal)
}
