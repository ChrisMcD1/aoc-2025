package main

import (
	"fmt"
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

type Move struct {
	Direction string
	Amount    int
}

func part1(filename string) (int, error) {
	moves, err := utils.ReadAndParse(filename, func(s string) (*Move, error) {
		direction := string(s[0])
		amount, err := strconv.Atoi(s[1:])
		if err != nil {
			return nil, err
		}
		return &Move{
			Direction: direction,
			Amount:    amount,
		}, nil
	})
	if err != nil {
		return 0, err
	}

	result := 0
	position := 50
	for _, line := range moves {
		for range line.Amount {
			if line.Direction == "L" {
				position = (position - 1) % 100
			} else {
				position = (position + 1) % 100
			}
			if position == 0 {
				result++
			}
		}
	}

	return result, nil
}
