package puzzle

import (
	"core"
	"fmt"
	"strconv"
	"strings"
)

type Day6B struct {
	core.Puzzle
}

func (p *Day6B) PartB(path string) int {
	str := p.GetData(path)
	lines := strings.Split(str, "\n")
	times := strings.Split(lines[0], " ")
	distances := strings.Split(lines[1], " ")

	time, _ := strconv.Atoi(times[0])
	distance, _ := strconv.Atoi(distances[0])

	total := 0

	for i := 1; i < time; i++ {
		currentDistance := (time - i) * i

		if currentDistance > distance {
			total++
		}
	}

	fmt.Println(total)

	return total
}
