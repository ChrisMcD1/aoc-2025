package main

import "testing"

func TestPart1(t *testing.T) {
	t.Run("Part 1 given", func(t *testing.T) {
		got, err := part1("./test_input.txt")
		if err != nil {
			t.Errorf("Got back an error! %s", err.Error())
		}
		if got != 4174379265 {
			t.Errorf("Expected got == 6, recieved %d", got)
		}
	})
}
