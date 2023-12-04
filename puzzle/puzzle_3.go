package puzzle

import (
	"core"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"golang.org/x/exp/slices"
)

type Day3 struct {
	core.Puzzle
}

func (p *Day3) Puzzle3(path string) ([]int, int) {
	str := p.GetData(path)
	lines := strings.Split(str, "\n")
	re1 := regexp.MustCompile(`[!@#\$%^&*()_+\-=\[\]{};':"\\|,<>\/?]`)
	re2 := regexp.MustCompile(`\d+`)
	var total []int
	
	for i, line := range lines {
		fmt.Println("Line #", i + 1)
		numbers := re2.FindAllStringIndex(line, len(line))
		var indexes [][]int
		var baseIndex [][]int

		for _, number := range numbers {
			indexes = append(indexes, generateIndexes(number[0] - 1, number[1] + 1))
			baseIndex = append(baseIndex, generateIndexes(number[0], number[1] - 1))
		}

		for j, indexNumber := range indexes {
			stop := false
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
						stop = true
						break
					}
				}
			}

			if (stop) {
				continue
			}

			lineCurrent := lines[i]	
			for _, specialCharacter := range re1.FindAllStringIndex(lineCurrent, len(lineCurrent)) {
				if slices.Index(indexNumber, specialCharacter[1]) == 0 {
					continue
				}

				if slices.Contains(indexNumber, specialCharacter[1]) {
					splittedLine := strings.Split(line, "")
					val := ""

					for _, foundNumber := range baseIndex[j] {
						val += splittedLine[foundNumber]
					}

					convertVal, _ := strconv.Atoi(val)
					total = append(total, convertVal)
					stop = true
					break
				}
			}

			if (stop) {
				continue
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
					}
				}
			}
		}		
	}

	arr := []int{}
	result := 0

	for _, r := range total {
		arr = append(arr, r)
		result += r
	}

	fmt.Println(arr)
	fmt.Println(result)

	return arr, result
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