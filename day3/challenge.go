package day3

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"golang.org/x/exp/slices"
)

func Challenge() {
	content, _ := ioutil.ReadFile("./data/day3_input.txt")

	str := string(content)
	lines := strings.Split(str, "\n")
	re1 := regexp.MustCompile(`[!@#\$%^&*()_+\-=\[\]{};':"\\|,<>\/?]`)
	re2 := regexp.MustCompile(`\d+`)
	var total []int
	
	for i, line := range lines {
		fmt.Println("Line #", i)
		numbers := re2.FindAllStringIndex(line, len(line))
		var indexes [][]int
		var baseIndex [][]int

		for _, number := range numbers {
			indexes = append(indexes, generateIndexes(number[0] - 1, number[1] + 1))
			baseIndex = append(baseIndex, generateIndexes(number[0], number[1] - 1))
		}

		for j, indexNumber := range indexes {
			if i > 0 {
				lineBefore := lines[i - 1]
				
				for _, specialCharacter := range re1.FindAllStringIndex(lineBefore, len(lineBefore)) {
					if slices.Contains(indexNumber, specialCharacter[0]) {
	
						splittedLine := strings.Split(line, "")
						val := ""

						for _, foundNumber := range baseIndex[j] {
							val += splittedLine[foundNumber]
						}

						convertVal, _ := strconv.Atoi(val)
						total = append(total, convertVal)
						fmt.Println(convertVal)
					}
				}
			}

			lineCurrent := lines[i]
				
			for _, specialCharacter := range re1.FindAllStringIndex(lineCurrent, len(lineCurrent)) {
				if slices.Contains(indexNumber, specialCharacter[0]) {
					splittedLine := strings.Split(line, "")
					val := ""

					for _, foundNumber := range baseIndex[j] {
						val += splittedLine[foundNumber]
					}

					convertVal, _ := strconv.Atoi(val)
					total = append(total, convertVal)
					fmt.Println(convertVal)
				}
			}

			if i < len(lines) {
				lineAfter := lines[i + 1]
				
				for _, specialCharacter := range re1.FindAllStringIndex(lineAfter, len(lineAfter)) {
					if slices.Contains(indexNumber, specialCharacter[0]) {
						splittedLine := strings.Split(line, "")
						val := ""

						for _, foundNumber := range baseIndex[j] {
							val += splittedLine[foundNumber]
						}

						convertVal, _ := strconv.Atoi(val)
						total = append(total, convertVal)
						fmt.Println(convertVal)
					}
				}
			}
		}		
	}

	arr, result := removeDuplicatesAndSum(total)
	fmt.Println(arr)
	fmt.Println(result)
}

func generateIndexes(start int, end int) []int {
	if start > end {
		return nil
	}

	result := make([]int, end - start + 1)
	for i := range result {
		result[i] = start + i
	}

	return result
}

func removeDuplicatesAndSum(slice []int) ([]int, int) {
	result := []int{}
	total := 0

	for i, r := range slice {
		if i == 0 {
			total += r
			result = append(result, r)
		}

		if i > 0 && slice[i] != slice[i - 1] {
			total += r
			result = append(result, r)
		}
	}
	
	return result, total
 }