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
	None Operation = iota
	Add
	Multiply
)

type Column struct {
	values    []int
	operation Operation
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
	var columns []Column
	columnsLen := len(fileData[0])
	// scanning horizontal right to left until we find an empty column, and then we start a new
	var currentColumn Column
	for j := columnsLen - 1; j >= 0; j-- {
		var currentNumber []rune
		var currentOperation Operation
		for i := 0; i < len(fileData)-1; i++ {
			if fileData[i][j] != ' ' {
				currentNumber = append(currentNumber, fileData[i][j])
			}
		}
		targetForOperation := fileData[len(fileData)-1][j]
		if targetForOperation == '+' {
			currentOperation = Add
		} else if targetForOperation == '*' {
			currentOperation = Multiply
		}
		if len(currentNumber) != 0 {
			parsedNum, err := strconv.Atoi(string(currentNumber))
			if err != nil {
				return 0, nil
			}
			currentColumn.values = append(currentColumn.values, parsedNum)
		}
		if currentOperation != None {
			currentColumn.operation = currentOperation
		}
		if currentOperation == None && len(currentNumber) == 0 {
			columns = append(columns, currentColumn)
			currentColumn = Column{}
		}
	}
	columns = append(columns, currentColumn)

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
