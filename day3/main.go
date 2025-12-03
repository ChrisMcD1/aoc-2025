package main

import (
	"fmt"
	"math"
	"strconv"

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

type Row struct {
	values []int
}

func part1(filename string) (int, error) {
	items, err := utils.ReadAndParse(filename, func(s string) (*Row, error) {
		var values []int
		for _, char := range s {
			value, _ := strconv.Atoi(string(char))
			values = append(values, value)
		}
		return &Row{values}, nil
	})
	if err != nil {
		return 0, err
	}

	result := 0
	for _, row := range items {
		result += maxBattery(row.values)
	}

	return result, nil
}

func maxBattery(row []int) int {
	selected := make([]int, 12)
	availableIndex := 0
	for selectedIndex := range 12 {
		maximumIndex := len(row) - (12 - selectedIndex)
		// Always start in the next free index after the one we gave to the previous digit
		// Stop when we would otherwise run out of digits to assign to the yet-to-be-selected
		for i := availableIndex; i <= maximumIndex; i++ {
			if row[i] > selected[selectedIndex] {
				selected[selectedIndex] = row[i]
				availableIndex = i + 1
			}
		}
	}

	var result int
	for i, val := range selected {
		result += val * int(math.Pow10(11-i))
	}
	// fmt.Printf("Selected %v and got result %d for row %v\n", selected, result, row)
	return result
}
