package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/chrismcd1/aoc2025/utils"
)

type FreshRange struct {
	Lower int
	Upper int
}

func main() {
	fmt.Println("Hello World Day 1")
	solution, err := part1("./full_input.txt")
	if err != nil {
		fmt.Println("Error: %w", err)
	}
	fmt.Println("The solution is:")
	fmt.Printf("%d", solution)
}

func part1(filename string) (int, error) {
	var all string
	lines, _ := utils.ReadLines(filename)
	for _, line := range lines {
		all += line + "\n"
	}
	ranges := strings.Split(all, "\n\n")
	freshRegion := ranges[0]

	freshRanges := make([]FreshRange, 0)
	for _, freshLine := range strings.Split(freshRegion, "\n") {
		foo := strings.Split(freshLine, "-")
		lower, _ := strconv.Atoi(foo[0])
		upper, _ := strconv.Atoi(foo[1])
		freshRanges = append(freshRanges, FreshRange{
			Lower: lower,
			Upper: upper,
		})
	}
	fmt.Println(freshRanges)

	var processedRanges []FreshRange

	for _, fresh := range freshRanges {
		var freshSplits []*FreshRange
		freshSplits = append(freshSplits, &fresh)
		// We need to check if it is against the processedRanges
		for _, processed := range processedRanges {
			// Do all of the splits we can on each of the current children
			for i, split := range freshSplits {
				fmt.Println("Considering split i", i, split, processed)
				if split == nil {
					continue
				}
				if split.Upper < processed.Lower {
					continue
				} else if split.Lower > processed.Upper {
					continue
				} else if split.Lower >= processed.Lower && split.Upper <= processed.Upper {
					// Fully contained, so we just remove it!
					freshSplits[i] = nil
				} else if (split.Lower >= processed.Lower && split.Lower <= processed.Upper) && split.Upper > processed.Upper {
					// Overlapping to the right
					foo := FreshRange{
						Lower: processed.Upper + 1,
						Upper: split.Upper,
					}
					fmt.Println("Overlapping to right", split, processed, foo)
					freshSplits[i] = nil
					freshSplits = append(freshSplits, &foo)
				} else if (split.Upper <= processed.Upper && split.Upper >= processed.Lower) && split.Lower < processed.Lower {
					foo := FreshRange{
						Lower: split.Lower,
						Upper: processed.Lower - 1,
					}
					fmt.Println("Overlapping to left", split, processed, foo)
					// Overlapping to the left
					freshSplits[i] = nil
					freshSplits = append(freshSplits, &foo)
				} else if split.Lower < processed.Lower && split.Upper > processed.Lower {
					// fully coverying this range
					lowerSection := FreshRange{
						Lower: split.Lower,
						Upper: processed.Lower - 1,
					}
					upperSection := FreshRange{
						Lower: processed.Upper + 1,
						Upper: split.Upper,
					}
					fmt.Println("Overlapping to bothSides", split, processed, lowerSection, upperSection)
					// Overlapping to the left
					freshSplits[i] = nil
					freshSplits = append(freshSplits, &lowerSection)
					freshSplits = append(freshSplits, &upperSection)
				}
			}
		}

		for _, guy := range freshSplits {
			if guy != nil {
				processedRanges = append(processedRanges, *guy)
			}
		}
	}

	fmt.Println(processedRanges)

	result := 0
	for _, processed := range processedRanges {
		result += processed.Upper - processed.Lower + 1
	}

	return result, nil
}
