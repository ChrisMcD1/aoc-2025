package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/chrismcd1/aoc2025/utils"
)

func main() {
	fmt.Println("Hello World Day 1")
	solution, err := part1("./full_input.txt")
	if err != nil {
		fmt.Println("Error: %w", err)
	}
	fmt.Println("The solution is:")
	fmt.Printf("%d", solution)
}

type Segment struct {
	Lower int
	Upper int
}

func part1(filename string) (int, error) {
	ranges, _ := utils.ReadAndParse(filename, func(s string) ([]Segment, error) {
		splits := strings.Split(s, ",")
		ranges := make([]Segment, 0)
		for _, split := range splits {
			secondSplit := strings.Split(split, "-")
			lower, _ := strconv.Atoi(secondSplit[0])
			upper, _ := strconv.Atoi(secondSplit[1])
			ranges = append(ranges, Segment{
				Lower: lower,
				Upper: upper,
			})
		}
		return ranges, nil
	})
	fmt.Printf("Parsed ranges %v\n", ranges)

	result := 0
	for _, segments := range ranges {
		for _, segment := range segments {
			invalidNumbers := invalidNumbersInRange(segment)
			for _, number := range invalidNumbers {
				result += number
			}
		}
	}

	return result, nil
}

func invalidNumbersInRange(r Segment) []int {
	result := make([]int, 0)

	for i := r.Lower; i <= r.Upper; i++ {
		if numberIsRepeatedPart2(i) {
			result = append(result, i)
		}
	}
	return result
}

func numberIsRepeated(i int) bool {
	intAsString := strconv.Itoa(i)
	if len(intAsString)%2 != 0 {
		return false
	}
	left := intAsString[:len(intAsString)/2]
	right := intAsString[len(intAsString)/2:]
	return left == right
}

func numberIsRepeatedPart2(i int) bool {
	intAsString := strconv.Itoa(i)
	for count := range len(intAsString) + 1 {
		// fmt.Printf("Considering count %d for string %s\n", count, intAsString)
		if numberIsRepeatedWithCount(intAsString, count) {
			return true
		}
	}
	return false
}

func numberIsRepeatedWithCount(s string, count int) bool {
	if count == 0 || count == 1 {
		return false
	}
	// if len(s) == count {
	// 	return false
	// }
	if len(s)%count != 0 {
		return false
	}
	segmentLength := len(s) / count
	segments := make([]string, 0)
	for i := range count {
		segment := s[i*segmentLength : (i+1)*segmentLength]
		// fmt.Printf("We get segment %s with segmentLength %d, and i %d\n", segment, segmentLength, i)
		segments = append(segments, segment)
	}
	// fmt.Printf("We parsed it as segments %#v\n", segments)

	initial := segments[0]
	for _, segment := range segments {
		if initial != segment {
			// fmt.Printf("Exiting with a difference %s %s\n", initial, segment)
			return false
		}
	}
	// fmt.Printf("SUCCESSS for %s, %d\n", s, count)
	return true
}
