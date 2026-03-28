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

func main() {
	sum := SumPairsFromFile("input.txt")
	fmt.Printf("Answer: %v\n", sum)
}
