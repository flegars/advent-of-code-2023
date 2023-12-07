package day7

import (
	"core"
	"fmt"
	"sort"
	"strings"
)

type Day7A struct {
	core.Puzzle
}

type Hand struct {
	Hand string
	Type int
}

func (p *Day7A) PartA(path string) int {
	str := p.GetData(path)
	hands := strings.Split(str, "\n")
	cards := map[string]int{"A": 1, "K": 2, "Q": 3, "J": 4, "T": 5, "9": 6, "8": 7, "7": 8, "6": 9, "5": 10, "4": 11, "3": 12, "2": 13}
	types := map[string]int{"Five": 1, "Four": 2, "Full": 3, "Three": 4, "Two Pair": 5, "One Pair": 6, "High": 7}
	var resultHands []Hand

	for _, hand := range hands {
		resultHands = append(resultHands, Hand{Hand: hand,Type: getType(strings.Split(hand, " ")[0], cards, types)})
	}

	sort.Slice(resultHands, func(i, j int) bool {
		return resultHands[i].Type < resultHands[j].Type
	})

	fmt.Println(resultHands)

	for _, hand := range resultHands {
		fmt.Println(hand.Hand)
		fmt.Println(hand.Type)
	}

	return 0
}

func getType(hand string, cards map[string]int, types map[string]int) int {
	cardsCount := make(map[string]int)
	for _, card := range strings.Split(hand, "") {
    if _, exists := cardsCount[card]; !exists {
			cardsCount[card] = 1
		} else {
			cardsCount[card]++
		}
	}

	if len(cardsCount) == 1 {
		return types["Five"]
	}

	if len(cardsCount) == 2 && getMaxValue(cardsCount) == 4 {
		return types["Four"]
	}

	if len(cardsCount) == 2 && getMaxValue(cardsCount) == 3 {
		return types["Full"]
	}

	if len(cardsCount) == 3 && getMaxValue(cardsCount) == 3 {
		return types["Three"]
	}

	if len(cardsCount) == 3 {
		return types["Two Pair"]
	}

	if len(cardsCount) == 4 && getMaxValue(cardsCount) == 2 {
		return types["One Pair"]
	}

	return types["High"]
}

func getMaxValue(cardsCount map[string]int) int {
	lastIndex := ""
	maxVal := 0
	for index, cardCount := range cardsCount {
		if index != lastIndex && cardCount > cardsCount[lastIndex] {
			maxVal = cardCount
		}

		lastIndex = index
	}

	return maxVal
}