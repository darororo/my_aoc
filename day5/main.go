package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

var DebugMode bool = false

// Part 1
func FindFreshIngredients(rangeInts [][2]uint64, ingredients []string) int {

	freshCounts := 0

	for _, ingredient := range ingredients {
		freshFound := false
		val, err := strconv.ParseUint(ingredient, 10, 64)
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

// Part 2
func FindRangesOfGoodIngredients(rangeInts [][2]uint64) ([][2]uint64, uint64) {
	var total uint64 = 0

	clone := rangeInts

	for {
		fmt.Println("CLONE")
		fmt.Println(clone)

		didMerge := false

		for i := 0; i < len(clone); i++ {

			for j := 0; j < len(clone); {
				if i >= len(clone)-1 {
					break
				}
				fmt.Println("LEN", len(clone))

				fmt.Printf("i %v, j %v R2 %v\n", i, j, clone[j])
				if i == j {
					fmt.Println("skipping ", j)
					j++
					continue
				}

				r1 := clone[i]
				r2 := clone[j]

				merged, ok := MergeRanges(r1, r2)
				if ok {
					if DebugMode {
						fmt.Printf("replacing %v with %v\n", r1, merged)
					}

					clone[i] = merged
					if j < len(clone)-1 {
						if DebugMode {
							fmt.Printf("removing %v at %v\n", r2, j)
						}

						clone = append(clone[:j], clone[j+1:]...)
					} else {
						if DebugMode {
							fmt.Printf("removing %v at %v\n", r2, j)
						}

						clone = clone[:j]
					}
					fmt.Println("AFTER REMOVED", clone)

					// j = j - 1

					// fmt.Printf("Delete %v\n", r2)

					didMerge = true
				} else {
					j++
					if DebugMode {
						fmt.Printf("Merge failed: %v\n", r2)
					}
					// failedCounts++
				}
			}

		}

		if !didMerge {
			break
		}

		if DebugMode {
			fmt.Printf("Clone: %v\n", clone)
			// fmt.Printf("Failed merges: %v\n", failedMerges)
		}

		if len(clone) == 0 {
			break
		}

		if DebugMode {
			fmt.Printf("Clone: %v\n", clone)
		}

	}

	for _, r := range clone {
		if DebugMode {
			fmt.Printf("Optimal: %v\n", r)
		}

		total += r[1] - r[0] + 1
	}

	fmt.Println("clone")
	fmt.Println(clone)

	return clone, total
}

func parseRange(r string) [2]uint64 {
	list := strings.Split(r, "-")
	// low, err := strconv.Atoi(list[0])
	low, err := strconv.ParseUint(list[0], 10, 64)
	if err != nil {
		log.Fatalln(err)
	}

	high, err := strconv.ParseUint(list[1], 10, 64)
	if err != nil {
		log.Fatalln(err)
	}

	return [2]uint64{low, high}
}

func MergeRanges(r1 [2]uint64, r2 [2]uint64) ([2]uint64, bool) {

	// 11-20
	// 10-12
	// => can merge [10, 20], true

	// 8-10
	// 12-20 => cannot merge => [0, 0], false

	// 12-20
	// 8-10 => cannot merge => [0, 0], false

	findSubset := func(r1 [2]uint64, r2 [2]uint64) ([2]uint64, bool) {
		isSubset := false
		l1 := r1[0]
		h1 := r1[1]

		l2 := r2[0]
		h2 := r2[1]

		var ld uint64
		var hd uint64

		if h1 <= h2 {
			// mergeable
			// r1: 10-15
			// r2: 16-20
			// merged: 10-20
			if l2-h1 <= 1 {
				isSubset = true

				ld = l1
				hd = h2
			}

			if h1 > l2 {
				isSubset = true
				hd = h2
				if l1 < l2 {
					// mergeable
					// r1: 10-15
					// r2: 12-20
					// merged: 10-20
					ld = l1
				} else {
					// mergeable
					// r1: 15-20
					// r2: 12-20
					// merged: 12-20
					ld = l2
				}
			}
		}

		if !isSubset {
			return [2]uint64{0, 0}, false
		}

		merged := [2]uint64{ld, hd}
		return merged, true
	}

	// try to merge r1 against r2
	merged, ok := findSubset(r1, r2)
	if ok {
		return merged, ok
	}
	// return merge result of r2 against r1
	return findSubset(r2, r1)
}

func parseInput(input string) ([][2]uint64, []string) {
	parts := strings.Split(input, "\n\n")

	rangeStrs := strings.Split(parts[0], "\n")
	var rangeInts [][2]uint64

	for _, str := range rangeStrs {
		r := parseRange(str)
		rangeInts = append(rangeInts, r)
	}
	ingredients := strings.Split(parts[1], "\n")

	return rangeInts, ingredients
}

func sortRanges(rangeInts [][2]uint64) [][2]uint64 {
	slices.SortFunc(rangeInts, func(a, b [2]uint64) int {
		return cmp.Compare(a[0], b[0])
	})

	return rangeInts
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

	rangeInts, ingredients := parseInput(input)

	part1 := FindFreshIngredients(rangeInts, ingredients)
	fmt.Printf("Part 1: %v\n", part1)

	// part2 := 0
	// var totalLast uint64 = 0

	fmt.Printf("length: %v\n", len(rangeInts))
	fmt.Printf("arr:\n%v\n", rangeInts)

	_, total := FindRangesOfGoodIngredients(rangeInts)

	fmt.Printf("Part 2: \n%v\n", total)

}
