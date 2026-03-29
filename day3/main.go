package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const DEBUG = false

func GetBestPair(input string) (int, int) {
	length := len(input)

	bestLeft := 0
	bestLeftIdx := 0
	for i := 0; i < length-1; i++ {
		c := input[i : i+1] // c is string type
		n, _ := strconv.Atoi(c)
		if n > bestLeft {
			bestLeft = n
			bestLeftIdx = i
		}
	}

	bestRight := 0
	// Iterate after bestLeft
	// to get bestRight
	for i := bestLeftIdx + 1; i < length; i++ {
		c := input[i : i+1] // c is string type
		n, _ := strconv.Atoi(c)
		if n > bestRight {
			bestRight = n
		}
	}
	return bestLeft, bestRight

}

func SumPairsFromFile(path string) int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		l, r := GetBestPair(line)
		pair := l*10 + r
		if DEBUG {
			log.Printf("line: %v", line)
			log.Printf("SUM = %v; adding %v", sum, pair)
		}
		sum += pair
	}

	return sum
}

func GetBest12(input string) int {

	buffer := input
	curIndex := 0
	for curIndex < len(buffer)-1 {

		if len(buffer) <= 12 {
			break
		}

		charCur := buffer[curIndex : curIndex+1]
		charNext := buffer[curIndex+1 : curIndex+2]

		nCur, _ := strconv.Atoi(charCur)
		nNext, _ := strconv.Atoi(charNext)

		if nCur < nNext {
			part1 := buffer[:curIndex]
			part2 := buffer[curIndex+1:]

			if DEBUG {
				charRemove := buffer[curIndex : curIndex+1]
				log.Printf("Buffer: %v", buffer)
				log.Printf("charCur: %v", charCur)
				log.Printf("charNext: %v", charNext)
				log.Printf("Removing: %v", charRemove)
				log.Printf("Part1: %v", part1)
				log.Printf("Part2: %v", part2)
			}

			buffer = part1 + part2
			// offset is the index
			// to resume the loop at.
			// we can also just set the offset at 0
			var offset int
			if len(part1) > 0 {
				offset = len(part1) - 1
			} else {
				offset = 0
			}
			curIndex = offset

		} else {
			curIndex = curIndex + 1
		}

	}

	// Trim the numbers to 12 digits.
	// Since none of left digit is less than the right,
	// the rightmost must consit of the lowest sequnce of numbers. e,g. 65432222
	if len(buffer) > 12 {
		buffer = buffer[:12]
	}

	best12, err := strconv.Atoi(buffer)
	if err != nil {
		panic(err)
	}

	return best12
}

func SumBest12FromFile(path string) int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		best12 := GetBest12(line)
		sum += best12
	}

	return sum
}

func main() {
	// Part 1
	sumPair := SumPairsFromFile("input.txt")
	fmt.Printf("Part 1's Answer: %v\n", sumPair)

	// Part 2
	sum12 := SumBest12FromFile("input.txt")
	fmt.Printf("Part 2's Answer: %v\n", sum12)
}
