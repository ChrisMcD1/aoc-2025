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
	var result int
	selected := make([]int, 12)
	// What if we assume that the first 12 are good, and then work from there?
	// I think we need to find the higest first
	securedIndexes := 0
	for target := 0; target < 12; target++ {
		for i := securedIndexes; i < len(row)-(11-target); i++ {
			rowValue := row[i]
			if rowValue > selected[target] {
				selected[target] = rowValue
				securedIndexes = i + 1
			}
		}
	}
	for i, val := range selected {
		result += val * int(math.Pow10(11-i))
	}
	// fmt.Printf("Selected %v and got result %d for row %v\n", selected, result, row)
	return result
}
