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

type Tile int

const (
	empty Tile = iota
	paper
)

func part1(filename string) (int, error) {
	grid, err := utils.ReadAndParse(filename, func(s string) ([]Tile, error) {
		var values []Tile
		for _, char := range s {
			var tile Tile
			switch char {
			case '.':
				tile = empty
			case '@':
				tile = paper
			}
			values = append(values, tile)
		}
		return values, nil
	})
	if err != nil {
		return 0, err
	}

	result := 0
	thisLoopChanged := true
	for thisLoopChanged {
		thisLoopChanged = false
		for i, row := range grid {
			for j, cell := range row {
				if cell == paper && canBeReached(grid, i, j) {
					result++
					thisLoopChanged = true
					grid[i][j] = empty
				}
			}
		}
	}

	return result, nil
}

// Checks if fewer than 4 paper are in the adjacent squares
func canBeReached(grid [][]Tile, row int, col int) bool {
	maxRow := len(grid) - 1
	maxCol := len((grid)[0]) - 1
	var tilesToConsider []Tile
	if row > 0 {
		tilesToConsider = append(tilesToConsider, grid[row-1][col])
	}
	if row < maxRow {
		tilesToConsider = append(tilesToConsider, grid[row+1][col])
	}
	if col > 0 {
		tilesToConsider = append(tilesToConsider, grid[row][col-1])
	}
	if col < maxCol {
		tilesToConsider = append(tilesToConsider, grid[row][col+1])
	}
	if row > 0 && col > 0 {
		tilesToConsider = append(tilesToConsider, grid[row-1][col-1])
	}
	if row < maxRow && col < maxCol {
		tilesToConsider = append(tilesToConsider, grid[row+1][col+1])
	}
	if col > 0 && row < maxRow {
		tilesToConsider = append(tilesToConsider, grid[row+1][col-1])
	}
	if col < maxCol && row > 0 {
		tilesToConsider = append(tilesToConsider, grid[row-1][col+1])
	}

	paperInRegion := 0
	for _, tile := range tilesToConsider {
		if tile == paper {
			paperInRegion++
		}
	}
	return paperInRegion < 4
}
