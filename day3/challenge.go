package day3

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func Challenge() {
	content, _ := ioutil.ReadFile("./day3/input.txt")

	str := string(content)
	lines := strings.Split(str, "\n")
	lineLength := len(lines[0])
	re := regexp.MustCompile(`[!@#\$%^&*()_+\-=\[\]{};':"\\|,<>\/?]`)
	var values []string

	for i, line := range lines {
		charactersIndex := re.FindStringIndex(line)

		for j, character := range charactersIndex {
			realPosition := i * lineLength + character
			
		}
	}

	total := 0

	for _, value := range values {
		if convertedValue, err := strconv.Atoi(value); err == nil {
			total += convertedValue
		}
	}

	fmt.Println(values)
	fmt.Println(total)
}