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

type Operation int

const (
	Add Operation = iota
	Multiply
)

type Column struct {
	values    []int
	operation Operation
}

func part1(filename string) (int, error) {
	lines, _ := utils.ReadLines(filename)
	dataRowsLength := len(lines)
	width := len(splitWhiteSpace(lines[0]))
	columns := make([]Column, width)
	for i, line := range lines {
		if i == dataRowsLength-1 {
			for j, val := range splitWhiteSpace(line) {
				column := columns[j]
				if val == "*" {
					column.operation = Multiply
				} else if val == "+" {
					column.operation = Add
				} else {
					return 0, fmt.Errorf("Not a valid operation %s", val)
				}
				columns[j] = column
			}
		} else {
			for j, val := range splitWhiteSpace(line) {
				column := columns[j]
				valAsInt, err := strconv.Atoi(val)
				if err != nil {
					return 0, fmt.Errorf("Not a valid number %w", err)
				}
				column.values = append(column.values, valAsInt)
				columns[j] = column
			}
		}
	}

	result := 0
	for _, column := range columns {
		if column.operation == Add {
			intermediate := 0
			for _, value := range column.values {
				intermediate += value
			}
			result += intermediate
		} else if column.operation == Multiply {
			intermediate := 1
			for _, value := range column.values {
				intermediate *= value
			}
			result += intermediate
		}
	}

	return result, nil
}

func splitWhiteSpace(s string) []string {
	initial := strings.Split(s, " ")
	var result []string
	for _, val := range initial {
		if val != " " && val != "" {
			result = append(result, val)
		}
	}

	return result
}
