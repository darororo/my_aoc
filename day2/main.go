package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"strconv"
	"strings"
)

const DEBUG = false

// Part One
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

// Part 2

// keep creating equal-size subsequences,
// and check if each subseq is the same.
// Input is not valid if it has repeating subsequences at least twice.
func ValidateIDv2(input string) bool {
	// contain a leading zero
	if input[0:1] == "0" {
		return false
	}

	// Check repeating pattern
	length := len(input)
	subLength := 1

	isRepeating := false
	for subLength < length {
		// can't have repeating pattern
		// if we can't create equal-size substrings
		if length%subLength != 0 {
			subLength++
			continue
		}

		subSeq := input[:subLength]
		subSeqNums := length / subLength

		for i := 1; i < subSeqNums; i++ {
			start := i * subLength
			end := i*subLength + subLength
			cur := input[start:end]

			if DEBUG {
				log.Printf("i=%v\n", i)
				log.Printf("cur=%v\n", cur)
			}

			if subSeq == cur {
				isRepeating = true
			} else {
				isRepeating = false
				break
			}
		}

		// found the repeating subsequence
		if isRepeating {
			break
		}
		// increase the substring size
		subLength++
	}

	// Not valid if it has repeating sequences
	return !isRepeating

}

func AddInvalidsFromRange(low, high int, validator func(string) bool) int {
	sum := 0
	for i := low; i <= high; i++ {
		s := strconv.Itoa(i)
		if ok := validator(s); !ok {
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

func SumInvalidsFromFile(path string, validator func(string) bool) int {
	// 1. Open the file
	file, err := os.Open(path)
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

		invalidCounts := AddInvalidsFromRange(low, high, validator)

		invalidTotal = invalidTotal + invalidCounts
	}

	return invalidTotal
}

func main() {
	part1 := SumInvalidsFromFile("input.txt", ValidateID)
	println("Part 1's Answer: %v", part1)
	part2 := SumInvalidsFromFile("input.txt", ValidateIDv2)
	println("Part 2's Answer: %v", part2)
}
