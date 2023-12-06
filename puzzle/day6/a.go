package puzzle

import (
	"core"
	"fmt"
	"strings"
)

type Day6 struct {
	core.Puzzle
}

func (p *Day6) PartA(path string) {
	str := p.GetData(path)
	lines := strings.Split(str, "\n")
	times := strings.Split(lines[0], " ")
	distances := strings.Split(lines[1], " ")

	fmt.Println(times)
	fmt.Println(distances)
}
