package main

import (
	"fmt"
	"io/fs"
	"os"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

type Number struct {
	value   int
	digits  []rune
	incides [][]int
}

func contains[T any](s []T, v T) bool {
	for _, vS := range s {
		if reflect.DeepEqual(vS, v) {
			return true
		}
	}
	return false
}

func (number *Number) AddIndex(index []int) {
	number.incides = append(number.incides, index)
}
func makeLines(data []string) [][]rune {
	retLines := [][]rune{}
	for _, line := range data {
		retLines = append(retLines, []rune(line))
	}
	return retLines
}
func makeNumber(lines [][]rune) []Number {
	isNewNumber := true
	var retNumbers []Number
	for rowIndex, row := range lines {
		for colIndex, cell := range row {
			if unicode.IsDigit(cell) {
				if isNewNumber {
					retNumbers = append(retNumbers, Number{value: 0, digits: []rune{cell}, incides: [][]int{{rowIndex, colIndex}}})
					isNewNumber = false
				} else {
					retNumbers[len(retNumbers)-1].digits = append(retNumbers[len(retNumbers)-1].digits, cell)
				}
				if rowIndex > 0 {
					retNumbers[len(retNumbers)-1].AddIndex([]int{rowIndex - 1, colIndex})
					if colIndex > 0 {
						retNumbers[len(retNumbers)-1].AddIndex([]int{rowIndex - 1, colIndex - 1})
					}
					if colIndex < len(lines[0])-1 {
						retNumbers[len(retNumbers)-1].AddIndex([]int{rowIndex - 1, colIndex + 1})
					}
				}
				if rowIndex < len(lines)-1 {
					retNumbers[len(retNumbers)-1].AddIndex([]int{rowIndex + 1, colIndex})
					if colIndex > 0 {
						retNumbers[len(retNumbers)-1].AddIndex([]int{rowIndex + 1, colIndex - 1})
					}
					if colIndex < len(lines[0])-1 {
						retNumbers[len(retNumbers)-1].AddIndex([]int{rowIndex + 1, colIndex + 1})
					}
				}
				if colIndex > 0 {
					retNumbers[len(retNumbers)-1].AddIndex([]int{rowIndex, colIndex - 1})
				}
				if colIndex < len(lines[0])-1 {
					retNumbers[len(retNumbers)-1].AddIndex([]int{rowIndex, colIndex + 1})
				}
			} else {
				if !isNewNumber {
					retNumbers[len(retNumbers)-1].value, _ = strconv.Atoi(string(retNumbers[len(retNumbers)-1].digits))
				}
				isNewNumber = true
			}
		}
	}
	return retNumbers
}
func partOne(numbers []Number, lines [][]rune) int {
	sum := 0
	for _, number := range numbers {
		isPart := false
		for _, index := range number.incides { /* indexIndex = index of index in the list of indices, wow */
			if !(unicode.IsDigit(lines[index[0]][index[1]]) || lines[index[0]][index[1]] == '.') {
				isPart = true
				break
			}
		}
		if isPart {
			sum += number.value
		}
	}
	return sum
}
func partTwo(numbers []Number, lines [][]rune) int {
	var gearRatios []int
	sum := 0
	for rowIndex, row := range lines {
		for colIndex, cell := range row {
			if cell == '*' {
				gearParts := []Number{}
				for _, number := range numbers {
					if contains(number.incides, []int{rowIndex, colIndex}) {
						gearParts = append(gearParts, number)
					}
				}
				if len(gearParts) == 2 {
					gearRatios = append(gearRatios, gearParts[0].value*gearParts[1].value)
				}
			}
		}
	}
	for _, ratio := range gearRatios {
		sum += ratio
	}
	return sum
}
func main() {
	fileData, err := fs.ReadFile(os.DirFS(".."), "input.txt")
	if err != nil {
		fmt.Println(err)
	}
	var lines [][]rune = makeLines(strings.Split(string(fileData), "\n"))
	var numbers []Number = makeNumber(lines)

	fmt.Println(partOne(numbers, lines))
	fmt.Println(partTwo(numbers, lines))
}
