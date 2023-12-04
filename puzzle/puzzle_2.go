package puzzle

import (
	"core"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type Day2 struct {
	core.Puzzle
}

func (p *Day2) Puzzle2(path string) {
	str := p.GetData(path)
	lines := strings.Split(str, "\n")
	games := 0

	for index, line := range lines {
		reg := regexp.MustCompile(`Game \d+:`)
		res := strings.Map(func(r rune) rune {
			if unicode.IsGraphic(r) {
				return r
			}

			return -1
		}, strings.ReplaceAll(strings.ReplaceAll(reg.ReplaceAllString(line, "${1}"), " ", ""), ";", ","))
		sets := strings.Split(res, ",")

		if isGameValid(sets) {
			games += index + 1
		}
	}

	fmt.Println(games)
}

func isGameValid(game []string) bool {
	cubeLimits := map[string]int{"red": 12, "green": 13, "blue": 14}

	for _, set := range game {
		colors := strings.Split(set, ",")

		for _, color := range colors {
			reg := regexp.MustCompile(`(\d+)(\w+)`)
			res := reg.FindStringSubmatch(color)

			if len(res) > 0 {
				number, _ := strconv.Atoi(res[1])
				cubeLimit, exists := cubeLimits[res[2]]

				if exists && number > cubeLimit {
					return false
				}
			}
		}
	}

	return true
}