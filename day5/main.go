package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func FindFreshIngredients(input string) int {

	parts := strings.Split(input, "\n\n")

	rangeStrs := strings.Split(parts[0], "\n")
	rangeInts := make([][2]int, len(rangeStrs))

	freshCounts := 0

	for _, str := range rangeStrs {
		rangeInts = append(rangeInts, parseRange(str))
	}

	ingredients := strings.SplitSeq(parts[1], "\n")

	for ingredient := range ingredients {
		freshFound := false
		val, err := strconv.Atoi(string(ingredient))
		if err != nil {
			log.Fatalln(err)
		}

		for _, r := range rangeInts {
			low := r[0]
			high := r[1]
			if val >= low && val <= high {
				freshCounts++
				break
			}
		}

		if freshFound {
			continue
		}
	}

	return freshCounts
}

func parseRange(r string) [2]int {
	list := strings.Split(r, "-")
	low, err := strconv.Atoi(list[0])
	if err != nil {
		log.Fatalln(err)
	}

	high, err := strconv.Atoi(list[1])
	if err != nil {
		log.Fatalln(err)
	}

	return [2]int{low, high}
}

func main() {
	fmt.Println("hello")

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	scan := bufio.NewScanner(file)

	var builder strings.Builder
	for scan.Scan() {
		builder.WriteString(scan.Text())
		builder.WriteString("\n")
	}
	// remove last \n
	input := builder.String()
	input = input[:len(input)-1]

	counts := FindFreshIngredients(input)

	fmt.Println(counts)

}
