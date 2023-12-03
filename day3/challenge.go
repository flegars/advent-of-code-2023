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
	re := regexp.MustCompile(`[!@#\$%^&*()_+\-=\[\]{};':"\\|,<>\/?]`)
	var values []string

	for i, line := range lines {
		characters := re.FindAllStringIndex(line, len(line))
		
		for _, character := range characters {
			value := strings.Split(lines[i - 1], "")[character[0] - 1]

			if convertedValue, err := strconv.Atoi(value); err == nil {
				if 
			}
		}
	}

	total := 0

	fmt.Println(values)
	fmt.Println(total)
}