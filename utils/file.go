package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func ReadAndParse[T any](filename string, parser func(string) (T, error)) ([]T, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result []T
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		item, err := parser(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("parse error on line %q: %w", scanner.Text(), err)
		}

		result = append(result, item)
	}

	return result, scanner.Err()
}

