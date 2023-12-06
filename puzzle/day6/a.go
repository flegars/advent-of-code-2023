package puzzle

import (
	"core"
	"fmt"
	"strconv"
	"strings"
)

type Day6A struct {
	core.Puzzle
}

func (p *Day6A) PartA(path string) int {
	str := p.GetData(path)
	lines := strings.Split(str, "\n")
	times := strings.Split(lines[0], " ")
	distances := strings.Split(lines[1], " ")
	var results []int

	for i := 0; i <= len(lines); i++ {
		time, _ := strconv.Atoi(times[i])
		distance, _ := strconv.Atoi(distances[i])

		total := 0
		for j := 1; j <= time; j++ {
			currentDistance := 0

			for k := 1; k <= time - j; k++ {
				currentDistance += j 
			}

			if currentDistance > distance {
				total += 1
			}
		}
		
		results = append(results, total)
	}


	sum := results[0]
	for i := 1; i < len(results); i++ {
		if results[i] > 0 {
			sum *= results[i]
		}
	}

	fmt.Println(sum)

	return sum
}
