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
	// A map from beamIndex to number of shared beams
	beams := make(map[int]int)
	for i, val := range fileData[0] {
		if val == 'S' {
			beams[i] = 1
		}
	}
	// Loop through each row, keeping track of a set of currently active beams
	for _, line := range fileData[1:] {
		newBeams := make(map[int]int)
		for beam, numberOfBeams := range beams {
			if line[beam] == '^' {
				newBeams[beam-1] += numberOfBeams
				newBeams[beam+1] += numberOfBeams
			} else if line[beam] == '.' {
				newBeams[beam] += numberOfBeams
			}
		}
		beams = newBeams
	}

	result := 0
	for _, beamCount := range beams {
		result += beamCount
	}

	return result, nil
}
