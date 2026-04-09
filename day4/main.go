package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const DEBUG = true

// paper roll is safe
// if there are fewer than four rolls of paper in
// the eight adjacent positions
func FindGoodPapers(input string) (int, string) {
	lines := strings.Split(input, "\n")

	counts := 0

	// final output with
	// safe "@" set to "-"
	var clone strings.Builder

	for lineIdx, line := range lines {
		fmt.Println(lineIdx)

		// write output with the
		// good paper rolls set to "x"
		var buffer strings.Builder

		for chIdx := range len(line) {
			curChr := line[chIdx : chIdx+1]
			if curChr != "@" {
				buffer.WriteString(".")
				continue
			}

			top := ""
			mid := "" // curChar is in the middle substr
			bot := ""

			// assuming each line has the same length
			var start int
			var end int

			if chIdx == 0 { // first char in a line
				start = 0
				end = chIdx + 2
			} else if chIdx == len(line)-1 { // last char in a line
				start = chIdx - 1
				end = chIdx + 1

			} else {
				start = chIdx - 1
				end = chIdx + 2
			}

			if start < 0 || end > len(line) || start > end {
				// Handle the "error" here
				fmt.Println("Invalid slice indices")
				fmt.Printf("line len: %v\n", len(line))
				fmt.Printf("line: %v\n", lineIdx)
				fmt.Printf("start: %v\n", start)
				fmt.Printf("end: %v\n", end)

				break
			} else {
				subString := line[start:end]
				fmt.Println(subString)
			}

			mid = lines[lineIdx][start:end]

			// has top line
			if lineIdx != 0 {
				top = lines[lineIdx-1][start:end]
			}

			// has bottom line
			if lineIdx != len(lines)-1 {
				bot = lines[lineIdx+1][start:end]
				// next := len(lines[lineIdx+1])
				// fmt.Printf("Line + 1: %v\n", next)
			}

			subStrings := top + bot + mid

			// Count all stars around the target '@'
			// substract 1 because we don't count the target '@'
			countStars := strings.Count(subStrings, "@") - 1

			// found safe paper
			if countStars < 4 {
				buffer.WriteString("x")
				counts++
			} else {
				buffer.WriteString("@")
			}
		}
		clone.WriteString(buffer.String())
		clone.WriteString("\n")
		buffer.Reset()
	}

	str := clone.String()
	// remove \n from the last line
	marked := str[:len(str)-1]

	fmt.Println(clone.String())
	fmt.Println()
	fmt.Println(input)

	return counts, marked
}

func writeStringToFile(path string, content string) {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close() // Ensure file is closed

	_, err = f.WriteString(content)
	if err != nil {
		panic(err)
	}
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var builder strings.Builder

	for scanner.Scan() {
		line := scanner.Text()
		builder.WriteString(line)
		builder.WriteString("\n")
	}

	str := builder.String()
	// remove last \n
	input := strings.Clone(str[:len(str)-1])

	counts, marked := FindGoodPapers(input)

	writeStringToFile("output.txt", marked)

	fmt.Printf("Answers: %v", counts)

}
