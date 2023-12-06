package puzzle

import (
	"core"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Day1 struct {
	core.Puzzle
}

func (p *Day1) PartA(path string) {
	str := p.GetData(path)
	lines := strings.Split(str, "\n")
	re := regexp.MustCompile(`\d`)
	results := 0

	for _, line := range lines {
		matches := re.FindAllString(line, -1)

		result := matches[0]
		result += matches[len(matches)-1]

		convert, _ := strconv.Atoi(result)
		results += convert
	}

	fmt.Println("Day 1: ", results)
}