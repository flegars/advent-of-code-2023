package day1

import (
	"core"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Day1 struct {
	core.Challenge
}

func (c *Day1) ChallengeDay1() {
	str := c.GetData("./data/day1_input.txt")
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