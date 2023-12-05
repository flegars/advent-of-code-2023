package tests

import (
	"puzzle"
	"testing"
)

func TestInput1(t *testing.T) {
	res := new(puzzle.Day4).Puzzle4("./data/day4/input_1.txt")
	
	if res != 13 {
		t.Errorf("%d founded, want 13", res)
	}
}