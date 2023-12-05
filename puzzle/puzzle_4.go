package puzzle

import (
	"core"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

type Day4 struct {
	core.Puzzle
}

func (p *Day4) Puzzle4(path string) int {
	str := p.GetData(path)
	lines := strings.Split(str, "\n")
	result := 0

	for _, line := range lines {
		var winNums []int
		var myNums []int
		lineTotal := 0
		splitted := strings.Split(strings.Split(line, ":")[1], "|")
		
		for _, winning := range strings.Split(splitted[0], " ") {
			if val, err := strconv.Atoi(strings.Trim(winning, " ")); err == nil {
				winNums = append(winNums, val)
			}
		}

		for _, myNum := range strings.Split(splitted[1], " ") {
			if val, err := strconv.Atoi(strings.Trim(myNum, " ")); err == nil {
				myNums = append(myNums, val)
			}
		}

		for _, winNum := range winNums {
			if slices.Contains(myNums, winNum) {
				if lineTotal == 0 {
					lineTotal = 1
				} else {
					lineTotal = lineTotal * 2
				}
			}
		}

		result += lineTotal
	}

	fmt.Println(result)

	return result
}