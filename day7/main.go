package main

import (
	"fmt"

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

func part1(filename string) (int, error) {
	lines, _ := utils.ReadLines(filename)
	fileData := make([][]rune, len(lines))
	for i, line := range lines {
		fileData[i] = make([]rune, len(line))
		for j, char := range line {
			fileData[i][j] = char
		}
	}
	beams := make(map[int]bool)
	for i, val := range fileData[0] {
		if val == 'S' {
			beams[i] = true
		}
	}
	splitCount := 0
	// Loop through each row, keeping track of a set of currently active beams
	for _, line := range fileData[1:] {
		fmt.Printf("Beginning line %v with beams %v\n", line, beams)
		newBeams := make(map[int]bool)
		for beam, ok := range beams {
			if !ok {
				continue
			}
			if line[beam] == '^' {
				splitCount++
				newBeams[beam-1] = true
				newBeams[beam+1] = true
			} else if line[beam] == '.' {
				newBeams[beam] = true
			}
		}
		beams = newBeams
	}

	result := 0
	for _, ok := range beams {
		if ok {
			result++
		}
	}

	return splitCount, nil
}
