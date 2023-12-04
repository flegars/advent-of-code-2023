package tests

import (
	"puzzle"
	"testing"
)

//[619 665 406 10 985 944 296 169 686 290 767 977 873 665 381 17 246 861 104 208 320 529 676 89 112] => 11895
func TestDay3Input1(t *testing.T) {
	_, res := new(puzzle.Day3).Puzzle3("./data/day3/input_1.txt")
	
	if res != 12560 {
		t.Errorf("%d founded, want 12560", res)
	}
}

func TestDay3Input2(t *testing.T) {
	_, res := new(puzzle.Day3).Puzzle3("./data/day3/input_2.txt")
	
	if res != 4361 {
		t.Errorf("%d founded, want 4361", res)
	}
}