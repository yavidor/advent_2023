package main

import (
	"fmt"
	"io/fs"
	"math"
	"os"
	"reflect"
	"strings"
)

type Card struct {
	winningNumbers []string
	myNumbers      []string
	value          int
}

func contains[T any](s []T, v T) bool {
	for _, vS := range s {
		if reflect.DeepEqual(vS, v) {
			return true
		}
	}
	return false
}

func NewCard(winningNumbers, myNumbers []string) *Card {
	card := Card{winningNumbers, myNumbers, 0}
	for _, v := range card.myNumbers {
		if contains(card.winningNumbers, v) {
			card.value += 1
		}
	}
	return &card
}

func makeLines(data []byte) [][][]string {
	inputLines := strings.Split(string(data), "\n")
	outputLines := make([][][]string, 0, len(inputLines))
	for _, line := range inputLines {
		lineContent := line[strings.Index(line, ":")+2:]
		fields := strings.Split(lineContent, "|")
		var subLine [][]string
		for _, field := range fields {
			subLine = append(subLine, strings.Fields(field))
		}
		outputLines = append(outputLines, subLine)

	}
	return outputLines
}

func makeCards(lines [][][]string) []Card {
	retCards := make([]Card, len(lines))
	for i, v := range lines {
		retCards[i] = *NewCard(v[0], v[1])
	}
	return retCards
}

func partOne(cards []Card) int {
	sum := 0
	for _, v := range cards {
		sum += int(math.Pow(2, float64(v.value)-1))
	}
	return sum
}

func partTwo(cards []Card) int {
	var processedCards []int = make([]int, len(cards))
	for i := range processedCards {
		processedCards[i] = 1
	}
	for i, v := range cards {
		for j := 0; j < v.value; j++ {
			processedCards[i+j+1] += processedCards[i]
		}
	}
	sum := 0
	for _, v := range processedCards {
		sum += v
	}
	return sum
}

func main() {
	fileData, err := fs.ReadFile(os.DirFS(".."), "input.txt")
	if err != nil {
		fmt.Println(err)
	}
	var lines [][][]string = makeLines(fileData)
	var cards []Card = makeCards(lines)
	fmt.Println(partOne(cards))
	fmt.Println(partTwo(cards))
}
