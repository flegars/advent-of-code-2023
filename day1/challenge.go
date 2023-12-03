package day1

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func Challenge() {
	content, _ := ioutil.ReadFile("./day1/input.txt")

	str := string(content)
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