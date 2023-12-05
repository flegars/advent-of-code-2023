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
	result := checkLines(lines)
	return result
}

func checkLines(lines []string) int {
	counts := make([]int, len(lines))
	total := 0

	for i := 0; i < len(lines); i++ {
		counts[i] = 1
	}

	for i, line := range lines {
		var winNums []int
		var myNums []int
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

		wins := 0
		for _, winNum := range winNums {
			if slices.Contains(myNums, winNum) {
				wins += 1
			}
		}

		for j := i + 1; j <= i + wins; j++ {
			counts[j] += counts[i]
		}

	}

	for _, val := range counts {
		total += val
	}

	fmt.Println(total)

	return total
}